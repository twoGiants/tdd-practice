package primesp_test

import (
	"slices"
	primesp "tdd/examples/primes-practice"
	"testing"
)

func TestFactorsOf(t *testing.T) {
	factors := primesp.FactorsOf(1)
	result := []int{}
	if !slices.Equal(factors, result) {
		t.Errorf("expected %v but got %v", result, factors)
	}

	factors = primesp.FactorsOf(2)
	result = []int{2}
	if !slices.Equal(factors, result) {
		t.Errorf("expected %v but got %v", result, factors)
	}

	factors = primesp.FactorsOf(3)
	result = []int{3}
	if !slices.Equal(factors, result) {
		t.Errorf("expected %v but got %v", result, factors)
	}

	factors = primesp.FactorsOf(4)
	result = []int{2, 2}
	if !slices.Equal(factors, result) {
		t.Errorf("expected %v but got %v", result, factors)
	}

	factors = primesp.FactorsOf(5)
	result = []int{5}
	if !slices.Equal(factors, result) {
		t.Errorf("expected %v but got %v", result, factors)
	}

	factors = primesp.FactorsOf(6)
	result = []int{2, 3}
	if !slices.Equal(factors, result) {
		t.Errorf("expected %v but got %v", result, factors)
	}

	factors = primesp.FactorsOf(7)
	result = []int{7}
	if !slices.Equal(factors, result) {
		t.Errorf("expected %v but got %v", result, factors)
	}

	factors = primesp.FactorsOf(8)
	result = []int{2, 2, 2}
	if !slices.Equal(factors, result) {
		t.Errorf("expected %v but got %v", result, factors)
	}

	factors = primesp.FactorsOf(9)
	result = []int{3, 3}
	if !slices.Equal(factors, result) {
		t.Errorf("expected %v but got %v", result, factors)
	}

	factors = primesp.FactorsOf(2 * 2 * 3 * 5 * 7 * 11 * 13 * 17)
	result = []int{2, 2, 3, 5, 7, 11, 13, 17}
	if !slices.Equal(factors, result) {
		t.Errorf("expected %v but got %v", result, factors)
	}
}
