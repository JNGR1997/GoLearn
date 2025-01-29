package main

import "fmt"

func main() {
	sieve(18, 323)
}

func boolSmooth(n int) (int, [6]int, [6]int) {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var a [6]int
	var c [6]int
	for b, s := range primes {
		for n%s == 0 {
			n = n / s
			c[b] = c[b] + 1
			if a[b] == 0 {
				a[b] = 1
			} else {
				a[b] = 0
			}
		}
	}
	return n, a, c
}

func sieve(a, b int) {
	e := 0
	i := a
	var f [7][6]int
	for e < 7 {
		c, d, g := boolSmooth(i * i % b)
		if c == 1 {
			fmt.Println(i, g)
			f[e] = d
			e++
		}
		i++
	}
}
