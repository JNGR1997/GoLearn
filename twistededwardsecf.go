package main

import "fmt"

func main() {
	fmt.Println(point{x: 1, y: 1}.tower(10))
}

const n = 30758

type point struct {
	x int
	y int
}

func (p point) tower(m int) point {
	for m%2 == 0 {
		p = p.double()
		m = m / 2
	}

	if m > 1 {
		p = p.add(p.tower(m - 1))
	}
	return p
}

func (f1 point) add(f2 point) point {
	denomx := (1 + 10*f1.x*f1.y*f2.x*f2.y) % n
	a, b := euclideanInverse(denomx)
	if !a {
		return point{x: 0, y: 1}
	}

	denomy := move(1 - 10*f1.x*f1.y*f2.x*f2.y)
	c, d := euclideanInverse(denomy)
	if !c {
		return point{x: 0, y: 1}
	}
	numx := (f1.x*f2.y + f1.y*f2.x) % n
	numy := move(f1.y*f2.y - 6*f1.x*f2.x)
	return point{x: numx * b % n, y: numy * d % n}
}

func (f point) double() point {
	a, b := euclideanInverse((6*f.x*f.x + f.y*f.y) % n)
	if !a {
		return point{x: 0, y: 1}
	}
	c, d := euclideanInverse(move(2 - 6*f.x*f.x - f.y*f.y))
	if !c {
		return point{x: 0, y: 1}
	}
	numx := (2 * f.x * f.y) % n
	numy := move(f.y*f.y - 6*f.x*f.x)
	return point{x: move(numx * b), y: move(numy * d)}
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
	q := (n - n%b) / b
	r := n % b
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
		fmt.Println(a, " does not have an inverse mod ", n, ".")
		return false, 0
	}
	return true, move(t2)
}
