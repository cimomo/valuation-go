package cmd

import (
	"github.com/spf13/cobra"
)

// NewValuationCmd creates a new root command for Valuation-Go.
func NewValuationCmd() *cobra.Command {
	valuationCmd := &cobra.Command{
		Use:   "value",
		Short: "Valuation-Go",
		Long:  ``,
	}

	return valuationCmd
}
