package main

import "fmt"

func main() {
  fmt.Println(doublepoint(1,1)
}

func doublepoint(x, y, n int) int {
  s := 3*x*x + 5
  a, b := euclideanInverse(2*y, n)
  if a == true {
    s=s*b
  } else {
    fmt.Println(2*y %n)
  }
  newx := s*s - 2*x
  newy := s*(x - newx) - y
  return newx, newy
}

func euclideanInverse(a, m int) (bool, int) {
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
  }
  if b != 1 {
    return false, 0
  } else {
    for t2 < 0 {
      t2 = t2 + m
    }
    return true, t2 % m
  }
}