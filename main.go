package main

import (
	"flag"

	"github.com/fer1santi/expense-tracker-cli/data"
)

var expenses []data.Expense

// Package main provides the entry point for the expense tracker CLI application.
func main() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	summaryCmd := flag.NewFlagSet("summary", flag.ExitOnError)
}
