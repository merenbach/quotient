package main

import (
	"flag"
	"fmt"
	"log"
)

// Radix holds the base in which division operations are performed.
const radix = 10

// Divide returns a generator that returns an endlessly-repeating quotient.
func Divide(a, b uint) func() uint {
	return func() uint {
		q, r := a/b, a%b
		a = radix * r
		return q
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
