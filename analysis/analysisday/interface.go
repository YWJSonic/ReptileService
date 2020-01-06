package analysisday

// GetPriceDeatil ...
type GetPriceDeatil interface {
	StockCode() string
	StockNmae() string
	Price() string
	Count() int64
	Time() string
	Date() string
	OpenPrice() string
	YesterdayPrice() string
	GetDiffStr() (string, error)
	GetDiffPreStr() (string, error)
	IsPriceUp() bool
}
