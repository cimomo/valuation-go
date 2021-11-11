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
	baseYear, err := output.computeBaseYear()
	if err != nil {
		return err
	}

	output.BaseYear = baseYear

	market := output.Market
	input := output.Input
	prevYear := baseYear
	startingRevenueGrowthRate := input.RevenueGrowthRate
	terminalRevenueGrowthRate := market.TerminalRiskFreeRate
	revenueGrowthRate := startingRevenueGrowthRate
	ebitMargin := input.StartingEBITMargin
	startingEBITMargin := input.StartingEBITMargin
	terminalEBITMargin := input.TerminalEBITMargin
	taxRate := input.EffectiveTaxRate
	costOfCapital := input.CostOfCapital
	startingCostOfCapital := input.CostOfCapital
	terminalCostOfCapital := input.TerminalCostOfCapital
	discountFactor := 1 / (1 + costOfCapital)

	for i := 0; i < yearsOfHighGrowth; i++ {

		if i != 0 {
			ebitMargin = terminalEBITMargin - ((terminalEBITMargin-startingEBITMargin)/float64((yearsOfHighGrowth*2-1)))*float64((yearsOfHighGrowth*2-i-1))
		}

		year, err := output.computeYearInGrowth(prevYear, revenueGrowthRate, ebitMargin, taxRate, costOfCapital, discountFactor)
		if err != nil {
			return err
		}

		output.HighGrowthYears[i] = *year
		prevYear = year
		discountFactor = discountFactor * (1 / (1 + costOfCapital))
	}

	for i := 0; i < yearsOfHighGrowth; i++ {
		revenueGrowthRate = startingRevenueGrowthRate - ((startingRevenueGrowthRate-terminalRevenueGrowthRate)/float64(yearsOfHighGrowth))*float64(i+1)

		ebitMargin = terminalEBITMargin - ((terminalEBITMargin-startingEBITMargin)/float64((yearsOfHighGrowth*2-1)))*float64((yearsOfHighGrowth-i-1))

		taxRate = input.EffectiveTaxRate + ((market.MarginalTaxRate-input.EffectiveTaxRate)/float64(yearsOfHighGrowth))*float64(i+1)

		costOfCapital = startingCostOfCapital - (startingCostOfCapital-terminalCostOfCapital)/float64(yearsOfHighGrowth)*float64(i+1)

		year, err := output.computeYearInGrowth(prevYear, revenueGrowthRate, ebitMargin, taxRate, costOfCapital, discountFactor)
		if err != nil {
			return err
		}

		output.LowGrowthYears[i] = *year
		prevYear = year
		discountFactor = discountFactor * (1 / (1 + costOfCapital))
	}

	terminalYear, err := output.computeTerminalYear()
	if err != nil {
		return err
	}

	output.TerminalYear = terminalYear

	return nil
}

func (output *Output) computeBaseYear() (*OutputYear, error) {
	baseYear := OutputYear{}
	input := output.Input

	baseYear.Revenue = input.Revenue
	baseYear.EBIT = input.EBIT
	baseYear.TaxRate = input.EffectiveTaxRate

	baseYear.EBITMargin = baseYear.EBIT / baseYear.Revenue

	if baseYear.EBIT > 0 {
		baseYear.AfterTaxEBIT = baseYear.EBIT * (1 - baseYear.TaxRate)
	} else {
		baseYear.AfterTaxEBIT = baseYear.EBIT
	}

	return &baseYear, nil
}

func (output *Output) computeYearInGrowth(previousYear *OutputYear, revenueGrowthRate float64, ebitMargin float64, taxRate float64, costOfCapital float64, discountFactor float64) (*OutputYear, error) {
	result := OutputYear{}

	result.RevenueGrowthRate = revenueGrowthRate
	result.Revenue = previousYear.Revenue * (1 + revenueGrowthRate)
	result.EBITMargin = ebitMargin
	result.EBIT = result.Revenue * ebitMargin
	result.TaxRate = taxRate
	result.AfterTaxEBIT = result.EBIT * (1 - taxRate)
	result.Reinvestment = (result.Revenue - previousYear.Revenue) / output.Input.SalesToCapital
	result.FCFF = result.AfterTaxEBIT - result.Reinvestment
	result.CostOfCapital = costOfCapital
	result.DiscountFactor = discountFactor
	result.PresentValueOfCashFlow = result.FCFF * discountFactor

	return &result, nil
}

func (output *Output) computeTerminalYear() (*OutputYear, error) {
	result := OutputYear{}

	market := output.Market
	input := output.Input
	previousYear := output.HighGrowthYears[yearsOfHighGrowth-1]

	result.RevenueGrowthRate = market.TerminalRiskFreeRate
	result.Revenue = previousYear.Revenue * (1 + result.RevenueGrowthRate)
	result.EBITMargin = input.TerminalEBITMargin
	result.EBIT = result.Revenue * result.EBITMargin
	result.TaxRate = market.MarginalTaxRate
	result.AfterTaxEBIT = result.EBIT * (1 - result.TaxRate)

	if result.RevenueGrowthRate > 0 {
		result.Reinvestment = result.AfterTaxEBIT * (result.RevenueGrowthRate / input.TerminalCostOfCapital)
	} else {
		result.Reinvestment = 0
	}

	result.FCFF = result.AfterTaxEBIT - result.Reinvestment
	result.CostOfCapital = input.TerminalCostOfCapital
	result.DiscountFactor = previousYear.DiscountFactor

	return &result, nil
}
