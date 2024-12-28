package primes_test

import (
	"slices"
	primes "tdd/examples/primes"
	"testing"
)

func TestPrimeFactorsOf(t *testing.T) {
	factors := primes.FactorsOf(1)
	result := []int{}
	if !slices.Equal(factors, result) {
		t.Errorf("expected %v but got %v", result, factors)
	}

	factors = primes.FactorsOf(2)
	result = []int{2}
	if !slices.Equal(factors, result) {
		t.Errorf("expected %v but got %v", result, factors)
	}

	factors = primes.FactorsOf(3)
	result = []int{3}
	if !slices.Equal(factors, result) {
		t.Errorf("expected %v but got %v", result, factors)
	}

	factors = primes.FactorsOf(4)
	result = []int{2, 2}
	if !slices.Equal(factors, result) {
		t.Errorf("expected %v but got %v", result, factors)
	}

	factors = primes.FactorsOf(5)
	result = []int{5}
	if !slices.Equal(factors, result) {
		t.Errorf("expected %v but got %v", result, factors)
	}

	factors = primes.FactorsOf(6)
	result = []int{2, 3}
	if !slices.Equal(factors, result) {
		t.Errorf("expected %v but got %v", result, factors)
	}

	factors = primes.FactorsOf(7)
	result = []int{7}
	if !slices.Equal(factors, result) {
		t.Errorf("expected %v but got %v", result, factors)
	}

	factors = primes.FactorsOf(8)
	result = []int{2, 2, 2}
	if !slices.Equal(factors, result) {
		t.Errorf("expected %v but got %v", result, factors)
	}

	factors = primes.FactorsOf(9)
	result = []int{3, 3}
	if !slices.Equal(factors, result) {
		t.Errorf("expected %v but got %v", result, factors)
	}

	factors = primes.FactorsOf(2 * 2 * 3 * 5 * 7 * 11 * 13 * 17)
	result = []int{2, 2, 3, 5, 7, 11, 13, 17}
	if !slices.Equal(factors, result) {
		t.Errorf("expected %v but got %v", result, factors)
	}
}
