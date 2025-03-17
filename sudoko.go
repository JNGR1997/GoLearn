package main

func solved(s [9][9][9]bool) bool {
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if numOptions(s[x][y]) > 1 {
				return false
			}
		}
	}
	return true
}


func turn(s [9][9][9]bool) [9][9][9]bool {
	new := [9][9][9]bool{}
	for k := 0; k < 9; k++ {
		for i := 0; i < 9; i++ {
			new[i][k] = s[k][i]
		}
	}
	return new
}

func unturn(s [9][9][9]bool) [9][9][9]bool {
	return turn(turn(turn(s)))
}

func twist(s [9][9][9]bool) [9][9][9]bool {
	new := [9][9][9]bool{}
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			new[x][y] = s[(x-x%3)+(y-(y%3))/3][3*(x%3)+(y%3)]
		}
	}
	return new
}

func untwist(s [9][9][9]bool) [9][9][9]bool {
	return twist(s)
}

func checkCells(a [9][9]bool) [9][9]bool {
	c := 0
	for d := 0; d < 9; d++ {
		b := 0
		for f, g := range a {
			if g[d] == true {
				b++
				c = f
			}
		}
		if b == 1 {
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

func fiveOptions(a [9][9]bool) [9][9]bool {
	for b := 0; b < 5; b++ {
		for c := b + 1; c < 6; c++ {
			for d := c + 1; d < 7; d++ {
				for e := d + 1; e < 8; e++ {
					for f := e + 1; f < 9; f++ {
						suds := sudUnion(a[b], sudUnion(a[c], sudUnion(a[d], sudUnion(a[e], a[f]))))
						if numOptions(suds) == 5 {
							for k, v := range suds {
								if v {
									for g := 0; g < 9; g++ {
										if g != b && g != c && g != d && g != e && g != f {
											a[g] = removeOption(k, a[g])
										}
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

func sixOptions(a [9][9]bool) [9][9]bool {
	for b := 0; b < 4; b++ {
		for c := b + 1; c < 5; c++ {
			for d := c + 1; d < 6; d++ {
				for e := d + 1; e < 7; e++ {
					for f := e + 1; f < 8; f++ {
						for g := f + 1; g < 9; g++ {
							suds := sudUnion(a[b], sudUnion(a[c], sudUnion(a[d], sudUnion(a[e], sudUnion(a[f], a[g])))))
							if numOptions(suds) == 6 {
								for k, v := range suds {
									if v {
										for h := 0; h < 9; h++ {
											if h != b && h != c && h != d && h != e && h != f && h != g {
												a[h] = removeOption(k, a[h])
											}
										}
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
