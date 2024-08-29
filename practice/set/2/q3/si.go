package main

import "fmt"

// Loan struct to hold loan information
type Loan struct {
	Principal          float64
	AnnualInterestRate float64
	NumberOfYears      int
	SimpleInterest     float64
}

func main() {
	var numSets int

	// Prompt the user for the number of sets
	fmt.Print("Enter the number of sets of data: ")
	fmt.Scanln(&numSets)

	// Create an array of Loan structs to store the input data and results
	// loans := make([]Loan, numSets)
	loans := [numSets]Loan

	// Input data for each set
	for i := 0; i < numSets; i++ {
		fmt.Printf("\nSet %d:\n", i+1)

		fmt.Print("Enter Principal Amount: ")
		fmt.Scanln(&loans[i].Principal)

		fmt.Print("Enter Annual Interest Rate: ")
		fmt.Scanln(&loans[i].AnnualInterestRate)

		fmt.Print("Enter Number of Years: ")
		fmt.Scanln(&loans[i].NumberOfYears)

		// Calculate simple interest for the current set
		loans[i].SimpleInterest = calculateSimpleInterest(loans[i].Principal, loans[i].AnnualInterestRate, loans[i].NumberOfYears)
	}

	// Display the results
	fmt.Println("\nResults:")
	for i := 0; i < numSets; i++ {
		fmt.Printf("\nSet %d:\n", i+1)
		fmt.Printf("Principal Amount: %.2f\n", loans[i].Principal)
		fmt.Printf("Annual Interest Rate: %.2f%%\n", loans[i].AnnualInterestRate)
		fmt.Printf("Number of Years: %d\n", loans[i].NumberOfYears)
		fmt.Printf("Simple Interest: %.2f\n", loans[i].SimpleInterest)
	}
}

// calculateSimpleInterest calculates the simple interest using the formula
func calculateSimpleInterest(principal float64, annualInterestRate float64, numberOfYears int) float64 {
	return (principal * annualInterestRate * float64(numberOfYears)) / 100
}
