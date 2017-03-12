#! /usr/bin/env python3

import itertools
import quotient
import unittest


class QuotientTest(unittest.TestCase):

    def join_and_test(self, i, s):
        sliced = itertools.islice(i, len(s))
        self.assertEqual(''.join(sliced), s)

    def test_divide(self):
        q = quotient.divide(10, 2)
        self.join_and_test(q, '5.00000')

        q = quotient.divide(10, 7)
        self.join_and_test(q, '1.428571428571428')

        q = quotient.divide(10, 0)
        with self.assertRaises(ZeroDivisionError):
            next(q)

    def test_negative_dividend(self):
        q = quotient.divide(-10, 2, scale=5)
        self.join_and_test(q, '-5.00000')

        q = quotient.divide(-10, 7, scale=15)
        self.join_and_test(q, '-2.571428571428571')

    def test_negative_divisor(self):
        q = quotient.divide(10, -2, scale=5)
        self.join_and_test(q, '-5.00000')

        q = quotient.divide(10, -7, scale=15)
        self.join_and_test(q, '-2.571428571428571')

    def test_negative_dividend_and_divisor(self):
        q = quotient.divide(-10, -2, scale=5)
        self.join_and_test(q, '5.00000')

        q = quotient.divide(-10, -7, scale=15)
        self.join_and_test(q, '1.428571428571428')


if __name__ == '__main__':
    unittest.main()
