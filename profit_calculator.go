package main

import "fmt"

func countProfit() {
	var revenue, expenses, taxRate float64

	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter taxRate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	fmt.Printf("EBT: %.2f\nProfit: %.2f\n", ebt, profit)

	ratio := ebt / profit
	fmt.Println("Ratio:", ratio)
}
