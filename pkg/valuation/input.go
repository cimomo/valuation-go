package valuation

// Input defines the company specific input data for the valuation
type Input struct {
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
