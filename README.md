# Quotient

Repeating decimal calculator written in Python

## How do I use it?

Just provide a dividend and a divisor.  For example, to divide 100 by 17:

    ./quotient 100 17

You can cancel by pressing `^C` (control-C) on your keyboard.  If you only want the first 40 digits, either of the following should suffice:

    ./quotient -s 40 100 17
    ./quotient 100 17 -s 40

(The `-s` argument can come before or after the dividend and divisor.)

## Does it do any other tricks?

You can also divide in other bases.  More info will be posted soon.

