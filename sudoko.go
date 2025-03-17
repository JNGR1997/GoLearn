package main

func main() {
}

func checkCells(a [9][9]bool) [9][9]bool {
 b := [9]int{}
	c := 0
	for d, _ := range b {
		for f, g := range a {
			if g[d] == true {
				b[d] = b[d] + 1
				c = f
			}
		}
		if b[d] == 1 {
			a[c] = removeOtherOptions(d)
		}
	}
	return a
}

func numOptions(a [9]bool) int {
	b := 0
	for d := 0; d < 9; d++ {
		if a[d] {
			b++
		}
	}
	return b
}

func compareOptions(a, b [9]bool) bool {
	for k := 0; k < 9; k++ {
		if a[k] != b[k] {
			return false
		}
	}
	return true
}

func removeOption(a int, b [9]bool) [9]bool {
	b[a] = false
	return b
}

func removeOtherOptions(a int) [9]bool {
	b := [9]bool{}
	b[a] = true
	return b
}

func oneOption(a [9][9]bool) [9][9]bool {
for k, l := range a {
		if numOptions(l) == 1 {
			for m := 0; m < 9; m++ {
				if m != k {
					a[m] = removeOption(k, a[m])
				}
			}
		}
	}
	return a
}

func twoOptions(a [9][9]bool) [9][9]bool {
	for b := 0; b < 8; b++ {
		for c := b + 1; c < 9; c++ {
			suds := sudUnion(a[b], a[c])
			if numOptions(suds) == 2 {
				for k, v := range suds {
					if v {
						for d := 0; d < 9; d++ {
							if d != b && d != c {
								a[d] = removeOption(k, a[d])
							}
						}
					}
				}
			}
		}
	}
	return a
}

func threeOptions(a [9][9]bool) [9][9]bool {
	for b := 0; b < 7; b++ {
		for c := b + 1; c < 8; c++ {
			for d := c + 1; d < 9; d++ {
				suds := sudUnion(a[b], sudUnion(a[c], a[d]))
				if numOptions(suds) == 3 {
					for k, v := range suds {
						if v {
							for e := 0; e < 9; e++ {
								if e != b && e != c && e != d {
									a[e] = removeOption(k, a[e])
								}
							}
						}
					}
				}
			}
		}
	}
	return a
}

func fourOptions(a [9][9]bool) [9][9]bool {
	for b := 0; b < 6; b++ {
		for c := b + 1; c < 7; c++ {
			for d := c + 1; d < 8; d++ {
				for e := d + 1; e < 9; e++ {
					suds := sudUnion(a[b], sudUnion(a[c], sudUnion(a[d], a[e])))
					if numOptions(suds) == 4 {
						for k, v := range suds {
							if v {
								for f := 0; f < 9; f++ {
									if f != b && f != c && f != d && f != e {
										a[f] = removeOption(k, a[f])
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return a
}

func sudUnion(a, b [9]bool) [9]bool {
	c := [9]bool{}
	for d := 0; d < 9; d++ {
		if a[d] || b[d] {
			c[d] = true
		} else {
			c[d] = false
		}
	}
	return c
}
