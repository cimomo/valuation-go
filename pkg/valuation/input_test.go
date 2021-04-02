package valuation

import "testing"

func TestComputeInput(t *testing.T) {
	company, err := NewCompany("MSFT")
	if err != nil {
		t.Error(err)
	}

	err = company.Fetch()
	if err != nil {
		t.Error(err)
	}

	input, err := NewInput(company)
	if err != nil {
		t.Error(err)
	}

	err = input.Compute()
	if err != nil {
		t.Error(err)
	}

	t.Logf("Revenue TTM: %f", input.Revenue)
	t.Logf("EBIT TTM: %f", input.EBIT)
	t.Logf("Total Equity: %f", input.TotalEquity)
}
