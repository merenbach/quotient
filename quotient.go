package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// Radix holds the base in which division operations are performed.
const radix = 10

// DivMod returns the integer quotient and remainder of a fraction.
func DivMod(a, b int64) (q, r int64) {
	q, r = a/b, a%b
	return
}

// Divide divides one number by another out to an specified number of decimal places.
func Divide(a, b, scale int64) {
	for i := int64(0); i < scale; i++ {
		q, r := DivMod(a, b)
		fmt.Printf("%d", q)
		if i == 0 {
			fmt.Print(".")
		}
		a = r * radix
	}
}

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: go run quotient.go DIVIDEND DIVISOR SCALE")
	}

	a, _ := strconv.ParseInt(os.Args[1], radix, 64)
	b, _ := strconv.ParseInt(os.Args[2], radix, 64)
	scale, _ := strconv.ParseInt(os.Args[3], radix, 64)
	Divide(a, b, scale)
}
