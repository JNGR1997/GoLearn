package main

import "fmt"

func main() {
	sieve(18, 323)
}

func boolSmooth(n int) (int, [6]int) {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var a [6]int
	for b, s := range primes {
		for n%s == 0 {
			n = n / s
			if a[b] == 0 {
				a[b] = 1
			} else {
				a[b] = 0
			}
		}
	}
	return n, a
}

func sieve(a, b int) {
	for i := a; i < b; i++ {
		c, d := boolSmooth(i * i % b)
		if c == 1 {
			fmt.Println(i, d)
		}
	}
}
