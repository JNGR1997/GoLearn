package main

import "fmt"

func main() {
}

func doublepoint(x, y, n int) int {
  s := 3*x*x + 5
  s = s/(2*y)
  newx := s*s - 2*x
  newy := s*(x - newx) - y
}