package cryptgo

import (
	"testing"
)

func TestExp(t *testing.T) {
	result := Pow(9, 2, 8)
	expected := 1
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestGcd(t *testing.T) {
	result := Gcd(27, 123)
	expected := 3
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestBezout(t *testing.T) {
	result1, result2 := Bezout(38371, 3813)
	expected1, expected2 := -443, 4458
	if result1 != expected1 || result2 != expected2 {
		t.Errorf("Expected (%d, %d) but got (%d, %d)", expected1, expected2, result1, result2)
	}
}

func TestInv(t *testing.T) {
	result := Inv(3813, 38371)
	expected := 4458
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}