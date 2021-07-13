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

	input, err := NewInput(company, 0.15, 0.07, 0.07, 0.2)
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
}
