#
# Copyright 2015-2016 Universidad Complutense de Madrid
#
# This file is part of Numina
#
# SPDX-License-Identifier: GPL-3.0+
# License-Filename: LICENSE.txt
#

from __future__ import division
from __future__ import print_function

import argparse
import numpy as np

from .pause_debugplot import pause_debugplot


def ximplotxy(x, y, plottype=None,
              xlim=None, ylim=None, 
              xlabel=None, ylabel=None, title=None,
              show=True, geometry=(0, 0, 640, 480), tight_layout=True,
              debugplot=0, **kwargs):
    """
    Parameters
    ----------
    x : 1d numpy array, float
        Array containing the X coordinate.
    y : 1d numpy array, float
        Array containing the Y coordinate.
    plottype : string
        Plot type. It can be 'semilog' or normal (default).
    xlim : tuple of floats
        Tuple defining the x-axis range.
    ylim : tuple of floats
        Tuple defining the y-axis range.
    xlabel : string
        X-axis label.
    ylabel : string
        Y-axis label.
    title : string
        Plot title.
    show : bool
        If True, the function shows the displayed image. Otherwise
        plt.show() is expected to be executed outside.
    geometry : tuple (4 integers) or None
        x, y, dx, dy values employed to set the Qt backend geometry.
    tight_layout : bool
        If True, and show=True, a tight display layout is set.
    debugplot : int
        Determines whether intermediate computations and/or plots
        are displayed. The valid codes are defined in
        numina.array.display.pause_debugplot.

    Returns
    -------
    ax : axes object
        Matplotlib axes instance. This value is returned only when
        'show' is False.

    """

    from numina.array.display.matplotlib_qt import plt
    fig = plt.figure()
    ax = fig.add_subplot(111)
    if plottype == 'semilog':
        ax.semilogy(x, y, **kwargs)
    else:
        ax.plot(x, y, **kwargs)

    if xlim is not None:
        ax.set_xlim(xlim[0], xlim[1])
    if ylim is not None:
        ax.set_ylim(ylim[0], ylim[1])

    if xlabel is not None:
        ax.set_xlabel(xlabel)
    if ylabel is not None:
        ax.set_ylabel(ylabel)
    if title is not None:
        ax.set_title(title)

    if geometry is not None:
        x_geom, y_geom, dx_geom, dy_geom = geometry
        mngr = plt.get_current_fig_manager()
        mngr.window.setGeometry(x_geom, y_geom, dx_geom, dy_geom)

    if show:
        pause_debugplot(debugplot, pltshow=show, tight_layout=tight_layout)
    else:
        # return axes
        return ax


def main(args=None):

    # parse command-line options
    parser = argparse.ArgumentParser(prog='ximplotxy')
    parser.add_argument("filename",
                        help="ASCII file with data in columns")
    parser.add_argument("col1",
                        help="Column number for X data",
                        type=int)
    parser.add_argument("col2",
                        help="Column number for Y data",
                        type=int)
    parser.add_argument("--kwargs",
                        help="Extra arguments for plot, e.g.: "
                             "\"{'marker':'o',"
                             " 'linestyle':'dotted',"
                             " 'xlabel':'x axis', 'ylabel':'y axis',"
                             " 'title':'sample plot',"
                             " 'xlim':[-1,1], 'ylim':[-2,2],"
                             " 'label':'sample data',"
                             " 'color':'magenta'}\"")
    args = parser.parse_args(args)

    # ASCII file
    filename = args.filename

    # columns to be plotted (first column will be number 1 and not 0)
    col1 = args.col1 - 1
    col2 = args.col2 - 1

    # read ASCII file
    bigtable = np.genfromtxt(filename)
    x = bigtable[:, col1]
    y = bigtable[:, col2]

    if args.kwargs is None:
        ximplotxy(x, y, debugplot=12, marker='o', linestyle='')
    else:
        ximplotxy(x, y, debugplot=12, **eval(args.kwargs))


if __name__ == '__main__':
    main()
