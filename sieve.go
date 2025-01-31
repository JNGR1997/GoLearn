package main

func main() {
	sieve(18, 323)
}

func boolSmooth(n int) (int, [6]int) {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var a [6]int
	for b, s := range primes {
		for n%s == 0 {
			n = n / s
			a[b] = (a[b] + 1) % 2
		}
	}
	return n, a
}

func sieve(a, b int) ([7]int, [7][6]int) {
	e := 0
	i := a
	var f [7][6]int
	var g [7]int
	for e < 7 {
		c, d := boolSmooth(i * i % b)
		if c == 1 {
			f[e] = d
			g[e] = i
			e++
		}
		i++
	}
	return g, f
}

func linalg(a [7]int, b [7][6]int) {
}