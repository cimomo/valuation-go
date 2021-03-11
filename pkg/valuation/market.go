package valuation

// Market defines the broader market and macro data used as valuation input
type Market struct {
	RiskFreeRate         float64
	TerminalRiskFreeRate float64
	MarginalTaxRate      float64
}
