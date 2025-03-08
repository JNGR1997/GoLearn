package main

func modInverse(a, n int) (bool, int) {
	for b := 1; b < n; b++ {
		if a*b%n == 1 {
			return true, b
		}
	}
	return false, 0
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
	}
	return true, move(t2, m)
}
