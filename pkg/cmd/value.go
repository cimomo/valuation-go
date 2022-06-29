package cmd

import (
	"github.com/spf13/cobra"
)

// NewStockCmd creates a new root command for Valuation-Go.
func NewValueCmd() *cobra.Command {
	valueCmd := &cobra.Command{
		Use:   "value",
		Short: "Value a stock",
		Long:  ``,
	}

	return valueCmd
}
