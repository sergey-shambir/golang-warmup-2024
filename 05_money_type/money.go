package main

import (
	"fmt"
	"math/big"
	"strconv"
)

type Money struct {
	amount   *big.Int
	currency string
}

func Zero(currency string) *Money {
	return FromInt64(0, currency)
}

func FromInt64(amount int64, currency string) *Money {
	return &Money{
		amount:   big.NewInt(amount),
		currency: currency,
	}
}

func (m *Money) Currency() string {
	return m.currency
}

func (m *Money) Amount() *big.Int {
	return m.amount
}

// Returns new object that represents money sum
// Panics on currency mismatch
func (m *Money) Sum(other *Money) *Money {
	if m.currency != other.currency {
		panic(fmt.Errorf("cannot add money: currency mismatch %s <> %s", m.currency, other.currency))
	}

	return &Money{
		amount:   big.NewInt(0).Add(m.amount, other.amount),
		currency: m.currency,
	}
}

func (m *Money) Diff(from *Money) *Money {
	if m.currency != from.currency {
		panic(fmt.Errorf("cannot substract money: currency mismatch %s <> %s", m.currency, from.currency))
	}

	return &Money{
		amount:   big.NewInt(0).Sub(m.amount, from.amount),
		currency: m.currency,
	}
}

func (m *Money) Times(factor int64) *Money {
	newValue := big.NewInt(factor)

	return &Money{
		amount:   newValue.Mul(newValue, m.amount),
		currency: m.currency,
	}
}

func (m *Money) Format(s fmt.State, ch rune) {
	traits := getCurrencyTraits(m.currency)
	integer, fraction := splitAmount(m.amount, traits.unitSize)

	var amountText string
	if fraction != nil {
		pattern := "%d.%0" + strconv.Itoa(traits.decimalPlaces) + "d"
		amountText = fmt.Sprintf(pattern, integer, fraction)
	} else {
		amountText = fmt.Sprintf("%d", integer)
	}

	// nolint:errcheck
	s.Write([]byte(amountText))

	// nolint:errcheck
	s.Write([]byte(" "))

	// nolint:errcheck
	s.Write([]byte(m.currency))
}

func ToString(m *Money) string {
	return fmt.Sprintf("%s", m)
}

// TODO: Implement Allocate and comparison operators
