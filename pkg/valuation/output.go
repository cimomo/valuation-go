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
	DiscountFactor         float64
	PresentValueOfCashFlow float64
}

// Output defines the valuation result
type Output struct {
	Market                      *Market
	Input                       *Input
	BaseYear                    *OutputYear
	HighGrowthYears             []OutputYear
	LowGrowthYears              []OutputYear
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

	highGrowthYears := make([]OutputYear, yearsOfHighGrowth)
	lowGrowthYears := make([]OutputYear, yearsOfHighGrowth)

	output.BaseYear = &OutputYear{}
	output.HighGrowthYears = highGrowthYears
	output.LowGrowthYears = lowGrowthYears
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

func (output *Output) computeYearInGrowth(previousYear *OutputYear, revenueGrowthRate float64, ebitMargin float64, taxRate float64, salesToCapital float64, discountFactor float64) (*OutputYear, error) {
	result := OutputYear{}

	result.RevenueGrowthRate = revenueGrowthRate
	result.Revenue = previousYear.Revenue * (1 + revenueGrowthRate)
	result.EBITMargin = ebitMargin
	result.EBIT = result.Revenue * ebitMargin
	result.AfterTaxEBIT = result.EBIT * (1 - taxRate)
	result.Reinvestment = (result.Revenue - previousYear.Revenue) / salesToCapital
	result.FCFF = result.AfterTaxEBIT - result.Reinvestment
	result.CostOfCapital = output.Input.CostOfCapital
	result.DiscountFactor = discountFactor
	result.PresentValueOfCashFlow = result.FCFF * discountFactor

	return &result, nil
}
