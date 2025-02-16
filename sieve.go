package main

func main() {
	linalg(sieve(18, 323))
}

func boolSmooth(n int) (int, []int) {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	a := []int{0, 0, 0, 0, 0, 0}
	for b, s := range primes {
		for n%s == 0 {
			n = n / s
			a[b] = (a[b] + 1) % 2
		}
	}
	return n, a
}

func sieve(a, b int) ([]int, [][]int) {
	e := 0
	i := a
	var f [][]int
	var g []int
	for e < 7 {
		c, d := boolSmooth(i * i % b)
		if c == 1 {
			f = append(f, d)
			g = append(g, i)
			e++
		}
		i++
	}
	return g, f
}

func linalg(a []int, b [][]int) {
	c := a
	d := b
	for !checkBlock(c, d) {

	}
}

func checkBlock(a []int, b [][]int) bool {
	for _, c := range b {
		if checkNtuple(c) {
			return true
		}
	}
	return false
}

func checkNtuple(a []int) bool {
	for _, b := range a {
		if b != 0 {
			return false
		}
	}
	return true
}

func removeEmptyColumns(a [][]int) [][]int {
}
