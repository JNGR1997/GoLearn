package main

import "fmt"

func main() {
  fmt.Println(rho(6))
}

func gcd(a, b int) int {
  for b!= 0 {
    a, b = b, a%b
  }
  return b
}

func rho(n int) int {
  x := 2
  y := 2
  d := 1
  for d == 1 {
    x = (x*x + 1) %n
    y = (y*y*y*y + 2*y*y + 2) %n
    fmt.Println(x,y,d)
    if x-y > 0 {
      d = gcd(n, x-y)
    } else {
      d = gcd(n, y-x)
    }
  }
  return d
}