package main

import "fmt"

func main() {
	fmt.Println(doublepoint(1, 1, 455839))
}

func doublepoint(x, y, n int) (int, int) {
	num := 3*x*x + 5
	denom := 2 * y
	comfact := gcd(num, denom)
	if gcd(num, denom) != 1 {
		num = num / comfact
		denom = denom / comfact
	}
	s := num % n
	if denom != 1 {
		a, b := euclideanInverse(denom, n)
		if a == true {
			s = s * b % n
		} else {
			fmt.Println(2 * y % n)
		}
	}
	newx := s*s - 2*x
	newy := s*(x-newx) - y
	return newx, newy
}

func addpoints(a,b,c,d,n int) (int, int) {
  if a == c {
    if b==d {
      x, y := doublepoint(a,b,n)
      return x, y
    }
  } else {
    num := move(d-b,n)
    denom := move(c-a,n)
    e, f := euclideanInverse(denom,n)
    if e != true {
      fmt.Println(denom)
      return 0,0
    } else {
      s := num*f%n
    }
  }
  return move(s*s - a -c,n), move(b - s*(a-move(s*s - a-c)),n)
}

func move(x,n int) int {
  for x < 0 {
    x = x+n
  }
  return x%n
}

func euclideanInverse(a, m int) (bool, int) {
	n := m
	b := a
	q := (n - n%b) / b
	r := n % b
	t1 := 0
	t2 := 1
	t3 := t1 - q*t2
	for r != 0 {
		n = b
		b = r
		q = (n - n%b) / b
		r = n % b
		t1 = t2
		t2 = t3
		t3 = t1 - q*t2
	}
	if b != 1 {
		return false, 0
	} else {
		for t2 < 0 {
			t2 = t2 + m
		}
		return true, t2 % m
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
