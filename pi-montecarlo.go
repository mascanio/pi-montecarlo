package main

import (
	"fmt"
	"math/rand"
	"time"
)

func pi[T float32 | float64](n int, out chan T, gen func() T) {
	var inside int = 0
	for i := 0; i < n; i++ {
		x := gen()
		y := gen()
		if x*x+y*y <= 1 {
			inside++
		}
	}
	out <- 4 * T(inside) / T(n)
}

func runAll[T float32 | float64](nItersThread, nGo int, gen func() T) T {
	ch := make(chan T, nGo)

	for i := 0; i < nGo; i++ {
		go pi(nItersThread, ch, gen)
	}

	var value T
	for i := 0; i < nGo; i++ {
		value += <-ch / T(nGo)
	}
	return value
}

func runAll32(nItersThread, nGo int) float32 {
	return runAll(nItersThread, nGo, rand.Float32)
}

func runAll64(nItersThread, nGo int) float64 {
	return runAll(nItersThread, nGo, rand.Float64)
}

func main() {
	threads := [...]int{4, 8, 16}
	N := 5
	nIters := 1 << 29
	for _, nGo := range threads {
		nItersThread := nIters / nGo
		start := time.Now()
		var value32 float32
		for i := 0; i < N; i++ {
			value32 += runAll32(nItersThread, nGo)
		}
		fmt.Printf("%v threads\t %v\n", nGo, time.Since(start))
	}
	for _, nGo := range threads {
		nItersThread := nIters / nGo
		start := time.Now()
		var value64 float64
		for i := 0; i < N; i++ {
			value64 += runAll64(nItersThread, nGo)
		}
		fmt.Printf("%v threads\t %v\n", nGo, time.Since(start))
	}
}
