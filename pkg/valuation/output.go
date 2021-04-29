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
	OutputYears                 []OutputYear
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

	years := make([]OutputYear, yearsOfHighGrowth*2+2)

	output.OutputYears = years

	return &output, nil
}
