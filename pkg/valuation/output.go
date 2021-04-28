package valuation

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
	TerminalCashFlow            float64
	TerminalValue               float64
	PresentValueOfTerminalValue float64
	PresentValueOfCashFlow      float64
	PresentValue                float64
	EquityValue                 float64
	ValuePerShare               float64
}
