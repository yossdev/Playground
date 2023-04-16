package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	var collections []uint64
	const x = 500_000
	const n int = 94 // close to max u64
	i := 0
	for i < x {
		res := fib(n)
		collections = append(collections, res)
		i++
	}
	fmt.Printf("Last collections: %d\n", collections[x-1])
	fmt.Printf("Time took: %s, n: %d\n", time.Since(start), n)
}

func fib(n int) uint64 {
	if n <= 1 {
		return uint64(n)
	}
	var f = make([]uint64, n)
	f[0] = 0
	f[1] = 1
	for i := 2; i < n; i++ {
		prev := i - 2
		current := i - 1
		f[i] = f[prev] + f[current]
	}
	return f[n-1]
}
