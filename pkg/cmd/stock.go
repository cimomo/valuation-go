package cmd

import (
	"github.com/spf13/cobra"
)

// NewStockCmd creates a new root command for Valuation-Go.
func NewStockCmd() *cobra.Command {
	stockCmd := &cobra.Command{
		Use:   "stock",
		Short: "Valuation-Go",
		Long:  ``,
	}

	return stockCmd
}
