#!/usr/bin/env python

# Copyright (c) 2018, DIANA-HEP
# All rights reserved.
# 
# Redistribution and use in source and binary forms, with or without
# modification, are permitted provided that the following conditions are met:
# 
# * Redistributions of source code must retain the above copyright notice, this
#   list of conditions and the following disclaimer.
# 
# * Redistributions in binary form must reproduce the above copyright notice,
#   this list of conditions and the following disclaimer in the documentation
#   and/or other materials provided with the distribution.
# 
# * Neither the name of the copyright holder nor the names of its
#   contributors may be used to endorse or promote products derived from
#   this software without specific prior written permission.
# 
# THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
# AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
# IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
# DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
# FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
# DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
# SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
# CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
# OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
# OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

import numbers

import numpy

import histbook.axis

class Channel(object):
    """Abstract class for graphical channels in a Vega-Lite plot."""

    def __init__(self, axis):
        self.axis = axis

class OverlayChannel(Channel):
    """Represents an overlayed axis in a Vega-Lite plot."""

    def __repr__(self):
        return ".overlay({0})".format(self.axis)

class StackChannel(Channel):
    """Represents a stacked axis in a Vega-Lite plot."""

    def __init__(self, axis, order):
        super(StackChannel, self).__init__(axis)
        self.order = order

    def __repr__(self):
        if self.order is None:
            return ".stack({0})".format(self.axis)
        else:
            return ".stack({0}, order={1})".format(self.axis, repr(self.order))

class BesideChannel(Channel):
    """Represents a side-by-side axis in a Vega-Lite plot."""

    def __repr__(self):
        return ".beside({0})".format(self.axis)

class BelowChannel(Channel):
    """Represents an above-and-below axis in a Vega-Lite plot."""

    def __repr__(self):
        return ".below({0})".format(self.axis)

class TerminalChannel(Channel):
    """Abstract class for the last graphical channel in a Vega-Lite plot."""

    def __init__(self, axis, profile, error, width, height, title, config, xscale, yscale, colorscale, shapescale):
        self.axis = axis
        self.profile = profile
        self.error = error
        self.width = width
        self.height = height
        self.title = title
        self.config = config
        self.xscale = xscale
        self.yscale = yscale
        self.colorscale = colorscale
        self.shapescale = shapescale

    def __repr__(self):
        args = [repr(self.axis)]
        if self.profile is not None:
            args.append("profile={0}".format(self.profile))
        if self.error is not False:
            args.append("error={0}".format(self.error))
        if self.width is not None:
            args.append("width={0}".format(repr(self.width)))
        if self.height is not None:
            args.append("height={0}".format(repr(self.height)))
        if self.title is not None:
            args.append("title={0}".format(repr(self.title)))
        if self.config is not None:
            args.append("config={0}".format(repr(self.config)))
        if self.xscale is not None:
            args.append("xscale={0}".format(repr(self.xscale)))
        if self.yscale is not None:
            args.append("yscale={0}".format(repr(self.yscale)))
        if self.colorscale is not None:
            args.append("colorscale={0}".format(repr(self.colorscale)))
        if self.shapescale is not None:
            args.append("shapescale={0}".format(repr(self.shapescale)))
        return ".{0}({1})".format(self._method, "".join(args))

class BarChannel(TerminalChannel):
    """Represents a bar axis in a Vega-Lite plot."""
    _method = "bar"

class StepChannel(TerminalChannel):
    """Represents a step axis in a Vega-Lite plot."""
    _method = "step"

class AreaChannel(TerminalChannel):
    """Represents an area axis in a Vega-Lite plot."""
    _method = "area"

class LineChannel(TerminalChannel):
    """Represents a line axis in a Vega-Lite plot."""
    _method = "line"

class MarkerChannel(TerminalChannel):
    """Represents a marker axis in a Vega-Lite plot."""
    _method = "marker"

class PlottingChain(object):
    """Mix-in for :py:class:`Hist <histbook.hist.Hist>` (as the first in a plotting chain) and :py:class:`Channels <histbook.vega.Channel>` in a plotting chain."""

    def __init__(self, source, item):
        if isinstance(source, PlottingChain):
            self._source = source._source
            self._chain = source._chain + (item,)
        else:
            self._source = source
            self._chain = (item,)

    def __repr__(self):
        return "".join(repr(x) for x in (self._source,) + self._chain)

    def __str__(self, indent="\n     ", paren=True):
        return ("(" if paren else "") + indent.join(repr(x) for x in (self._source,) + self._chain) + (")" if paren else "")

    def _singleaxis(self, axis):
        if axis is None:
            if len(self._source._group + self._source._fixed) == 1:
                axis, = self._source._group + self._source._fixed
            else:
                raise TypeError("histogram has more than one axis; one must be specified for plotting")
        return axis

    def _asaxis(self, axis):
        if axis is None:
            return None
        elif isinstance(axis, histbook.axis.Axis):
            return axis
        else:
            return self._source.axis[axis]

    def overlay(self, axis):
        """
        Display bins in ``axis`` overlaid on each other in different colors.

        Parameters
        ----------
        axis : :py:class:`Axis <histbook.axis.Axis>`, algebraic expression (lambda or string), or index position (integer)
            the axis to overlay

        Returns
        -------
        :py:class:`PlottingChain <histbook.vega.PlottingChain>`
        """
        if any(isinstance(x, OverlayChannel) for x in self._chain):
            raise TypeError("cannot overlay an overlay")
        return PlottingChain(self, OverlayChannel(self._asaxis(axis)))

    def stack(self, axis, order=None):
        """
        Display bins in ``axis`` stacked on one another in an area plot.

        Parameters
        ----------
        axis : :py:class:`Axis <histbook.axis.Axis>`, algebraic expression (lambda or string), or index position (integer)
            the axis to overlay

        order : iterable of strings
            stacking order of bins

        Returns
        -------
        :py:class:`PlottingChain <histbook.vega.PlottingChain>`
        """
        if any(isinstance(x, StackChannel) for x in self._chain):
            raise TypeError("cannot stack a stack")
        return PlottingChain(self, StackChannel(self._asaxis(axis), order))

    def beside(self, axis):
        """
        Display bins in ``axis`` next to each other horizontally.

        Parameters
        ----------
        axis : :py:class:`Axis <histbook.axis.Axis>`, algebraic expression (lambda or string), or index position (integer)
            the axis to overlay

        Returns
        -------
        :py:class:`PlottingChain <histbook.vega.PlottingChain>`
        """
        if any(isinstance(x, BesideChannel) for x in self._chain):
            raise TypeError("cannot split plots beside each other that are already split with beside (can do beside and below)")
        return PlottingChain(self, BesideChannel(self._asaxis(axis)))

    def below(self, axis):
        """
        Display bins in ``axis`` next to each other vertically.

        Parameters
        ----------
        axis : :py:class:`Axis <histbook.axis.Axis>`, algebraic expression (lambda or string), or index position (integer)
            the axis to overlay

        Returns
        -------
        :py:class:`PlottingChain <histbook.vega.PlottingChain>`
        """
        if any(isinstance(x, BelowChannel) for x in self._chain):
            raise TypeError("cannot split plots below each other that are already split with below (can do beside and below)")
        return PlottingChain(self, BelowChannel(self._asaxis(axis)))

    def bar(self, axis=None, profile=None, error=False, width=None, height=None, title=None, config=None, xscale=None, yscale=None, colorscale=None, shapescale=None):
        """
        Display bins in ``axis`` (if not the only axis) as bars on the horizontal axis.

        Parameters
        ----------
        axis : ``None``, :py:class:`Axis <histbook.axis.Axis>`, algebraic expression (lambda or string), or index position (integer)
            the axis to overlay; if ``None`` *(default)*, use the only axis in this :py:class:`Hist <histbook.hist.Hist>`

        profile : ``None``, :py:class:`profile <histbook.axis.profile>`, algebraic expression (lambda or string) or index position (integer)
            if ``None`` *(default)*, display bin counts; otherwise, display profile means (and errors on the mean)

        error : bool
            if ``True``, overlay error bars

        width, height, title, config, xscale, yscale, colorscale, shapescale : ``None`` or JSON
            graphical directives to pass to Vega-Lite

        Returns
        -------
        :py:class:`Plotable <histbook.vega.Plotable>`
        """
        if error and any(isinstance(x, (BesideChannel, BelowChannel)) for x in self._chain):
            raise NotImplementedError("error bars are currently incompatible with splitting beside or below")
        return Plotable(self, BarChannel(self._asaxis(self._singleaxis(axis)), self._asaxis(profile), error, width, height, title, config, xscale, yscale, colorscale, shapescale))

    def step(self, axis=None, profile=None, error=False, width=None, height=None, title=None, config=None, xscale=None, yscale=None, colorscale=None, shapescale=None):
        """
        Display bins in ``axis`` (if not the only axis) as steps on the horizontal axis.

        Parameters
        ----------
        axis : ``None``, :py:class:`Axis <histbook.axis.Axis>`, algebraic expression (lambda or string), or index position (integer)
            the axis to overlay; if ``None`` *(default)*, use the only axis in this :py:class:`Hist <histbook.hist.Hist>`

        profile : ``None``, :py:class:`profile <histbook.axis.profile>`, algebraic expression (lambda or string) or index position (integer)
            if ``None`` *(default)*, display bin counts; otherwise, display profile means (and errors on the mean)

        error : bool
            if ``True``, overlay error bars

        width, height, title, config, xscale, yscale, colorscale, shapescale : ``None`` or JSON
            graphical directives to pass to Vega-Lite

        Returns
        -------
        :py:class:`Plotable <histbook.vega.Plotable>`
        """
        if any(isinstance(x, StackChannel) for x in self._chain):
            raise TypeError("only area and bar can be stacked")
        if error and any(isinstance(x, (BesideChannel, BelowChannel)) for x in self._chain):
            raise NotImplementedError("error bars are currently incompatible with splitting beside or below")
        return Plotable(self, StepChannel(self._asaxis(self._singleaxis(axis)), self._asaxis(profile), error, width, height, title, config, xscale, yscale, colorscale, shapescale))

    def area(self, axis=None, profile=None, error=False, width=None, height=None, title=None, config=None, xscale=None, yscale=None, colorscale=None, shapescale=None):
        """
        Display bins in ``axis`` (if not the only axis) as areas on the horizontal axis.

        Parameters
        ----------
        axis : ``None``, :py:class:`Axis <histbook.axis.Axis>`, algebraic expression (lambda or string), or index position (integer)
            the axis to overlay; if ``None`` *(default)*, use the only axis in this :py:class:`Hist <histbook.hist.Hist>`

        profile : ``None``, :py:class:`profile <histbook.axis.profile>`, algebraic expression (lambda or string) or index position (integer)
            if ``None`` *(default)*, display bin counts; otherwise, display profile means (and errors on the mean)

        error : bool
            if ``True``, overlay error bars

        width, height, title, config, xscale, yscale, colorscale, shapescale : ``None`` or JSON
            graphical directives to pass to Vega-Lite

        Returns
        -------
        :py:class:`Plotable <histbook.vega.Plotable>`
        """
        if error and any(isinstance(x, (BesideChannel, BelowChannel)) for x in self._chain):
            raise NotImplementedError("error bars are currently incompatible with splitting beside or below")
        return Plotable(self, AreaChannel(self._asaxis(self._singleaxis(axis)), self._asaxis(profile), error, width, height, title, config, xscale, yscale, colorscale, shapescale))

    def line(self, axis=None, profile=None, error=False, width=None, height=None, title=None, config=None, xscale=None, yscale=None, colorscale=None, shapescale=None):
        """
        Display bins in ``axis`` (if not the only axis) as lines on the horizontal axis.

        Parameters
        ----------
        axis : ``None``, :py:class:`Axis <histbook.axis.Axis>`, algebraic expression (lambda or string), or index position (integer)
            the axis to overlay; if ``None`` *(default)*, use the only axis in this :py:class:`Hist <histbook.hist.Hist>`

        profile : ``None``, :py:class:`profile <histbook.axis.profile>`, algebraic expression (lambda or string) or index position (integer)
            if ``None`` *(default)*, display bin counts; otherwise, display profile means (and errors on the mean)

        error : bool
            if ``True``, overlay error bars

        width, height, title, config, xscale, yscale, colorscale, shapescale : ``None`` or JSON
            graphical directives to pass to Vega-Lite

        Returns
        -------
        :py:class:`Plotable <histbook.vega.Plotable>`
        """
        if any(isinstance(x, StackChannel) for x in self._chain):
            raise TypeError("only area and bar can be stacked")
        if error and any(isinstance(x, (BesideChannel, BelowChannel)) for x in self._chain):
            raise NotImplementedError("error bars are currently incompatible with splitting beside or below")
        return Plotable(self, LineChannel(self._asaxis(self._singleaxis(axis)), self._asaxis(profile), error, width, height, title, config, xscale, yscale, colorscale, shapescale))

    def marker(self, axis=None, profile=None, error=True, width=None, height=None, title=None, config=None, xscale=None, yscale=None, colorscale=None, shapescale=None):
        """
        Display bins in ``axis`` (if not the only axis) as markers on the horizontal axis.

        Parameters
        ----------
        axis : ``None``, :py:class:`Axis <histbook.axis.Axis>`, algebraic expression (lambda or string), or index position (integer)
            the axis to overlay; if ``None`` *(default)*, use the only axis in this :py:class:`Hist <histbook.hist.Hist>`

        profile : ``None``, :py:class:`profile <histbook.axis.profile>`, algebraic expression (lambda or string) or index position (integer)
            if ``None`` *(default)*, display bin counts; otherwise, display profile means (and errors on the mean)

        error : bool
            if ``True``, overlay error bars

        width, height, title, config, xscale, yscale, colorscale, shapescale : ``None`` or JSON
            graphical directives to pass to Vega-Lite

        Returns
        -------
        :py:class:`Plotable <histbook.vega.Plotable>`
        """
        if any(isinstance(x, StackChannel) for x in self._chain):
            raise TypeError("only area and bar can be stacked")
        if error and any(isinstance(x, (BesideChannel, BelowChannel)) for x in self._chain):
            raise NotImplementedError("error bars are currently incompatible with splitting beside or below")
        return Plotable(self, MarkerChannel(self._asaxis(self._singleaxis(axis)), self._asaxis(profile), error, width, height, title, config, xscale, yscale, colorscale, shapescale))

class Plotable(object):
    """Mix-in for :py:class:`Channels <histbook.vega.Channel>` and :py:class:`Combinations <histbook.vega.Combination>` that can be plotted."""

    def __init__(self, source, item):
        if isinstance(source, PlottingChain):
            self._source = source._source
            self._chain = source._chain + (item,)
        else:
            self._source = source
            self._chain = (item,)

    def __repr__(self):
        return "".join(repr(x) for x in (self._source,) + self._chain)

    def __str__(self, indent="\n     ", paren=True):
        return ("(" if paren else "") + indent.join(repr(x) for x in (self._source,) + self._chain) + (")" if paren else "")

    @property
    def _last(self):
        return self._chain[-1]

    def _data(self, prefix, varname):
        error = self._last.error
        baseline = isinstance(self._last, (StepChannel, AreaChannel))
        if isinstance(self._last.axis, (histbook.axis.bin, histbook.axis.intbin, histbook.axis.split)):
            if isinstance(self._last, BarChannel):
                xtype = "ordinal"
            else:
                xtype = "quantitative"
        else:
            xtype = "nominal"

        profile = self._last.profile
        if profile is None:
            profiles = ()
        else:
            profiles = (profile,)

        projected = self._source.project(*(x.axis for x in self._chain))
        table = projected.table(*profiles, count=(profile is None), error=error, recarray=False)

        projectedorder = [x for x in projected.axis if not isinstance(x, histbook.axis.ProfileAxis)]
        lastj = projectedorder.index(self._last.axis)

        data = []
        domains = {}
        if error and baseline:
            shifts = {}

        def recurse(j, content, row, base):
            if j == len(projectedorder):
                if base:
                    row = row + ((0.0, 0.0) if error else (0.0,))
                else:
                    row = row + tuple(float(x) for x in content)

                datum = dict(prefix + tuple(zip([varname + str(i) for i in range(len(row))], row)))
                if datum[varname + str(j)] > 0 or (self._last.yscale != "log" and not (isinstance(self._last.yscale, dict) and self._last.yscale.get("type", None) == "log")):
                    data.append(datum)

            else:
                axis = projectedorder[j]
                if axis not in domains:
                    domains[axis] = set()
                domains[axis].update(axis.keys(content))

                if isinstance(axis, histbook.axis.intbin):
                    axis = axis.bin()

                for i, (n, x) in enumerate(axis.items(content)):
                    if isinstance(n, histbook.axis.Interval):
                        if j == lastj and xtype == "quantitative":
                            if numpy.isfinite(n.low) and numpy.isfinite(n.high):
                                low = n.low
                                if baseline and isinstance(axis, (histbook.axis.bin, histbook.axis.split)) and n.low == axis.low:
                                    recurse(j + 1, x, row + (n.low,), True)
                                    low += 1e-10*(axis.high - axis.low)

                                recurse(j + 1, x, row + (low,), base)

                                if error and baseline:
                                    shifts[low] = 0.5*(n.low + n.high)

                                if baseline and isinstance(axis, (histbook.axis.bin, histbook.axis.split)) and n.high == axis.high:
                                    recurse(j + 1, x, row + (n.high,), True)

                        else:
                            recurse(j + 1, x, row + (str(n),), base)

                    elif isinstance(n, (bool, numpy.bool_, numpy.bool)):
                        recurse(j + 1, x, row + ("pass" if n else "fail",), base)

                    else:
                        recurse(j + 1, x, row + (str(n),), base)

        recurse(0, table, (), False)
        if error and baseline:
            for x in data:
                x[varname + str(lastj) + "c"] = shifts.get(x[varname + str(lastj)], x[varname + str(lastj)])

        return projectedorder, data, domains

    def _vegalite(self, axis, domains, varname):
        error = self._last.error
        baseline = isinstance(self._last, (StepChannel, AreaChannel))
        if isinstance(self._last.axis, (histbook.axis.bin, histbook.axis.intbin, histbook.axis.split)):
            if isinstance(self._last, BarChannel):
                xtype = "ordinal"
            else:
                xtype = "quantitative"
        else:
            xtype = "nominal"

        if isinstance(self._last, BarChannel):
            mark = "bar"
        elif isinstance(self._last, StepChannel):
            if xtype == "nominal":
                mark = "bar"
            else:
                mark = {"type": "line", "interpolate": "step-before"}
        elif isinstance(self._last, AreaChannel):
            if xtype == "nominal":
                mark = "bar"
            else:
                mark = {"type": "area", "interpolate": "step-before"}
        elif isinstance(self._last, LineChannel):
            mark = {"type": "line"}
        elif isinstance(self._last, MarkerChannel):
            mark = {"type": "point"}
        else:
            raise AssertionError(self._last)

        xtitle = self._last.axis.expr
        if self._last.profile is None:
            ytitle = "entries per bin"
        else:
            ytitle = self._last.profile.expr

        transform = []
        def makeorder(i, var, values):
            if len(values) == 1:
                return "if(datum.{0} === {1}, {2}, {3})".format(var, repr(values[0]), i, i + 1)
            elif len(values) > 1:
                return "if(datum.{0} === {1}, {2}, {3})".format(var, repr(values[0]), i, makeorder(i + 1, var, values[1:]))
            else:
                raise AssertionError(values)

        encoding = {"x": {"field": varname + str(axis.index(self._last.axis)), "type": xtype, "scale": {"zero": False}, "axis": {"title": xtitle}},
                    "y": {"field": varname + str(len(axis)), "type": "quantitative", "axis": {"title": ytitle}}}
        for channel in self._chain[:-1]:
            if isinstance(channel, OverlayChannel):
                overlayorder = [str(x) for x in sorted(domains[channel.axis])]
                encoding["color"] = {"field": varname + str(axis.index(channel.axis)), "type": "nominal", "legend": {"title": channel.axis.expr}, "scale": {"domain": overlayorder}}

            elif isinstance(channel, StackChannel):
                if channel.order is None:
                    stackorder = [str(x) for x in sorted(domains[channel.axis])]
                else:
                    stackorder = channel.order
                encoding["color"] = {"field": varname + str(axis.index(channel.axis)), "type": "nominal", "legend": {"title": channel.axis.expr}, "scale": {"domain": list(reversed(stackorder))}}
                encoding["y"]["aggregate"] = "sum"
                encoding["order"] = {"field": "stackorder", "type": "nominal"}
                transform.append({"calculate": makeorder(0, varname + str(axis.index(channel.axis)), stackorder), "as": "stackorder"})

            elif isinstance(channel, BesideChannel):
                # FIXME: sorting doesn't work??? https://github.com/vega/vega-lite/issues/2176
                encoding["column"] = {"field": varname + str(axis.index(channel.axis)), "type": "nominal", "header": {"title": channel.axis.expr}}

            elif isinstance(channel, BelowChannel):
                # FIXME: sorting doesn't work??? https://github.com/vega/vega-lite/issues/2176
                encoding["row"] = {"field": varname + str(axis.index(channel.axis)), "type": "nominal", "header": {"title": channel.axis.expr}}

            else:
                raise AssertionError(channel)

        if self._last.xscale is not None and "x" in encoding:
            if isinstance(self._last.xscale, dict):
                encoding["x"]["scale"] = self._last.xscale
            else:
                encoding["x"]["scale"] = {"type": self._last.xscale}
        if self._last.yscale is not None and "y" in encoding:
            if isinstance(self._last.yscale, dict):
                encoding["y"]["scale"] = self._last.yscale
            else:
                encoding["y"]["scale"] = {"type": self._last.yscale}
        if self._last.colorscale is not None and "color" in encoding:
            if isinstance(self._last.colorscale, dict):
                encoding["color"]["scale"] = self._last.colorscale
            else:
                encoding["color"]["scale"] = {"type": self._last.colorscale}
        if self._last.shapescale is not None and "shape" in encoding:
            if isinstance(self._last.shapescale, dict):
                encoding["shape"]["scale"] = self._last.shapescale
            else:
                encoding["shape"]["scale"] = {"type": self._last.shapescale}

        if not error:
            return [mark], [encoding], [transform]

        else:
            if error and baseline:
                errx = varname + str(axis.index(self._last.axis)) + "c"
            else:
                errx = varname + str(axis.index(self._last.axis))

            encoding2 = {"x": {"field": errx, "type": "quantitative"},
                         "y": {"field": "error-down", "type": "quantitative"},
                         "y2": {"field": "error-up", "type": "quantitative"}}

            transform2 = [{"calculate": "datum.{0} - datum.{1}".format(varname + str(len(axis)), varname + str(len(axis) + 1)), "as": "error-down"},
                          {"calculate": "datum.{0} + datum.{1}".format(varname + str(len(axis)), varname + str(len(axis) + 1)), "as": "error-up"}]

            return [mark, "rule"], [encoding, encoding2], [transform, transform2]
        
    def vegalite(self):
        """Return the Vega-Lite JSON for this plot."""

        axis, data, domains = self._data((), "a")
        marks, encodings, transforms = self._vegalite(axis, domains, "a")

        if len(marks) == 1:
            return {"$schema": "https://vega.github.io/schema/vega-lite/v2.json",
                    "width": self._last.width,
                    "height": self._last.height,
                    "title": self._last.title,
                    "config": self._last.config,
                    "data": {"values": data},
                    "mark": marks[0],
                    "encoding": encodings[0],
                    "transform": transforms[0]}

        else:
            return {"$schema": "https://vega.github.io/schema/vega-lite/v2.json",
                    "width": self._last.width,
                    "height": self._last.height,
                    "title": self._last.title,
                    "config": self._last.config,
                    "data": {"values": data},
                    "layer": [{"mark": m, "encoding": e, "transform": t} for m, e, t in zip(marks, encodings, transforms)]}

    def to(self, fcn):
        """Call ``fcn`` on the Vega-Lite JSON for this plot."""
        return fcn(self.vegalite())

class Combination(object):
    """Abstract class for :py:class:`Plotables <histbook.vega.Plotable>` that have been combined as an overlay or side-by-side plots."""

    def __init__(self, *plotables):
        self._plotables = []
        for arg in plotables:           # first level: for varargs
            try:
                iter(arg)
            except:
                arg = [arg]
            for plotable in arg:        # second level: for iterators
                if not isinstance(plotable, Plotable):
                    raise TypeError("only Plotables can be combined with {0}".format(self.__class__.__name__))
                self._plotables.append(plotable)

    def __repr__(self):
        return "{0}({1})".format(self.__class__.__name__, ", ".join(repr(x) for x in self._plotables))

    def __str__(self, indent="\n    ", paren=False):
        return "{0}({1})".format(self.__class__.__name__, "".join(indent + x.__str__(indent + "    ", False) for x in self._plotables))

    def _varname(self, i):
        out = []
        i += 1
        while i > 0:
            out.append("abcdefghijklmnopqrstuvwxyz"[i % 26])
            i //= 26
        return "".join(reversed(out))

    def to(self, fcn):
        return fcn(self.vegalite())

class overlay(Combination):
    """:py:class:`Plotable <histbook.vega.Plotable>` overlaying two or more independently produced :py:class:`Plotables <histbook.vega.Plotable>`."""

    def vegalite(self):
        allaxis = []
        alldata = []
        alldomains = []
        for i, plotable in enumerate(self._plotables):
            varname = self._varname(i)
            axis, data, domains = plotable._data((("id", varname),), varname)
            allaxis.append(axis)
            alldata.extend(data)
            alldomains.append(domains)

        out = {"$schema": "https://vega.github.io/schema/vega-lite/v2.json",
               "data": {"values": alldata},
               "layer": []}

        for i, plotable in enumerate(self._plotables):
            varname = self._varname(i)
            marks, encodings, transforms = plotable._vegalite(allaxis[i], alldomains[i], varname)
            thislayer = [{"filter": {"field": "id", "equal": varname}}]
            
            if len(marks) == 1:
                out["layer"].append({"mark": marks[0],
                                     "encoding": encodings[0],
                                     "transform": transforms[0] + thislayer})

            else:
                for m, e, t in zip(marks, encodings, transforms):
                    out["layer"].append({"mark": m, "encoding": e, "transform": t + thislayer})

        return out

class beside(Combination):
    """:py:class:`Plotable <histbook.vega.Plotable>` displaying two or more independently produced :py:class:`Plotables <histbook.vega.Plotable>` beside each other horizontally."""

    def __init__(self, *plotables):
        super(beside, self).__init__(*plotables)
        if any(isinstance(x, BesideChannel) for x in self._plotables):
            raise TypeError("cannot place plots beside each other that are already split with beside (can do beside and below)")

    def vegalite(self):
        allaxis = []
        alldata = []
        alldomains = []
        for i, plotable in enumerate(self._plotables):
            varname = self._varname(i)
            axis, data, domains = plotable._data((("id", varname),), varname)
            allaxis.append(axis)
            alldata.extend(data)
            alldomains.append(domains)

        out = {"$schema": "https://vega.github.io/schema/vega-lite/v2.json",
               "data": {"values": alldata},
               "hconcat": []}

        for i, plotable in enumerate(self._plotables):
            varname = self._varname(i)
            marks, encodings, transforms = plotable._vegalite(allaxis[i], alldomains[i], varname)
            thislayer = [{"filter": {"field": "id", "equal": varname}}]

            if len(marks) == 1:
                out["hconcat"].append({"mark": marks[0], "encoding": encodings[0], "transform": transforms[0] + thislayer})

            else:
                out["hconcat"].append({"layer": [{"mark": m, "encoding": e, "transform": t} for m, e, t in zip(marks, encodings, transforms)]})

        return out

class below(Combination):
    """:py:class:`Plotable <histbook.vega.Plotable>` displaying two or more independently produced :py:class:`Plotables <histbook.vega.Plotable>` below each other vertically."""

    def __init__(self, *plotables):
        super(below, self).__init__(*plotables)
        if any(isinstance(x, BelowChannel) for x in self._plotables):
            raise TypeError("cannot place plots below each other that are already split with below (can do beside and below)")

    def vegalite(self):
        allaxis = []
        alldata = []
        alldomains = []
        for i, plotable in enumerate(self._plotables):
            varname = self._varname(i)
            axis, data, domains = plotable._data((("id", varname),), varname)
            allaxis.append(axis)
            alldata.extend(data)
            alldomains.append(domains)

        out = {"$schema": "https://vega.github.io/schema/vega-lite/v2.json",
               "data": {"values": alldata},
               "vconcat": []}

        for i, plotable in enumerate(self._plotables):
            varname = self._varname(i)
            marks, encodings, transforms = plotable._vegalite(allaxis[i], alldomains[i], varname)
            thislayer = [{"filter": {"field": "id", "equal": varname}}]

            if len(marks) == 1:
                out["vconcat"].append({"mark": marks[0], "encoding": encodings[0], "transform": transforms[0] + thislayer})

            else:
                out["vconcat"].append({"layer": [{"mark": m, "encoding": e, "transform": t} for m, e, t in zip(marks, encodings, transforms)]})

        return out
