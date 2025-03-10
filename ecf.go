package main

// We are using the curve y^2 = x^3 + 5x -5. (1,1) is a non-trivial point on this curve.

import "fmt"

type point interface {
	add(point, int) point
	double(int) point
	tower(int, int) point
}

type identity struct{}

type finite struct {
	x int
	y int
}

func (id identity) add(p point, n int) point {
	return p
}

func (id identity) double(n int) point {
	return id
}

func (id identity) tower(m, n int) point {
	return id
}

func (fin1 finite) add(p point, n int) point {
	fin2, ok := p.(finite)
	if !ok {
		// P is the identity element.
		return fin1
	}

	if fin1.x == fin2.x {
		if fin1.y == fin2.y {
			// Fin1 and fin2 are the same point.
			return fin1.double(n)
		}
		// Fin1 is the inverse of fin2.
		return identity{}
	}

	// Fin1 and fin2 are distinct finite points.
	num := move(fin2.y-fin1.y, n)
	denom := move(fin2.x-fin1.x, n)
	e, f := euclideanInverse(denom, n)
	if f {
		s := e * num % n
		newx := move(s*s-fin1.x-fin2.x, n)
		newy := move(fin1.y-s*(fin1.x-newx), n)
		return finite{x: newx, y: newy}
	}
	fmt.Println(denom, " does not have an inverse mod ", n, ".")
	return identity{}
}

func (fin finite) double(n int) point {
	num := (3*fin.x*fin.x + 5) % n
	denom := (2 * fin.y) % n
	a, b := euclideanInverse(denom, n)
	if b {
		s := (num * a) % n
		newx := move(s*s-2*fin.x, n)
		newy := move(fin.y-s*(fin.x-newx), n)
		return finite{x: newx, y: newy}
	}
	fmt.Println(denom, " does not have an inverse mod ", n, ".")
	return identity{}
}

func (fin finite) tower(m, n int) point {
	var pt point = fin
	for m%2 == 0 {
		pt = pt.double(n)
		m = m / 2
	}
	if m > 1 {
		pt = pt.add(pt.tower(m-1, n), n)
	}
	return pt
}

func move(x, n int) int {
	if x < 0 {
		return n - (-x)%n
	}
	return x % n
}

func euclideanInverse(a, n int) (int, bool) {
	m := n
	b := a
	q := (m - m%b) / b
	r := m % b
	t1 := 0
	t2 := 1
	t3 := t1 - q*t2
	for r != 0 {
		m = b
		b = r
		q = (m - m%b) / b
		r = m % b
		t1 = t2
		t2 = t3
		t3 = t1 - q*t2
	}
	if b != 1 {
		return 0, false
	}
	return move(t2, n), true
}
