package main

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
		n = b
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
