"""
Functions and classes for parsing and executing PAC files.
"""
import warnings
from contextlib import contextmanager

import sys

from pypac.parser_functions import function_injections


ARBITRARY_HIGH_RECURSION_LIMIT = 10000


@contextmanager
def _temp_recursion_limit(limit):
    """Context manager to temporarily adjust Python's recursion limit."""
    prev_limit = sys.getrecursionlimit()
    sys.setrecursionlimit(limit)
    yield
    sys.setrecursionlimit(prev_limit)


class PACFile(object):
    """
    Represents a PAC file.

    JavaScript parsing and execution is handled by the `Js2Py`_ library.

    .. _Js2Py: https://github.com/PiotrDabkowski/Js2Py
    """

    def __init__(self, pac_js, recursion_limit=ARBITRARY_HIGH_RECURSION_LIMIT):
        """
        Load a PAC file from a given string of JavaScript.
        Errors during parsing and validation may raise a specialized exception.
        
        :param str pac_js: JavaScript that defines the FindProxyForURL() function.
        :param int recursion_limit: Python recursion limit when executing JavaScript.
            PAC files are often complex enough to need this to be higher than the interpreter default.
        :raises MalformedPacError: If the JavaScript could not be parsed, does not define FindProxyForURL(),
            or is otherwise invalid.
        :raises PyImportError: If the JavaScript tries to use Js2Py's `pyimport` keyword,
            which is not legitimate in the context of a PAC file.
        :raises PacComplexityError: If the JavaScript was complex enough that the
            Python recursion limit was hit during parsing.
        """
        self._recursion_limit = recursion_limit

        # Defer importing of js2py, as it greatly increases memory use (~120MB). See #20.
        import js2py

        if 'pyimport' in pac_js:
            raise PyimportError()
        # Disallow parsing of the unsafe 'pyimport' statement in Js2Py.
        js2py.disable_pyimport()
        try:
            with _temp_recursion_limit(self._recursion_limit):
                context = js2py.EvalJs(function_injections)
                context.execute(pac_js)
                # A test call to weed out errors like unimplemented functions.
                context.FindProxyForURL('/', '0.0.0.0')

            self._context = context
            self._func = context.FindProxyForURL
        except js2py.base.PyJsException:  # as e:
            raise MalformedPacError()  # from e
        except RuntimeError:
            # RecursionError in PY >= 3.5.
            raise PacComplexityError()

    def find_proxy_for_url(self, url, host):
        """
        Call ``FindProxyForURL()`` in the PAC file with the given arguments.

        :param str url: The full URL.
        :param str host: The URL's host.
        :return: Result of evaluating the ``FindProxyForURL()`` JavaScript function in the PAC file.
        :rtype: str
        """
        with _temp_recursion_limit(self._recursion_limit):
            return self._func(url, host)


class MalformedPacError(Exception):
    def __init__(self, msg=None):
        if not msg:
            msg = "Malformed PAC file"
        super(MalformedPacError, self).__init__(msg)


class PyimportError(MalformedPacError):
    def __init__(self):
        super(PyimportError, self).__init__("PAC file contains pyimport statement. "
                                            "Ensure that the source of your PAC file is trustworthy")


class PacComplexityError(RuntimeError):
    def __init__(self):
        super(PacComplexityError, self).__init__(
            "Maximum recursion depth exceeded while parsing PAC file. "
            "Raise it using sys.setrecursionlimit()")


def parse_pac_value(value, socks_scheme=None):
    """
    Parse the return value of ``FindProxyForURL()`` into a list.
    List elements will either be the string "DIRECT" or a proxy URL.

    For example, the result of parsing ``PROXY example.local:8080; DIRECT``
    is a list containing strings ``http://example.local:8080`` and ``DIRECT``.

    :param str value: Any value returned by ``FindProxyForURL()``.
    :param str socks_scheme: Scheme to assume for SOCKS proxies. ``socks5`` by default.
    :returns: Parsed output, with invalid elements ignored. Warnings are logged for invalid elements.
    :rtype: list[str]
    """
    config = []
    for element in value.split(';'):
        element = element.strip()
        if not element:
            continue
        try:
            config.append(proxy_url(element, socks_scheme))
        except ValueError as e:
            warnings.warn(str(e))
    return config


def proxy_url(value, socks_scheme=None):
    """
    Parse a single proxy config value from FindProxyForURL() into a more usable element.

    :param str value: Value to parse, e.g.: ``DIRECT``, ``PROXY example.local:8080``, or ``SOCKS example.local:8080``.
    :param str socks_scheme: Scheme to assume for SOCKS proxies. ``socks5`` by default.
    :returns: Parsed value, e.g.: ``DIRECT``, ``http://example.local:8080``, or ``socks5://example.local:8080``.
    :rtype: str
    :raises ValueError: If input value is invalid.
    """
    if value.upper() == 'DIRECT':
        return 'DIRECT'
    parts = value.split()

    if len(parts) == 2:
        keyword, proxy = parts[0].upper(), parts[1]
        if keyword == 'PROXY':
            return 'http://' + proxy
        if keyword == 'SOCKS':
            if not socks_scheme:
                socks_scheme = 'socks5'
            return '{0}://{1}'.format(socks_scheme, proxy)

    raise ValueError("Unrecognized proxy config value '{}'".format(value))
