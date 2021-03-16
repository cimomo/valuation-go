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
