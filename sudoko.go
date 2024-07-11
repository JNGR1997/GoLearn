package main

func main() {
}

func checkCells(a,b,c,d,e,f,g,h,i [9]bool) ([9]bool, [9]bool, [9]bool, [9]bool, [9]bool, [9]bool, [9]bool, [9]bool, [9]bool) {
  j := [9][9]bool{a,b,c,d,e,f,g,h,i}
  for {
    for k,l := range j {
      if numOptions(l)==1 {
      }
    }
  }
}

func numOptions(a [9]bool) int {
  b := 0
  for _,d := range a {
    if d==true {
      b++
    }
  }
  return b
}

func compareOptions(a, b [9]bool) bool {
  for k := 0; k<9; k++ {
    if a[k] != b[k] {
      return false
    }
  }
  return true
}

func removeOption(a int, b [9]bool) [9]bool {
  b[a-1]=false
  return b
}