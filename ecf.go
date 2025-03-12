package main

// We are dealing with curves of the form y*y = x*x*x + bx + c mod n.
const (
	b = 5
	c = -5
	n = 455389
)

type point interface {
	add(point) point
	double() point
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
	numerator := move(fin2.y - fin1.y)
	denominator := move(fin2.x - fin1.x)
	inverse, exists := euclideanInverse(denominator)
	if exists {
		s := (numerator * inverse) % n
		newx := move(s*s - fin1.x - fin2.x)
		newy := move(fin1.y - s*(fin1.x-newx))
		return finite{x: newx, y: newy}
	}
	return identity{}
}

func (fin finite) double() point {
	if fin.y == 0 {
		return identity{}
	}
	numerator := move(3*fin.x*fin.x + b)
	denominator := (2 * fin.y) % n
	if denominator != 0 {
		inverse, exists := euclideanInverse(denominator)
		if exists {
			s := (numerator * inverse) % n
			newx := move(s*s - 2*fin.x)
			newy := move(fin.y - s*(fin.x-newx))
			return finite{x: newx, y: newy}
		}
	}
	return identity{}
}

func move(x int) int {
	if x < 0 {
		return n - (-x)%n
	}
	return x % n
}

func euclideanInverse(a int) (int, bool) {
	m := n
	j := a
	r := n % a
	q := (m - r) / j
	t1 := 0
	t2 := 1
	t3 := -q
	for r != 0 {
		m = j
		j = r
		r = m % j
		q = (m - r) / j
		t1 = t2
		t2 = t3
		t3 = t1 - q*t2
	}
	if j != 1 {
		return 0, false
	}
	return move(t2), true
}
