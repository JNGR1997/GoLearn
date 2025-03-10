package main

func isSmooth(n int, primes []int) bool {
	for _, s := range primes {
		for n%s == 0 {
			n = n / s
		}
	}
	if n == 1 {
		return true
	}
	return false
}

func smooth(n int, primes []int) (int, []int) {
	var a []int
	for b, s := range primes {
		a = append(a, 0)
		for n%s == 0 {
			n = n / s
			a[b] = a[b] + 1
		}
	}
	return n, a
}