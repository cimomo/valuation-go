package cmd

import (
	"errors"
	"fmt"

	"github.com/cimomo/valuation-go/pkg/valuation"
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

			return doValue(args[0])
		},
	}

	return valueCmd
}

func doValue(ticker string) error {
	fmt.Println("[stock] valuing", ticker)

	company, err := valuation.NewCompany(ticker)
	if err != nil {
		return err
	}

	err = company.Fetch()
	if err != nil {
		return err
	}

	market, err := valuation.NewMarket(0.02, 0.02, 0.25)
	if err != nil {
		return err
	}

	input, err := valuation.NewInput(company, 0.15, 0.07, 0.07, 0.12, 0.4, 0.45, 1.17)
	if err != nil {
		return err
	}

	err = input.Compute()
	if err != nil {
		return err
	}

	output, err := valuation.NewOutput(market, input)
	if err != nil {
		return err
	}

	err = output.Compute()
	if err != nil {
		return err
	}

	fmt.Println("Value per share:", output.ValuePerShare)
	return nil
}
