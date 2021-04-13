package valuation

// Market defines the broader market and macro data used as valuation input
type Market struct {
	RiskFreeRate         float64
	TerminalRiskFreeRate float64
	MarginalTaxRate      float64
}

// NewMarket returns a new market object
func NewMarket(riskFreeRate float64, terminalRiskFreeRate float64, marginalTaxRate float64) (*Market, error) {
	market := Market{
		RiskFreeRate:         riskFreeRate,
		TerminalRiskFreeRate: terminalRiskFreeRate,
		MarginalTaxRate:      marginalTaxRate,
	}

	return &market, nil
}
