package analysisday

// NewAnalysisDay return new create info
func NewAnalysisDay() *Info {
	info := &Info{}
	info.PriceDeatil = make(map[string]int64)
	info.PriceDetailLog = make([]string, 0)
	return info
}
