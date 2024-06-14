package main 

import "fmt"

func main() {
  fmt.Println(lcm(18,14))
}

func lcm(a, b int) int {
  return a*b/gcd(a,b)
}

func gcd(a, b int) int {
  for b != 0 {
    a, b = b, a%b
  }
  return a
}
