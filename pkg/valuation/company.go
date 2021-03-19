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

// Fetch retrieves company data from financial data provider
func (company *Company) Fetch() error {
	overview, err := company.Client.GetCompanyOverview(company.Symbol)
	if err != nil {
		return err
	}

	incomeStatement, err := company.Client.GetIncomeStatement(company.Symbol)
	if err != nil {
		return err
	}

	balanceSheet, err := company.Client.GetBalanceSheet(company.Symbol)
	if err != nil {
		return err
	}

	cashFlow, err := company.Client.GetCashFlow(company.Symbol)
	if err != nil {
		return err
	}

	company.Overview = overview
	company.IncomeStatement = incomeStatement
	company.BalanceSheet = balanceSheet
	company.CashFlow = cashFlow

	return nil
}
