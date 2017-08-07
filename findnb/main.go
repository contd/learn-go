package main

import (
	"fmt"
	"math"
)

func FindNb(m int) int {
	// Just go backwards until running total is equal or
	// return if it becomes > before being equal
	rtot := 1
	n := 2
	for m >= rtot {
		if rtot == m {
			return n - 1
		}
		N := float64(n)
		new := math.Pow(N, 3)
		rtot = rtot + int(new)
		n = n + 1
	}
	return -1
}

func main() {
	ans := FindNb(1071225)
	fmt.Printf("%d\n", ans)
}
