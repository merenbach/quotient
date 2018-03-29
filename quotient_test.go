package main

import (
	"fmt"
	"testing"
)

func TestDivMod(t *testing.T) {
	tables := []struct {
		a int64
		b int64
		q int64
		r int64
	}{
		{1, 7, 0, 1},
		{10, 7, 1, 3},
		{17, 43, 0, 17},
		{43, 17, 2, 9},
	}
	for _, table := range tables {
		q, r := DivMod(table.a, table.b)
		fmt.Println("q and r = ", q, r)
		if q != table.q || r != table.r {
			t.Fatalf("Expected %d/%d=%dr%d; got %dr%d", table.a, table.b, table.q, table.r, q, r)
		}
	}
}
