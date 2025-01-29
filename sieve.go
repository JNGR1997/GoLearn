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
			a[b] = (a[b]+1)%2
		}
	}
	return n, a
}

func sieve(a, b int) {
	e := 0
	i := a
	var f [7][6]int
	for e < 7 {
		c, d := boolSmooth(i * i % b)
		if c == 1 {
			fmt.Println(i, d)
			f[e] = d
			e++
		}
		i++
	}
}
