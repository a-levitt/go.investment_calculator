package main

import (
	"fmt"
	"math"
)

func countInvestment() {
	const inflationRate = 2.5
	var investmentAmount, expectedReturnRate, years float64

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Enter return rate: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Enter period years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Printf("Feature value: %.2f\n", futureValue)
	fmt.Printf("Real future value: %.2f", futureRealValue)
}
