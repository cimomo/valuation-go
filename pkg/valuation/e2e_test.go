package valuation

import "testing"

func TestValuation(t *testing.T) {
	company, err := NewCompany("MSFT")
	if err != nil {
		t.Error(err)
	}

	err = company.Fetch()
	if err != nil {
		t.Error(err)
	}

	market, err := NewMarket(0.02, 0.02, 0.25)
	if err != nil {
		t.Error(err)
	}

	input, err := NewInput(company, 0.15, 0.07, 0.07, 0.12, 0.4, 0.45, 1.17)
	if err != nil {
		t.Error(err)
	}

	err = input.Compute()
	if err != nil {
		t.Error(err)
	}

	output, err := NewOutput(market, input)
	if err != nil {
		t.Error(err)
	}

	err = output.Compute()
	if err != nil {
		t.Error(err)
	}

	t.Logf("Revenue TTM: %f", input.Revenue)
	t.Logf("EBIT TTM: %f", input.EBIT)
	t.Logf("Total Equity: %f", input.TotalEquity)
	t.Logf("Total Debt: %f", input.TotalDebt)
	t.Logf("Total Cash: %f", input.TotalCash)

	t.Logf("Base year revenue growth rate: %f", output.BaseYear.RevenueGrowthRate)
	t.Logf("Base year revenue: %f", output.BaseYear.Revenue)
	t.Logf("Base year EBIT margin: %f", output.BaseYear.EBITMargin)
	t.Logf("Base year EBIT: %f", output.BaseYear.EBIT)
	t.Logf("Base year tax rate: %f", output.BaseYear.TaxRate)
	t.Logf("Base year NOPAT: %f", output.BaseYear.AfterTaxEBIT)

	for i, year := range output.HighGrowthYears {
		t.Logf("Year %d revenue growth rate: %f", i, year.RevenueGrowthRate)
		t.Logf("Year %d revenue: %f", i, year.Revenue)
		t.Logf("Year %d EBIT margin: %f", i, year.EBITMargin)
		t.Logf("Year %d EBIT: %f", i, year.EBIT)
		t.Logf("Year %d tax rate: %f", i, year.TaxRate)
		t.Logf("Year %d NOPAT: %f", i, year.AfterTaxEBIT)
		t.Logf("Year %d reinvestment: %f", i, year.Reinvestment)
		t.Logf("Year %d FCFF: %f", i, year.FCFF)
		t.Logf("Year %d cost of capital: %f", i, year.CostOfCapital)
		t.Logf("Year %d discount factor: %f", i, year.DiscountFactor)
		t.Logf("Year %d PV(FCFF): %f", i, year.PresentValueOfCashFlow)
	}

	offset := len(output.HighGrowthYears)

	for i, year := range output.LowGrowthYears {
		t.Logf("Year %d revenue growth rate: %f", i+offset, year.RevenueGrowthRate)
		t.Logf("Year %d revenue: %f", i+offset, year.Revenue)
		t.Logf("Year %d EBIT margin: %f", i+offset, year.EBITMargin)
		t.Logf("Year %d EBIT: %f", i+offset, year.EBIT)
		t.Logf("Year %d tax rate: %f", i+offset, year.TaxRate)
		t.Logf("Year %d NOPAT: %f", i+offset, year.AfterTaxEBIT)
		t.Logf("Year %d reinvestment: %f", i+offset, year.Reinvestment)
		t.Logf("Year %d FCFF: %f", i+offset, year.FCFF)
		t.Logf("Year %d cost of capital: %f", i+offset, year.CostOfCapital)
		t.Logf("Year %d discount factor: %f", i+offset, year.DiscountFactor)
		t.Logf("Year %d PV(FCFF): %f", i+offset, year.PresentValueOfCashFlow)
	}

	t.Logf("Terminal year revenue growth rate: %f", output.TerminalYear.RevenueGrowthRate)
	t.Logf("Terminal year revenue: %f", output.TerminalYear.Revenue)
	t.Logf("Terminal year EBIT margin: %f", output.TerminalYear.EBITMargin)
	t.Logf("Terminal year EBIT: %f", output.TerminalYear.EBIT)
	t.Logf("Terminal year tax rate: %f", output.TerminalYear.TaxRate)
	t.Logf("Terminal year NOPAT: %f", output.TerminalYear.AfterTaxEBIT)
	t.Logf("Terminal year reinvestment: %f", output.TerminalYear.Reinvestment)
	t.Logf("Terminal year FCFF: %f", output.TerminalYear.FCFF)
	t.Logf("Terminal year cost of capital: %f", output.TerminalYear.CostOfCapital)

	t.Logf("Terminal cash flow: %f", output.TerminalCashFlow)
	t.Logf("Terminal value: %f", output.TerminalValue)
	t.Logf("Present value of terminal value: %f", output.PresentValueOfTerminalValue)
	t.Logf("Present value of cash flows: %f", output.PresentValueOfCashFlow)
	t.Logf("Present value: %f", output.PresentValue)
	t.Logf("Equity value: %f", output.EquityValue)
	t.Logf("Value per share: %f", output.ValuePerShare)
}
