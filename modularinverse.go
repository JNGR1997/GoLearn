package main

import "fmt"

func main() {
}

func modInverse(a,n int) (bool, int) {
  for b := 1, b < n, b++ {
    if a*b % n == 1 {
      return true, b
    }
  }
  return false, 0
}