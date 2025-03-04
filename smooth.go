package main

import "fmt"

func main() {
	fmt.Println(smooth(34))
}

func isSmooth(n int) bool {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	for _, s := range primes {
		for n%s == 0 {
			n = n / s
		}
	}
	if n == 1 {
		return true
	}
	return false
}

func smooth(n int) (int, [6]int) {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var a [6]int
	for b, s := range primes {
		for n%s == 0 {
			n = n / s
			a[b] = a[b] + 1
		}
	}
	return n, a
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