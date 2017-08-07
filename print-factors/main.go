package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Comes from http://exercism.io/exercises/go/raindrops/readme
	fmt.Println("Raindrop speak for 28:")
	printFactors(28)
	fmt.Println("Raindrop speak for 30:")
	printFactors(30)
	fmt.Println("Raindrop speak for 34:")
	printFactors(34)
}

func printFactors(nr int64) {
	if nr < 1 {
		return
	}

	fs := make([]int64, 1)
	fs[0] = 1

	apf := func(p int64, e int) {
		n := len(fs)
		for i, pp := 0, p; i < e; i, pp = i+1, pp*p {
			for j := 0; j < n; j++ {
				fs = append(fs, fs[j]*pp)
			}
		}
	}

	e := 0

	for ; nr&1 == 0; e++ {
		nr >>= 1
	}

	apf(2, e)

	for d := int64(3); nr > 1; d += 2 {
		if d*d > nr {
			d = nr
		}
		for e = 0; nr%d == 0; e++ {
			nr /= d
		}
		if e > 0 {
			apf(d, e)
		}
	}

	out := strconv.FormatInt(fs[len(fs)-1], 10)
	sout := ""

	for _, f := range fs {
		if f == 3 {
			sout += "Pling"
		}
		if f == 5 {
			sout += "Plang"
		}
		if f == 7 {
			sout += "Plong"
		}
	}
	if sout != "" {
		fmt.Println(sout)
	} else {
		fmt.Println(out)
	}
}
