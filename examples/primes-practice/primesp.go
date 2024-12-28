package primesp

func FactorsOf(n int) []int {
	factors := []int{}
	for divisor := 2; n > 1; divisor++ {
		for ; n%divisor == 0; n /= divisor {
			factors = append(factors, divisor)
		}
	}

	return factors
}
