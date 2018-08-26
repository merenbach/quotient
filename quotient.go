package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

// Radix holds the base in which division operations are performed.
const radix = 10

// DivMod returns the integer quotient and remainder of a fraction.
func DivMod(a, b int64) (q, r int64) {
	q, r = a/b, a%b
	return
}

// Divide divides one number by another out to an specified number of decimal places.
func Divide(w io.Writer, a, b int64, scale int, newline bool) {
	q, r := DivMod(a*radix, b*radix)
	fmt.Fprintf(w, "%d.", q)
	for i := 0; i < scale; i++ {
		q, r = DivMod(r*radix, b*radix)
		fmt.Fprintf(w, "%d", q)
	}
	if newline {
		fmt.Fprintln(w)
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

	Divide(os.Stdout, *a, *b, *scale, true)
}
