package main

func main() {
}

func checkCells(a [9][9]bool) [9][9]bool {
  for {
    for k,l := range a {
      if numOptions(l)==1 {
        for m,n := range a {
          if m != k {
            a[m] = removeOption(k,a[m])
          }
        }
      }
    }
  }
  return a
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

func oneOption(a [9][9]bool) [9][9]bool {
  b := [9]int
  for c,d := range a {
    for e,f := range d {
      if f == true {
        b[e-1] = b[e-1]+1
      }
    }
  }
  for g,h := range b {
    if h==1 {
      for i,j := range a {
        if j[g]==true {
          a[i] = {false, false, false, false, false, false, false,  false, false}
          a[i][g]=true
        }
      }
    }
  }
  return a
}