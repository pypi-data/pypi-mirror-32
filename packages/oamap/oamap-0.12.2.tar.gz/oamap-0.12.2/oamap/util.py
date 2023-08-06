#!/usr/bin/env python

# Copyright (c) 2017, DIANA-HEP
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

import ast
import math
import sys
import types

import numpy

if sys.version_info[0] > 2:
    basestring = str
    unicode = str
    def MethodType(function, instance, cls):
        if instance is None:
            return function
        else:
            return types.MethodType(function, instance)
else:
    MethodType = types.MethodType

try:
    from collections import OrderedDict
except ImportError:
    # simple OrderedDict implementation for Python 2.6
    class OrderedDict(dict):
        def __init__(self, items=(), **kwds):
            items = list(items)
            self._order = [k for k, v in items] + [k for k, v in kwds.items()]
            super(OrderedDict, self).__init__(items)
        def keys(self):
            return self._order
        def values(self):
            return [self[k] for k in self._order]
        def items(self):
            return [(k, self[k]) for k in self._order]
        def __setitem__(self, name, value):
            if name not in self._order:
                self._order.append(name)
            super(OrderedDict, self).__setitem__(name, value)
        def __delitem__(self, name):
            if name in self._order:
                self._order.remove(name)
            super(OrderedDict, self).__delitem__(name)
        def __repr__(self):
            return "OrderedDict([{0}])".format(", ".join("({0}, {1})".format(repr(k), repr(v)) for k, v in self.items()))

try:
    from UserDict import DictMixin as MutableMapping
except ImportError:
    from collections import MutableMapping

try:
    from importlib import import_module
except ImportError:
    def import_module(modulename):
        module = __import__(modulename)
        for name in modulename.split(".")[1:]:
            module = module.__dict__[name]
        return module

def slice2sss(index, length):
    step = 1 if index.step is None else index.step

    if step == 0:
        raise ValueError("slice step cannot be zero")

    elif step > 0:
        if index.start is None:
            start = 0                                         # in-range
        elif index.start >= 0:
            start = min(index.start, length)                  # in-range or length
        else:
            start = max(0, index.start + length)              # in-range

        if index.stop is None:
            stop = length                                     # length
        elif index.stop >= 0:
            stop = max(start, min(length, index.stop))        # in-range or length
        else:
            stop = max(start, index.stop + length)            # in-range or length

    else:
        if index.start is None:
            start = length - 1                                # in-range
        elif index.start >= 0:
            start = min(index.start, length - 1)              # in-range
        else:
            start = max(index.start + length, -1)             # in-range or -1

        if index.stop is None:
            stop = -1                                         # -1
        elif index.stop >= 0:
            stop = min(start, index.stop)                     # in-range or -1
        else:
            stop = min(start, max(-1, index.stop + length))   # in-range or -1

    return start, stop, step

def json2python(value):
    def recurse(value):
        if isinstance(value, dict) and len(value) == 2 and set(value.keys()) == set(["real", "imag"]) and all(isinstance(x, (int, float)) for x in value.values()):
            return value["real"] + value["imag"]*1j
        elif value == "inf":
            return float("inf")
        elif value == "-inf":
            return float("-inf")
        elif value == "nan":
            return float("nan")
        elif isinstance(value, list):
            return [recurse(x) for x in value]
        elif isinstance(value, dict):
            return dict((n, recurse(x)) for n, x in value.items())
        else:
            return value
    return recurse(value)

def python2json(value, allowlinks=False):
    def recurse(value, memo):
        if id(value) in memo:
            if allowlinks:
                return memo[id(value)]
            else:
                raise TypeError("cross-linking within an object is not allowed")

        if value is None:
            memo[id(value)] = None

        elif isinstance(value, (numbers.Integral, numpy.integer)):
            memo[id(value)] = int(value)

        elif isinstance(value, (numbers.Real, numpy.floating)):
            if math.isnan(value):
                memo[id(value)] = "nan"
            elif math.isinf(value) and value > 0:
                memo[id(value)] = "inf"
            elif math.isinf(value):
                memo[id(value)] = "-inf"
            else:
                memo[id(value)] = float(value)

        elif isinstance(value, (numbers.Complex, numpy.complex)):
            memo[id(value)] = {"real": float(value.real), "imag": float(value.imag)}

        elif isinstance(value, basestring):
            memo[id(value)] = value

        elif isinstance(value, dict):
            memo[id(value)] = {}
            for n, x in value.items():
                if not isinstance(n, basestring):
                    raise TypeError("dict keys for JSON must be strings")
                memo[id(value)][n] = recurse(x, memo)

        else:
            memo[id(value)] = []
            for x in value:
                memo[id(value)].append(recurse(x, memo))

        return memo[id(value)]

    return recurse(value, {})

def python2hashable(value):
    def recurse(value):
        if isinstance(value, dict):
            return tuple((n, recurse(value[n])) for n in sorted(value))
        elif isinstance(value, list):
            return tuple(recurse(x) for x in value)
        else:
            return value
    return recurse(python2json(value))

def varname(avoid, trial=None):
    while trial is None or trial in avoid:
        trial = "v" + str(len(avoid))
    avoid.add(trial)
    return trial

def paramtypes(args):
    try:
        import numba as nb
    except ImportError:
        return None
    else:
        return tuple(nb.typeof(x) for x in args)

def doexec(module, env):
    exec(module, env)

def stringfcn(fcn):
    if isinstance(fcn, basestring):
        parsed = ast.parse(fcn).body
        if isinstance(parsed[-1], ast.Expr):
            parsed[-1] = ast.Return(parsed[-1].value)
            parsed[-1].lineno = parsed[-1].value.lineno
            parsed[-1].col_offset = parsed[-1].value.col_offset

        env = dict(math.__dict__)
        env.update(globals())

        free = set()
        defined = set(["None", "False", "True"])
        defined.update(env)
        def recurse(node):
            if isinstance(node, ast.Name):
                if isinstance(node.ctx, ast.Store):
                    defined.add(node.id)
                elif isinstance(node.ctx, ast.Load) and node.id not in defined:
                    free.add(node.id)
            elif isinstance(node, ast.AST):
                for n in node._fields:
                    recurse(getattr(node, n))
            elif isinstance(node, list):
                for x in node:
                    recurse(x)
        recurse(parsed)

        avoid = free.union(defined)
        fcnname = varname(avoid, "fcn")

        module = ast.parse("""
def {fcn}({params}):
    REPLACEME
""".format(fcn=fcnname, params=",".join(free)))
        module.body[0].body = parsed
        module = compile(module, "<fcn string>", "exec")

        doexec(module, env)
        fcn = env[fcnname]

    return fcn

def trycompile(fcn, paramtypes=None, numba=True):
    fcn = stringfcn(fcn)

    if numba is None or numba is False:
        return fcn

    try:
        import numba as nb
    except ImportError:
        return fcn

    if numba is True:
        numbaopts = {}
    else:
        numbaopts = numba

    if isinstance(fcn, nb.dispatcher.Dispatcher):
        fcn = fcn.py_fcn

    if paramtypes is None:
        return nb.jit(**numbaopts)(fcn)
    else:
        return nb.jit(paramtypes, **numbaopts)(fcn)

def returntype(fcn, paramtypes):
    try:
        import numba as nb
    except ImportError:
        return None

    if isinstance(fcn, nb.dispatcher.Dispatcher):
        overload = fcn.overloads.get(paramtypes, None)
        if overload is None:
            return None
        else:
            return overload.signature.return_type
