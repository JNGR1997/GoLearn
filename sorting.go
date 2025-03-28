package main

import "math/rand/v2"

func sort(list []int, a, b int) {
	if b-a > 0 {
		k := a + rand.IntN(b-a)
		list[k], list[b] = list[b], list[k]
		z, k := b-1, b
		for z > a-1 {
			if list[z] < list[k] {
				z--
			} else {
				list[z], list[k-1], list[k] = list[k-1], list[k], list[z]
				k--
				z--
			}
		}
		sort(list, a, k-1)
		sort(list, k+1, b)
	}
}
