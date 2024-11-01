package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func countInvestment() {

	var investmentAmount, expectedReturnRate, years float64

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Enter return rate: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Enter period years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	fmt.Printf("Feature value: %.2f\n", futureValue)
	fmt.Printf("Real future value: %.2f", futureRealValue)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
