package main

import (
	"errors"
	"fmt"
)

func countProfit() {
	revenue, err := getUserInput("Enter revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	expenses, err := getUserInput("Enter expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	taxRate, err := getUserInput("Enter taxRate: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, profit := countEbtProfit(revenue, expenses, taxRate)
	fmt.Printf("EBT: %.2f\nProfit: %.2f\n", ebt, profit)

	ratio := ebt / profit
	fmt.Printf("Ratio: %.2f", ratio)
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("Input must be positive number.")
	}
	return userInput, nil
}

func countEbtProfit(revenue, expenses, taxRate float64) (float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	return ebt, profit
}
