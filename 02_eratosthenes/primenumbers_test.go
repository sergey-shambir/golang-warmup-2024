package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListPrimeNumbers(t *testing.T) {
	assert.Equal(t, ListPrimeNumbers(1), []uint{})

	assert.Equal(t, ListPrimeNumbers(2), []uint{2})

	assert.Equal(t, ListPrimeNumbers(3), []uint{2, 3})

	assert.Equal(t, ListPrimeNumbers(11), []uint{2, 3, 5, 7, 11})

	assert.Equal(t, ListPrimeNumbers(20), []uint{2, 3, 5, 7, 11, 13, 17, 19})

	assert.Equal(t, ListPrimeNumbers(100), []uint{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97})
}

func TestListPrimeNumbersCalculatedUpToMillion(t *testing.T) {
	primes := ListPrimeNumbers((1_000_000))
	assert.Equal(t, primes[0], uint(2))
	assert.Equal(t, primes[len(primes)-1], uint(999983))
	assert.Equal(t, len(primes), 78498)
}
