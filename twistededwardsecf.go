package main

import "fmt"

func main() {
	fmt.Println(towerpoint(63, 2914, 455839, 100))
}

func towerpoint(x, y, n, m int) (int, int) {
	for m%2 == 0 {
		x, y = doublepoint(x, y, n)
		m = m / 2
	}

	if m > 1 {
		w, z := towerpoint(x, y, n, m-1)
		x, y = addpoints(x, y, w, z, n)
	}
	return x, y
}

func addpoints(e, f, g, h, n int) (int, int) {
	numx := (e*h + f*g) % n
	numy := move(f*h-6*e*g, n)
	denomx := (1 + 10*e*f*g*h) % n
	denomy := move(1-10*e*f*g*h, n)
	a, b := euclideanInverse(denomx, n)
	c, d := euclideanInverse(denomy, n)
	if !a {
		fmt.Println(denomx, " does not have an inverse mod ", n, ".")
	}
	if !c {
		fmt.Println(denomy, " does not have an inverse mod ", n, ".")
	}
	return numx * b % n, numy * d % n
}

func doublepoint(x, y, n int) (int, int) {
	numx := (2 * x * y) % n
	numy := move(y*y-6*x*x, n)
	denomx := (6*x*x + y*y) % n
	denomy := move(2-6*x*x-y*y, n)
	a, b := euclideanInverse(denomx, n)
	c, d := euclideanInverse(denomy, n)
	if !a {
		fmt.Println(denomx, " does not have an inverse mod ", n, ".")
	}
	if !c {
		fmt.Println(denomy, " does not have an inverse mod ", n, ".")
	}
	return move(numx*b, n), move(numy*d, n)
}

func move(x, n int) int {
	if x < 0 {
		return n - (-x)%n
	}
	return x % n
}

func euclideanInverse(a, m int) (bool, int) {
	n := m
	b := a
	q := (n - n%b) / b
	r := n % b
	t1 := 0
	t2 := 1
	t3 := t1 - q*t2
	for r != 0 {
		n = b
		b = r
		q = (n - n%b) / b
		r = n % b
		t1 = t2
		t2 = t3
		t3 = t1 - q*t2
	}
	if b != 1 {
		return false, 0
	}
	return true, move(t2, m)
}