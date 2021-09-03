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

	input, err := NewInput(company, 0.15, 0.07, 0.07, 0.2, 0.4, 0.45, 2.0)
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

	t.Logf("Base year revenue: %f", output.BaseYear.Revenue)
	t.Logf("Base year EBIT margin: %f", output.BaseYear.EBITMargin)
	t.Logf("Base year EBIT: %f", output.BaseYear.EBIT)
	t.Logf("Base year tax rate: %f", output.BaseYear.TaxRate)
	t.Logf("Base year NOPAT: %f", output.BaseYear.AfterTaxEBIT)

	for i, year := range output.HighGrowthYears {
		t.Logf("Year %d revenue: %f", i, year.Revenue)
		t.Logf("Year %d EBIT margin: %f", i, year.EBITMargin)
		t.Logf("Year %d EBIT: %f", i, year.EBIT)
		t.Logf("Year %d tax rate: %f", i, year.TaxRate)
		t.Logf("Year %d NOPAT: %f", i, year.AfterTaxEBIT)
	}
}
