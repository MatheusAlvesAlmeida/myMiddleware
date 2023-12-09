package invoker

import (
	"errors"
)

type PercentageCalculator struct{}

func NewPercentageCalculatorInvoker() PercentageCalculator {
	return PercentageCalculator{}
}

func (PercentageCalculator) GetValueOf(percentage int, totalValue int) (float64, error) {
	if totalValue <= 0 {
		return 0, errors.New("totalValue must be a positive non-zero number")
	}

	if percentage < 0 || percentage > 100 {
		return 0, errors.New("percentage must be between 0 and 100")
	}

	return float64(percentage*totalValue) / 100, nil
}

func (PercentageCalculator) GetPercentageOf(partialValue int, totalValue int) (float64, error) {
	if totalValue <= 0 {
		return 0, errors.New("totalValue must be a positive non-zero number")
	}

	if partialValue < 0 || partialValue > totalValue {
		return 0, errors.New("partialValue must be between 0 and totalValue")
	}

	return (float64(partialValue) / float64(totalValue)) * 100, nil
}
