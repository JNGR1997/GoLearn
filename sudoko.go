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

func compareOptions(a, b [9]bool) bool {
  for k =: 0; k<9; k++ {
    if a[k] != b[k] {
      return false
    }
  }
  return true
}