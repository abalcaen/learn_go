package main

import (
	"errors"
	"fmt"
	"os"
)

func getUserInput(inputTxt string) (float64, error) {
	var inputVar float64
	fmt.Print(inputTxt)
	fmt.Scan(&inputVar)

	if inputVar <= 0 {
		return 0, errors.New("inputted value cannot zero or a negative value")
	}

	return inputVar, nil
}

func performCalculations(revenue float64, expenses float64, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate)
	ratio = ebt / profit
	return ebt, profit, ratio
}

func main() {
	revenue, err := getUserInput("Input revenue:")
	if err1 != nil {
		fmt.Printf("==========ERROR==========\n")
		fmt.Println(err)
		fmt.Printf("=========================\n")
		return
	}
	expenses, err := getUserInput("Input expenses:")
	if err != nil {
		fmt.Printf("==========ERROR==========\n")
		fmt.Println(err)
		fmt.Printf("=========================\n")
		return
	}
	taxRate, err := getUserInput("Input tax rate:")
	if err != nil {
		fmt.Printf("==========ERROR==========\n")
		fmt.Println(err)
		fmt.Printf("=========================\n")
		return
	}

	ebt, profit, ratio := performCalculations(revenue, expenses, taxRate)

	fmt.Printf("EBT is %0.2f\n", ebt)
	fmt.Printf("Profit is %0.2f\n", profit)
	fmt.Printf("Ratio is %0.2f\n", ratio)

	outputTxt := fmt.Sprintf("EBT,%0.2f\nProfit,%0.2f\nRatio,%0.2f\n", ebt, profit, ratio)
	os.WriteFile("calculator_output.txt", []byte(outputTxt), 0644)
}

// ask for revenue, expenses, tax rate
// calculate earnings before tax and after tax (profit)
// calculate ratio (ebt/profit)
// output ebt, profit and ratio

// have asking for inputs as a function
// have calculations as a function

// validate user inputs by showing error message and exit if there is an invalid input. Those being no negative numbers and no zero
// store the calculated results into a file
