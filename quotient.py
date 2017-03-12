#! /usr/bin/env python3

import argparse
import itertools
import sys


def divide(dividend, divisor, radix=10):
    """ Convert to the specified base and divide.

    Parameters
    ----------
    dividend : int (or coercible to int)
        The numerator, `a`, in `a divided by b`.
    divisor : int (or coercible to int)
        The denominator, `b`, in `a divided by b`.
    radix : int, optional
        The numeric base in which to divide.  Default `10` for decimal.

    Yields
    ------
    A single character of the extended quotient.

    """
    dividend = int(dividend, radix)
    divisor = int(divisor, radix)
    while True:
        quotient, remainder = divmod(dividend, divisor)
        dividend = remainder * radix
        yield quotient

if __name__ == '__main__':
    parser = argparse.ArgumentParser(
        description='Calculate an extended quotient of two integers.')
    parser.add_argument('dividend', action='store', type=str,
                        help='a number to divide')
    parser.add_argument('divisor', action='store', type=str,
                        help='a number by which to divide')
    parser.add_argument('-r', '--radix', dest='radix', action='store',
                        type=int, default=10,
                        help='a numeric base in which to calculate')
    parser.add_argument('-s', '--scale', dest='scale', action='store',
                        type=int, default=None,
                        help='how long the fractional output should be')

    args = parser.parse_args()

    scale = args.scale
    quotient= divide(args.dividend, args.divisor, radix=args.radix)

    # get the part before the decimal point
    q = next(quotient)
    print(q, end='')

    # only print decimal point if scale is non-zero
    if scale != 0:
        print('.', end='')

    # now display the part after the decimal point
    for q in itertools.islice(quotient, scale):
        print(q, end='')
