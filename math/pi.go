package math

import (
	"math"
)

// from: https://golang.org/
func pi(n int) float64 {
	ch := make(chan float64)
	for k := 0; k <= n; k++ {
		go term(ch, float64(k))
	}
	f := 0.0
	for k := 0; k <= n; k++ {
		f += <-ch
	}
	return f
}

func term(ch chan float64, k float64) {
	ch <- 4 * math.Pow(-1, k) / (2*k + 1)
}


// another way to compute pi
func pi1(n int) float64 {
	ch := make(chan float64)
	for k := 1; k <= n; k++ {
		go term1(ch, float64(k))
	}
	f := 0.0
	for k := 1; k <= n; k++ {
		f += <-ch
	}
	return math.Sqrt(f)
}

func term1(ch chan float64, k float64) {
	ch <- 6 * 1/ math.Pow( k, 2)
}
