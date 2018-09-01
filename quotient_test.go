package main

import (
	"testing"
)

func TestDivide(t *testing.T) {
	tables := []struct {
		a uint
		b uint
		c []uint
	}{
		{1, 7, []uint{0, 1, 4, 2, 8, 5, 7, 1, 4, 2, 8, 5, 7, 1}},
		{10, 7, []uint{1, 4, 2, 8, 5, 7, 1, 4, 2, 8, 5, 7, 1}},
		{1, 19, []uint{0, 0, 5, 2, 6, 3, 1, 5, 7, 8, 9, 4, 7, 3, 6, 8, 4, 2, 1, 0, 5, 2, 6, 3, 1, 5, 7, 8, 9, 4, 7, 3, 6, 8, 4, 2, 1, 0}},
	}
	for _, table := range tables {
		f := Divide(table.a, table.b)
		for i, expected := range table.c {
			g := f()
			if expected != g {
				t.Fatalf("Dividing %d by %d, expected value at position %d to be %d; got %d instead", table.a, table.b, i, expected, g)
			}
		}
	}
}
