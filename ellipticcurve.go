package main

const (
	b = 5
	c = -5
	n = 455839
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

	if move(fin1.x) == move(fin2.x) {
		if move(fin1.y) == move(fin2.y) {
			// Fin1 and fin2 are the same point.
			return fin1.double()
		}
		// Fin1 is the inverse of fin2.
		return identity{}
	}

	// Fin1 and fin2 are distinct finite points.
	num := move(fin2.y - fin1.y)
	denom := move(fin2.x - fin1.x)
	inverse, exists := euclideanInverse(denom)
	if exists {
		s := inverse * num % n
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
	num := move(3*fin.x*fin.x + b)
	denom := (2 * fin.y) % n
	if denom != 0 {
		inv, exists := euclideanInverse(denom)
		if exists {
			s := (num * inv) % n
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
