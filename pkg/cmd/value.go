package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// NewStockCmd creates a new root command for Valuation-Go.
func NewValueCmd() *cobra.Command {
	valueCmd := &cobra.Command{
		Use:   "value",
		Short: "Value a stock",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("stock value expects TICKER")
			}

			fmt.Println("[stock] valuing", args[0])
			return nil
		},
	}

	return valueCmd
}
