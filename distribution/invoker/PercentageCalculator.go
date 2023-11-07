package invoker

type PercentageCalculator struct{}

func NewPercentageCalculatorInvoker() PercentageCalculator {
	return PercentageCalculator{}
}

func (PercentageCalculator) GetValueOf(percentage int, totalValue int) float64 {
	return float64(percentage*totalValue) / 100
}

func (PercentageCalculator) GetPercentageOf(partialValue int, totalValue int) float64 {
	return float64(partialValue) / float64(totalValue) * 100
}
