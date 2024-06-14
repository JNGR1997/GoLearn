package main 

import "fmt"

func main() {
  fmt.Println(lcm(18,14))
}

func lcm(a, b int) int {
  c := a*b
  for b != 0 {
    a, b = b, a%b
  }
  return c/a
}
