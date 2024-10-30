package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Println("Enter investment amount:")
	fmt.Scan(&investmentAmount)
	fmt.Println("Enter return rate:")
	fmt.Scan(&expectedReturnRate)
	fmt.Println("Enter period years:")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println("Feature value:", futureValue)
	fmt.Println("Real future value:", futureRealValue)
}
