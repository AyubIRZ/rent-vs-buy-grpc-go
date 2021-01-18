package rent_guess

import (
	"github.com/ayubirz/rent-vs-buy-grpc-go/internal/protos/v1/breakeven"
)

//Guess calculates a Break-even point of renting and buying with equal sense
func Guess(input *breakeven.BreakevenRequest) float64 {
	mortgage := float64(input.HouseValue) - downPayment(input.HouseValue, input.DownPayment)
	mortgageInterest := mortgageInterest(mortgage, input.MortgageInterestRate)

	propertyTax := propertyTax(input.HouseValue, input.PropertyTaxRate)
	propertyTransferTax := propertyTransferTax(input.HouseValue, input.PropertyTransferTaxRate, input.LivingYears)

	annualCosts := mortgageInterest + propertyTax + float64(input.MonthlyCommonCost * 12) + propertyTransferTax

	return annualCosts / 12
}

func downPayment(propertyValue int64, downPaymentRate float32) float64 {
	return float64(float32(propertyValue) * (downPaymentRate / 100))
}

func mortgageInterest(mortgage float64, mortgageInterestRate int32) float64 {
	return mortgage * (float64(mortgageInterestRate) / 100)
}

func propertyTax(propertyValue int64, propertyTaxRate float32) float64 {
	return float64(float32(propertyValue) * (propertyTaxRate / 100))
}

func propertyTransferTax(propertyValue int64, propertyTransferTaxRate float32, livingYears int32) float64 {
	return float64(float32(propertyValue) * (propertyTransferTaxRate / 100)) / float64(livingYears)
}

