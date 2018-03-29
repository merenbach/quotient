package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Radix holds the base in which division operations are performed.
const radix = 10

// DivMod returns the integer quotient and remainder of a fraction.
func DivMod(a, b int64) (q, r int64) {
	q, r = a/b, a%b
	return
}

// Divide divides one number by another out to an infinite number of decimal places.
func Divide(a, b int64) {
	for i := 0; ; i++ {
		q, r := DivMod(a, b)
		fmt.Printf("%d", q)
		if i == 0 {
			fmt.Print(".")
		}
		a = r * radix
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please supply a fraction to calculate.")
	}

	args := strings.Split(os.Args[1], "/")
	a, _ := strconv.ParseInt(args[0], radix, 64)
	b, _ := strconv.ParseInt(args[1], radix, 64)
	Divide(a, b)
}
