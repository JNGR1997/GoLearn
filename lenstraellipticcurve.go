package main

import "fmt"

func main() {
	fmt.Println(towerpoint(1, 1, 455839, 30300))
}

func doublepoint(x, y, n int) (int, int) {
	s := (3*x*x + 5) % n
	denom := (2 * y) % n
	if denom != 1 {
		a, b := euclideanInverse(denom, n)
		if a {
			s = (s * b) % n
		} else {
			fmt.Println(denom, " does not have an inverse mod ", n, ".")
			return 0, 0
		}
	}
	newx := move(s*s-2*x, n)
	newy := move(s*(x-newx)-y, n)
	return newx, newy
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
	if a == c {
		if b == d {
			return doublepoint(a, b, n)
		}
	} else {
		num := move(d-b, n)
		denom := move(c-a, n)
		e, f := euclideanInverse(denom, n)
		if e {
			s := num * f % n
			return move(s*s-a-c, n), move(b-s*(a-move(s*s-a-c, n)), n)
		}
		fmt.Println(denom, " does not have an inverse mod ", n, ".")
	}
	return 0, 0
}

func move(x, n int) int {
	if x < 0 {
		x = n - (-x)%n
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
	} else {
		for t2 < 0 {
			t2 = t2 + m
		}
		return true, t2 % m
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
