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

	fmt.Println("Revenue TTM:", input.Revenue)
	fmt.Println("EBIT TTM:", input.EBIT)
	fmt.Println("Total Equity:", input.TotalEquity)
	fmt.Println("Total Debt:", input.TotalDebt)
	fmt.Println("Total Cash:", input.TotalCash)

	fmt.Println("Base year revenue growth rate:", output.BaseYear.RevenueGrowthRate)
	fmt.Println("Base year revenue:", output.BaseYear.Revenue)
	fmt.Println("Base year EBIT margin:", output.BaseYear.EBITMargin)
	fmt.Println("Base year EBIT:", output.BaseYear.EBIT)
	fmt.Println("Base year tax rate:", output.BaseYear.TaxRate)
	fmt.Println("Base year NOPAT:", output.BaseYear.AfterTaxEBIT)

	for i, year := range output.HighGrowthYears {
		fmt.Printf("Year %d revenue growth rate: %f\n", i, year.RevenueGrowthRate)
		fmt.Printf("Year %d revenue: %f\n", i, year.Revenue)
		fmt.Printf("Year %d EBIT margin: %f\n", i, year.EBITMargin)
		fmt.Printf("Year %d EBIT: %f\n", i, year.EBIT)
		fmt.Printf("Year %d tax rate: %f\n", i, year.TaxRate)
		fmt.Printf("Year %d NOPAT: %f\n", i, year.AfterTaxEBIT)
		fmt.Printf("Year %d reinvestment: %f\n", i, year.Reinvestment)
		fmt.Printf("Year %d FCFF: %f\n", i, year.FCFF)
		fmt.Printf("Year %d cost of capital: %f\n", i, year.CostOfCapital)
		fmt.Printf("Year %d discount factor: %f\n", i, year.DiscountFactor)
		fmt.Printf("Year %d PV(FCFF): %f\n", i, year.PresentValueOfCashFlow)
	}

	offset := len(output.HighGrowthYears)

	for i, year := range output.LowGrowthYears {
		fmt.Printf("Year %d revenue growth rate: %f\n", i+offset, year.RevenueGrowthRate)
		fmt.Printf("Year %d revenue: %f\n", i+offset, year.Revenue)
		fmt.Printf("Year %d EBIT margin: %f\n", i+offset, year.EBITMargin)
		fmt.Printf("Year %d EBIT: %f\n", i+offset, year.EBIT)
		fmt.Printf("Year %d tax rate: %f\n", i+offset, year.TaxRate)
		fmt.Printf("Year %d NOPAT: %f\n", i+offset, year.AfterTaxEBIT)
		fmt.Printf("Year %d reinvestment: %f\n", i+offset, year.Reinvestment)
		fmt.Printf("Year %d FCFF: %f\n", i+offset, year.FCFF)
		fmt.Printf("Year %d cost of capital: %f\n", i+offset, year.CostOfCapital)
		fmt.Printf("Year %d discount factor: %f\n", i+offset, year.DiscountFactor)
		fmt.Printf("Year %d PV(FCFF): %f\n", i+offset, year.PresentValueOfCashFlow)
	}

	fmt.Println("Terminal year revenue growth rate:", output.TerminalYear.RevenueGrowthRate)
	fmt.Println("Terminal year revenue:", output.TerminalYear.Revenue)
	fmt.Println("Terminal year EBIT margin:", output.TerminalYear.EBITMargin)
	fmt.Println("Terminal year EBIT:", output.TerminalYear.EBIT)
	fmt.Println("Terminal year tax rate:", output.TerminalYear.TaxRate)
	fmt.Println("Terminal year NOPAT:", output.TerminalYear.AfterTaxEBIT)
	fmt.Println("Terminal year reinvestment:", output.TerminalYear.Reinvestment)
	fmt.Println("Terminal year FCFF:", output.TerminalYear.FCFF)
	fmt.Println("Terminal year cost of capital:", output.TerminalYear.CostOfCapital)

	fmt.Println("Terminal cash flow:", output.TerminalCashFlow)
	fmt.Println("Terminal value:", output.TerminalValue)
	fmt.Println("Present value of terminal value:", output.PresentValueOfTerminalValue)
	fmt.Println("Present value of cash flows:", output.PresentValueOfCashFlow)
	fmt.Println("Present value:", output.PresentValue)
	fmt.Println("Equity value:", output.EquityValue)

	fmt.Println("Value per share:", output.ValuePerShare)
	return nil
}
