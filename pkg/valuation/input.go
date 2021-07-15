package valuation

import (
	"errors"
	"strconv"

	"github.com/cimomo/alphavantage-go"
)

// Input defines the company specific input data for the valuation
type Input struct {
	Company               *Company
	Revenue               float64
	EBIT                  float64
	TotalEquity           float64
	TotalDebt             float64
	TotalCash             float64
	EffectiveTaxRate      float64
	CostOfCapital         float64
	TerminalCostOfCapital float64
	RevenueGrowthRate     float64
	StartingEBITMargin    float64
	TerminalEBITMargin    float64
	SalesToCapital        float64
}

// NewInput returns a new valuation input object
func NewInput(company *Company, effectiveTaxRate float64, costOfCapital float64, terminalCostOfcapital float64, revenueGrowthRate float64, startingEBITMargin float64, terminalEBITMargin float64, salesToCapital float64) (*Input, error) {
	input := Input{
		Company:               company,
		EffectiveTaxRate:      effectiveTaxRate,
		CostOfCapital:         costOfCapital,
		TerminalCostOfCapital: terminalCostOfcapital,
		RevenueGrowthRate:     revenueGrowthRate,
		StartingEBITMargin:    startingEBITMargin,
		TerminalEBITMargin:    terminalEBITMargin,
		SalesToCapital:        salesToCapital,
	}

	return &input, nil
}

// Compute computes the valuation input from company fundamentals data
func (input *Input) Compute() error {
	err := input.computeRevenue()
	if err != nil {
		return err
	}

	err = input.computeEBIT()
	if err != nil {
		return err
	}

	err = input.computeTotalEquity()
	if err != nil {
		return err
	}

	err = input.computeTotalDebt()
	if err != nil {
		return err
	}

	err = input.computeTotalCash()
	if err != nil {
		return err
	}

	return nil
}

func (input *Input) getQuarterlyIncomeStatementsTTM() ([]alphavantage.IncomeStatementReport, error) {
	quarterly := input.Company.IncomeStatement.QuarterlyReports
	if quarterly == nil {
		return nil, errors.New("No quarterly income statement found")
	}

	if len(quarterly) < 4 {
		return nil, errors.New("Need at least four quarters of results")
	}

	ttm := quarterly[:4]

	return ttm, nil
}

func (input *Input) computeRevenue() error {
	ttm, err := input.getQuarterlyIncomeStatementsTTM()
	if err != nil {
		return err
	}

	revenue := 0.0
	for _, v := range ttm {
		r, err := strconv.ParseFloat(v.TotalRevenue, 64)
		if err != nil {
			return err
		}
		revenue += r
	}

	input.Revenue = revenue

	return nil
}

func (input *Input) computeEBIT() error {
	ttm, err := input.getQuarterlyIncomeStatementsTTM()
	if err != nil {
		return err
	}

	ebit := 0.0
	for _, v := range ttm {
		r, err := strconv.ParseFloat(v.EBIT, 64)
		if err != nil {
			return err
		}
		ebit += r
	}

	input.EBIT = ebit

	return nil
}

func (input *Input) computeTotalEquity() error {
	quarterly := input.Company.BalanceSheet.QuarterlyReports
	if quarterly == nil {
		return errors.New("No quarterly balance sheet found")
	}

	if len(quarterly) == 0 {
		return errors.New("Need at least one quarter of results")
	}

	balanceSheet := quarterly[0]

	totalEquity, err := strconv.ParseFloat(balanceSheet.TotalShareholderEquity, 64)
	if err != nil {
		return err
	}

	input.TotalEquity = totalEquity

	return nil
}

func (input *Input) computeTotalDebt() error {
	quarterly := input.Company.BalanceSheet.QuarterlyReports
	if quarterly == nil {
		return errors.New("No quarterly balance sheet found")
	}

	if len(quarterly) == 0 {
		return errors.New("Need at least one quarter of results")
	}

	balanceSheet := quarterly[0]

	// ShortTermDebt includes short term borrowing and current portion of long term debt
	shortTermDebt, err := strconv.ParseFloat(balanceSheet.ShortTermDebt, 64)
	if err != nil {
		return err
	}

	longTermDebt, err := strconv.ParseFloat(balanceSheet.LongTermDebtNoncurrent, 64)
	if err != nil {
		return err
	}

	capitalLease, err := strconv.ParseFloat(balanceSheet.CapitalLeaseObligations, 64)
	if err != nil {
		return err
	}

	input.TotalDebt = shortTermDebt + longTermDebt + capitalLease

	return nil
}

func (input *Input) computeTotalCash() error {
	quarterly := input.Company.BalanceSheet.QuarterlyReports
	if quarterly == nil {
		return errors.New("No quarterly balance sheet found")
	}

	if len(quarterly) == 0 {
		return errors.New("Need at least one quarter of results")
	}

	balanceSheet := quarterly[0]

	totalCash, err := strconv.ParseFloat(balanceSheet.CashAndShortTermInvestments, 64)
	if err != nil {
		return err
	}

	input.TotalCash = totalCash

	return nil
}
