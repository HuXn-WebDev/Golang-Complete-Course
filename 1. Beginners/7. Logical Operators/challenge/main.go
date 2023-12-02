package main

import "fmt"

func main() {
	highIncome := false
	goodCreditScore := false
	eligibleForLoan := highIncome || goodCreditScore
	applicationRefused := !eligibleForLoan
	fmt.Println("Eligible:", eligibleForLoan)
	fmt.Println("Application Refused:", applicationRefused)
}
