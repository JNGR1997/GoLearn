package main

import "fmt"

func main() {
	fmt.Println(towerpoint(1, 1, 455839, 30300))
}

func doublepoint(x, y, n int) (int, int) {
	num := (3*x*x + 5) % n
	denom := (2 * y) % n
	a, b := euclideanInverse(denom, n)
	if b {
		s := (num * a) % n
		newx := move(s*s-2*x, n)
		newy := move(s*(x-newx)-y, n)
		return newx, newy
	}
	fmt.Println(denom, " does not have an inverse mod ", n, ".")
	return 0, 0
}

func towerpoint(x, y, n, m int) (int, int) {
	for m%2 == 0 {
		x, y = doublepoint(x, y, n)
		m = m / 2
	}
	if m > 1 {
		w, z := towerpoint(x, y, n, m-1)
		return addpoints(x, y, w, z, n)
	}
	return x, y
}

func addpoints(a, b, c, d, n int) (int, int) {
	if a == c && b == d {
		return doublepoint(a, b, n)
	}
	num := move(d-b, n)
	denom := move(c-a, n)
	e, f := euclideanInverse(denom, n)
	if f {
		s := (num * e) % n
		newx := move(s*s-a-c, n)
		newy := move(b-s*(a-newx), n)
		return newx, newy
	}
	fmt.Println(denom, " does not have an inverse mod ", n, ".")
	return 0, 0
}

func move(x, n int) int {
	if x < 0 {
		return n - (-x)%n
	}
	return x % n
}

func euclideanInverse(a, m int) (int, bool) {
	n := m
	b := a
	r := n % b
	q := (n - r) / b
	t1 := 0
	t2 := 1
	t3 := -q
	for r != 0 {
		n = b
		b = r
		r = n % b
		q = (n - r) / b
		t1 = t2
		t2 = t3
		t3 = t1 - q*t2
	}
	if b != 1 {
		return 0, false
	}
	return move(t2, m), true
}
