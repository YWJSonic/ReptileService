package analysisday

// GetPriceDeatil ...
type GetPriceDeatil interface {
	Price() string
	Count() int64
	Time() string
	Date() string
	OpenPrice() string
}
