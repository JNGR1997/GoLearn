package main

func main() {
}

func checkCells(a,b,c,d,e,f,g,h,i [9]bool) ([9]bool, [9]bool, [9]bool, [9]bool, [9]bool, [9]bool, [9]bool, [9]bool, [9]bool) {
}

func numOptions(a [9]bool) int {
  b := 0
  for c,d := range a {
    if d==true {
      b++
    }
  }
  return b
}