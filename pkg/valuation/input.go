package valuation

import (
	"errors"
	"strconv"
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

	return nil
}

func (input *Input) computeRevenue() error {
	quarterly := input.Company.IncomeStatement.QuarterlyReports
	if quarterly == nil {
		return errors.New("No quarterly income statement found")
	}

	if len(quarterly) < 4 {
		return errors.New("Need at least four quarters of results")
	}

	ttm := quarterly[:4]

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
