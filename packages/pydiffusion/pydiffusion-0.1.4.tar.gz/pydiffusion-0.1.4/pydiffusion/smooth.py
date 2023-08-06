"""
The smooth module provides tools for data smoothing of original diffusion
profile data.
"""

import numpy as np
import matplotlib.pyplot as plt
from scipy.interpolate import splev, splrep
from pydiffusion.core import DiffProfile
from pydiffusion.io import ask_input, ita_start, ita_finish


def movingradius(dis, X, r):
    """
    Data smooth using moving average within a radius range
    The first and last data point is unchanged, data at d is averaged by data
    in the range of (d-r, d+r)

    Parameters
    ----------
    dis, X : numpy.array
        Diffusion profile required to be smoothed
    r : float
        Radius (micron) of moving average
    """
    dmin, dmax = dis[0], dis[-1]
    n = np.size(dis)
    f = splrep(dis, X, k=1)
    Xnew = np.zeros(n)
    for i in range(n):
        h = min(abs(dmin-dis[i]), abs(dmax-dis[i]), r)
        Xnew[i] = np.mean(splev(np.linspace(dis[i]-h, dis[i]+h, 100), f))
    return Xnew


def phasesmooth(dis, X, ax):
    """
    Data smooth of a single phase, Using movingradius method.

    Parameters
    ----------
    dis, X : numpy.array
        Diffusion profile within a single phase
    """
    Xsm = X.copy()
    smoo = True
    while smoo:
        ax.cla()
        ax.plot(dis, Xsm, 'bo')
        plt.draw()

        # Zoom in or not
        ipt = ask_input('Enter the zoomed in start and end location, if no, enter nothing\n')
        zm = [float(i) for i in ipt.split(' ')] if ipt != '' else [dis[0], dis[-1]]
        zmid = np.where((dis >= zm[0]) & (dis <= zm[1]))[0]

        sm = True
        while sm:
            ax.cla()
            ax.plot(dis[zmid], Xsm[zmid], 'bo')
            plt.draw()
            Xsmn = np.copy(Xsm[zmid])
            msg = 'Enter Start and End Composition for this region: ['
            msg += str(Xsmn[0])+' '+str(Xsmn[-1])+']\n'
            ipt = ask_input(msg)
            if ipt != '':
                Xsmn[0], Xsmn[-1] = [float(i) for i in ipt.split(' ')]
            msg = 'Smooth Radius and Times: [1 1]\n'
            ipt = ask_input(msg)
            if ipt != '':
                ipt = ipt.split(' ')
                r, t = float(ipt[0]), int(ipt[1])
            else:
                r, t = 1.0, 1
            for i in range(t):
                Xsmn = movingradius(dis[zmid], Xsmn, r)
            ax.cla()
            ax.plot(dis[zmid], Xsm[zmid], 'bo', dis[zmid], Xsmn, 'ro')
            plt.draw()
            ipt = ask_input('Redo this smooth? (y/[n])')
            sm = True if 'y' in ipt or 'Y' in ipt else False
            if not sm:
                Xsm[zmid] = Xsmn
        ax.cla()
        ax.plot(dis, X, 'bo', dis, Xsm, 'ro')
        plt.draw()
        ipt = ask_input('Further smooth for this phase? (y/[n])')
        smoo = True if 'y' in ipt or 'Y' in ipt else False
    return Xsm


def datasmooth(profile, interface=[], n=2000):
    """
    Data smooth of diffusion profile. The functions use moving radius method
    on each phase.

    Parameters
    ----------
    profile: DiffProfile
        Diffusion profile data.
    interface : list of float
        Np-1 locations of interfaces for a Np-phase system
    n : int
        Interpolation number of the smoothed profile

    Returns
    -------
    profile : pydiffusion.diffusion.DiffProfile
    """
    dis, X = profile.dis, profile.X
    if len(dis) != len(X):
        raise ValueError('Nonequal length of distance and composition data')
    try:
        If = np.array(interface)
    except TypeError:
        print('interface must be array_like')
    if If.ndim != 1:
        raise ValueError('interface must be 1d-array')

    fig = plt.figure()
    ax = fig.add_subplot(111)

    If = [dis[0]-0.5] + list(If) + [dis[-1]+0.5]
    Np = len(If)-1
    Ip = [0]*(Np+1)
    disn, Xn = dis.copy(), X.copy()

    # Two distance value is the same
    for i in range(len(disn)-1):
        if disn[i] == disn[i+1]:
            disn[i+1] += 1e-5

    ita_start()

    # Apply phasesmooth to each phase
    for i in range(Np):
        pid = np.where((disn > If[i]) & (disn < If[i+1]))[0]
        Xn[pid] = phasesmooth(disn[pid], Xn[pid], ax)
    plt.close()

    # Create a sharp interface at interface locations
    # Solubility of each phase will be extended a little
    for i in range(1, Np):
        pid = np.where(disn > If[i])[0][0]
        start = max(pid-5, np.where(disn > If[i-1])[0][0])
        end = min(pid+5, np.where(disn < If[i+1])[0][-1])
        fX1 = splrep(disn[start:pid], Xn[start:pid], k=2)
        fX2 = splrep(disn[pid:end], Xn[pid:end], k=2)
        disn = np.insert(disn, pid, [If[i], If[i]])
        Xn = np.insert(Xn, pid, [splev(If[i], fX1), splev(If[i], fX2)])
        Ip[i] = pid+1
    Ip[-1] = len(Xn)
    disni, Xni = disn.copy(), Xn.copy()

    # Interpolation
    if n > 0:
        ni = [int(n*(If[i]-If[0])//(If[-1]-If[0])) for i in range(Np)]+[n]
        disni, Xni = np.zeros(n), np.zeros(n)
        for i in range(Np):
            fX = splrep(disn[Ip[i]:Ip[i+1]], Xn[Ip[i]:Ip[i+1]], k=1)
            disni[ni[i]:ni[i+1]] = np.linspace(disn[Ip[i]], disn[Ip[i+1]-1], ni[i+1]-ni[i])
            Xni[ni[i]:ni[i+1]] = splev(disni[ni[i]:ni[i+1]], fX)

    ita_finish()

    print('Data smoothing finished')

    return DiffProfile(disni, Xni, If[1:-1])
