package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSqrt(t *testing.T) {
	assertSqrtInDelta(t, 0.0, 0.0)
	assertSqrtInDelta(t, 0.49, 0.7)

	assertSqrtInDelta(t, 1.0, 1.0)
	assertSqrtInDelta(t, 4.0, 2.0)
	assertSqrtInDelta(t, 11.56, 3.4)
}

func TestSqrtWithNegative(t *testing.T) {
	assertSqrtErrNegativeSqrt(t, -0.01)
	assertSqrtErrNegativeSqrt(t, -1)
}

func assertSqrtInDelta(t *testing.T, value, expected float64) {
	epsilon := 0.1e-8
	actual, err := Sqrt(value)
	assert.Nil(t, err)
	assert.InDelta(t, actual, expected, epsilon)
}

func assertSqrtErrNegativeSqrt(t *testing.T, value float64) {
	actual, err := Sqrt(value)
	assert.Equal(t, actual, 0.0)
	assert.EqualError(t, err, fmt.Sprintf("cannot get square root of negative number %v", value))
}
