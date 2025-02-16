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
	c := []int{}
	var d [][]int
	for _, v := range a {
		c = append(c, v)
	}
	for _, v := range b {
		j := []int{}
		for _, k := range v {
			j = append(j, k)
		}
		d = append(d, j)
	}
	for !checkNtuple(c, d) {

	}
}

func checkNtuple(a []int, b [][]int) bool {
	for _, c := range b {
		d := 0
		for _, e := range c {
			d = e + d
		}
		if d == 0 {
			return true
		}
	}
	return false
}
