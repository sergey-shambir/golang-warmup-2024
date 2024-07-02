package main

import "math"

func estimatePrimeNumbersCount(maxNumber uint) uint {
	if maxNumber < 4 {
		return 2
	}

	// According to the Prime Number Theorem, prime numbers count in range [1...n] asymptotically tends to `n / ln(n)`
	return maxNumber / uint(math.Log(float64(maxNumber)))
}

func ListPrimeNumbers(maxNumber uint) []uint {
	if maxNumber <= 1 {
		return []uint{}
	}

	results := make([]uint, 0, estimatePrimeNumbersCount(maxNumber))

	// In this sieve, `index = number - 1` and `true` means "may be prime number"
	sieve := make([]bool, maxNumber)
	sieveSize := uint(len(sieve))
	for i := range sieveSize {
		sieve[i] = true
	}

	maxNumberRoot := uint(math.Sqrt(float64(maxNumber)))
	for number := uint(2); number <= maxNumberRoot; number++ {
		index := number - 1
		if sieve[index] {
			results = append(results, number)
		}
		for ; index < sieveSize; index += number {
			sieve[index] = false
		}
	}

	for index := maxNumberRoot; index < sieveSize; index++ {
		if sieve[index] {
			results = append(results, index+1)
		}
	}

	return results
}
