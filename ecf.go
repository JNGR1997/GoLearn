package main

import "fmt"

const n = 455839

func main() {
	start := finite{x: 1, y: 1}
	fmt.Println(start.tower(30300))
}

type point interface {
	add(point) point
	double() point
	tower(int) point
}

type identity struct{}

type finite struct {
	x int
	y int
}

func (id identity) add(p point) point {
	return p
}

func (id identity) double() point {
	return id
}

func (id identity) tower(m int) point {
	return id
}

func (fin1 finite) add(p point) point {
	fin2, ok := p.(finite)
	if !ok {
		// P is the identity element.
		return fin1
	}

	if fin1.x == fin2.x {
		if fin1.y == fin2.y {
			// Fin1 and fin2 are the same point.
			return fin1.double()
		}
		// Fin1 is the inverse of fin2.
		return identity{}
	}

	// Fin1 and fin2 are distinct finite points.
	e, f := euclideanInverse(move(fin2.x - fin1.x))
	if f {
		s := e * move(fin2.y-fin1.y) % n
		newx := move(s*s - fin1.x - fin2.x)
		newy := move(fin1.y - s*(fin1.x-newx))
		return finite{x: newx, y: newy}
	}
	fmt.Println(move(fin2.x-fin1.x), " does not have an inverse mod ", n, ".")
	return identity{}
}

func (fin finite) double() point {
	num := (3*fin.x*fin.x + 5) % n
	denom := (2 * fin.y) % n
	a, b := euclideanInverse(denom)
	if b {
		s := (num * a) % n
		newx := move(s*s - 2*fin.x)
		newy := move(s*(fin.x-newx) - fin.y)
		return finite{x: newx, y: newy}
	}
	fmt.Println(denom, " does not have an inverse mod ", n, ".")
	return identity{}
}

func (fin finite) tower(m int) point {
	var pt point = fin
	for m%2 == 0 {
		pt = pt.double()
		m = m / 2
	}
	if m > 1 {
		rem := pt.tower(m - 1)
		pt = pt.add(rem)
	}
	return pt
}

func move(x int) int {
	if x < 0 {
		return n - (-x)%n
	}
	return x % n
}

func euclideanInverse(a int) (int, bool) {
	m := n
	b := a
	if b == 0 {
		return 0, false
	}
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
	return move(t2), true
}
