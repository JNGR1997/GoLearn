package main

import (
	"math/rand/v2"
)

func sort(list []int) []int {
	if len(list) > 1 {
		var left []int
		var right []int
		k := rand.IntN(len(list) - 1)
		pivot := list[k]
		list = append(list[:k], list[k+1:]...)
		for _, v := range list {
			if v < pivot {
				left = append(left, v)
			} else {
				right = append(right, v)
			}
		}
		return append(append(sort(left), pivot), sort(right)...)
	}
	return list
}
