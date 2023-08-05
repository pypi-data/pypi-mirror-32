import pytest
import numpy as np
import scipy.sparse as spp
import numpy.testing as npt
from itertools import product
from scipy.signal import fftconvolve

import indigo
from indigo.backends import available_backends
BACKENDS = available_backends()

@pytest.mark.parametrize("backend,M,N,K,density,alpha,beta",
    product( BACKENDS, [23,45], [45,23], [1,8,9,17], [0.01,0.1,0.5,1], [0,.5,1], [0,.5,1] ))
def test_SpMatrix(backend, M, N, K, density, alpha, beta):
    b = backend()
    A_h = indigo.util.randM(M, N, density)
    A = b.SpMatrix(A_h)

    # forward
    x = b.rand_array((N,K))
    y = b.rand_array((M,K))
    y_exp = beta * y.to_host() + alpha * A_h * x.to_host()
    A.eval(y, x, alpha=alpha, beta=beta)
    npt.assert_allclose(y.to_host(), y_exp, rtol=1e-5)

    # adjoint
    x = b.rand_array((M,K))
    y = b.rand_array((N,K))
    y_exp = beta * y.to_host() + alpha * A_h.H * x.to_host()
    A.H.eval(y, x, alpha=alpha, beta=beta)
    npt.assert_allclose(y.to_host(), y_exp, rtol=1e-5)

    # double adjoint
    x = b.rand_array((N,K))
    y = b.rand_array((M,K))
    y_exp = beta * y.to_host() + alpha * A_h.H.H * x.to_host()
    A.H.H.eval(y, x, alpha=alpha, beta=beta)
    npt.assert_allclose(y.to_host(), y_exp, rtol=1e-5)

    # triple adjoint
    x = b.rand_array((M,K))
    y = b.rand_array((N,K))
    y_exp = beta * y.to_host() + alpha * A_h.H.H.H * x.to_host()
    A.H.H.H.eval(y, x, alpha=alpha, beta=beta)
    npt.assert_allclose(y.to_host(), y_exp, rtol=1e-5)

    # shape
    assert A.shape == (M,N)
    assert A.H.shape == (N,M)

    # dtype
    assert A.dtype == np.dtype('complex64')


@pytest.mark.parametrize("backend,L,M,N,K,density,alpha,beta",
    product( BACKENDS, [3,4], [5,6], [7,8], [1,8,9,17], [0.01,0.1,0.5,1], [0,.5,1], [0,.5,1] ))
def test_Product(backend, L, M, N, K, density, alpha, beta):
    b = backend()
    A0_h = indigo.util.randM(L, M, density)
    A1_h = indigo.util.randM(M, N, density)
    A0 = b.SpMatrix(A0_h, name='A0')
    A1 = b.SpMatrix(A1_h, name='A1')
    A = A0 * A1

    # forward
    x = b.rand_array((N,K))
    y = b.rand_array((L,K))
    y_exp = beta * y.to_host() + alpha * A0_h @ (A1_h @ x.to_host())
    A.eval(y, x, alpha=alpha, beta=beta)
    npt.assert_allclose(y.to_host(), y_exp, rtol=1e-5)

    # adjoint
    x = b.rand_array((L,K))
    y = b.rand_array((N,K))
    y_exp = beta * y.to_host() + alpha * A1_h.H @ (A0_h.H @ x.to_host())
    A.H.eval(y, x, alpha=alpha, beta=beta)
    npt.assert_allclose(y.to_host(), y_exp, rtol=1e-5)

    # shape
    assert A.shape == (L,N)
    assert A.H.shape == (N,L)

    # dtype
    assert A.dtype == np.dtype('complex64')


@pytest.mark.parametrize("backend,stack,M,N,K,density,alpha,beta",
    product( BACKENDS, [1,2,3], [5,6], [7,8], [1,4,8,9,17], [0.1,0.5,1], [0,.5,1], [0,1,0.5] ))
def test_VStack(backend, stack, M, N, K, density, alpha, beta):
    b = backend()
    mats_h = [indigo.util.randM(M,N,density) for i in range(stack)]
    A_h = spp.vstack(mats_h)

    mats_d = [b.SpMatrix(m) for m in mats_h]
    A = b.VStack(mats_d)

    # forward
    x = b.rand_array((A.shape[1],K))
    y = b.rand_array((A.shape[0],K))
    y_exp = beta * y.to_host() + alpha * A_h @ x.to_host()
    A.eval(y, x, alpha=alpha, beta=beta)
    npt.assert_allclose(y.to_host(), y_exp, rtol=1e-5)

    # adjoint
    x = b.rand_array((A.shape[0],K))
    y = b.rand_array((A.shape[1],K))
    y_exp = beta * y.to_host() + alpha * A_h.H @ x.to_host()
    A.H.eval(y, x, alpha=alpha, beta=beta)
    npt.assert_allclose(y.to_host(), y_exp, rtol=1e-5)

    # shape
    assert A.shape == (M*stack,N)
    assert A.H.shape == (N,M*stack)

    # dtype
    assert A.dtype == np.dtype('complex64')


@pytest.mark.parametrize("backend,stack,M,N,K,density,alpha,beta",
    product( BACKENDS, [1,2,3], [5,6], [7,8], [1,8,9,17], [0.01,0.1,0.5,1], [0,.5,1], [0,.5,1] ))
def test_HStack(backend, stack, M, N, K, density, alpha, beta):
    b = backend()
    mats_h = [indigo.util.randM(M,N,density) for i in range(stack)]
    A_h = spp.hstack(mats_h)

    mats_d = [b.SpMatrix(m) for m in mats_h]
    A = b.HStack(mats_d)

    # forward
    x = b.rand_array((A.shape[1],K))
    y = b.rand_array((A.shape[0],K))
    y_exp = beta * y.to_host() + alpha * A_h @ x.to_host()
    A.eval(y, x, alpha=alpha, beta=beta)
    npt.assert_allclose(y.to_host(), y_exp, rtol=1e-5)

    # adjoint
    x = b.rand_array((A.shape[0],K))
    y = b.rand_array((A.shape[1],K))
    y_exp = beta * y.to_host() + alpha * A_h.H @ x.to_host()
    A.H.eval(y, x, alpha=alpha, beta=beta)
    npt.assert_allclose(y.to_host(), y_exp, rtol=1e-5)

    # shape
    assert A.shape == (M,N*stack)
    assert A.H.shape == (N*stack,M)

    # dtype
    assert A.dtype == np.dtype('complex64')


@pytest.mark.parametrize("backend,stack,M,N,K,density,alpha,beta",
    product( BACKENDS, [1,2,3], [5,6], [7,8], [1,8,9,17], [0.01,0.1,0.5,1], [0,.5,1], [0,.5,1] ))
def test_BlockDiag(backend, stack, M, N, K, density, alpha, beta):
    b = backend()
    mats_h = [indigo.util.randM(M,N,density) for i in range(stack)]
    A_h = spp.block_diag(mats_h)

    mats_d = [b.SpMatrix(m) for m in mats_h]
    A = b.BlockDiag(mats_d)

    # forward
    x = b.rand_array((A.shape[1],K))
    y = b.rand_array((A.shape[0],K))
    y_exp = beta * y.to_host() + alpha * A_h @ x.to_host()
    A.eval(y, x, alpha=alpha, beta=beta)
    npt.assert_allclose(y.to_host(), y_exp, rtol=1e-5)

    # adjoint
    x = b.rand_array((A.shape[0],K))
    y = b.rand_array((A.shape[1],K))
    y_exp = beta * y.to_host() + alpha * A_h.H @ x.to_host()
    A.H.eval(y, x, alpha=alpha, beta=beta)
    npt.assert_allclose(y.to_host(), y_exp, rtol=1e-5)

    # shape
    assert A.shape == (M*stack,N*stack)
    assert A.H.shape == (N*stack,M*stack)

    # dtype
    assert A.dtype == np.dtype('complex64')


@pytest.mark.parametrize("backend,M,K",
    product( BACKENDS, [23,45], [45,23] ))
def test_Eye(backend, M, K):
    b = backend()
    A = b.Eye(M, dtype=np.complex64)
    x = b.rand_array((M,K))
    y = b.rand_array((M,K))
    y_exp = x.to_host()
    A.eval(y, x)
    npt.assert_allclose(y.to_host(), y_exp, rtol=1e-5)


@pytest.mark.parametrize("backend,M,N,K,density,alpha,beta",
    product( BACKENDS, [5,6], [7,8], [1,8,9,17], [0.9,0.5,1], [0,.5,1], [0,.5,1] ))
def test_KronI(backend, M, N, K, density, alpha, beta):
    b = backend()
    mat_h = indigo.util.randM(M,N,density)
    A_h = spp.kron( spp.eye(K), mat_h )

    mat_d = b.SpMatrix(mat_h)
    A = b.KronI(K, mat_d)

    # forward
    x = b.rand_array((A.shape[1],K))
    y = b.rand_array((A.shape[0],K))
    y_exp = beta * y.to_host() + alpha * A_h @ x.to_host()
    A.eval(y, x, alpha=alpha, beta=beta)
    npt.assert_allclose(y.to_host(), y_exp, rtol=1e-5)

    # adjoint
    x = b.rand_array((A.shape[0],K))
    y = b.rand_array((A.shape[1],K))
    y_exp = beta * y.to_host() + alpha * np.conj(A_h.T) @ x.to_host()
    A.H.eval(y, x, alpha=alpha, beta=beta)
    npt.assert_allclose(y.to_host(), y_exp, rtol=1e-5)

    # shape
    assert A.shape == (M*K,N*K)
    assert A.H.shape == (N*K,M*K)

    # dtype
    assert A.dtype == np.dtype('complex64')


@pytest.mark.parametrize("backend,M,N,K,B",
    product( BACKENDS, [22,23,24], [22,23,24], [22,23,24], [1,2,3,8] )
)
def test_UnscaledFFT_3d(backend, M, N, K, B ):
    b = backend()

    # forward
    x = b.rand_array( (M*N*K, B) )
    y = b.rand_array( (M*N*K, B) )
    x_h = x.to_host().reshape( (M,N,K,B), order='F' )

    A = b.UnscaledFFT( (M,N,K), dtype=x.dtype )

    A.eval(y, x)

    y_exp = np.fft.fftn( x_h, axes=(0,1,2) )
    y_act = y.to_host().reshape( (M,N,K,B), order='F' )
    npt.assert_allclose(y_act, y_exp, rtol=1e-2)

    # adjoint
    x = b.rand_array( (M*N*K, B) )
    y = b.rand_array( (M*N*K, B) )
    x_h = x.to_host().reshape( (M,N,K,B), order='F' )

    A.H.eval(y, x)

    y_exp = np.fft.ifftn( x_h, axes=(0,1,2) ) * (M*N*K)
    y_act = y.to_host().reshape( (M,N,K,B), order='F' )
    npt.assert_allclose(y_act, y_exp, rtol=1e-2)

@pytest.mark.parametrize("backend,M,N,B",
    product( BACKENDS, [22,23,24], [22,23,24], [1,2,3,8] )
)
def test_UnscaledFFT_2d(backend, M, N, B ):
    b = backend()

    # forward
    x = b.rand_array( (M*N, B) )
    y = b.rand_array( (M*N, B) )
    x_h = x.to_host().reshape( (M,N,B), order='F' )

    A = b.UnscaledFFT( (M,N), dtype=x.dtype )

    A.eval(y, x)

    y_exp = np.fft.fftn( x_h, axes=(0,1) )
    y_act = y.to_host().reshape( (M,N,B), order='F' )
    npt.assert_allclose(y_act, y_exp, rtol=1e-2)

    # adjoint
    x = b.rand_array( (M*N, B) )
    y = b.rand_array( (M*N, B) )
    x_h = x.to_host().reshape( (M,N,B), order='F' )

    A.H.eval(y, x)

    y_exp = np.fft.ifftn( x_h, axes=(0,1) ) * (M*N)
    y_act = y.to_host().reshape( (M,N,B), order='F' )
    npt.assert_allclose(y_act, y_exp, rtol=1e-2)


@pytest.mark.parametrize("backend,M,B",
    product( BACKENDS, [22,23,24], [1,2,3,8] )
)
def test_UnscaledFFT_1d(backend, M, B ):
    b = backend()

    # forward
    x = b.rand_array( (M, B) )
    y = b.rand_array( (M, B) )
    x_h = x.to_host().reshape( (M,B), order='F' )

    A = b.UnscaledFFT( (M,), dtype=x.dtype )

    A.eval(y, x)

    y_exp = np.fft.fftn( x_h, axes=(0,) )
    y_act = y.to_host().reshape( (M,B), order='F' )
    npt.assert_allclose(y_act, y_exp, rtol=1e-2)

    # adjoint
    x = b.rand_array( (M, B) )
    y = b.rand_array( (M, B) )
    x_h = x.to_host().reshape( (M,B), order='F' )

    A.H.eval(y, x)

    y_exp = np.fft.ifftn( x_h, axes=(0,) ) * M
    y_act = y.to_host().reshape( (M,B), order='F' )
    npt.assert_allclose(y_act, y_exp, rtol=1e-2)


@pytest.mark.parametrize("backend,M,N,K,B",
    product( BACKENDS, [22,23,24], [22,23,24], [22,23,24], [1,2,3,8])
)
def test_UnitaryFFT(backend, M, N, K, B ):
    b = backend()
    A = b.FFT( (M,N,K), dtype=np.dtype('complex64') )

    x = b.rand_array( (M*N*K, B) )
    y = b.rand_array( (M*N*K, B) )
    x_exp = x.to_host().reshape( (M,N,K,B), order='F' )

    A.eval(y, x)
    A.H.eval(x, y)

    x_act = x.to_host().reshape( (M,N,K,B), order='F' )
    npt.assert_allclose(x_act, x_exp, rtol=1e-2)


@pytest.mark.parametrize("backend,M,N,K,B",
    product( BACKENDS, [22,23,24], [22,23,24], [22,23,24], [1,2,3,8])
)
def test_CenteredFFT(backend, M, N, K, B ):
    from numpy.fft import fftshift, ifftshift, fftn, ifftn

    b = backend()
    A = b.FFTc( (M,N,K), dtype=np.dtype('complex64') )

    # forward
    ax = (0,1,2)
    x = b.rand_array( (M*N*K,B) )
    y = b.rand_array( (M*N*K,B) )
    x_h = x.to_host().reshape( (M,N,K,B), order='F' )

    A.eval(y, x)

    y_act = y.to_host().reshape( (M,N,K,B), order='F' )
    y_exp = fftshift( fftn( ifftshift(x_h, axes=ax), axes=ax, norm='ortho'), axes=ax)
    npt.assert_allclose(y_act, y_exp, rtol=1e-2)

    # adjoint
    x = b.rand_array( (M*N*K,B) )
    y = b.rand_array( (M*N*K,B) )
    x_h = x.to_host().reshape( (M,N,K,B), order='F' )

    A.H.eval(y, x)

    y_act = y.to_host().reshape( (M,N,K,B), order='F' )
    y_exp = fftshift( ifftn( ifftshift(x_h, axes=ax), axes=ax, norm='ortho'), axes=ax)
    npt.assert_allclose(y_act, y_exp, rtol=1e-2)


@pytest.mark.parametrize("backend,batch,x,y,z,px,py,pz",
    product( BACKENDS, [1,2,4,8], [3,4],[3,4],[3,4],
                                  [0,1,2],[0,1,2],[0,1,2] )
)
def test_zpad(backend, batch, x, y, z, px, py, pz):
    N = (z, y, x, batch)
    M = (z+2*pz, y+2*py, x+2*px, batch)
    u = np.random.rand(*N) + 1j*np.random.rand(*N)
    v = np.random.rand(*M) + 1j*np.random.rand(*M)
    u = np.require(u, dtype=np.dtype('complex64'), requirements='F')
    v = np.require(v, dtype=np.dtype('complex64'), requirements='F')

    b = backend()
    A = b.Zpad( M[:3], N[:3], dtype=v.dtype )

    # check adjoint
    u_d = b.copy_array(u.reshape((-1,batch), order='F'))
    v_d = b.copy_array(v.reshape((-1,batch), order='F'))
    A.H.eval(u_d, v_d)
    u_h = u_d.to_host().reshape(N, order='F')
    v_inner = v[ pz:(-pz or None), py:(-py or None), px:(-px or None) ]
    np.testing.assert_allclose(v_inner, u_h)

    # check forward
    u_d = b.copy_array(u.reshape((-1,batch), order='F'))
    v_d = b.copy_array(v.reshape((-1,batch), order='F'))
    A.eval(v_d, u_d)
    v_h = v_d.to_host().reshape(v.shape, order='F')
    v_inner = v_h[ pz:(-pz or None), py:(-py or None), px:(-px or None) ]
    np.testing.assert_allclose(v_inner, u)


@pytest.mark.parametrize("backend,batch,x,y,z,ro,tr",
    product( BACKENDS, [1,2,4,8], [3,4],[3,4],[3,4],[5,6],[7,8] )
)
def test_interp(backend, batch, x, y, z, ro, tr):
    N = (z, y, x, batch)
    M = (1, ro, tr, batch)
    u = np.random.rand(*N) + 1j*np.random.rand(*N)
    v = np.random.rand(*M) + 1j*np.random.rand(*M)
    u = np.require(u, dtype=np.dtype('complex64'), requirements='F')
    v = np.require(v, dtype=np.dtype('complex64'), requirements='F')

    b = backend()
    width = 3
    table = np.random.random(128)
    coord = np.random.rand(3, *M[1:3])
    A = b.Interp( N[:3], coord, width, table, dtype=u.dtype )

    u_d = b.copy_array(u.reshape((-1,batch), order='F'))
    v_d = b.copy_array(v.reshape((-1,batch), order='F'))
    A.H.eval(u_d, v_d)
    AHv = u_d.to_host().flatten(order='F')

    u_d = b.copy_array(u.reshape((-1,batch), order='F'))
    v_d = b.copy_array(v.reshape((-1,batch), order='F'))
    A.eval(v_d, u_d)
    Au = v_d.to_host().flatten(order='F')

    # adjointness test
    u = u.flatten(order='F')
    v = v.flatten(order='F')
    np.testing.assert_allclose( np.dot(Au,v), np.dot(u,AHv), atol=1e-3 )


@pytest.mark.parametrize("backend,batch,M,N,c1,c2",
    product( BACKENDS, [1,2,4,8], [3,4],[3,4], [4,5],[3,6] )
)
def test_nested_kroni(backend, batch, M, N, c1, c2):
    A00_h = indigo.util.randM(M, N, 0.9)
    A0_h = spp.kron( spp.eye(c1), A00_h )
    A_h = spp.kron( spp.eye(c2), A0_h )

    b = backend()
    A00 = b.SpMatrix(A00_h)
    A0 = b.KronI(c1, A00)
    A = b.KronI(c2, A0)

    V, U = A.shape
    u = np.random.rand(U) + 1j*np.random.rand(U)
    u = np.require(u, dtype=np.dtype('complex64'), requirements='F')
    v = np.random.rand(V) + 1j*np.random.rand(V)
    v = np.require(v, dtype=np.dtype('complex64'), requirements='F')

    # forward
    u_d = b.copy_array(u)
    v_d = b.copy_array(v)
    A.eval(v_d, u_d)
    v_act = v_d.to_host()
    v_exp = A_h @ u
    np.testing.assert_allclose( v_act, v_exp, rtol=1e-6 )

    # adjoint
    u_d = b.copy_array(u)
    v_d = b.copy_array(v)
    A.H.eval(u_d, v_d)
    u_act = u_d.to_host()
    u_exp = A_h.H @ v
    np.testing.assert_allclose( u_act, u_exp, rtol=1e-6 )


@pytest.mark.parametrize("backend,M,N,K,forward,alpha,beta",
    product( BACKENDS, [23,24,45], [45,24,23], [1,8,9,17], [True,False], [0,.5,1], [0,.5,1] ))
def test_DenseMatrix(backend, M, N, K, forward, alpha, beta):
    b = backend()
    A_h = indigo.util.rand64c(M,N)
    A = b.DenseMatrix(A_h)

    if forward:
        x = b.rand_array((N,K))
        y = b.rand_array((M,K))
        y_exp = beta * y.to_host() + alpha * A_h.dot(x.to_host())
        A.eval(y, x, alpha=alpha, beta=beta)
    else:
        x = b.rand_array((M,K))
        y = b.rand_array((N,K))
        y_exp = beta * y.to_host() + alpha * np.conj(A_h.T).dot(x.to_host())
        A.H.eval(y, x, alpha=alpha, beta=beta)

    npt.assert_allclose(y.to_host(), y_exp, rtol=1e-5)


@pytest.mark.parametrize("backend,M,N,K,density,alpha,beta,batch",
    product( BACKENDS, [23,45], [45,23], [1,8,9,17], [0.01,0.1,0.5,1], [0,.5,1], [0,.5,1],
             [None,1,2,10] ))
def test_batch(backend, M, N, K, density, alpha, beta, batch):
    b = backend()
    A_h = indigo.util.randM(M, N, density)
    A = b.SpMatrix(A_h, batch=batch)

    # forward
    x = b.rand_array((N,K))
    y = b.rand_array((M,K))
    y_exp = beta * y.to_host() + alpha * A_h * x.to_host()
    A.eval(y, x, alpha=alpha, beta=beta)
    npt.assert_allclose(y.to_host(), y_exp, rtol=1e-5)

    # adjoint
    x = b.rand_array((M,K))
    y = b.rand_array((N,K))
    y_exp = beta * y.to_host() + alpha * A_h.H * x.to_host()
    A.H.eval(y, x, alpha=alpha, beta=beta)
    npt.assert_allclose(y.to_host(), y_exp, rtol=1e-5)


@pytest.mark.parametrize("backend,M,N,gamma",
    product( BACKENDS, [11,12,13], [15,16], [0.0, 0.5, 1.0, 1.5] ))
def test_Eye_Scale(backend, M, N, gamma):
    B = backend()
    A = gamma * B.Eye(M, dtype=np.complex64)
    x = indigo.util.rand64c(M,N)
    y_exp = gamma * x
    y_act = A * x
    np.testing.assert_allclose(y_act, y_exp, rtol=1e-5)


@pytest.mark.parametrize("backend,M,N,gamma",
    product( BACKENDS, [11,12,13], [15,16], [0.0, 0.5, 1.0, 1.5] ))
def test_Sum(backend, M, N, gamma):
    x = indigo.util.rand64c(M,N)
    B = backend()
    S = B.Eye(M, dtype=np.complex64) * 2
    E = B.Eye(M, dtype=np.complex64) * 4
    y_exp = (2+4)*x
    y_act = (S+E) * x
    np.testing.assert_allclose(y_act, y_exp, rtol=1e-5)


@pytest.mark.parametrize("backend,M,N,gamma",
    product( BACKENDS, [11,12,13], [15,16], [0.0, 0.5, 1.0, 1.5] ))
def test_Difference(backend, M, N, gamma):
    x = indigo.util.rand64c(M,N)
    B = backend()
    S = B.Eye(M, dtype=np.complex64) * 2
    E = B.Eye(M, dtype=np.complex64) * 4
    y_exp = (2-4)*x
    y_act = (S-E) * x
    np.testing.assert_allclose(y_act, y_exp, rtol=1e-5)

@pytest.mark.parametrize("backend,M,N,K,alpha,beta,forward",
    product( BACKENDS, [11,12,13], [15,16], [1,2,3],
             [0.0, 0.5, 1.0, 1.5], [0.0, 0.5, 1.0, 1.5], [True, False] ))
def test_One(backend, M, N, K, alpha, beta, forward):
    x = indigo.util.rand64c(K,N)
    y = indigo.util.rand64c(M,N)
    B = backend()
    if getattr(B.onemm, '__isabstractmethod__', False):
        pytest.skip("backed <%s> doesn't implement onemm" % backend.__name__)
    if not hasattr(B, 'onemm'):
        pytest.skip("backend doesn't implement onemm")
    O = B.One((M,K), dtype=np.complex64)

    if forward:
        u, v = x, y
    else:
        v, u = x, y

    u_d = B.copy_array(u)
    v_d = B.copy_array(v)
    exp = beta * v + \
        np.broadcast_to(alpha*u.sum(axis=0,keepdims=True), v.shape)
    O.eval(v_d, u_d, alpha=alpha, beta=beta, forward=forward)
    act = v_d.to_host()
    np.testing.assert_allclose(act, exp, rtol=1e-5)


@pytest.mark.parametrize("backend,sym,m,k",
    product(BACKENDS, [True,False], [2,4,5,6], [1,2,3])
)
def test_DenseMatrix_symmetric(backend, sym, m, k):
    b = backend()
    data = indigo.util.rand64c(m,m)
    if sym:
        data = np.asfortranarray(data + data.T) # make symmetric
        data.imag = 0 # make imaginary
    M = b.DenseMatrix(data)
    assert (M._real_symmetric if sym else not M._real_symmetric)
    x = indigo.util.rand64c(m,k)
    y_exp = np.dot(data, x)
    y_act = M * x
    np.testing.assert_allclose(y_exp, y_act, atol=1e-5)


@pytest.mark.parametrize("backend,L,Q,alpha,beta,eyeL,eyeR",
    product( BACKENDS, [3,4], [5,6],
        [0,.5,1], [0,.5,1], [True, False], [True, False] ))
def test_Kron_general(backend, L, Q, alpha, beta, eyeL, eyeR):
    b = backend()

    if eyeL:
        A = np.eye(L, dtype=np.complex64)
        A_d = b.Eye(L)
        if not eyeR: A_d._eval = lambda unreachable: ()
    else:
        A = indigo.util.rand64c(L,L)
        A_d = b.DenseMatrix(A)

    if eyeR:
        B = np.eye(Q, dtype=np.complex64)
        B_d = b.Eye(Q)
        if not eyeL: B_d._eval = lambda unreachable: ()
    else:
        B = indigo.util.rand64c(Q,Q)
        B = B + B.T
        B.imag = 0
        B_d = b.DenseMatrix(B)

    K = b.Kron( B_d, A_d )
    vec = lambda arr: arr.reshape((-1,1), order='F')

    # check
    X = indigo.util.rand64c(L,Q)
    Y = indigo.util.rand64c(L,Q)
    y_exp  = alpha * np.kron(B.T, A) @ vec(X) + beta * vec(Y)
    y_exp2 = alpha * vec(A @ X @ B) + beta * vec(Y)
    npt.assert_allclose(y_exp2, y_exp, rtol=1e-5)

    # forward
    x_d = b.copy_array(X)
    y_d = b.copy_array(Y)
    K.eval(y_d, x_d, alpha=alpha, beta=beta)
    y_act = y_d.to_host()
    npt.assert_allclose(vec(y_act), y_exp, rtol=1e-5)

    # adjoint
    X = indigo.util.rand64c(Q,L)
    Y = indigo.util.rand64c(Q,L)
    y_exp = alpha * np.conj(np.kron(B.T, A).T) @ vec(X) + beta * vec(Y)
    x_d = b.copy_array(X)
    y_d = b.copy_array(Y)
    K.eval(y_d, x_d, alpha=alpha, beta=beta, forward=False)
    y_act = y_d.to_host()
    npt.assert_allclose(vec(y_act), y_exp, rtol=1e-5)


@pytest.mark.parametrize("backend,M,N,P",
    product( BACKENDS, [23,45], [45,23], [1,2,3] ))
def test_Convolution(backend, M, N, P):
    from scipy.signal import fftconvolve

    b = backend()

    #k = indigo.util.rand64c(M,N)
    k = np.zeros((M,N), dtype=np.complex64)
    k[M//2,N//2] = 1

    C = b.Convolution(k)

    x = indigo.util.rand64c(M,N,P)

    y_act = C * x
    y_act = y_act.reshape((M,N,-1), order='F')

    y_exp = np.zeros((M,N,P), np.complex64)
    for p in range(P):
        y_exp[:,:,p] = fftconvolve( x[:,:,p], k, mode='same' )

    pytest.xfail("under development")
    #npt.assert_allclose(y_act, y_exp, rtol=1e-5)
