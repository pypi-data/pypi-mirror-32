import logging
import abc, time
import numpy as np
import scipy.sparse as spp
from contextlib import contextmanager

import indigo.operators as op
from indigo.util import profile

log = logging.getLogger(__name__)

class Backend(object):
    """
    Provides the routines and data structures necessary to implement
    a linear operator chain on different platforms.
    """
    __metaclass__ = abc.ABCMeta

    def __init__(self, device_id=0):
        profile._backend = self

    class dndarray(object):
        """
        N-dimensional array in device memory.

        Parameters
        ----------
        backend : indigo.backends.Backend
            Backend instance.
        shape : tuple
            Array shape, a la numpy.
        dtype : numpy.dtype
            Datatype.
        ld : tuple
            Shape of array before slicing, used for ldb/ldc values.
        own : bool
            True if this object malloc'ed the underlying memory.
        data : ?
            Handle to underlying memory.
        """
        __metaclass__ = abc.ABCMeta
        _memory = dict()

        def __init__(self, backend, shape, dtype,
                     ld=None, own=True, data=None, name=''):
            assert isinstance(shape, (tuple,list))
            self.dtype = dtype
            self.shape = shape
            self._backend = backend
            self._leading_dim = ld or shape[0]
            self._own = own
            #assert isinstance(backend, Backend), (backend, type(backend))
            if data is None:
                self._arr = self._malloc(shape, dtype)
                self._memory[ id(self._arr) ] = (name, shape, dtype)
            else:
                self._arr = data

        def reshape(self, new_shape):
            old_shape = self.shape
            old_leading_dim = self._leading_dim
            if (-1) in new_shape:
                one = new_shape.index(-1)
                new_size = -int(np.prod(new_shape))
                old_size = self.size
                factor = old_size // new_size
                assert new_size * factor == old_size, \
                    "Cannot reshape {} into {}. (size mismatch)".format(old_shape, new_shape)
                new_shape = list(new_shape)
                new_shape[one] = factor
                new_shape = tuple(new_shape)
            if new_shape[0] > old_shape[0]:
                contig = old_shape[0] == self._leading_dim
                assert contig, "Cannot stack non-contiguous columns."
            assert np.prod(new_shape) == self.size
            # min for Kron -- make new lda
            # max for VStack -- preserve original lda
            #new_leading_dim = min(new_shape[0], old_leading_dim) # FIXME: need consistent semantics for reshape
            #new_leading_dim = old_leading_dim # works with VStack
            #new_leading_dim = new_shape[0] # works with Kron

            if new_shape[0] < old_shape[0]:
                #assert self.contiguous, "Cannot stack vectors of non-contiguous matrix."
                new_leading_dim = new_shape[0]
            else:
                new_leading_dim = old_leading_dim

            return self._backend.dndarray( self._backend,
                new_shape, dtype=self.dtype, ld=new_leading_dim, own=False, data=self._arr)

        @property
        def size(self):
            return np.prod(self.shape)

        @property
        def itemsize(self):
            return self.dtype.itemsize

        @property
        def nbytes(self):
            return self.size * np.dtype(self.dtype).itemsize

        @property
        def ndim(self):
            return len(self.shape)

        @property
        def contiguous(self):
            if self.ndim == 1:
                return True
            else:
                return self._leading_dim == self.shape[0]

        def copy_from(self, arr):
            ''' copy from device when both arrays exist '''
            assert isinstance(arr, np.ndarray)
            if self.size != arr.size:
                raise ValueError("size mismatch, expected {} got {}" \
                    .format(self.shape, arr.shape))
            if self.dtype != arr.dtype:
                raise TypeError("dtype mismatch, expected {} got {}" \
                    .format(self.dtype, arr.dtype))
            if not arr.flags['F_CONTIGUOUS']:
                raise TypeError("order mismatch, expected 'F' got {}" \
                    .format(arr.flags['F_CONTIGUOUS']))
            self._copy_from(arr)

        def copy_to(self, arr):
            ''' copy to device when both arrays exist '''
            assert isinstance(arr, np.ndarray)
            if self.size != arr.size:
                raise ValueError("size mismatch, expected {} got {}" \
                    .format(self.shape, arr.shape))
            if self.dtype != arr.dtype:
                raise TypeError("dtype mismatch, expected {} got {}" \
                    .format(self.dtype, arr.dtype))
            self._copy_to(arr)

        def to_host(self):
            ''' copy from device when host array doesnt exist '''
            arr = np.ndarray(self.shape, self.dtype, order='F')
            self.copy_to(arr)
            return arr

        @contextmanager
        def on_host(self):
            arr_h = self.to_host()
            yield arr_h
            self.copy_from(arr_h)

        def copy(self, other=None, name=''):
            ''' copy array on device'''
            if other:
                assert isinstance(other, self._backend.dndarray)
                self._copy(other)
            else:
                other = self._backend.zero_array(self.shape, self.dtype, name=name)
                other._copy(self)
                return other

        @classmethod
        def to_device(cls, backend, arr, name=''):
            ''' copy to device when device array doesnt exist '''
            arr_f = np.require(arr, requirements='F')
            d_arr = cls(backend, arr.shape, arr.dtype, name=name)
            d_arr.copy_from(arr_f)
            return d_arr

        def __del__(self):
            """ destructor """
            if self._own and hasattr(self, '_arr'):
                self._memory.pop( id(self._arr) )
                self._free()

        def __setitem__(self, slc, other):
            #FIXME don't ignore slc
            assert not(slc.start or slc.stop), "dndarray setitem cant slice"
            self._copy(other)

        @abc.abstractmethod
        def __getitem__(self, slc):
            """
            Slice notation. Slices must be contiguous in memory. Returns a view.
            """
            raise NotImplementedError()

        @abc.abstractmethod
        def _copy_from(self, arr):
            """ copy HtoD implementation """
            raise NotImplementedError()

        @abc.abstractmethod
        def _copy_to(self, arr):
            """ copy DtoH implementation """
            raise NotImplementedError()

        @abc.abstractmethod
        def _copy(self, arr):
            """ copy DtoD implementation """
            raise NotImplementedError()

        @abc.abstractmethod
        def _malloc(self, shape, dtype):
            """ malloc implementation """
            raise NotImplementedError()

        @abc.abstractmethod
        def _free(self):
            """ malloc implementation """
            raise NotImplementedError()

        @abc.abstractmethod
        def _zero(self):
            """ set to zero """
            raise NotImplementedError()

        @staticmethod
        def from_param(obj):
            """ convert _arr into ctypes object """
            raise NotImplementedError()

    def copy_array(self, arr, name=''):
        return self.dndarray.to_device(self, arr, name=name)

    def zero_array(self, shape, dtype, name=''):
        d_arr = self.empty_array(shape, dtype, name=name)
        d_arr._zero()
        return d_arr

    def zeros_like(self, other, name=''):
        return self.zero_array(other.shape, other.dtype, name=name)

    def empty_array(self, shape, dtype, name=''):
        d_arr = self.dndarray(self, shape, dtype, name=name)
        return d_arr

    def rand_array(self, shape, dtype=np.dtype('complex64'), name=''):
        x = np.random.random(shape) + 1j*np.random.random(shape)
        x = np.require(x, dtype=np.dtype('complex64'), requirements='F')
        x_d = self.copy_array(x, name=name)
        return x_d

    def get_max_threads(self):
        return 1

    def barrier(self):
        pass

    def mem_usage(self):
        nbytes = 0
        log.info("Memory report:")
        table = []
        for name, shape, dtype in self.dndarray._memory.values():
            n = np.prod(shape) * dtype.itemsize
            table.append( (name, n, shape, dtype) )
            nbytes += n
        for name, n, shape, dtype in sorted(table, key=lambda tup: tup[1]):
            if n > 1e6:
                log.info("  %40s: % 3.0f MB, %20s, %15s", name, n/1e6, shape, dtype)
        return nbytes

    @contextmanager
    def scratch(self, shape=None, nbytes=None):
        assert not (shape is not None and nbytes is not None), \
            "Specify either shape or nbytes to backend.scratch()."
        if nbytes is not None:
            shape = (nbytes//np.dtype('complex64').itemsize,)
        size = np.prod(shape)
        if hasattr(self, '_scratch'):
            pos = self._scratch_pos
            total = self._scratch.size
            assert pos + size <= total, "Not enough scratch memory (wanted %d MB, but only have %d MB available of %d MB total)." % (size/1e6, (total-pos)/1e6, total/1e6)
            mem = self._scratch[pos:pos+size].reshape(shape)
            self._scratch_pos += size
            yield mem
            self._scratch_pos -= size
        else:
            log.debug("dynamically allocating scratch space in shape %s", shape)
            mem = self.zero_array(shape, dtype=np.complex64)
            yield mem
            del mem

    # -----------------------------------------------------------------------
    # Operator Building Interface 
    # -----------------------------------------------------------------------

    def SpMatrix(self, M, **kwargs):
        """ A := M """
        assert isinstance(M, spp.spmatrix)
        return op.SpMatrix(self, M, **kwargs)

    def DenseMatrix(self, M, **kwargs):
        """ A := M """
        assert isinstance(M, np.ndarray)
        assert M.ndim == 2
        return op.DenseMatrix(self, M, **kwargs)

    def Diag(self, v, **kwargs):
        """ A := diag(v) """
        v = np.require(v, requirements='F')
        if v.ndim > 1:
            v = v.flatten(order='A')
        dtype = kwargs.get('dtype', np.dtype('complex64'))
        M = spp.diags( v, offsets=0 ).astype(dtype)
        return self.SpMatrix(M, **kwargs)

    def Adjoint(self, A, **kwargs):
        """ C := A^H """
        return op.Adjoint(self, A, **kwargs)

    def KronI(self, c, B, **kwargs):
        """ C := I_c (KRON) B """
        I = self.Eye(c)
        return op.Kron(self, I, B, **kwargs)

    def Kron(self, A, B, **kwargs):
        """ C := A (KRON) B """
        return op.Kron(self, A, B, **kwargs)

    def BlockDiag(self, Ms, **kwargs):
        return op.BlockDiag(self, *Ms, **kwargs)

    def VStack(self, Ms, **kwargs):
        return op.VStack(self, *Ms, **kwargs)

    def HStack (self, Ms, **kwargs):
        return op.HStack(self, *Ms, **kwargs)

    def UnscaledFFT(self, shape, dtype, **kwargs):
        """ A := FFT{ . } """
        return op.UnscaledFFT(self, shape, dtype, **kwargs)

    def Eye(self, n, dtype=np.dtype('complex64'), **kwargs):
        """ A := I_n """
        return op.Eye(self, n, dtype=dtype, **kwargs)

    def One(self, shape, dtype=np.dtype('complex64'), **kwargs):
        """ A := [1] (matrix of ones) """
        return op.One(self, shape, dtype=dtype, **kwargs)

    def CopyIn(self, shape, dtype, **kwargs):
        return op.CopyIn(self, shape, dtype)

    def CopyOut(self, shape, dtype, **kwargs):
        return op.CopyOut(self, shape, dtype)

    def FFT(self, shape, dtype, **kwargs):
        """ Unitary FFT """
        n = np.prod(shape)
        s = np.ones(n, order='F', dtype=dtype) / np.sqrt(n)
        S = self.Diag(s, name='scale')
        F = self.UnscaledFFT(shape, dtype, **kwargs)
        return S*F

    def FFTc(self, ft_shape, dtype, normalize=True, **kwargs):
        """ Centered, Unitary FFT """
        mod_slice = [ slice(d) for d in ft_shape ]
        idx = np.mgrid[mod_slice]
        mod = 0
        for i in range(len(ft_shape)):
            c = ft_shape[i] // 2
            mod += (idx[i] - c / 2.0) * (c / ft_shape[i])
        mod = np.exp(1j * 2.0 * np.pi * mod).astype(dtype)
        M = self.Diag(mod, name='mod')
        if normalize:
            F = self.FFT(ft_shape, dtype=dtype, **kwargs)
        else:
            F = self.UnscaledFFT(ft_shape, dtype=dtype, **kwargs)
        return M*F*M

    def Zpad(self, M, N, mode='center', dtype=np.dtype('complex64'), **kwargs):
        slc = []
        if mode == 'center':
            for m, n in zip(M, N):
                slc += [slice(m // 2 + int(np.ceil(-n / 2)),
                              m // 2 + int(np.ceil( n / 2))), ]
        elif mode == 'edge':
            for m, n in zip(M, N):
                slc.append(slice(n))
            pass
        x = np.arange( np.prod(M), dtype=int ).reshape(M, order='F')
        rows = x[slc].flatten(order='F')
        cols = np.arange(rows.size)
        ones = np.ones_like(cols)
        shape = np.prod(M), np.prod(N)
        M = spp.coo_matrix( (ones, (rows,cols)), shape=shape, dtype=dtype )
        return self.SpMatrix(M, **kwargs)

    def Crop(self, M, N, dtype=np.dtype('complex64'), **kwargs):
        return self.Zpad(N, M, dtype=dtype, **kwargs).H

    def Interp(self, N, coord, width, table, dtype=np.dtype('complex64'), **kwargs):
        assert len(N) == 3
        ndim  = coord.shape[0]
        npts = np.prod( coord.shape[1:] )
        coord = coord.reshape((ndim,-1), order='F')

        from indigo.interp import interp_mat
        M = interp_mat(npts, N, width, table, coord, 1).astype(dtype)

        return self.SpMatrix(M, **kwargs)

    def NUFFT(self, M, N, coord, width=3, n=128, oversamp=None, dtype=np.dtype('complex64'), **kwargs):
        assert len(M) == 3
        assert len(N) == 3
        assert M[1:] == coord.shape[1:]

        # target 448 x 270 x 640
        #   448 x 270 x 640   mkl-batch: 170.83 ms, 237.51 gflop/s  back-to-back: 121.76 ms, 333.23 gflop/s
        #   1.45  1.30  1.33
        #   432 x 280 x 640   mkl-batch: 183.85 ms  220.7 gflop/s   back-to-back: 149.62 ms  271.19 gflop/s
        #   1.40  1.35  1.33
        #   432 x 270 x 640   mkl-batch: 168.62 ms  231.57 gflop/s  back-to-back: 118.31 ms  330.05 gflop/s
        #   1.40  1.30  1.33

        if isinstance(oversamp, tuple):
            omin = min(oversamp)
        else:
            omin = oversamp
            oversamp = (omin, omin, omin)

        import scipy.signal as signal
        from indigo.noncart import rolloff3
        ndim  = coord.shape[0]
        npts  = np.prod( coord.shape[1:] )

        oN = list(N)
        for i in range(3):
            oN[i] *= oversamp[i]
        oN = tuple(int(on) for on in oN)

        Z = self.Zpad(oN, N, dtype=dtype, name='zpad')
        F = self.FFTc(oN, dtype=dtype, name='fft')

        beta = np.pi * np.sqrt(((width * 2. / omin) * (omin- 0.5)) ** 2 - 0.8)
        kb = signal.kaiser(2 * n + 1, beta)[n:]
        G = self.Interp(oN, coord, width, kb, dtype=np.float32, name='interp')

        r = rolloff3(omin, width, beta, N)
        R = self.Diag(r, name='apod')

        return G*F*Z*R

    def Convolution(self, kernel, normalize=True, name='noname'):
        F = self.FFTc(kernel.shape, name='%s.convF' % name, normalize=normalize, dtype=np.complex64)
        K = self.Diag(F * kernel, name='%s.convK' % name)
        I = self.Eye(F.shape[0])
        return F.H * K * F

    # -----------------------------------------------------------------------
    # BLAS Routines
    # -----------------------------------------------------------------------
    def axpby(self, beta, y, alpha, x):
        """ y = beta * y + alpha * x """
        raise NotImplementedError()

    def dot(self, x, y):
        """ returns x^T * y """
        raise NotImplementedError()

    def norm2(self, x):
        """ returns ||x||_2"""
        raise NotImplementedError()

    def scale(self, x, alpha):
        """ x *= alpha """
        raise NotImplementedError()

    def pdot(self, x, y, comm):
        xHy = self.dot(x, y)
        if comm is not None:
            xHy = comm.allreduce( xHy )
        return xHy

    def pnorm2(self, x, comm):
        xTx = self.norm2(x)
        if comm is not None:
            xTx = comm.allreduce( xTx )
        return xTx

    def cgemm(self, y, M, x, alpha, beta, forward):
        """
        Peform a dense matrix-matrix multiplication.
        """
        raise NotImplementedError()

    def csymm(self, y, M, x, alpha, beta, left=True):
        """
        Peform a symmetric dense matrix-matrix multiplication for real symmetric matrices.
        """
        raise NotImplementedError()

    # -----------------------------------------------------------------------
    # FFT Routines
    # -----------------------------------------------------------------------

    @abc.abstractmethod
    def fftn(self, y, x):
        """
        Peform an unscaled multidimensional forward FFT on x.
        """
        raise NotImplementedError()

    @abc.abstractmethod
    def ifftn(self, y, x):
        """
        Peform an unscaled multidimensional inverse FFT on x.
        """
        raise NotImplementedError()

    def _fft_workspace_size(self, x_shape):
        return 0

    @abc.abstractmethod
    def ccsrmm(self, y, A_shape, A_indx, A_ptr, A_vals, x, alpha=1, beta=0, adjoint=False, exwrite=False):
        """
        Computes Y[:] = A * X.
        """
        raise NotImplementedError()

    @abc.abstractmethod
    def cdiamm(self, y, shape, offsets, data, x, alpha=1.0, beta=0.0, adjoint=True):
        """
        Computes Y[:] = A * X.
        """
        raise NotImplementedError()

    @abc.abstractmethod
    def onemm(self, y, x, alpha=1, beta=0):
        """
        Computes Y[:] = beta * Y + alpha * [1] * X.
        """
        raise NotImplementedError()

    class csr_matrix(object):
        """
        A device-resident sparse matrix in CSR format.
        """
        _index_base = 0

        def __init__(self, backend, A, name='mat'):
            """
            Create a matrix from the given `scipy.sparse.sppmatrix`.
            """
            if not isinstance(A, spp.csr_matrix):
                A = A.tocsr()
            A = self._type_correct(A)
            self._backend = backend
            self.rowPtrs = backend.copy_array(A.indptr + self._index_base, name=name+".rowPtrs")
            self.colInds = backend.copy_array(A.indices + self._index_base, name=name+".colInds")
            self.values  = backend.copy_array(A.data, name=name+".data")
            self.shape = A.shape
            self.dtype = A.dtype

            # fraction of nonzero rows/columns
            try:
                from indigo.backends._customcpu import inspect
                nzrow, nzcol, self._exwrite = inspect(A.shape[0], A.shape[1], A.indices, A.indptr)
                self._row_frac = nzrow / A.shape[0]
                self._col_frac = nzcol / A.shape[1]
                log.debug("matrix %s has %2d%% nonzero rows and %2d%% nonzero columns",
                    name, 100*self._row_frac, 100*self._col_frac)
                log.debug("matrix %s supports exwrite: %s", name, self._exwrite)
            except ImportError:
                self._row_frac = 1.0
                self._col_frac = 1.0
                log.debug("skipping exwrite inspection. Is CustomCPU backend available?")

        def forward(self, y, x, alpha=1, beta=0):
            """ y[:] = A * x """
            assert x.dtype == np.dtype("complex64"), "Bad dtype: expected compelx64, got %s" % x.dtype
            assert y.dtype == np.dtype("complex64"), "Bad dtype: expected compelx64, got %s" % y.dtype
            assert self.values.dtype == np.dtype("complex64")
            self._backend.ccsrmm(y,
                self.shape, self.colInds, self.rowPtrs, self.values,
                x, alpha=alpha, beta=beta, adjoint=False, exwrite=True)

        def adjoint(self, y, x, alpha=1, beta=0):
            """ y[:] = A.H * x """
            assert x.dtype == np.dtype("complex64"), "Bad dtype: expected compelx64, got %s" % x.dtype
            assert y.dtype == np.dtype("complex64"), "Bad dtype: expected compelx64, got %s" % y.dtype
            assert self.values.dtype == np.dtype("complex64")
            self._backend.ccsrmm(y,
                self.shape, self.colInds, self.rowPtrs, self.values,
                x, alpha=alpha, beta=beta, adjoint=True, exwrite=self._exwrite)

        @property
        def nbytes(self):
            return self.rowPtrs.nbytes + self.colInds.nbytes + self.values.nbytes

        @property
        def nnz(self):
            return self.values.size

        def _type_correct(self, A):
            return A.astype(np.complex64)


    class dia_matrix(object):
        """
        A device-resident sparse matrix in DIA format.
        """
        def __init__(self, backend, A, name='mat'):
            """
            Create a matrix from the given `scipy.sparse.sppmatrix`.
            """
            assert isinstance(A, spp.dia_matrix)
            A = A.astype(np.complex64)
            self._backend = backend
            self.data = backend.copy_array(A.data.T, name=name+".data")
            self.offsets = backend.copy_array(A.offsets, name=name+".data")
            self.shape = A.shape
            self.dtype = A.dtype
            self._row_frac = 1
            self._col_frac = 1

        def forward(self, y, x, alpha=1, beta=0):
            """ y[:] = A * x """
            self._backend.cdiamm(y, self.shape, self.offsets, self.data,
                x, alpha=alpha, beta=beta, adjoint=False)

        def adjoint(self, y, x, alpha=1, beta=0):
            """ y[:] = A.H * x """
            self._backend.cdiamm(y, self.shape, self.offsets, self.data,
                x, alpha=alpha, beta=beta, adjoint=True)

        @property
        def nbytes(self):
            return self.offsets.nbytes + self.data.nbytes

        @property
        def nnz(self):
            return self.data.size

    # -----------------------------------------------------------------------
    # Algorithms
    # -----------------------------------------------------------------------

    def cg(self, A, b_h, x_h, lamda=0.0, tol=1e-10, maxiter=100, team=None):
        """
        Conjugate gradient. Solves for A x = b, where A is positive semi-definite.

        Parameters
        ----------
        A : function to perform A(x)
        y : 1D array
        x : 1D array, initial solution
        maxiter : int, optional
        {IterPrint, IterPlot, IterWrite, IterCompare}
        """
        x  = self.copy_array( x_h, name='x' )
        b  = self.copy_array( b_h, name='b' )
        Ap = x.copy()

        # r = b - A(x) - lamda * x
        r = b
        A.eval(Ap, x)

        self.axpby(1, r, -1, Ap)
        self.axpby(1, r, -lamda, x)

        p = r.copy(name='p')
        rr = self.pnorm2(r, team)
        r0 = rr

        for it in range(maxiter):
            profile.extra['it'] = it
            with profile("iter"):
                A.eval(Ap, p)
                self.axpby(1, Ap, lamda, p)
                alpha = rr / self.pdot(p, Ap, team)
                self.axpby(1, x, alpha, p)
                self.axpby(1, r, -alpha, Ap)

                r2 = self.pnorm2(r, team)
                beta = r2 / rr
                self.scale(p, beta)
                self.axpby(1, p, 1, r)
                rr = r2

                resid = np.sqrt(rr / r0)
                log.info("iter %d, residual %g", it, resid.real)

                if resid < tol:
                    log.info("cg reached tolerance")
                    break
        else:
            log.info("cg reached maxiter")
        x.copy_to(x_h)

    def apgd(self, gradf, proxg, alpha, x_h, maxiter=100, team=None):
        '''Accelerated proximal gradient descent.
        Solves for min_x f(x) + g(x)

        Parameters
        ----------
        gradf : Gradient of f
        proxg : Proximal of g
        alpha : Step size
        x0 : 1D array, initial solution
        maxiter : int, optional
        '''
        x_k = self.copy_array(x_h)
        y_k = x_k.copy()
        y_k1 = x_k.copy()
        x_k1 = x_k.copy()

        gf = x_k.copy()

        t_k = 1

        for it in range(1,maxiter+1):
            profile.extra['it'] = it

            with profile("iter"):
                gradf(gf, y_k)
                self.axpby(1, x_k, -alpha, gf)

                proxg(x_k, alpha)

                t_k1 = (1.0 + np.sqrt(1.0 + 4.0 * t_k**2)) / 2.0

                t_ratio = (t_k - 1) / t_k1
                self.axpby(0, y_k1, 1+t_ratio, x_k)
                self.axpby(1, y_k1,  -t_ratio, x_k1)

                x_k1.copy(x_k)
                y_k.copy(y_k1)

            log.info("iter %d", it)

        x_k.copy_to(x_h)

    def max(self, val, arr):
        """ Computes elementwise maximum: arr[:] = max(arr, val). """
        raise NotImplementedError()
