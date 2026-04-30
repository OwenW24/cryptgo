package cryptgo

import (
	"fmt"
	"math"
)

func Hello(name string) string {
	var message string = fmt.Sprintf("Welcome to cryptgo, %v", name)
	return message
}


// variadic 
// Pow(...args) [0]base, [1]exp, [2]m
func Pow(arg ...int) int {
	sz := len(arg)
	if !(sz == 2 || sz == 3) {
		panic("invalid args for Pow")
	}

	base := arg[0]
	exp := arg[1]
	if exp == 0 {
		return 1
	}
	if sz == 3 {
		p := arg[2]
		base = base % p 
		if base < 0 {
			base += p	
		}
		// exp = exp % Order(base, p) recursive call fucks itself, checkout order...

		fmt.Println(base, exp, p, Order(base, p))
		if exp < 0 {
			return Inv(int(math.Pow(float64(base), float64(-1 * exp))) % p, p) % p // this works somehow
		}
		return int(math.Pow(float64(base), float64(exp))) % p
	}

	return int(math.Pow(float64(base), float64(exp)))

	
}


// public functions, types, etc should be uppercase
// recursive euclid
func Gcd(u, v int) int { 
	if v != 0 {
		return Gcd(v, u % v)
	}
	return u
}

// recursive extended euclid
func Bezout(u, v int) (int, int) {  
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

func Order(x, m int) int {
	var k int = 1
	for Pow(x, k, m) != 1 {
		k++
	}
	return k
}

func PrimitiveRoots(m int) []int {
	var roots []int
	for i := 1; i < m; i++ {
		if Order(i, m) == m - 1 {
			roots = append(roots, i)
		}
	}

	return roots
}

// I'll do Polhig-Hellman and index calc l8r

// primality testing
// beware of Carmichael :/ how the hell did i mess this up
func FermatTest(p int) bool { 
	for i:=1; i < p; i++ {
		if Pow(i, p-1, p) != 1 {  // large powers breaking, i believe we are exceeding integer bounds
			fmt.Printf("%d^%d != 1 mod %d", i, p-1, p)
			return false
		}
	}
	return true
}

// Shanks baby-giant
func Shanks(g, h, p int) int {

	N := Order(g, p)
	b := int(math.Floor(math.Sqrt(float64(N)))) + 1
	// fmt.Printf("order = %d, b = %d\n", N, b) // cool

	// not efficient, for educational purposes, would like to make this more efficient with my own set

	var baby, giant []int

	for i := range (b + 1) {
		baby = append(baby, Pow(g, i, p))
		giant = append(giant, h * (Pow(g,-1 * i * b, p)) % p)
	}

	for i := range baby {
		for j := range giant {
			if baby[i] == giant[j] {
				return i + j * b
			}
		}
	}

	return 0
}


// FACTORING !!!
// just returns one of the factors of some composoite
func Pollard_Factor(N, bound int) int { 
	var a int = 2
	var i int = 2
	var curr int = a
	var d int = 1
	for i <= bound {
		curr = Pow(curr, i, N)
		d = Gcd(curr - 1, N)
		if 1 < d && d < N {
			return d
		} else if d == N {
			i = 2
			a++
		}
		i++
	}
	return d
}





// elliptic curve crypto

// addition using variadic funcs
func Ec_add(a, b float64, pq ...float64,) (float64, float64) {

	sz := len(pq)
	if sz != 2 && sz != 4 {
		return 0, 0
	}

	if sz == 2 {
		x1 := pq[0]
		y1 := pq[1]
		lam := (3*math.Pow((x1), 2) + a) / (2 * y1)
		x3 := math.Pow(lam, 2) - (2 * x1)
		y3 := lam * (x1 - x3) - y1
		return x3, y3
	}

	if sz == 4 {
		x1, y1, x2, y2 := pq[0], pq[1], pq[2], pq[3]
		lam := (y2 - y1) / (x2 - x1)
		x3 := math.Pow(lam, 2) - x1 - x2
		y3 := lam * (x1 - x3) - y1
		return x3, y3
	}

	return 0, 0 
}


func Ecff_add(a, b, p int, pq ...int) (int, int) {
	sz := len(pq)

	x3, y3 := 0
	if sz != 2 && sz != 4 {
		return x3, y3
	}


	if sz == 2 {
		x1, y1 := pq[0], pq[1]
		lam = (3 * Pow(x1, 2, p) + a) * Pow(2*y1, -1, p)
		x3 = (Pow(lam, 2, p) - 2*x1) % p
		y3 = lam*(x1 - x3) - y1
		return x3, y3
	}


	if sz == 4 {
		x1, y1, x2, y2 := pq[0], pq[1], pq[2], pq[3]
		lam = (y2 - y1) * pow((x2 - x1), -1, p)  // not sure if pow works for negatives, i should update inputs...
		x3 = (Pow(lam, 2, p) - 2 * x1) % p
		y3 = lam*(x1 - x3) - y1
		return x3, y3
	}

	return x3, y3
}