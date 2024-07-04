package main

import (
	"fmt"
	"math/big"
)

type currencyTraits struct {
	decimalPlaces int
	unitSize      int
}

func tryGetCurrencyTraits(currency string) (currencyTraits, error) {
	// NOTE: Currencies list is far from complete
	switch currency {
	case "RUB", "CNY", "AED", "INR", "USD", "GBP", "EUR":
		return currencyTraits{
			decimalPlaces: 2,
			unitSize:      100,
		}, nil
	case "JPY", "KRW":
		return currencyTraits{
			decimalPlaces: 0,
			unitSize:      0,
		}, nil
	default:
		return currencyTraits{}, fmt.Errorf("unsupported currency %s", currency)
	}
}

func getCurrencyTraits(currency string) currencyTraits {
	result, err := tryGetCurrencyTraits(currency)
	if err != nil {
		panic(err)
	}
	return result
}

func splitAmount(amount *big.Int, unitSize int) (*big.Int, *big.Int) {
	if unitSize == 0 {
		return amount, nil
	}

	unit := big.NewInt(int64(unitSize))
	integer := big.NewInt(0).Div(amount, unit)
	fraction := big.NewInt(0).Mod(amount, unit)

	return integer, fraction
}
