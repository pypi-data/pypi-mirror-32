#!/usr/bin/env python
# -*- coding: utf-8 -*-

"""
Two coupled decays
------------------

:download:`examples/decay.py` demonstrates accuracy
by comparison with analytic solution for a simple system
of two coupled decays

::

 $ python decay.py --help

.. exec::
   echo "::\\n\\n"
   python examples/examples/decay.py --help | sed "s/^/   /"


Here is an example generated by:

::

 $ python decay.py --plot --savefig decay.png


.. image:: ../_generated/decay.png


Motivation for sigmoid damped exp(); vary tend: 5, 700, 1700.
Never mind 700 not being correctly represented, the problem
is 1700 completely ruining the integration (NaN's due to overflow).

::

 $ python decay.py --plot --savefig decay_long.png --rates 1.0 --logy --logt \
   --rtol 1e-13 --atol 1e-6 --scale-err 100.0 --plotlogy --nt 1024 --tend 1700


.. image:: ../_generated/decay_long.png


::

 $ python decay.py --plot --savefig decay_long_damp.png --rates 1.0 --logy \
   --logt --rtol 1e-13 --atol 1e-6 --scale-err 100.0 --plotlogy --nt 1024 \
   --tend 1700 --sigm-damp


.. image:: ../_generated/decay_long_damp.png

"""

from __future__ import absolute_import, division, print_function

import argh
import numpy as np

from chemreac import ReactionDiffusion
from chemreac.integrate import run
from chemreac.util.analysis import solver_linear_error
from chemreac.util.plotting import save_and_or_show_plot


analytic = {
    0: lambda y0, k, t: (
        y0[0] * np.exp(-k[0]*t)),
    1: lambda y0, k, t: (
        y0[1] * np.exp(-k[1] * t) + y0[0] * k[0] / (k[1] - k[0]) *
        (np.exp(-k[0]*t) - np.exp(-k[1]*t))),
    2: lambda y0, k, t: (
        y0[2] * np.exp(-k[2] * t) + y0[1] * k[1] / (k[2] - k[1]) *
        (np.exp(-k[1]*t) - np.exp(-k[2]*t)) +
        k[1] * k[0] * y0[0] / (k[1] - k[0]) *
        (1 / (k[2] - k[0]) * (np.exp(-k[0]*t) - np.exp(-k[2]*t)) -
         1 / (k[2] - k[1]) * (np.exp(-k[1]*t) - np.exp(-k[2]*t))))
}


def get_Cref(k, y0, tout):
    coeffs = k + [0]*(3-len(k))
    return np.column_stack([
        analytic[i](y0, coeffs, tout) for i in range(
            min(3, len(k)+1))])


def integrate_rd(tend=2.0, A0=1.0, nt=67, t0=0.0,
                 rates='3.40715,4.0', logy=False, logt=False,
                 plot=False, savefig='None', method='bdf',
                 atol='1e-7,1e-6,1e-5', rtol='1e-6', sigm_damp=False,
                 num_jac=False, scale_err=1.0, small='None', use_log2=False,
                 plotlogy=False, plotlogt=False, verbose=False):
    """
    Analytic solution through Bateman equation =>
    ensure :math:`|k_i - k_j| \gg eps`
    """
    k = list(map(float, rates.split(',')))
    n = len(k)+1
    if n > 4:
        raise ValueError("Max 3 consequtive decays supported at the moment.")

    atol = list(map(float, atol.split(',')))
    if len(atol) == 1:
        atol = atol[0]
    rtol = float(rtol)

    rd = ReactionDiffusion(
        n, [[i] for i in range(n-1)], [[i] for i in range(1, n)],
        k, logy=logy, logt=logt, use_log2=use_log2)

    y0 = np.zeros(n)
    y0[0] = A0
    if small == 'None':
        tiny = None
    else:
        tiny = 0
        y0 += float(small)
    tout = np.linspace(t0, tend, nt)
    integr = run(rd, y0, tout, atol=atol, rtol=rtol, method=method,
                 with_jacobian=not num_jac, sigm_damp=sigm_damp, tiny=tiny)
    Cout, yout, info = integr.Cout, integr.yout, integr.info
    Cref = get_Cref(k, y0, tout - tout[0]).reshape((nt, 1, n))

    if verbose:
        print('rate: ', k)
        print(info)

    if plot:
        nshow = min(n, 3)
        try:
            min_atol = min(info['atol'])
        except:
            min_atol = info['atol']

        import matplotlib.pyplot as plt
        plt.figure(figsize=(6, 10))
        c = 'rgb'
        for i, l in enumerate('ABC'[:nshow]):
            ax = plt.subplot(nshow+1, 1, 1)
            if plotlogy:
                ax.set_yscale('log')
            if plotlogt:
                ax.set_xscale('log')
            ax.plot(tout, Cout[:, 0, i], label=l, color=c[i])

            ax = plt.subplot(nshow+1, 1, 2+i)
            if plotlogy:
                ax.set_yscale('symlog')  # abs error might be < 0
            if plotlogt:
                ax.set_xscale('log')
            ax.plot(tout, (Cout[:, 0, i]-Cref[:, 0, i])/min_atol,
                    label=l, color=c[i])

            try:
                atol = info['atol'][i]
            except:
                atol = info['atol']

            try:
                rtol = info['rtol'][i]
            except:
                rtol = info['rtol']

            le_l, le_u = solver_linear_error(
                yout[:, 0, i], rtol, atol, rd.logy, scale_err=scale_err, expb=rd.expb)
            plt.fill_between(tout, (le_l - Cout[:, 0, i])/min_atol,
                             (le_u - Cout[:, 0, i])/min_atol,
                             color=c[i], alpha=0.2)

            # Print indices and values of violations of (scaled) error bounds
            def _print(violation):
                print(violation)
                print(le_l[violation],
                      Cref[violation, 0, i],
                      le_u[violation])
            l_viols = np.where(le_l > Cref[:, 0, i])[0]
            u_viols = np.where(le_u < Cref[:, 0, i])[0]
            if verbose and (len(l_viols) > 0 or len(u_viols) > 0):
                print("Outside error bounds for rtol, atol:", rtol, atol)
                # for violation in chain(l_viols, u_viols):
                #     _print(violation)

        plt.subplot(nshow+1, 1, 1)
        plt.title('Concentration vs. time')
        plt.legend(loc='best', prop={'size': 11})
        plt.xlabel('t')
        plt.ylabel('[X]')
        for i in range(nshow):
            plt.subplot(nshow+1, 1, 2+i)
            plt.title('Absolute error in [{}](t) / min(atol)'.format('ABC'[i]))
            plt.legend(loc='best')
            plt.xlabel('t')
            plt.ylabel('|E[{0}]| / {1:7.0g}'.format('ABC'[i], min_atol))
        plt.tight_layout()
        save_and_or_show_plot(savefig=savefig)

    return integr.yout, Cref, rd, info

if __name__ == '__main__':
    argh.dispatch_command(integrate_rd, output_file=None)
