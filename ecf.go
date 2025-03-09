package main

import "fmt"

const n = 455839

// We are using the curve y*y = x*x*x + 5*x - 5.

type point interface {
	add(point) point
	double() point
	tower(int) point
}

type identity struct{}

type infinity struct{}

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

func (inf infinity) add(p point) point {
	return inf
}

func (inf infinity) double() point {
	return inf
}

func (inf infinity) tower(m int) point {
	return inf
}

func (fin1 finite) add(p point) point {
	// Check if p is the identity element.
	_, ok := p.(identity)
	if ok {
		return fin1
	}

	// Check if p is the point at infinity.
	_, ok = p.(infinity)
	if ok {
		return p
	}

	// P is not the point at infinity or the identity element, so it is a finite point.
	fin2, _ := p.(finite)

	// Check if fin1 and fin2 are the same point.
	if fin1.x == fin2.x && fin1.y == fin2.y {
		return fin1.double()
	}

	// Fin1 and fin2 are distinct finite points.
	e, f := euclideanInverse(move(fin2.x - fin1.x))
	if e {
		s := f * move(fin2.y-fin1.y) % n
		newx := move(s*s - fin1.x - fin2.x)
		newy := move(fin1.y - s*(fin1.x-newx))
		return finite{x: newx, y: newy}
	}
	fmt.Println(move(fin2.x-fin1.x), " does not have an inverse mod ", n, ".")
	return infinity{}
}

func (fin finite) double() point {
	num := (3*fin.x*fin.x + 5) % n
	denom := (2 * fin.y) % n
	a, b := euclideanInverse(denom)
	if a {
		s := (num * b) % n
		newx := move(s*s - 2*fin.x)
		newy := move(s*(fin.x-newx) - fin.y)
		return finite{x: newx, y: newy}
	}
	fmt.Println(denom, " does not have an inverse mod ", n, ".")
	return infinity{}
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

func euclideanInverse(a int) (bool, int) {
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
		return false, 0
	}
	return true, move(t2)
}

func main() {
	start := finite{x: 1, y: 1}
	fmt.Println(start.tower(30300))
}
