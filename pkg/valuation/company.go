package valuation

import "github.com/cimomo/alphavantage-go"

// Company defines valuation input data for a company
type Company struct {
	Symbol          string
	Client          *alphavantage.Client
	Overview        *alphavantage.Company
	IncomeStatement *alphavantage.IncomeStatement
	BalanceSheet    *alphavantage.BalanceSheet
	CashFlow        *alphavantage.CashFlow
}

// NewCompany returns a new Company object
func NewCompany(symbol string) (*Company, error) {
	company := Company{
		Symbol: symbol,
	}

	client, err := alphavantage.NewClient("thisshouldnotwork")
	if err != nil {
		return nil, err
	}

	company.Client = client

	return &company, nil
}
