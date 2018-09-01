package main

import (
	"flag"
	"fmt"
	"log"
)

// Radix holds the base in which division operations are performed.
const radix = 10

// // DivMod returns the integer quotient and remainder of a fraction.
// func DivMod(a, b int64) (q, r int64) {
// 	q, r = a/b, a%b
// 	return
// }

// Divide returns a generator that returns an endlessly-repeating quotient.
func Divide(a, b uint) func() uint {
	// var q int64
	return func() uint {
		// q, a = DivMod(a*radix, b*radix)
		a = radix * (a % b)
		return a / b
	}
}

func main() {
	a := flag.Uint("dividend", 0, "a number to divide")
	b := flag.Uint("divisor", 0, "a number by which to divide")
	scale := flag.Int("scale", 0, "the number of places to calculate after the decimal point")

	flag.Parse()

	if *a < 1 || *b < 1 || *scale < 1 {
		log.Fatal("The dividend, divisor, and scale must each be at least one.")
	}

	quotient := Divide(*a, *b)
	fmt.Printf("%d.", quotient())
	for i := 0; i < *scale; i++ {
		fmt.Printf("%d", quotient())
	}
	fmt.Println()
}
