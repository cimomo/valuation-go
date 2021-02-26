package cmd

import (
	"github.com/spf13/cobra"
)

// NewValueCmd creates a new root command for Valuation-Go.
func NewValueCmd() *cobra.Command {
	valueCmd := &cobra.Command{
		Use:   "value",
		Short: "Valuation-Go",
		Long:  ``,
	}

	return valueCmd
}
