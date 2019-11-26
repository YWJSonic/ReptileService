package analysisday

// NewAnalysisDay return new create info
func NewAnalysisDay() *Info {
	info := &Info{}
	info.PriceDeatil = make(map[string]int64)
	info.PriceDetailLog = make([]string, 0)
	info.countLogLimit = 500000
	info.countDeatilTop = 5
	info.CountDeatil = make(map[string]*CountDeatilInfo)
	return info
}
