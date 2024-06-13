package main

import "fmt"

func main() {
  fmt.Println(gcd(26,6))
}

func gcd(a, b int) int {
  for b != 0 {
    a, b = b, a%b
  }
  return a
}