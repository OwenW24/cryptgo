package cryptgo

import (
	"fmt"
)

func Hello(name string) string {
	var message string = fmt.Sprintf("Welcome to cryptgo, %v", name)
	return message
}

// public functions, types, etc should be uppercase
// recursive euclid
func Gcd(u int, v int) int { 
	if v != 0 {
		return Gcd(v, u % v)
	}
	return u
}

// recursive extended euclid
func Bezout(u int, v int) (int, int) {  
	var q, r int = 0, 1
	var s1, s2, s3, t1, t2, t3 int = 1, 0, 1, 0, 1, 0  
	for r > 0 {  // while does not exist
		q = u / v
		r = u - q * v
		s3 = s1 - q * s2; t3 = t1 - q * t2
		if r > 0 {
			u = v; v = r
			s1 = s2; s2 = s3
			t1 = t2; t2 = t3
		}
	}
	return s2, t2
}

func Inv(u int, v int) (int) {  // inverse of u mod v
	var i, _ int = Bezout(u, v)
	if i < 0 {
		i += v
	}
	return i
}

