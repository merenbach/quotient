package main

import (
	"flag"
	"fmt"
	"log"
)

// Radix holds the base in which division operations are performed.
const radix = 10

// DivMod returns the integer quotient and remainder of a fraction.
func DivMod(a, b int64) (q, r int64) {
	q, r = a/b, a%b
	return
}

// Divide divides one number by another out to an specified number of decimal places.
func Divide(a, b int64, scale int) {
	for i := 0; i <= scale; i++ {
		q, a_ := DivMod(a*radix, b*radix)
		a = a_
		fmt.Printf("%d", q)
		if i == 0 && scale > 0 {
			fmt.Print(".")
		}
		// a = r * radix
	}
}

func main() {
	a := flag.Int64("dividend", -1, "a number to divide")
	b := flag.Int64("divisor", -1, "a number by which to divide")
	scale := flag.Int("scale", 0, "the number of places to calculate after the decimal point")

	flag.Parse()

	if *a < 0 || *b < 0 || *scale < 0 {
		log.Fatal("Dividend, divisor, and scale must be non-negative.")
	}

	Divide(*a, *b, *scale)
}
