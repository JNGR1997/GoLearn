package main 

import "fmt"

func main() {
  fmt.Println(isSmooth(34))
}

func isSmooth(n int) bool {
  primes := [6]int{2,3,5,7,11,13}
  for _, s := range primes {
    for n%s == 0 {
      n = n/s
    }
  }
  if n == 1 {
    return true
  }
  return false
}