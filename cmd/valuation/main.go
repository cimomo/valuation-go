package main

import (
	"fmt"

	"github.com/cimomo/valuation-go/pkg/cmd"
)

func main() {
	fmt.Println("Hello, valuation.")

	stockCmd := cmd.NewStockCmd()
	stockCmd.Execute()
}
