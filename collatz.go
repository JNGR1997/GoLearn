package main 

import "fmt"

func main() {
  total := 0
  for i := 0; i<32; i++ {
    if col(32,i) == true {
      total++
    }
  }
  fmt.Println(total)
}

func col(a,b int) bool {
  c := a
  d := b
  for c >= a {
    if c%2 == 0 {
      if d%2 == 0 {
        c = c/2
        d = d/2
      } else {
        c = 3*c
        d = 3*d + 1
      }
    } else {
      return true
    }
  }
  return false
}