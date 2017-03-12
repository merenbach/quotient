#! /usr/bin/env python3

import argparse
import itertools
import locale
import sys


def divide_(dividend, divisor):
    """ Convert to the specified base and divide.

    Parameters
    ----------
    dividend : int
        The numerator, `a`, in `a divided by b`.
    divisor : int
        The denominator, `b`, in `a divided by b`.

    Yields
    ------
    out : int
        The whole number portion of the quotient (first yield only), followed
        by successive post-decimal-point volues.

    Notes
    -----
    The first result yielded is the sole whole-number portion of the quotient.

    """
    while True:
        quotient, remainder = divmod(dividend, divisor)
        dividend = remainder * 10
        yield quotient

def divide(dividend, divisor, scale=None):
    """ Convert to the specified base and divide.

    Parameters
    ----------
    dividend : int (or coercible to int)
        The numerator, `a`, in `a divided by b`.
    divisor : int (or coercible to int)
        The denominator, `b`, in `a divided by b`.
    scale : int, optional
        An integer representing the number of digits after the decimal point.
        Default `None`, which leads to infinite output (until interrupted).
        If this is set to zero or less, neither a decimal point nor anything
        after will be output.

    Yields
    ------
    out : str
        A string representation of the repeating decimal.

    """
    quotient = divide_(int(dividend), int(divisor))

    # get the part before the decimal point
    yield str(next(quotient))

    # only print decimal point if scale is `None` or non-zero
    if scale != 0:
        yield locale.nl_langinfo(locale.RADIXCHAR)
        for q in itertools.islice(quotient, scale):
            yield str(q)

if __name__ == '__main__':
    parser = argparse.ArgumentParser(
        description='Calculate an extended quotient of two integers.')
    parser.add_argument('dividend', action='store', type=str,
                        help='a number to divide')
    parser.add_argument('divisor', action='store', type=str,
                        help='a number by which to divide')
    parser.add_argument('-s', '--scale', dest='scale', action='store',
                        type=int, default=None,
                        help='how long the fractional output should be')

    args = parser.parse_args()

    for q in divide(args.dividend, args.divisor, args.scale):
        print(q, end='')
