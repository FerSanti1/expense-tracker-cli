package main

import (
	"encoding/json"
	"flag"
	"os"
	"time"

	"github.com/fer1santi/expense-tracker-cli/data"
)

var expenses []data.Expense

// Package main provides the entry point for the expense tracker CLI application.
func main() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	descriptionPtr := addCmd.String("description", "", "description of the expense")
	amountPtr := addCmd.Int("amount", 0, "amount of the expense")

	if len(os.Args) <= 2 {
		println("# Usage: expense-tracker-cli [add|delete|summary|list] [options]")
		return
	}
	addCmd.Parse(os.Args[2:])

	if *descriptionPtr == "" || *amountPtr <= 0 {
		println("# Please provide a valid description and amount.")
		return
	}

	// Initialize expenses slice if it's nil
	if expenses == nil {
		expenses = []data.Expense{}
	}

	// Check if the description and amount are provided
	if *descriptionPtr == "" || *amountPtr <= 0 {
		println("# Error: Description and amount must be provided.")
		return
	}

	// Create a new expense and add it to the expenses slice
	expense := data.Expense{
		ID:          len(expenses) + 1, // Simple ID generation
		Description: *descriptionPtr,
		Amount:      *amountPtr,
		Date:        time.Now().Format("2006-01-02"),
	}
	expenses = append(expenses, expense)
	println("# Expense added:", expense.Description, "Amount:", expense.Amount)
	output, err := json.Marshal(expenses)
	if err != nil {
		println("Error marshalling expenses:", output)
	}
	println("# Current expenses:", string(output))

}
