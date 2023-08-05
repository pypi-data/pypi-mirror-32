# -*- coding: utf-8 -*-
from __future__ import division
import numbers
from math import pi
from math import factorial as fac
from distutils.version import LooseVersion

import numpy as np
from ase.units import Ha, Bohr
from ase.utils.timing import timer

import gpaw.fftw as fftw
from gpaw.band_descriptor import BandDescriptor
from gpaw.blacs import BlacsGrid, BlacsDescriptor, Redistributor
from gpaw.density import Density
from gpaw.lfc import BaseLFC
from gpaw.lcao.overlap import fbt
from gpaw.hamiltonian import Hamiltonian
from gpaw.matrix_descriptor import MatrixDescriptor
from gpaw.spherical_harmonics import Y, nablarlYL
from gpaw.spline import Spline
from gpaw.utilities import unpack
from gpaw.utilities.blas import rk, r2k, gemm, axpy
from gpaw.utilities.progressbar import ProgressBar
from gpaw.wavefunctions.fdpw import FDPWWaveFunctions
from gpaw.wavefunctions.mode import Mode
from gpaw.wavefunctions.arrays import PlaneWaveExpansionWaveFunctions
import _gpaw


class PW(Mode):
    name = 'pw'

    def __init__(self, ecut=340, fftwflags=fftw.ESTIMATE, cell=None,
                 pulay_stress=None, dedecut=None,
                 force_complex_dtype=False):
        """Plane-wave basis mode.

        ecut: float
            Plane-wave cutoff in eV.
        dedecut: float or None or 'estimate'
            Estimate of derivative of total energy with respect to
            plane-wave cutoff.  Used to calculate pulay_stress.
        pulay_stress: float or None
            Pulay-stress correction.
        fftwflags: int
            Flags for making FFTW plan (default is ESTIMATE).
        cell: 3x3 ndarray
            Use this unit cell to chose the planewaves.

        Only one of dedecut and pulay_stress can be used.
        """

        self.ecut = ecut / Ha
        self.fftwflags = fftwflags
        self.dedecut = dedecut
        self.pulay_stress = (None
                             if pulay_stress is None
                             else pulay_stress * Bohr**3 / Ha)

        assert pulay_stress is None or dedecut is None

        if cell is None:
            self.cell_cv = None
        else:
            self.cell_cv = cell / Bohr

        Mode.__init__(self, force_complex_dtype)

    def __call__(self, parallel, initksl, gd, **kwargs):
        dedepsilon = 0.0
        volume = abs(np.linalg.det(gd.cell_cv))

        if self.cell_cv is None:
            ecut = self.ecut
        else:
            volume0 = abs(np.linalg.det(self.cell_cv))
            ecut = self.ecut * (volume0 / volume)**(2 / 3.0)

        if self.pulay_stress is not None:
            dedepsilon = self.pulay_stress * volume
        elif self.dedecut is not None:
            if self.dedecut == 'estimate':
                dedepsilon = 'estimate'
            else:
                dedepsilon = self.dedecut * 2 / 3 * ecut

        wfs = PWWaveFunctions(ecut, self.fftwflags, dedepsilon,
                              parallel, initksl, gd=gd,
                              **kwargs)

        return wfs

    def todict(self):
        dct = Mode.todict(self)
        dct['ecut'] = self.ecut * Ha
        if self.cell_cv is not None:
            dct['cell'] = self.cell_cv * Bohr
        if self.pulay_stress is not None:
            dct['pulay_stress'] = self.pulay_stress * Ha / Bohr**3
        if self.dedecut is not None:
            dct['dedecut'] = self.dedecut
        return dct


class PWDescriptor:
    ndim = 1  # all 3d G-vectors are stored in a 1d ndarray

    def __init__(self, ecut, gd, dtype=None, kd=None,
                 fftwflags=fftw.ESTIMATE):

        assert gd.pbc_c.all()
        assert gd.comm.size == 1

        self.gd = gd
        self.fftwflags = fftwflags

        N_c = gd.N_c
        self.comm = gd.comm

        ecutmax = 0.5 * pi**2 / (self.gd.h_cv**2).sum(1).max()

        if ecut is None:
            ecut = ecutmax * 0.9999
        else:
            if ecut > ecutmax:
                raise ValueError(
                    'You have a weird unit cell!  '
                    'Try to use the maximally reduced Niggli cell.  '
                    'See the ase.build.niggli_reduce() function.')

        self.ecut = ecut

        if dtype is None:
            if kd is None or kd.gamma:
                dtype = float
            else:
                dtype = complex
        self.dtype = dtype

        if dtype == float:
            Nr_c = N_c.copy()
            Nr_c[2] = N_c[2] // 2 + 1
            i_Qc = np.indices(Nr_c).transpose((1, 2, 3, 0))
            i_Qc[..., :2] += N_c[:2] // 2
            i_Qc[..., :2] %= N_c[:2]
            i_Qc[..., :2] -= N_c[:2] // 2
            self.tmp_Q = fftw.empty(Nr_c, complex)
            self.tmp_R = self.tmp_Q.view(float)[:, :, :N_c[2]]
        else:
            i_Qc = np.indices(N_c).transpose((1, 2, 3, 0))
            i_Qc += N_c // 2
            i_Qc %= N_c
            i_Qc -= N_c // 2
            self.tmp_Q = fftw.empty(N_c, complex)
            self.tmp_R = self.tmp_Q

        self.nbytes = self.tmp_R.nbytes

        self.fftplan = fftw.FFTPlan(self.tmp_R, self.tmp_Q, -1, fftwflags)
        self.ifftplan = fftw.FFTPlan(self.tmp_Q, self.tmp_R, 1, fftwflags)

        # Calculate reciprocal lattice vectors:
        B_cv = 2.0 * pi * gd.icell_cv
        i_Qc.shape = (-1, 3)
        self.G_Qv = np.dot(i_Qc, B_cv)
        self.nbytes += self.G_Qv.nbytes

        self.kd = kd
        if kd is None:
            self.K_qv = np.zeros((1, 3))
        else:
            self.K_qv = np.dot(kd.ibzk_qc, B_cv)

        # Map from vectors inside sphere to fft grid:
        self.Q_qG = []
        self.G2_qG = []
        Q_Q = np.arange(len(i_Qc), dtype=np.int32)

        self.ngmin = 100000000
        self.ngmax = 0
        for q, K_v in enumerate(self.K_qv):
            G2_Q = ((self.G_Qv + K_v)**2).sum(axis=1)
            mask_Q = (G2_Q <= 2 * ecut)
            if self.dtype == float:
                mask_Q &= ((i_Qc[:, 2] > 0) |
                           (i_Qc[:, 1] > 0) |
                           ((i_Qc[:, 0] >= 0) & (i_Qc[:, 1] == 0)))
            Q_G = Q_Q[mask_Q]
            self.Q_qG.append(Q_G)
            self.G2_qG.append(G2_Q[Q_G])
            ng = len(Q_G)
            self.ngmin = min(ng, self.ngmin)
            self.ngmax = max(ng, self.ngmax)
            self.nbytes += Q_G.nbytes + self.G2_qG[q].nbytes

        if kd is not None:
            self.ngmin = kd.comm.min(self.ngmin)
            self.ngmax = kd.comm.max(self.ngmax)

        self.n_c = np.array([self.ngmax])  # used by hs_operators.py XXX

    def get_reciprocal_vectors(self, q=0, add_q=True):
        """ Returns reciprocal lattice vectors plus q, G + q,
        in xyz coordinates.
        """

        assert q < len(self.K_qv), ('Choose a q-index belonging to ' +
                                    'the irreducible Brillouin zone.')
        q_v = self.K_qv[q]

        if add_q:
            G_Gv = self.G_Qv[self.Q_qG[q]] + q_v
        else:
            G_Gv = self.G_Qv[self.Q_qG[q]]

        return G_Gv

    def __getstate__(self):
        return (self.ecut, self.gd, self.dtype, self.kd, self.fftwflags)

    def __setstate__(self, state):
        self.__init__(*state)

    def estimate_memory(self, mem):
        mem.subnode('Arrays', self.nbytes)

    def bytecount(self, dtype=float):
        return self.ngmax * 16

    def zeros(self, x=(), dtype=None, q=-1):
        a_xG = self.empty(x, dtype, q)
        a_xG.fill(0.0)
        return a_xG

    def empty(self, x=(), dtype=None, q=-1):
        if dtype is not None:
            assert dtype == self.dtype
        if isinstance(x, numbers.Integral):
            x = (x,)
        if q == -1:
            shape = x + (self.ngmax,)
        else:
            shape = x + self.Q_qG[q].shape
        return np.empty(shape, complex)

    def fft(self, f_R, q=-1, Q_G=None):
        """Fast Fourier transform.

        Returns c(G) for G<Gc::

                   __
                  \        -iG.R
            c(G) = ) f(R) e
                  /__
                   R
        """

        self.tmp_R[:] = f_R

        self.fftplan.execute()
        if Q_G is None:
            Q_G = self.Q_qG[q]
        return self.tmp_Q.ravel()[Q_G]

    def ifft(self, c_G, q=-1):
        """Inverse fast Fourier transform.

        Returns::

                      __
                   1 \        iG.R
            f(R) = -  ) c(G) e
                   N /__
                      G
        """

        self.tmp_Q[:] = 0.0
        self.tmp_Q.ravel()[self.Q_qG[q]] = c_G
        if self.dtype == float:
            t = self.tmp_Q[:, :, 0]
            n, m = self.gd.N_c[:2] // 2 - 1
            t[0, -m:] = t[0, m:0:-1].conj()
            t[n:0:-1, -m:] = t[-n:, m:0:-1].conj()
            t[-n:, -m:] = t[n:0:-1, m:0:-1].conj()
            t[-n:, 0] = t[n:0:-1, 0].conj()
        self.ifftplan.execute()
        return self.tmp_R * (1.0 / self.tmp_R.size)

    def integrate(self, a_xg, b_yg=None,
                  global_integral=True, hermitian=False,
                  _transposed_result=None):
        """Integrate function(s) over domain.

        a_xg: ndarray
            Function(s) to be integrated.
        b_yg: ndarray
            If present, integrate a_xg.conj() * b_yg.
        global_integral: bool
            If the array(s) are distributed over several domains, then the
            total sum will be returned.  To get the local contribution
            only, use global_integral=False.
        hermitian: bool
            Result is hermitian.
        _transposed_result: ndarray
            Long story.  Don't use this unless you are a method of the
            MatrixOperator class ..."""

        if b_yg is None:
            # Only one array:
            assert self.dtype == float
            return a_xg[..., 0].real * self.gd.dv

        A_xg = a_xg.reshape((-1, a_xg.shape[-1]))
        B_yg = b_yg.reshape((-1, b_yg.shape[-1]))

        alpha = self.gd.dv / self.gd.N_c.prod()

        if self.dtype == float:
            alpha *= 2
            A_xg = A_xg.view(float)
            B_yg = B_yg.view(float)

        if _transposed_result is None:
            result_yx = np.zeros((len(B_yg), len(A_xg)), self.dtype)
        else:
            result_yx = _transposed_result

        if a_xg is b_yg:
            rk(alpha, A_xg, 0.0, result_yx)
        elif hermitian:
            r2k(0.5 * alpha, A_xg, B_yg, 0.0, result_yx)
        else:
            gemm(alpha, A_xg, B_yg, 0.0, result_yx, 'c')

        if self.dtype == float:
            correction_yx = np.outer(B_yg[:, 0], A_xg[:, 0])
            if hermitian:
                result_yx -= 0.25 * alpha * (correction_yx + correction_yx.T)
            else:
                result_yx -= 0.5 * alpha * correction_yx

        xshape = a_xg.shape[:-1]
        yshape = b_yg.shape[:-1]
        result = result_yx.T.reshape(xshape + yshape)

        if result.ndim == 0:
            return result.item()
        else:
            return result

    def interpolate(self, a_R, pd, q=-1):
        a_Q = self.tmp_Q
        b_Q = pd.tmp_Q

        e0, e1, e2 = 1 - self.gd.N_c % 2  # even or odd size
        a0, a1, a2 = pd.gd.N_c // 2 - self.gd.N_c // 2
        b0, b1, b2 = self.gd.N_c + (a0, a1, a2)

        if self.dtype == float:
            b2 = (b2 - a2) // 2 + 1
            a2 = 0
            axes = (0, 1)
        else:
            axes = (0, 1, 2)

        self.tmp_R[:] = a_R
        self.fftplan.execute()
        b_Q[:] = 0.0
        b_Q[a0:b0, a1:b1, a2:b2] = np.fft.fftshift(a_Q, axes=axes)

        if e0:
            b_Q[a0, a1:b1, a2:b2] *= 0.5
            b_Q[b0, a1:b1, a2:b2] = b_Q[a0, a1:b1, a2:b2]
            b0 += 1
        if e1:
            b_Q[a0:b0, a1, a2:b2] *= 0.5
            b_Q[a0:b0, b1, a2:b2] = b_Q[a0:b0, a1, a2:b2]
            b1 += 1
        if self.dtype == complex:
            if e2:
                b_Q[a0:b0, a1:b1, a2] *= 0.5
                b_Q[a0:b0, a1:b1, b2] = b_Q[a0:b0, a1:b1, a2]
        else:
            if e2:
                b_Q[a0:b0, a1:b1, b2 - 1] *= 0.5

        b_Q[:] = np.fft.ifftshift(b_Q, axes=axes)
        pd.ifftplan.execute()
        return pd.tmp_R * (1.0 / self.tmp_R.size), a_Q.ravel()[self.Q_qG[q]]

    def restrict(self, a_R, pd, q=-1):
        a_Q = pd.tmp_Q
        b_Q = self.tmp_Q

        e0, e1, e2 = 1 - pd.gd.N_c % 2  # even or odd size
        a0, a1, a2 = self.gd.N_c // 2 - pd.gd.N_c // 2
        b0, b1, b2 = pd.gd.N_c // 2 + self.gd.N_c // 2 + 1

        if self.dtype == float:
            b2 = pd.gd.N_c[2] // 2 + 1
            a2 = 0
            axes = (0, 1)
        else:
            axes = (0, 1, 2)

        self.tmp_R[:] = a_R
        self.fftplan.execute()
        b_Q[:] = np.fft.fftshift(b_Q, axes=axes)

        if e0:
            b_Q[a0, a1:b1, a2:b2] += b_Q[b0 - 1, a1:b1, a2:b2]
            b_Q[a0, a1:b1, a2:b2] *= 0.5
            b0 -= 1
        if e1:
            b_Q[a0:b0, a1, a2:b2] += b_Q[a0:b0, b1 - 1, a2:b2]
            b_Q[a0:b0, a1, a2:b2] *= 0.5
            b1 -= 1
        if self.dtype == complex and e2:
            b_Q[a0:b0, a1:b1, a2] += b_Q[a0:b0, a1:b1, b2 - 1]
            b_Q[a0:b0, a1:b1, a2] *= 0.5
            b2 -= 1

        a_Q[:] = b_Q[a0:b0, a1:b1, a2:b2]
        a_Q[:] = np.fft.ifftshift(a_Q, axes=axes)
        a_G = a_Q.ravel()[pd.Q_qG[q]] / 8
        pd.ifftplan.execute()
        return pd.tmp_R * (1.0 / self.tmp_R.size), a_G

    def map(self, pd, q=-1):
        N_c = np.array(self.tmp_Q.shape)
        N3_c = pd.tmp_Q.shape
        Q2_G = self.Q_qG[q]
        Q2_Gc = np.empty((len(Q2_G), 3), int)
        Q2_Gc[:, 0], r_G = divmod(Q2_G, N_c[1] * N_c[2])
        Q2_Gc.T[1:] = divmod(r_G, N_c[2])
        if self.dtype == float:
            C = 2
        else:
            C = 3
        Q2_Gc[:, :C] += N_c[:C] // 2
        Q2_Gc[:, :C] %= N_c[:C]
        Q2_Gc[:, :C] -= N_c[:C] // 2
        Q2_Gc[:, :C] %= N3_c[:C]
        Q3_G = Q2_Gc[:, 2] + N3_c[2] * (Q2_Gc[:, 1] + N3_c[1] * Q2_Gc[:, 0])
        G3_Q = np.empty(N3_c, int).ravel()
        G3_Q[pd.Q_qG[q]] = np.arange(len(pd.Q_qG[q]))
        return G3_Q[Q3_G]


def count_reciprocal_vectors(ecut, gd, q_c):
    N_c = gd.N_c
    i_Qc = np.indices(N_c).transpose((1, 2, 3, 0))
    i_Qc += N_c // 2
    i_Qc %= N_c
    i_Qc -= N_c // 2

    B_cv = 2.0 * pi * gd.icell_cv
    i_Qc.shape = (-1, 3)
    Gpq_Qv = np.dot(i_Qc, B_cv) + np.dot(q_c, B_cv)

    G2_Q = (Gpq_Qv**2).sum(axis=1)
    return (G2_Q <= 2 * ecut).sum()


class Preconditioner:
    """Preconditioner for KS equation.

    From:

      Teter, Payne and Allen, Phys. Rev. B 40, 12255 (1989)

    as modified by:

      Kresse and Furthmüller, Phys. Rev. B 54, 11169 (1996)
    """

    def __init__(self, G2_qG, pd):
        self.G2_qG = G2_qG
        self.pd = pd

    def calculate_kinetic_energy(self, psit_xG, kpt):
        if psit_xG.ndim == 1:
            return self.calculate_kinetic_energy(psit_xG[np.newaxis], kpt)[0]
        G2_G = self.G2_qG[kpt.q]
        return np.array([self.pd.integrate(0.5 * G2_G * psit_G, psit_G)
                         for psit_G in psit_xG])

    def __call__(self, R_xG, kpt, ekin_x):
        if R_xG.ndim == 1:
            return self.__call__(R_xG[np.newaxis], kpt, [ekin_x])[0]
        G2_G = self.G2_qG[kpt.q]
        PR_xG = np.empty_like(R_xG)
        for PR_G, R_G, ekin in zip(PR_xG, R_xG, ekin_x):
            x_G = 1 / ekin / 3 * G2_G
            a_G = 27.0 + x_G * (18.0 + x_G * (12.0 + x_G * 8.0))
            PR_G[:] = -4.0 / 3 / ekin * R_G * a_G / (a_G + 16.0 * x_G**4)
        return PR_xG


class NonCollinearPreconditioner(Preconditioner):
    def calculate_kinetic_energy(self, psit_xsG, kpt):
        shape = psit_xsG.shape
        ekin_xs = Preconditioner.calculate_kinetic_energy(
            self, psit_xsG.reshape((-1, shape[-1])), kpt)
        return ekin_xs.reshape(shape[:-1]).sum(-1)

    def __call__(self, R_sG, kpt, ekin):
        return Preconditioner.__call__(self, R_sG, kpt, [ekin, ekin])


class PWWaveFunctions(FDPWWaveFunctions):
    mode = 'pw'

    def __init__(self, ecut, fftwflags, dedepsilon,
                 parallel, initksl,
                 reuse_wfs_method, collinear,
                 gd, nvalence, setups, bd, dtype,
                 world, kd, kptband_comm, timer):
        self.ecut = ecut
        self.fftwflags = fftwflags
        self.dedepsilon = dedepsilon  # Pulay correction for stress tensor

        self.ng_k = None  # number of G-vectors for all IBZ k-points

        FDPWWaveFunctions.__init__(self, parallel, initksl,
                                   reuse_wfs_method=reuse_wfs_method,
                                   collinear=collinear,
                                   gd=gd, nvalence=nvalence, setups=setups,
                                   bd=bd, dtype=dtype, world=world, kd=kd,
                                   kptband_comm=kptband_comm, timer=timer)

    def empty(self, n=(), global_array=False, realspace=False,
              q=-1):
        if realspace:
            return self.gd.empty(n, self.dtype, global_array)
        else:
            return self.pd.empty(n, self.dtype, q)

    def integrate(self, a_xg, b_yg=None, global_integral=True):
        return self.pd.integrate(a_xg, b_yg, global_integral)

    def bytes_per_wave_function(self):
        return 16 * self.pd.ngmax

    def set_setups(self, setups):
        self.timer.start('PWDescriptor')
        self.pd = PWDescriptor(self.ecut, self.gd, self.dtype, self.kd,
                               self.fftwflags)
        self.timer.stop('PWDescriptor')

        # Build array of number of plane wave coefficiants for all k-points
        # in the IBZ:
        self.ng_k = np.zeros(self.kd.nibzkpts, dtype=int)
        for kpt in self.kpt_u:
            if kpt.s != 1:  # avoid double counting (only sum over s=0 or None)
                self.ng_k[kpt.k] = len(self.pd.Q_qG[kpt.q])
        self.kd.comm.sum(self.ng_k)

        self.pt = PWLFC([setup.pt_j for setup in setups], self.pd)

        FDPWWaveFunctions.set_setups(self, setups)

        if self.dedepsilon == 'estimate':
            dedecut = self.setups.estimate_dedecut(self.ecut)
            self.dedepsilon = dedecut * 2 / 3 * self.ecut

    def get_pseudo_partial_waves(self):
        return PWLFC([setup.get_actual_atomic_orbitals()
                      for setup in self.setups], self.pd)

    def __str__(self):
        s = 'Wave functions: Plane wave expansion\n'
        s += '  Cutoff energy: %.3f eV\n' % (self.pd.ecut * Ha)

        if self.dtype == float:
            s += ('  Number of coefficients: %d (reduced to %d)\n' %
                  (self.pd.ngmax * 2 - 1, self.pd.ngmax))
        else:
            s += ('  Number of coefficients (min, max): %d, %d\n' %
                  (self.pd.ngmin, self.pd.ngmax))

        stress = self.dedepsilon / self.gd.volume * Ha / Bohr**3
        dedecut = 1.5 * self.dedepsilon / self.ecut
        s += ('  Pulay-stress correction: {:.6f} eV/Ang^3 '
              '(de/decut={:.6f})\n'.format(stress, dedecut))

        if fftw.FFTPlan is fftw.NumpyFFTPlan:
            s += "  Using Numpy's FFT\n"
        else:
            s += '  Using FFTW library\n'
        return s + FDPWWaveFunctions.__str__(self)

    def make_preconditioner(self, block=1):
        if self.collinear:
            return Preconditioner(self.pd.G2_qG, self.pd)
        return NonCollinearPreconditioner(self.pd.G2_qG, self.pd)

    def apply_pseudo_hamiltonian(self, kpt, ham, psit_xG, Htpsit_xG):
        """Apply the non-pseudo Hamiltonian i.e. without PAW corrections."""
        Htpsit_xG[:] = 0.5 * self.pd.G2_qG[kpt.q] * psit_xG
        if self.collinear:
            for psit_G, Htpsit_G in zip(psit_xG, Htpsit_xG):
                psit_R = self.pd.ifft(psit_G, kpt.q)
                Htpsit_G += self.pd.fft(psit_R * ham.vt_sG[kpt.s], kpt.q)
        else:
            v, x, y, z = ham.vt_xG
            iy = y * 1j
            for psit_sG, Htpsit_sG in zip(psit_xG, Htpsit_xG):
                a = self.pd.ifft(psit_sG[0], kpt.q)
                b = self.pd.ifft(psit_sG[1], kpt.q)
                Htpsit_sG[0] += self.pd.fft(a * (v + z) + b * (x - iy), kpt.q)
                Htpsit_sG[1] += self.pd.fft(a * (x + iy) + b * (v - z), kpt.q)
        ham.xc.apply_orbital_dependent_hamiltonian(
            kpt, psit_xG, Htpsit_xG, ham.dH_asp)

    def add_orbital_density(self, nt_G, kpt, n):
        axpy(1.0, abs(self.pd.ifft(kpt.psit_nG[n], kpt.q))**2, nt_G)

    def add_to_density_from_k_point_with_occupation(self, nt_xR, kpt, f_n):
        if self.collinear:
            nt_R = nt_xR[kpt.s]
            for f, psit_G in zip(f_n, kpt.psit_nG):
                nt_R += f * abs(self.pd.ifft(psit_G, kpt.q))**2
        else:
            for f, psit_sG in zip(f_n, kpt.psit.array):
                p1 = self.pd.ifft(psit_sG[0], kpt.q)
                p2 = self.pd.ifft(psit_sG[1], kpt.q)
                p11 = p1.real**2 + p1.imag**2
                p22 = p2.real**2 + p2.imag**2
                p12 = p1.conj() * p2
                nt_xR[0] += f * (p11 + p22)
                nt_xR[1] += 2 * f * p12.real
                nt_xR[2] += 2 * f * p12.imag
                nt_xR[3] += f * (p11 - p22)

    def calculate_kinetic_energy_density(self):
        if self.kpt_u[0].f_n is None:
            return None

        taut_sR = self.gd.zeros(self.nspins)
        for kpt in self.kpt_u:
            G_Gv = self.pd.get_reciprocal_vectors(q=kpt.q)
            for f, psit_G in zip(kpt.f_n, kpt.psit_nG):
                for v in range(3):
                    taut_sR[kpt.s] += 0.5 * f * abs(
                        self.pd.ifft(1j * G_Gv[:, v] * psit_G, kpt.q))**2

        self.kptband_comm.sum(taut_sR)
        return taut_sR

    def apply_mgga_orbital_dependent_hamiltonian(self, kpt, psit_xG,
                                                 Htpsit_xG, dH_asp,
                                                 dedtaut_R):
        G_Gv = self.pd.get_reciprocal_vectors(q=kpt.q)
        for psit_G, Htpsit_G in zip(psit_xG, Htpsit_xG):
            for v in range(3):
                a_R = self.pd.ifft(1j * G_Gv[:, v] * psit_G, kpt.q)
                axpy(-0.5, 1j * G_Gv[:, v] *
                     self.pd.fft(dedtaut_R * a_R, kpt.q),
                     Htpsit_G)

    def _get_wave_function_array(self, u, n, realspace=True, periodic=False):
        psit_G = self.kpt_u[u].psit_nG[n]

        if not realspace:
            if self.collinear:
                zeropadded_G = np.zeros(self.pd.ngmax, complex)
                zeropadded_G[:len(psit_G)] = psit_G
                return zeropadded_G
            zeropadded_sG = np.zeros((2, self.pd.ngmax), complex)
            zeropadded_sG[:, :len(psit_G[0])] = psit_G
            return zeropadded_sG

        kpt = self.kpt_u[u]
        if self.kd.gamma or periodic:
            return self.pd.ifft(psit_G, kpt.q)

        k_c = self.kd.ibzk_kc[kpt.k]
        eikr_R = self.gd.plane_wave(k_c)
        return self.pd.ifft(psit_G, kpt.q) * eikr_R

    def get_wave_function_array(self, n, k, s, realspace=True,
                                cut=True, periodic=False):
        psit_G = FDPWWaveFunctions.get_wave_function_array(self, n, k, s,
                                                           realspace, periodic)
        if cut and psit_G is not None and not realspace:
            psit_G = psit_G[:self.ng_k[k]].copy()

        return psit_G

    def write(self, writer, write_wave_functions=False):
        FDPWWaveFunctions.write(self, writer)

        if not write_wave_functions:
            return

        if self.collinear:
            shape = (self.nspins,
                     self.kd.nibzkpts, self.bd.nbands, self.pd.ngmax)
        else:
            shape = (self.kd.nibzkpts, self.bd.nbands, 2, self.pd.ngmax)

        writer.add_array('coefficients', shape, complex)

        c = Bohr**-1.5
        for s in range(self.nspins):
            for k in range(self.kd.nibzkpts):
                for n in range(self.bd.nbands):
                    psit_G = self.get_wave_function_array(n, k, s,
                                                          realspace=False,
                                                          cut=False)
                    writer.fill(psit_G * c)

        writer.add_array('indices', (self.kd.nibzkpts, self.pd.ngmax),
                         np.int32)

        if self.bd.comm.rank > 0:
            return

        Q_G = np.empty(self.pd.ngmax, np.int32)
        kk = 0
        for r in range(self.kd.comm.size):
            for q, ks in enumerate(self.kd.get_indices(r)):
                s, k = divmod(ks, self.kd.nibzkpts)
                ng = self.ng_k[k]
                if s == 1:
                    return
                if r == self.kd.comm.rank:
                    Q_G[:ng] = self.pd.Q_qG[q]
                    if r > 0:
                        self.kd.comm.send(Q_G, 0)
                if self.kd.comm.rank == 0:
                    if r > 0:
                        self.kd.comm.receive(Q_G, r)
                    Q_G[ng:] = -1
                    writer.fill(Q_G)
                    assert k == kk
                    kk += 1

    def read(self, reader):
        FDPWWaveFunctions.read(self, reader)

        if 'coefficients' not in reader.wave_functions:
            return

        Q_kG = reader.wave_functions.indices
        for kpt in self.kpt_u:
            if kpt.s == 0:
                Q_G = Q_kG[kpt.k]
                ng = self.ng_k[kpt.k]
                assert (Q_G[:ng] == self.pd.Q_qG[kpt.q]).all()
                assert (Q_G[ng:] == -1).all()

        c = reader.bohr**1.5
        if reader.version < 0:
            c = 1  # old gpw file
        for kpt in self.mykpts:
            ng = self.ng_k[kpt.k]
            index = (kpt.s, kpt.k) if self.collinear else (kpt.k,)
            psit_nG = reader.wave_functions.proxy('coefficients', *index)
            psit_nG.scale = c
            psit_nG.length_of_last_dimension = ng

            kpt.psit = PlaneWaveExpansionWaveFunctions(
                self.bd.nbands, self.pd, self.dtype, psit_nG,
                kpt=kpt.q, dist=(self.bd.comm, self.bd.comm.size),
                spin=kpt.s, collinear=self.collinear)

        if self.world.size > 1:
            # Read to memory:
            for kpt in self.kpt_u:
                kpt.psit.read_from_file()

    def hs(self, ham, q=-1, s=0, md=None):
        npw = len(self.pd.Q_qG[q])
        N = self.pd.tmp_R.size

        if md is None:
            H_GG = np.zeros((npw, npw), complex)
            S_GG = np.zeros((npw, npw), complex)
            G1 = 0
            G2 = npw
        else:
            H_GG = md.zeros(dtype=complex)
            S_GG = md.zeros(dtype=complex)
            if S_GG.size == 0:
                return H_GG, S_GG
            G1, G2 = next(md.my_blocks(S_GG))[:2]

        H_GG.ravel()[G1::npw + 1] = (0.5 * self.pd.gd.dv / N *
                                     self.pd.G2_qG[q][G1:G2])

        for G in range(G1, G2):
            x_G = self.pd.zeros(q=q)
            x_G[G] = 1.0
            H_GG[G - G1] += (self.pd.gd.dv / N *
                             self.pd.fft(ham.vt_sG[s] *
                                         self.pd.ifft(x_G, q), q))

        S_GG.ravel()[G1::npw + 1] = self.pd.gd.dv / N

        f_IG = self.pt.expand(q)
        nI = len(f_IG)
        dH_II = np.zeros((nI, nI))
        dS_II = np.zeros((nI, nI))
        I1 = 0
        for a in self.pt.my_atom_indices:
            dH_ii = unpack(ham.dH_asp[a][s])
            dS_ii = self.setups[a].dO_ii
            I2 = I1 + len(dS_ii)
            dH_II[I1:I2, I1:I2] = dH_ii / N**2
            dS_II[I1:I2, I1:I2] = dS_ii / N**2
            I1 = I2

        H_GG += np.dot(f_IG.T[G1:G2].conj(), np.dot(dH_II, f_IG))
        S_GG += np.dot(f_IG.T[G1:G2].conj(), np.dot(dS_II, f_IG))

        return H_GG, S_GG

    @timer('Full diag')
    def diagonalize_full_hamiltonian(self, ham, atoms, occupations, log,
                                     nbands=None, ecut=None, scalapack=None,
                                     expert=False):

        if self.dtype != complex:
            raise ValueError('Your wavefunctions are not complex as '
                             'required by the PW diagonalization routine.\n'
                             'Please supply GPAW(..., dtype=complex, ...) '
                             'as an argument to the calculator to enforce '
                             'complex wavefunctions.')

        S = self.bd.comm.size

        if nbands is None and ecut is None:
            nbands = self.pd.ngmin // S * S
        elif nbands is None:
            ecut /= Ha
            vol = abs(np.linalg.det(self.gd.cell_cv))
            nbands = int(vol * ecut**1.5 * 2**0.5 / 3 / pi**2)

        if nbands % S != 0:
            nbands += S - nbands % S

        assert nbands <= self.pd.ngmin

        if expert:
            iu = nbands
        else:
            iu = None

        self.bd = bd = BandDescriptor(nbands, self.bd.comm)

        log('Diagonalizing full Hamiltonian ({0} lowest bands)'.format(nbands))
        log('Matrix size (min, max): {0}, {1}'.format(self.pd.ngmin,
                                                      self.pd.ngmax))
        mem = 3 * self.pd.ngmax**2 * 16 / S / 1024**2
        log('Approximate memory used per core to store H_GG, S_GG: {0:.3f} MB'
            .format(mem))
        log('Notice: Up to twice the amount of memory might be allocated\n'
            'during diagonalization algorithm.')
        log('The least memory is required when the parallelization is purely\n'
            'over states (bands) and not k-points, set '
            "GPAW(..., parallel={'kpt': 1}, ...).")

        if S > 1:
            if isinstance(scalapack, (list, tuple)):
                nprow, npcol, b = scalapack
            else:
                nprow = int(round(S**0.5))
                while S % nprow != 0:
                    nprow -= 1
                npcol = S // nprow
                b = 64
            log('ScaLapack grid: {0}x{1},'.format(nprow, npcol),
                'block-size:', b)
            bg = BlacsGrid(bd.comm, S, 1)
            bg2 = BlacsGrid(bd.comm, nprow, npcol)
            scalapack = True
        else:
            nprow = npcol = 1
            scalapack = False

        self.set_positions(atoms.get_scaled_positions())
        self.mykpts[0].P = None
        self.allocate_arrays_for_projections(self.pt.my_atom_indices)

        myslice = bd.get_slice()

        pb = ProgressBar(log.fd)
        nkpt = len(self.kpt_u)

        for u, kpt in enumerate(self.kpt_u):
            pb.update(u / nkpt)
            npw = len(self.pd.Q_qG[kpt.q])
            if scalapack:
                mynpw = -(-npw // S)
                md = BlacsDescriptor(bg, npw, npw, mynpw, npw)
                md2 = BlacsDescriptor(bg2, npw, npw, b, b)
            else:
                md = md2 = MatrixDescriptor(npw, npw)

            with self.timer('Build H and S'):
                H_GG, S_GG = self.hs(ham, kpt.q, kpt.s, md)

            if scalapack:
                r = Redistributor(bd.comm, md, md2)
                H_GG = r.redistribute(H_GG)
                S_GG = r.redistribute(S_GG)

            psit_nG = md2.empty(dtype=complex)
            eps_n = np.empty(npw)

            with self.timer('Diagonalize'):
                if not scalapack:
                    md2.general_diagonalize_dc(H_GG, S_GG, psit_nG, eps_n,
                                               iu=iu)
                else:
                    md2.general_diagonalize_dc(H_GG, S_GG, psit_nG, eps_n)
            del H_GG, S_GG

            kpt.eps_n = eps_n[myslice].copy()

            if scalapack:
                md3 = BlacsDescriptor(bg, npw, npw, bd.maxmynbands, npw)
                r = Redistributor(bd.comm, md2, md3)
                psit_nG = r.redistribute(psit_nG)

            kpt.psit = PlaneWaveExpansionWaveFunctions(
                self.bd.nbands, self.pd, self.dtype,
                psit_nG[:bd.mynbands].copy(),
                kpt=kpt.q, dist=(self.bd.comm, self.bd.comm.size),
                spin=kpt.s, collinear=self.collinear)
            del psit_nG

            with self.timer('Projections'):
                self.pt.integrate(kpt.psit_nG, kpt.P_ani, kpt.q)

            kpt.f_n = None

        pb.finish()

        occupations.calculate(self)

        return nbands

    def initialize_from_lcao_coefficients(self, basis_functions):
        N_c = self.gd.N_c

        N = len(self.mykpts[0].C_nM)
        psit_nR = self.gd.empty(N, self.dtype)

        for kpt in self.mykpts:
            if self.kd.gamma:
                emikr_R = 1.0
            else:
                k_c = self.kd.ibzk_kc[kpt.k]
                emikr_R = np.exp(-2j * pi *
                                 np.dot(np.indices(N_c).T, k_c / N_c).T)

            psit_nR[:] = 0.0
            basis_functions.lcao_to_grid(kpt.C_nM, psit_nR, kpt.q)
            kpt.C_nM = None

            kpt.psit = PlaneWaveExpansionWaveFunctions(
                self.bd.nbands, self.pd, self.dtype, kpt=kpt.q,
                dist=(self.bd.comm, -1, 1),
                spin=kpt.s, collinear=self.collinear)

            psit_nG = kpt.psit.array
            for psit_G, psit_R in zip(psit_nG.reshape((-1, psit_nG.shape[-1])),
                                      psit_nR):
                psit_G[:] = self.pd.fft(psit_R * emikr_R, kpt.q)

    def random_wave_functions(self, mynao):
        rs = np.random.RandomState(self.world.rank)
        for kpt in self.kpt_u:
            if kpt.psit is None:
                kpt.psit = PlaneWaveExpansionWaveFunctions(
                    self.bd.nbands, self.pd, self.dtype, kpt=kpt.q,
                    dist=(self.bd.comm, -1, 1),
                    spin=kpt.s, collinear=self.collinear)

            array = kpt.psit.array[mynao:]
            weight_G = 1.0 / (1.0 + self.pd.G2_qG[kpt.q])
            array.real = rs.uniform(-1, 1, array.shape) * weight_G
            array.imag = rs.uniform(-1, 1, array.shape) * weight_G

    def estimate_memory(self, mem):
        FDPWWaveFunctions.estimate_memory(self, mem)
        self.pd.estimate_memory(mem.subnode('PW-descriptor'))

    def get_kinetic_stress(self):
        sigma_vv = np.zeros((3, 3), dtype=complex)
        pd = self.pd
        dOmega = pd.gd.dv / pd.gd.N_c.prod()
        if pd.dtype == float:
            dOmega *= 2
        K_qv = self.pd.K_qv
        for kpt in self.kpt_u:
            G_Gv = pd.get_reciprocal_vectors(q=kpt.q, add_q=False)
            psit2_G = 0.0
            for n, f in enumerate(kpt.f_n):
                psit2_G += f * np.abs(kpt.psit_nG[n])**2
            for alpha in range(3):
                Ga_G = G_Gv[:, alpha] + K_qv[kpt.q, alpha]
                for beta in range(3):
                    Gb_G = G_Gv[:, beta] + K_qv[kpt.q, beta]
                    sigma_vv[alpha, beta] += (psit2_G * Ga_G * Gb_G).sum()

        sigma_vv *= -dOmega

        self.bd.comm.sum(sigma_vv)
        self.kd.comm.sum(sigma_vv)
        return sigma_vv


def ft(spline):
    l = spline.get_angular_momentum_number()
    rc = 50.0
    N = 2**10
    assert spline.get_cutoff() <= rc

    dr = rc / N
    r_r = np.arange(N) * dr
    dk = pi / 2 / rc
    k_q = np.arange(2 * N) * dk
    f_r = spline.map(r_r) * (4 * pi)

    f_q = fbt(l, f_r, r_r, k_q)
    f_q[1:] /= k_q[1:]**(2 * l + 1)
    f_q[0] = (np.dot(f_r, r_r**(2 + 2 * l)) *
              dr * 2**l * fac(l) / fac(2 * l + 1))

    return Spline(l, k_q[-1], f_q)


class PWLFC(BaseLFC):
    def __init__(self, spline_aj, pd, blocksize=5000, comm=None):
        """Reciprocal-space plane-wave localized function collection.

        spline_aj: list of list of spline objects
            Splines.
        pd: PWDescriptor
            Plane-wave descriptor object.
        blocksize: int
            Block-size to use when looping over G-vectors.  Use None for
            doing all G-vectors in one big block.
        comm: communicator
            Communicator for operations that support parallelization
            over planewaves (only integrate so far)."""

        self.pd = pd
        self.spline_aj = spline_aj

        self.dtype = pd.dtype

        self.initialized = False

        # These will be filled in later:
        self.lf_aj = []
        self.Y_qLG = []

        if blocksize is not None:
            if pd.ngmax <= blocksize:
                # No need to block G-vectors
                blocksize = None
        self.blocksize = blocksize

        # These are set later in set_potitions():
        self.eikR_qa = None
        self.my_atom_indices = None
        self.indices = None
        self.pos_av = None
        self.nI = None

        if comm is None:
            comm = pd.gd.comm
        self.comm = comm

    def initialize(self):
        if self.initialized:
            return

        cache = {}
        lmax = -1

        # Fourier transform radial functions:
        for a, spline_j in enumerate(self.spline_aj):
            self.lf_aj.append([])
            for spline in spline_j:
                l = spline.get_angular_momentum_number()
                if spline not in cache:
                    f = ft(spline)
                    f_qG = []
                    for G2_G in self.pd.G2_qG:
                        G_G = G2_G**0.5
                        f_qG.append(f.map(G_G))
                    cache[spline] = f_qG
                else:
                    f_qG = cache[spline]
                self.lf_aj[a].append((l, f_qG))
                lmax = max(lmax, l)

        # Spherical harmonics:
        for q, K_v in enumerate(self.pd.K_qv):
            G_Gv = self.pd.get_reciprocal_vectors(q=q)
            Y_LG = np.empty(((lmax + 1)**2, len(G_Gv)))
            for L in range((lmax + 1)**2):
                Y_LG[L] = Y(L, *G_Gv.T)
            self.Y_qLG.append(Y_LG)

        self.initialized = True

    def estimate_memory(self, mem):
        splines = set()
        lmax = -1
        for spline_j in self.spline_aj:
            for spline in spline_j:
                splines.add(spline)
                l = spline.get_angular_momentum_number()
                lmax = max(lmax, l)
        nbytes = ((len(splines) + (lmax + 1)**2) *
                  sum(G2_G.nbytes for G2_G in self.pd.G2_qG))
        mem.subnode('Arrays', nbytes)

    def get_function_count(self, a):
        return sum(2 * l + 1 for l, f_qG in self.lf_aj[a])

    def __iter__(self):
        I1 = 0
        for a in self.my_atom_indices:
            j = 0
            i1 = 0
            for l, f_qG in self.lf_aj[a]:
                i2 = i1 + 2 * l + 1
                I2 = I1 + 2 * l + 1
                yield a, j, i1, i2, I1, I2
                i1 = i2
                I1 = I2
                j += 1

    def set_positions(self, spos_ac):
        self.initialize()
        kd = self.pd.kd
        if kd is None or kd.gamma:
            self.eikR_qa = np.ones((1, len(spos_ac)))
        else:
            self.eikR_qa = np.exp(2j * pi * np.dot(kd.ibzk_qc, spos_ac.T))

        self.pos_av = np.dot(spos_ac, self.pd.gd.cell_cv)

        self.my_atom_indices = np.arange(len(spos_ac))
        self.indices = []
        I1 = 0
        for a in self.my_atom_indices:
            I2 = I1 + self.get_function_count(a)
            self.indices.append((a, I1, I2))
            I1 = I2
        self.nI = I1

    def old_expand(self, q=-1, G1=0, G2=None):
        # Pure-Python version of expand().  Left here for testing.
        if G2 is None:
            G2 = self.Y_qLG[q].shape[1]
        f_IG = np.empty((self.nI, G2 - G1), complex)
        emiGR_Ga = np.exp(-1j * np.dot(self.pd.G_Qv[self.pd.Q_qG[q][G1:G2]],
                                       self.pos_av.T))
        for a, j, i1, i2, I1, I2 in self:
            l, f_qG = self.lf_aj[a][j]
            f_IG[I1:I2] = (emiGR_Ga[:, a] * f_qG[q][G1:G2] * (-1.0j)**l *
                           self.Y_qLG[q][l**2:(l + 1)**2, G1:G2])
        return f_IG

    def expand(self, q=-1, G1=0, G2=None):
        if G2 is None:
            G2 = self.Y_qLG[q].shape[1]
        G_Qv = self.pd.G_Qv[self.pd.Q_qG[q][G1:G2]]
        f_IG = np.empty((self.nI, G2 - G1), complex)
        emiGRbuf_G = np.empty(len(G_Qv), complex)

        Y_LG = self.Y_qLG[q]

        _gpaw.pwlfc_expand(G_Qv, self.pos_av,
                           self.lf_aj, Y_LG, q, G1, G2,
                           f_IG, emiGRbuf_G)
        return f_IG

    def block(self, q=-1, serial=True):
        nG = self.Y_qLG[q].shape[1]
        iblock = 0
        if self.blocksize:
            G1 = 0
            while G1 < nG:
                G2 = min(G1 + self.blocksize, nG)
                if serial or iblock % self.comm.size == self.comm.rank:
                    yield G1, G2
                iblock += 1
                G1 = G2
        else:
            if serial or self.comm.rank == 0:
                yield 0, nG
            else:
                yield 0, 0

    def add(self, a_xG, c_axi=1.0, q=-1, f0_IG=None, serial=True):
        if isinstance(c_axi, float):
            assert q == -1, a_xG.dims == 1
            a_xG += (c_axi / self.pd.gd.dv) * self.expand(-1).sum(0)
            return

        c_xI = np.empty(a_xG.shape[:-1] + (self.nI,), self.pd.dtype)
        for a, I1, I2 in self.indices:
            c_xI[..., I1:I2] = c_axi[a] * self.eikR_qa[q][a].conj()
        c_xI = c_xI.reshape((np.prod(c_xI.shape[:-1], dtype=int), self.nI))

        a_xG = a_xG.reshape((-1, a_xG.shape[-1])).view(self.pd.dtype)

        for G1, G2 in self.block(q, serial=serial):
            if f0_IG is None:
                f_IG = self.expand(q, G1, G2)
            else:
                f_IG = f0_IG

            if self.pd.dtype == float:
                f_IG = f_IG.view(float)
                G1 *= 2
                G2 *= 2

            gemm(1.0 / self.pd.gd.dv, f_IG, c_xI, 1.0, a_xG[:, G1:G2])

    def integrate(self, a_xG, c_axi=None, q=-1):
        c_xI = np.zeros(a_xG.shape[:-1] + (self.nI,), self.pd.dtype)

        b_xI = c_xI.reshape((np.prod(c_xI.shape[:-1], dtype=int), self.nI))
        a_xG = a_xG.reshape((-1, a_xG.shape[-1]))

        alpha = 1.0 / self.pd.gd.N_c.prod()
        if self.pd.dtype == float:
            alpha *= 2
            a_xG = a_xG.view(float)

        if c_axi is None:
            c_axi = self.dict(a_xG.shape[:-1])

        for G1, G2 in self.block(q, serial=False):
            if G1 > 0:
                x = 1.0
            else:
                x = 0.0
            f_IG = self.expand(q, G1, G2)
            if self.pd.dtype == float:
                if G1 == 0:
                    f_IG[:, 0] *= 0.5
                f_IG = f_IG.view(float)
                G1 *= 2
                G2 *= 2

            gemm(alpha, f_IG, a_xG[:, G1:G2], x, b_xI, 'c')
        self.comm.sum(b_xI)

        for a, I1, I2 in self.indices:
            c_axi[a][:] = self.eikR_qa[q][a] * c_xI[..., I1:I2]

        return c_axi

    def matrix_elements(self, psit, out):
        P_ani = {a: P_in.T for a, P_in in out.items()}
        self.integrate(psit.array, P_ani, psit.kpt)

    def derivative(self, a_xG, c_axiv, q=-1):
        c_vxI = np.zeros((3,) + a_xG.shape[:-1] + (self.nI,), self.pd.dtype)
        b_vxI = c_vxI.reshape((3, np.prod(c_vxI.shape[1:-1], dtype=int),
                               self.nI))
        a_xG = a_xG.reshape((-1, a_xG.shape[-1])).view(self.pd.dtype)

        alpha = 1.0 / self.pd.gd.N_c.prod()

        K_v = self.pd.K_qv[q]

        serial = False

        x = 0.0
        for G1, G2 in self.block(q, serial=serial):
            f_IG = self.expand(q, G1, G2)
            G_Gv = self.pd.G_Qv[self.pd.Q_qG[q][G1:G2]]
            if self.pd.dtype == float:
                for v in range(3):
                    gemm(2 * alpha,
                         (f_IG * 1.0j * G_Gv[:, v]).view(float),
                         a_xG[:, 2 * G1:2 * G2],
                         x, b_vxI[v], 'c')
            else:
                for v in range(3):
                    gemm(-alpha,
                         f_IG * (G_Gv[:, v] + K_v[v]),
                         a_xG[:, G1:G2],
                         x, b_vxI[v], 'c')
            x = 1.0

        if not serial:
            self.comm.sum(c_vxI)

        for v in range(3):
            if self.pd.dtype == float:
                for a, I1, I2 in self.indices:
                    c_axiv[a][..., v] = c_vxI[v, ..., I1:I2]
            else:
                for a, I1, I2 in self.indices:
                    c_axiv[a][..., v] = (1.0j * self.eikR_qa[q][a] *
                                         c_vxI[v, ..., I1:I2])

    def stress_tensor_contribution(self, a_xG, c_axi=1.0, q=-1):
        cache = {}
        for a, j, i1, i2, I1, I2 in self:
            spline = self.spline_aj[a][j]
            if spline not in cache:
                s = ft(spline)
                G_G = self.pd.G2_qG[q]**0.5
                f_G = []
                dfdGoG_G = []
                for G in G_G:
                    f, dfdG = s.get_value_and_derivative(G)
                    if G < 1e-10:
                        G = 1.0
                    f_G.append(f)
                    dfdGoG_G.append(dfdG / G)
                f_G = np.array(f_G)
                dfdGoG_G = np.array(dfdGoG_G)
                cache[spline] = (f_G, dfdGoG_G)
            else:
                f_G, dfdGoG_G = cache[spline]

        if isinstance(c_axi, float):
            c_axi = dict((a, c_axi) for a in range(len(self.pos_av)))

        G0_Gv = self.pd.get_reciprocal_vectors(q=q)

        serial = False

        stress_vv = np.zeros((3, 3))
        for G1, G2 in self.block(q, serial=serial):
            G_Gv = G0_Gv[G1:G2]
            aa_xG = a_xG[..., G1:G2]
            for v1 in range(3):
                for v2 in range(3):
                    stress_vv[v1, v2] += self._stress_tensor_contribution(
                        v1, v2, cache, G1, G2, G_Gv, aa_xG, c_axi, q)

        if not serial:
            self.comm.sum(stress_vv)

        return stress_vv

    def _stress_tensor_contribution(self, v1, v2, cache, G1, G2,
                                    G_Gv, a_xG, c_axi, q):
        f_IG = np.empty((self.nI, G2 - G1), complex)
        K_v = self.pd.K_qv[q]
        for a, j, i1, i2, I1, I2 in self:
            l = self.lf_aj[a][j][0]
            spline = self.spline_aj[a][j]
            f_G, dfdGoG_G = cache[spline]

            emiGR_G = np.exp(-1j * np.dot(G_Gv, self.pos_av[a]))
            f_IG[I1:I2] = (emiGR_G * (-1.0j)**l *
                           np.exp(1j * np.dot(K_v, self.pos_av[a])) * (
                               dfdGoG_G[G1:G2] * G_Gv[:, v1] * G_Gv[:, v2] *
                               self.Y_qLG[q][l**2:(l + 1)**2, G1:G2] +
                               f_G[G1:G2] * G_Gv[:, v1] *
                               [nablarlYL(L, G_Gv.T)[v2]
                                for L in range(l**2, (l + 1)**2)]))

        c_xI = np.zeros(a_xG.shape[:-1] + (self.nI,), self.pd.dtype)

        b_xI = c_xI.reshape((np.prod(c_xI.shape[:-1], dtype=int), self.nI))
        a_xG = a_xG.reshape((-1, a_xG.shape[-1]))

        alpha = 1.0 / self.pd.gd.N_c.prod()
        if self.pd.dtype == float:
            alpha *= 2
            if G1 == 0:
                f_IG[:, 0] *= 0.5
            f_IG = f_IG.view(float)
            a_xG = a_xG.copy().view(float)

        gemm(alpha, f_IG, a_xG, 0.0, b_xI, 'c')

        stress = 0.0
        for a, I1, I2 in self.indices:
            stress -= self.eikR_qa[q][a] * (c_axi[a] * c_xI[..., I1:I2]).sum()
        return stress.real


class PseudoCoreKineticEnergyDensityLFC(PWLFC):
    def add(self, tauct_R):
        tauct_R += self.pd.ifft(1.0 / self.pd.gd.dv * self.expand(-1).sum(0))

    def derivative(self, dedtaut_R, dF_aiv):
        PWLFC.derivative(self, self.pd.fft(dedtaut_R), dF_aiv)


class ReciprocalSpaceDensity(Density):
    def __init__(self, gd, finegd, nspins, collinear, charge, redistributor,
                 background_charge=None):
        assert gd.comm.size == 1
        serial_finegd = finegd.new_descriptor(comm=gd.comm)

        from gpaw.utilities.grid import GridRedistributor
        noredist = GridRedistributor(redistributor.comm,
                                     redistributor.broadcast_comm, gd, gd)
        Density.__init__(self, gd, serial_finegd, nspins, collinear, charge,
                         redistributor=noredist,
                         background_charge=background_charge)

        self.pd2 = PWDescriptor(None, gd)
        self.pd3 = PWDescriptor(None, serial_finegd)

        self.G3_G = self.pd2.map(self.pd3)

        self.xc_redistributor = GridRedistributor(redistributor.comm,
                                                  redistributor.comm,
                                                  serial_finegd, finegd)
        self.nct_q = None
        self.nt_Q = None
        self.rhot_q = None

    def initialize(self, setups, timer, magmom_av, hund):
        Density.initialize(self, setups, timer, magmom_av, hund)

        spline_aj = []
        for setup in setups:
            if setup.nct is None:
                spline_aj.append([])
            else:
                spline_aj.append([setup.nct])
        self.nct = PWLFC(spline_aj, self.pd2)

        self.ghat = PWLFC([setup.ghat_l for setup in setups], self.pd3,
                          blocksize=256, comm=self.xc_redistributor.comm)

    def set_positions(self, spos_ac, atom_partition):
        Density.set_positions(self, spos_ac, atom_partition)
        self.nct_q = self.pd2.zeros()
        self.nct.add(self.nct_q, 1.0 / self.nspins)
        self.nct_G = self.pd2.ifft(self.nct_q)

    def interpolate_pseudo_density(self, comp_charge=None):
        """Interpolate pseudo density to fine grid."""
        if comp_charge is None:
            comp_charge, _Q_aL = self.calculate_multipole_moments()

        if self.nt_xg is None:
            self.nt_xg = self.finegd.empty(self.ncomponents)
            self.nt_sg = self.nt_xg[:self.nspins]
            self.nt_vg = self.nt_xg[self.nspins:]
            self.nt_Q = self.pd2.empty()

        self.nt_Q[:] = 0.0

        x = 0
        for nt_G, nt_g in zip(self.nt_xG, self.nt_xg):
            nt_g[:], nt_Q = self.pd2.interpolate(nt_G, self.pd3)
            if x < self.nspins:
                self.nt_Q += nt_Q
            x += 1

    def interpolate(self, in_xR, out_xR=None):
        """Interpolate array(s)."""
        if out_xR is None:
            out_xR = self.finegd.empty(in_xR.shape[:-3])

        a_xR = in_xR.reshape((-1,) + in_xR.shape[-3:])
        b_xR = out_xR.reshape((-1,) + out_xR.shape[-3:])

        for in_R, out_R in zip(a_xR, b_xR):
            out_R[:] = self.pd2.interpolate(in_R, self.pd3)[0]

        return out_xR

    distribute_and_interpolate = interpolate

    def calculate_pseudo_charge(self):
        self.rhot_q = self.pd3.zeros()
        Q_aL = self.Q.calculate(self.D_asp)
        self.ghat.add(self.rhot_q, Q_aL, serial=False)
        self.ghat.comm.sum(self.rhot_q)
        self.rhot_q[self.G3_G] += self.nt_Q * 8
        self.background_charge.add_fourier_space_charge_to(self.pd3,
                                                           self.rhot_q)
        self.rhot_q[0] = 0.0

    def get_pseudo_core_kinetic_energy_density_lfc(self):
        return PseudoCoreKineticEnergyDensityLFC(
            [[setup.tauct] for setup in self.setups], self.pd2)

    def calculate_dipole_moment(self):
        if LooseVersion(np.__version__) < '1.6.0':
            raise NotImplementedError
        pd = self.pd3
        N_c = pd.tmp_Q.shape

        m0_q, m1_q, m2_q = [i_G == 0
                            for i_G in np.unravel_index(pd.Q_qG[0], N_c)]
        rhot_q = self.rhot_q.imag
        rhot_cs = [rhot_q[m1_q & m2_q],
                   rhot_q[m0_q & m2_q],
                   rhot_q[m0_q & m1_q]]
        d_c = [np.dot(rhot_s[1:], 1.0 / np.arange(1, len(rhot_s)))
               for rhot_s in rhot_cs]
        return -np.dot(d_c, pd.gd.cell_cv) / pi * pd.gd.dv


class ReciprocalSpacePoissonSolver:
    def __init__(self, pd, realpbc_c):
        self.pd = pd
        self.realpbc_c = realpbc_c
        self.G2_q = pd.G2_qG[0][1:]

    def initialize(self):
        pass

    def get_stencil(self):
        return '????'

    def estimate_memory(self, mem):
        pass

    def todict(self):
        return {}

    def solve(self, vHt_q, dens):
        vHt_q[:] = 4 * pi * dens.rhot_q
        vHt_q[1:] /= self.G2_q


class ReciprocalSpaceHamiltonian(Hamiltonian):
    def __init__(self, gd, finegd, pd2, pd3, nspins, collinear,
                 setups, timer, xc, world, vext=None,
                 psolver=None, redistributor=None, realpbc_c=None):

        assert gd.comm.size == 1
        assert finegd.comm.size == 1
        assert redistributor is not None  # XXX should not be like this
        Hamiltonian.__init__(self, gd, finegd, nspins, collinear, setups,
                             timer, xc, world, vext=vext,
                             redistributor=redistributor)

        self.vbar = PWLFC([[setup.vbar] for setup in setups], pd2)
        self.pd2 = pd2
        self.pd3 = pd3

        self.vHt_q = pd3.empty()

        if psolver is None:
            psolver = ReciprocalSpacePoissonSolver(pd3, realpbc_c)
        elif isinstance(psolver, dict):
            direction = psolver['dipolelayer']
            assert len(psolver) == 1
            from gpaw.dipole_correction import DipoleCorrection
            psolver = DipoleCorrection(
                ReciprocalSpacePoissonSolver(pd3, realpbc_c), direction)
        self.poisson = psolver
        self.npoisson = 0

        self.vbar_Q = None
        self.vt_Q = None
        self.ebar = None
        self.epot = None
        self.exc = None

    def set_positions(self, spos_ac, atom_partition):
        Hamiltonian.set_positions(self, spos_ac, atom_partition)
        self.vbar_Q = self.pd2.zeros()
        self.vbar.add(self.vbar_Q)

    def update_pseudo_potential(self, dens):
        self.ebar = self.pd2.integrate(self.vbar_Q, dens.nt_Q)

        with self.timer('Poisson'):
            self.poisson.solve(self.vHt_q, dens)
            self.epot = 0.5 * self.pd3.integrate(self.vHt_q, dens.rhot_q)

        self.vt_Q = self.vbar_Q + self.vHt_q[dens.G3_G] / 8
        self.e_external = 0.0

        if self.vext is not None:
            gd = self.finegd
            vext_q = self.vext.get_potentialq(gd, self.pd3)
            self.vt_Q += vext_q[dens.G3_G] / 8
            self.e_external = self.pd3.integrate(vext_q, dens.rhot_q)

        self.vt_sG[:] = self.pd2.ifft(self.vt_Q)

        self.timer.start('XC 3D grid')
        nt_dist_xg = dens.xc_redistributor.distribute(dens.nt_xg)
        vxct_dist_xg = np.zeros_like(nt_dist_xg)
        self.exc = self.xc.calculate(dens.xc_redistributor.aux_gd,
                                     nt_dist_xg, vxct_dist_xg)
        vxct_xg = dens.xc_redistributor.collect(vxct_dist_xg)

        x = 0
        for vt_G, vxct_g in zip(self.vt_xG, vxct_xg):
            vxc_G, vxc_Q = self.pd3.restrict(vxct_g, self.pd2)
            if x < self.nspins:
                vt_G += vxc_G
                self.vt_Q += vxc_Q / self.nspins
            else:
                vt_G[:] = vxc_G
            x += 1

        self.timer.stop('XC 3D grid')

        return np.array([self.epot, self.ebar, self.e_external, self.exc])

    def calculate_atomic_hamiltonians(self, density):
        W_aL = {}
        for a in density.D_asp:
            W_aL[a] = np.empty((self.setups[a].lmax + 1)**2)
        if self.vext:
            vext_q = self.vext.get_potentialq(self.finegd, self.pd3)
            density.ghat.integrate(self.vHt_q+vext_q, W_aL)
        else:
            density.ghat.integrate(self.vHt_q, W_aL)
        return W_aL

    def calculate_kinetic_energy(self, density):
        ekin = 0.0
        for vt_G, nt_G in zip(self.vt_xG, density.nt_xG):
            ekin -= self.gd.integrate(vt_G, nt_G)
        ekin += self.gd.integrate(self.vt_sG, density.nct_G).sum()
        return ekin

    def restrict(self, in_xR, out_xR=None):
        """Restrict array."""
        if out_xR is None:
            out_xR = self.gd.empty(in_xR.shape[:-3])

        a_xR = in_xR.reshape((-1,) + in_xR.shape[-3:])
        b_xR = out_xR.reshape((-1,) + out_xR.shape[-3:])

        for in_R, out_R in zip(a_xR, b_xR):
            out_R[:] = self.pd3.restrict(in_R, self.pd2)[0]

        return out_xR

    restrict_and_collect = restrict

    def calculate_forces2(self, dens, ghat_aLv, nct_av, vbar_av):
        if self.vext:
            vext_q = self.vext.get_potentialq(self.finegd, self.pd3)
            dens.ghat.derivative(self.vHt_q+vext_q, ghat_aLv)
        else:
            dens.ghat.derivative(self.vHt_q, ghat_aLv)
        dens.nct.derivative(self.vt_Q, nct_av)
        self.vbar.derivative(dens.nt_Q, vbar_av)

    def get_electrostatic_potential(self, dens):
        self.poisson.solve(self.vHt_q, dens)
        return self.pd3.ifft(self.vHt_q)
