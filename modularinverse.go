package main

import "fmt"

func main() {
}

func modInverse(a,n int) (bool, int) {
  for b := 1; b < n; b++ {
    if a*b % n == 1 {
      return true, b
    }
  }
  return false, 0
}

func euclideanInverse(a, m int) {
  n := m
  b := a
  q := (n - n%b)/b
  r := n%b
  t1 := 0
  t2 := 1
  t3 := t1 - q*t2
  for r != 0 {
    n = b
    b = r
    q = (n - n%b)/b
    r = n%b
    t1 = t2
    t2 = t3
    t3 = t1 - q*t2
    fmt.Println(n,b,q,r,t1,t2,t3)
  }
  for t2 < 0 {
    t2 = t2 + m
  }
  return t2 % m
}