package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateMoney(t *testing.T) {
	assert.Equal(t, "0.00 USD", ToString(Zero("USD")))
	assert.Equal(t, "100.00 USD", ToString(FromInt64(10000, "USD")))
	assert.Equal(t, "100.87 USD", ToString(FromInt64(10087, "USD")))
	assert.Equal(t, "19900.00 RUB", ToString(FromInt64(1990000, "RUB")))

	assert.Equal(t, "8117 JPY", ToString(FromInt64(8117, "JPY")))
	assert.Equal(t, "256 KRW", ToString(FromInt64(256, "KRW")))
}

func TestMoneySum(t *testing.T) {
	m1 := FromInt64(8700, "USD")
	m2 := FromInt64(2050, "USD")
	sum := m1.Sum(m2)

	assert.Equal(t, "107.50 USD", ToString(sum))
	assert.Equal(t, "87.00 USD", ToString(m1))
	assert.Equal(t, "20.50 USD", ToString(m2))
}

func TestMoneyDiff(t *testing.T) {
	m1 := FromInt64(8700, "USD")
	m2 := FromInt64(2050, "USD")
	diff := m1.Diff(m2)

	assert.Equal(t, "66.50 USD", ToString(diff))
	assert.Equal(t, "87.00 USD", ToString(m1))
	assert.Equal(t, "20.50 USD", ToString(m2))
}

func TestMoneyTimes(t *testing.T) {
	m1 := FromInt64(8700, "USD")
	m2 := m1.Times(2)
	m3 := m1.Times(5)

	assert.Equal(t, "87.00 USD", ToString(m1))
	assert.Equal(t, "174.00 USD", ToString(m2))
	assert.Equal(t, "435.00 USD", ToString(m3))
}
