package valuation

const yearsOfHighGrowth = 5

// OutputYear defines the cash flow calculation of one particular year in the future
type OutputYear struct {
	RevenueGrowthRate      float64
	Revenue                float64
	EBITMargin             float64
	EBIT                   float64
	TaxRate                float64
	AfterTaxEBIT           float64
	Reinvestment           float64
	FCFF                   float64
	CostOfCapital          float64
	PresentValueOfCashFlow float64
}

// Output defines the valuation result
type Output struct {
	Market                      *Market
	Input                       *Input
	BaseYear                    *OutputYear
	OutputYears                 []OutputYear
	TerminalYear                *OutputYear
	TerminalCashFlow            float64
	TerminalValue               float64
	PresentValueOfTerminalValue float64
	PresentValueOfCashFlow      float64
	PresentValue                float64
	EquityValue                 float64
	ValuePerShare               float64
}

// NewOutput returns a new valuation output object
func NewOutput(market *Market, input *Input) (*Output, error) {
	output := Output{
		Market: market,
		Input:  input,
	}

	years := make([]OutputYear, yearsOfHighGrowth*2)

	output.BaseYear = &OutputYear{}
	output.OutputYears = years
	output.TerminalYear = &OutputYear{}

	return &output, nil
}

// Compute calculates the valuation output
func (output *Output) Compute() error {
	err := output.computeBaseYear()
	if err != nil {
		return err
	}

	return nil
}

func (output *Output) computeBaseYear() error {
	baseYear := OutputYear{}
	input := output.Input

	baseYear.Revenue = input.Revenue
	baseYear.EBIT = input.EBIT
	baseYear.TaxRate = input.EffectiveTaxRate

	baseYear.EBITMargin = baseYear.EBIT / baseYear.Revenue
	baseYear.AfterTaxEBIT = baseYear.EBIT * baseYear.TaxRate

	output.BaseYear = &baseYear
	return nil
}
