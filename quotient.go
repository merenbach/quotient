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

// Divide returns a generator that returns an endlessly-repeating quotient.
func Divide(a, b int64) func() int64 {
	var q int64
	return func() int64 {
		q, a = DivMod(a*radix, b*radix)
		return q
	}
}

func main() {
	a := flag.Int64("dividend", -1, "a number to divide")
	b := flag.Int64("divisor", -1, "a number by which to divide")
	scale := flag.Int("scale", 0, "the number of places to calculate after the decimal point")

	flag.Parse()

	if *a < 1 || *b < 1 || *scale < 1 {
		log.Fatal("Dividend, divisor, and scale must each be at least one.")
	}

	quotient := Divide(*a, *b)
	fmt.Printf("%d.", quotient())
	for i := 0; i < *scale; i++ {
		fmt.Printf("%d", quotient())
	}
	fmt.Println()
}
