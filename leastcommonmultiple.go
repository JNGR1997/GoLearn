package main 

func lcm(a, b int) int {
  return a*b/gcd(a,b)
}

func gcd(a, b int) int {
  for b != 0 {
    a, b = b, a%b
  }
  return a
}
