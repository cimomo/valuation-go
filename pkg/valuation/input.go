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
}

// NewInput returns a new valuation input object
func NewInput(company *Company) (*Input, error) {
	input := Input{
		Company: company,
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

	return nil
}

func (input *Input) getTTM() ([]alphavantage.IncomeStatementReport, error) {
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
	ttm, err := input.getTTM()
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
	ttm, err := input.getTTM()
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
