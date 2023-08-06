# -*- coding: utf-8 -*-

"""
Helpful utility functions.
"""


__all__ = ["no_value", "rel_path", "law_src_path", "law_home_path", "print_err", "abort", "colored",
           "uncolored", "query_choice", "multi_match", "is_lazy_iterable", "make_list", "flatten",
           "which", "map_verbose", "map_struct", "mask_struct", "tmp_file", "interruptable_popen",
           "create_hash", "copy_no_perm", "makedirs_perm", "user_owns_file", "iter_chunks",
           "human_bytes", "is_file_exists_error", "check_bool_flag", "ShorthandDict"]


import os
import sys
import types
import re
import math
import fnmatch
import tempfile
import subprocess
import signal
import hashlib
import shutil
import copy
import collections
import contextlib

import six


class NoValue(object):

    def __bool__(self):
        return False

    def __nonzero__(self):
        return False


#: Unique dummy value that evaluates to *False*.
no_value = NoValue()


def rel_path(anchor, *paths):
    """
    Returns a path made of framgment *paths* relativ to an *anchor* path. When *anchor* is a file,
    its absolute directory is used instead.
    """
    anchor = os.path.expandvars(os.path.expanduser(os.path.abspath(anchor)))
    if os.path.exists(anchor) and os.path.isfile(anchor):
        anchor = os.path.dirname(anchor)
    return os.path.normpath(os.path.join(anchor, *paths))


def law_src_path(*paths):
    """
    Returns the law installation directory, optionally joined with *paths*.
    """
    return rel_path(__file__, *paths)


def law_home_path(*paths):
    """
    Returns the law home directory (``$LAW_HOME``) that defaults to ``"$HOME/.law"``, optionally
    joined with *paths*.
    """
    home = os.getenv("LAW_HOME", "$HOME/.law")
    home = os.path.expandvars(os.path.expanduser(home))
    return os.path.normpath(os.path.join(home, *paths))


def print_err(*args, **kwargs):
    """ print_err(*args, flush=False)
    Same as *print*, but outputs to stderr. If *flush* is *True*, stderr is flushed after printing.
    """
    sys.stderr.write(" ".join(str(arg) for arg in args) + "\n")
    if kwargs.get("flush", False):
        sys.stderr.flush()


def abort(msg=None, exitcode=1):
    """
    Aborts the process (*sys.exit*) with an *exitcode*. If *msg* is not *None*, it is printed first
    to stdout if *exitcode* is 0 or *None*, and to stderr otherwise.
    """
    if msg is not None:
        if exitcode in (None, 0):
            print(msg)
        else:
            print_err(msg)
    sys.exit(exitcode)


colors = {
    "default": 39,
    "black": 30,
    "red": 31,
    "green": 32,
    "yellow": 33,
    "blue": 34,
    "magenta": 35,
    "cyan": 36,
    "light_gray": 37,
    "dark_gray": 90,
    "light_red": 91,
    "light_green": 92,
    "light_yellow": 93,
    "light_blue": 94,
    "light_magenta": 95,
    "light_cyan": 96,
    "white": 97,
}

backgrounds = {
    "default": 49,
    "black": 40,
    "red": 41,
    "green": 42,
    "yellow": 43,
    "blue": 44,
    "magenta": 45,
    "cyan": 46,
    "light_gray": 47,
    "dark_gray": 100,
    "light_red": 101,
    "light_green": 102,
    "light_yellow": 103,
    "light_blue": 104,
    "light_magenta": 105,
    "light_cyan": 106,
    "white": 107,
}

styles = {
    "default": 0,
    "bright": 1,
    "dim": 2,
    "underlined": 4,
    "blink": 5,
    "inverted": 7,
    "hidden": 8,
}

uncolor_cre = re.compile(r"\x1B\[[0-?]*[ -/]*[@-~]")


def colored(msg, color=None, background=None, style=None, force=False):
    """
    Return the colored version of a string *msg*. For *color*, *background* and *style* options, see
    https://misc.flogisoft.com/bash/tip_colors_and_formatting. Unless *force* is *True*, the *msg*
    string is returned unchanged in case the output is not a tty.
    """
    try:
        if not force and not os.isatty(sys.stdout.fileno()):
            return msg
    except:
        return msg

    color = colors.get(color, colors["default"])
    background = backgrounds.get(background, backgrounds["default"])

    if not isinstance(style, (tuple, list, set)):
        style = (style,)
    style = ";".join(str(styles.get(s, styles["default"])) for s in style)

    return "\033[{};{};{}m{}\033[0m".format(style, background, color, msg)


def uncolored(s):
    """
    Returns color codes from a string *s* and returns it.
    """
    return uncolor_cre.sub("", s)


def query_choice(msg, choices, default=None, descriptions=None, lower=True):
    """
    Interactively query a choice from the prompt until the input matches one of the *choices*. The
    prompt can be configured using *msg* and *descriptions*, which, if set, must have the same
    length as *choices*. When *default* is not *None* it must be one of the choices and is used when
    the input is empty. When *lower* is *True*, the input is compared to the choices in lower case.
    """
    choices = _choices = [str(c) for c in choices]
    if lower:
        _choices = [c.lower() for c in choices]

    if default is not None:
        if default not in choices:
            raise Exception("default must be one of the choices")

    hints = [(choice if choice != default else choice + "*") for choice in choices]
    if descriptions is not None:
        if len(descriptions) != len(choices):
            raise ValueError("length of descriptions must match length of choices")
        hints = ["{}({})".format(*tpl) for tpl in zip(hints, descriptions)]
    msg += " [{}] ".format(", ".join(hints))

    choice = None
    while choice not in _choices:
        if choice is not None:
            print("invalid choice: '{}'\n".format(choice))
        choice = six.moves.input(msg)
        if default is not None and choice == "":
            choice = default
        if lower:
            choice = choice.lower()

    return choice


def multi_match(name, patterns, mode=any, regex=False):
    """
    Compares *name* to multiple *patterns* and returns *True* in case of at least one match (*mode*
    = *any*, the default), or in case all patterns matched (*mode* = *all*). Otherwise, *False* is
    returned. When *regex* is *True*, *re.match* is used instead of *fnmatch.fnmatch*.
    """
    if not regex:
        return mode(fnmatch.fnmatch(name, pattern) for pattern in patterns)
    else:
        return mode(re.match(pattern, name) for pattern in patterns)


def is_lazy_iterable(obj):
    """
    Returns whether *obj* is iterable lazily, such as generators, range objects, etc.
    """
    return isinstance(obj,
        (types.GeneratorType, collections.MappingView, six.moves.range, enumerate))


def make_list(obj, cast=True):
    """
    Converts an object *obj* to a list and returns it. Objects of types *tuple* and *set* are
    converted if *cast* is *True*. Otherwise, and for all other types, *obj* is put in a new list.
    """
    if isinstance(obj, list):
        return list(obj)
    elif is_lazy_iterable(obj):
        return list(obj)
    elif isinstance(obj, (tuple, set)) and cast:
        return list(obj)
    else:
        return [obj]


def flatten(struct):
    """
    Flattens and returns a complex structured object *struct*.
    """
    if isinstance(struct, dict):
        return flatten(struct.values())
    elif isinstance(struct, (list, tuple, set)) or is_lazy_iterable(struct):
        objs = []
        for obj in struct:
            objs.extend(flatten(obj))
        return objs
    else:
        return [struct]


def which(prog):
    """
    Pythonic ``which`` implementation. Returns the path to an executable *prog* by searching in
    *PATH*, or *None* when it could not be found.
    """
    executable = lambda path: os.path.isfile(path) and os.access(path, os.X_OK)

    # prog can also be a path
    dirname, _ = os.path.split(prog)
    if dirname:
        if executable(prog):
            return prog
    elif "PATH" in os.environ:
        for search_path in os.environ["PATH"].split(os.pathsep):
            path = os.path.join(search_path.strip('"'), prog)
            if executable(path):
                return path

    return None


def map_verbose(func, seq, msg="{}", every=25, start=True, end=True, offset=0, callback=None):
    """
    Same as the built-in map function but prints a *msg* after chunks of size *every* iterations.
    When *start* (*stop*) is *True*, the *msg* is also printed after the first (last) iteration.
    Note that *msg* is supposed to be a template string that will be formatted with the current
    iteration number (starting at 0) plus *offset* using ``str.format``. When *callback* is
    callable, it is invoked instead of the default print method with the current iteration number
    (without *offset*) as the only argument. Example:

    .. code-block:: python

       func = lambda x: x ** 2
       msg = "computing square of {}"
       squares = map_verbose(func, range(7), msg, every=3)
       # ->
       # computing square of 0
       # computing square of 2
       # computing square of 5
       # computing square of 6
    """
    # default callable
    if not callable(callback):
        def callback(i):
            print(msg.format(i + offset))

    results = []
    for i, obj in enumerate(seq):
        results.append(func(obj))
        do_call = (start and i == 0) or (i + 1) % every == 0
        if do_call:
            callback(i)
    else:
        if end and results and not do_call:
            callback(i)

    return results


def map_struct(func, struct, cls=None, map_dict=True, map_list=True, map_tuple=False,
        map_set=False):
    """
    Applies a function *func* to each value of a complex structured object *struct* and returns the
    output in the same structure. Example:

    .. code-block:: python

       struct = {"foo": [123, 456], "bar": [{"1": 1}, {"2": 2}]}
       def times_two(i):
           return i * 2

       map_struct(struct, times_two)
       # -> {"foo": [246, 912], "bar": [{"1": 2}, {"2": 4}]}

    When *cls* is not *None*, it exclusively defines the class of objects that *func* is applied on.
    All other objects are unchanged. *map_dict*, *map_list*, *map_tuple* and *map_set* configure if
    objects of the respective types are traversed or mapped. The can be booleans or integer values
    that define the depth of that setting in the struct.
    """
    # interpret generators and views as lists
    if is_lazy_iterable(struct):
        struct = list(struct)

    valid_types = tuple()
    if map_dict:
        valid_types += (dict,)
        if isinstance(map_dict, int) and not isinstance(map_dict, bool):
            map_dict -= 1
    if map_list:
        valid_types += (list,)
        if isinstance(map_list, int) and not isinstance(map_list, bool):
            map_list -= 1
    if map_tuple:
        valid_types += (tuple,)
        if isinstance(map_tuple, int) and not isinstance(map_tuple, bool):
            map_tuple -= 1
    if map_set:
        valid_types += (set,)
        if isinstance(map_set, int) and not isinstance(map_set, bool):
            map_set -= 1

    # is an instance of cls?
    if cls is not None and isinstance(struct, cls):
        return func(struct)

    # traverse?
    elif isinstance(struct, valid_types):
        # create a new struct, treat tuples as lists for itertative item appending
        new_struct = struct.__class__() if not isinstance(struct, tuple) else []

        # create type-dependent generator and addition callback
        if isinstance(struct, (list, tuple)):
            gen = enumerate(struct)
            add = lambda _, value: new_struct.append(value)
        elif isinstance(struct, set):
            gen = enumerate(struct)
            add = lambda _, value: new_struct.add(value)
        else:  # dict
            gen = six.iteritems(struct)
            add = lambda key, value: new_struct.__setitem__(key, value)

        # recursively fill the new struct
        for key, value in gen:
            value = map_struct(func, value, cls=cls, map_dict=map_dict, map_list=map_list,
                map_tuple=map_tuple, map_set=map_set)
            add(key, value)

        # convert tuples
        if isinstance(struct, tuple):
            new_struct = struct.__class__(new_struct)

        return new_struct

    # when cls is set, just return
    elif cls is not None:
        return struct

    # apply func
    else:
        return func(struct)


def mask_struct(mask, struct, replace=no_value):
    """
    Masks a complex structured object *struct* with a *mask* and returns the remaining values. When
    *replace* is set, masked values are replaced with that value instead of being removed. The
    *mask* can have a complex structure as well. Examples:

    .. code-block:: python

        struct = {"a": [1, 2], "b": [3, ["foo", "bar"]]}

        # simple example
        mask_struct({"a": [False, True], "b": False}, struct)
        # => {"a": [2]}

        # omitting mask information results in keeping values
        mask_struct({"a": [False, True]}, struct)
        # => {"a": [2], "b": [3, ["foo", "bar"]]}
    """
    # interpret lazy iterables lists
    if is_lazy_iterable(struct):
        struct = list(struct)

    # when mask is a bool, or struct is not a dict or sequence, apply the mask immediately
    if isinstance(mask, bool) or not isinstance(struct, (list, tuple, dict)):
        return struct if mask else replace

    # check list and tuple types
    elif isinstance(struct, (list, tuple)) and isinstance(mask, (list, tuple)):
        new_struct = []
        for i, val in enumerate(struct):
            if i >= len(mask):
                new_struct.append(val)
            else:
                repl = replace
                if isinstance(replace, (list, tuple)) and len(replace) > i:
                    repl = replace[i]
                val = mask_struct(mask[i], val, replace=repl)
                if val != no_value:
                    new_struct.append(val)

        return struct.__class__(new_struct) if new_struct else replace

    # check dict types
    elif isinstance(struct, dict) and isinstance(mask, dict):
        new_struct = struct.__class__()
        for key, val in six.iteritems(struct):
            if key not in mask:
                new_struct[key] = val
            else:
                repl = replace
                if isinstance(replace, dict) and key in replace:
                    repl = replace[key]
                val = mask_struct(mask[key], val, replace=repl)
                if val != no_value:
                    new_struct[key] = val
        return new_struct or replace

    # when this point is reached, mask and struct have incompatible types
    raise TypeError("mask and struct must have the same type, got '{}' and '{}'".format(type(mask),
            type(struct)))


@contextlib.contextmanager
def tmp_file(*args, **kwargs):
    """
    Context manager that generates a temporary file, yields the file descriptor number and temporary
    path, and eventually removes the files. All *args* and *kwargs* are passed to
    :py:meth:`tempfile.mkstemp`.
    """
    fileno, path = tempfile.mkstemp(*args, **kwargs)

    # create the file
    with open(path, "w") as f:
        f.write("")

    # yield it
    try:
        yield fileno, path
    finally:
        if os.path.exists(path):
            os.remove(path)


def interruptable_popen(*args, **kwargs):
    """
    Shorthand to :py:class:`Popen` followed by :py:meth:`Popen.communicate`. All *args* and *kwargs*
    are forwatded to the :py:class:`Popen` constructor. The return code, standard output and
    standard error are returned in a tuple. The call :py:meth:`Popen.communicate` is interruptable
    by the user.
    """
    kwargs["preexec_fn"] = os.setsid

    p = subprocess.Popen(*args, **kwargs)

    try:
        out, err = p.communicate()
    except KeyboardInterrupt:
        os.killpg(os.getpgid(p.pid), signal.SIGTERM)
        raise

    if six.PY3:
        if out is not None:
            out = out.decode("utf-8")
        if err is not None:
            err = err.decode("utf-8")

    return p.returncode, out, err


def create_hash(inp, l=10, algo="sha256"):
    """
    Takes an input *inp* and creates a hash based on an algorithm *algo*. For valid algorithms, see
    python's hashlib. *l* corresponds to the maximum length of the returned hash. Internally, the
    string representation of *inp* is used.
    """
    return getattr(hashlib, algo)(six.b(str(inp))).hexdigest()[:l]


def copy_no_perm(src, dst):
    """
    Copies a file from *src* to *dst* including meta data except for permission bits.
    """
    shutil.copy(src, dst)
    perm = os.stat(dst).st_mode
    shutil.copystat(src, dst)
    os.chmod(dst, perm)


def makedirs_perm(path, perm=None):
    """
    Recursively creates directory up to *path*. If *perm* is set, the permissions of all newly
    created directories are set to its value.
    """
    if not os.path.exists(path):
        if perm is None:
            os.makedirs(path)
        else:
            umask = os.umask(0)
            try:
                os.makedirs(path, perm)
            finally:
                os.umask(umask)


def user_owns_file(path, uid=None):
    """
    Returns whether a file located at *path* is owned by the user with *uid*. When *uid* is *None*,
    the user id of the current process is used.
    """
    if uid is None:
        uid = os.getuid()
    path = os.path.expandvars(os.path.expanduser(path))
    return os.stat(path).st_uid == uid


def iter_chunks(l, size):
    """
    Returns a generator containing chunks of *size* of a list, integer or generator *l*. A *size*
    smaller than 1 results in no chunking at all.
    """
    if isinstance(l, six.integer_types):
        l = six.moves.range(l)

    if is_lazy_iterable(l):
        if size < 1:
            yield list(l)
        else:
            chunk = []
            for elem in l:
                if len(chunk) < size:
                    chunk.append(elem)
                else:
                    yield chunk
                    chunk = []
            else:
                if chunk:
                    yield chunk

    else:
        if size < 1:
            yield l
        else:
            for i in six.moves.range(0, len(l), size):
                yield l[i:i + size]


byte_units = ["bytes", "kB", "MB", "GB", "TB", "PB", "EB"]


def human_bytes(n, unit=None):
    """
    Takes a number of bytes *n*, assigns the best matching unit and returns the respective number
    and unit string in a tuple. When *unit* is set, that unit is used. Example:

    .. code-block:: python

        human_bytes(3407872)
        # -> (3.25, "MB")

        human_bytes(3407872, "kB")
        # -> (3328.0, "kB")
    """
    if n == 0:
        idx = 0
    elif unit:
        idx = byte_units.index(unit)
    else:
        idx = int(math.floor(math.log(abs(n), 1024)))
        idx = min(idx, len(byte_units))
    return n / 1024. ** idx, byte_units[idx]


def is_file_exists_error(e):
    """
    Returns whether the exception *e* was raised due to an already existing file or directory.
    """
    if six.PY3:
        return isinstance(e, FileExistsError)  # noqa: F821
    else:
        return isinstance(e, OSError) and e.errno == 17


def check_bool_flag(s):
    """
    Takes a string flag *s* and returns whether it evaluates to *True* (values ``"1"``, ``"true"``
    and ``"yes"``, case-insensitive) or *False* (any other value). When *s* is not a string, *s* is
    returned unchanged.
    """
    return s.lower() in ("1", "yes", "true") if isinstance(s, six.string_types) else s


class ShorthandDict(collections.OrderedDict):
    """
    Subclass of *OrderedDict* that implements ``__getattr__`` and ``__setattr__`` for a configurable
    list of attributes. Example:

    .. code-block:: python

        MyDict(ShorthandDict):

            attributes = {"foo": 1, "bar": 2}

        d = MyDict(foo=9)

        print(d.foo)
        # => 9

        print(d.bar)
        # => 2

        d.foo = 3
        print(d.foo)
        # => 3

    .. py:classattribute: attributes
       type: dict

       Mapping of attribute names to default values. ``__getattr__`` and ``__setattr__`` support is
       provided for these attributes.
    """

    attributes = {}

    def __init__(self, **kwargs):
        super(ShorthandDict, self).__init__()

        for attr, default in six.iteritems(self.attributes):
            self[attr] = kwargs.pop(attr, copy.deepcopy(default))

        self.update(kwargs)

    def copy(self):
        """"""
        kwargs = {key: copy.deepcopy(value) for key, value in six.iteritems(self)}
        return self.__class__(**kwargs)

    def __getattr__(self, attr):
        if attr in self.attributes:
            return self[attr]
        else:
            return super(ShorthandDict, self).__getattr__(attr)

    def __setattr__(self, attr, value):
        if attr in self.attributes:
            self[attr] = value
        else:
            super(ShorthandDict, self).__setattr__(attr, value)
