package main

import (
	"github.com/cimomo/valuation-go/pkg/cmd"
)

func main() {
	stockCmd := cmd.NewStockCmd()
	stockCmd.Execute()
}
