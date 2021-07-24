package legalperson

const (
	CollectionFlagkey string = "legalperson"
)

// Result ...
type Result struct {
	Original string
	Data     [][]interface{} `json:"data"`
	Date     string          `json:"date"`
	Fields   []string        `json:"fields"`
	Notes    []string        `json:"notes"`
	Stat     string          `json:"stat"`
	Title    string          `json:"title"`
}

// GetInfos 將interface 轉成Info
func (R *Result) GetInfos() []Info {
	var infos []Info

	for _, data := range R.Data {
		info := ConvertToInfo(data)
		info.Date = R.Date
		infos = append(infos, info)
	}
	return infos
}

// Info 2017 3月以後的結構(不含3月)
type Info struct {
	Date           string // 日期
	StockCode      string // 證卷代號
	StockName      string // 證卷名稱
	OutCountryBuy  string // "外資買進股數"
	OutCountrySell string // "外資賣出股數"
	OutCountryDiff string // "外資買賣超股數" = Buy - Sell
	BankEtfBuy     string // "投信買進股數"
	BankEtfSell    string // "投信賣出股數"
	BankEtfDiff    string // "投信買賣超股數"
	BankStockBuy   string // "自營商買進股數"
	BankStockSell  string // "自營商賣出股數"
	BankStockDiff  string // "自營商買賣超股數"
	TotalDiff      string // "三大法人買賣超股數" = OutCountryDiff + BankEtfDiff + BankStockDiff
}
