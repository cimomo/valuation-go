package valuation

import "testing"

func TestFetchCompany(t *testing.T) {
	company, err := NewCompany("IBM")
	if err != nil {
		t.Error(err)
	}

	err = company.Fetch()
	if err != nil {
		t.Error(err)
	}

	t.Logf("Number of annual reports: %d", len(company.IncomeStatement.AnnualReports))
	t.Logf("Number of quarterly reports: %d", len(company.IncomeStatement.QuarterlyReports))
}
