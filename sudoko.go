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
  c := 0
  for d,_ := range b {
    for f,g := range a {
      if g[d]==true {
        b[d] = b[d]+1
        c = f
      }
    }
    if b[d]==1 {
      for h,i := range a {
        if h != c {
          i[d]=false
        }
      }
    }
  }
  return a
}

func twoOptions(a [9][9]bool) [9][9]bool {
  for b,c := range a {
    for d,e := range a {
      if b != d {
        if numOptions(sudUnion(c,e))==2 {
          for f,g := range a {
            if f != b {
              if f != d {
                for h,i := range sudUnion(c,e) {
                  if i==true {
                    a[f]=removeOption(i,a[f])
                  }
                }
              }
            }
          }
        }
      }
    }
  }
  return a
}

func sudUnion(a,b [9]bool) [9]bool {
  c := [9]bool
  for d,_ := range c {
    if a[d]==true {
      c[d]=true
    } elseif b[d]==true {
      c[d]=true
    } else {
      c[d]=false
    }
  }
  return c
}