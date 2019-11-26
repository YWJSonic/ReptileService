package bwibbu

// Result ...
type Result struct {
	Original  string
	StockCode string          // 個股編號
	Stat      string          `json:"stat"`
	Title     string          `json:"title"`
	Fields    []string        `json:"fields"`
	Data      [][]interface{} `json:"data"`
	Notes     []string        `json:"notes"`
}

// GetInfos 將interface 轉成Info
func (R *Result) GetInfos() []Info {
	var infos []Info
	var info Info
	for _, data := range R.Data {
		info = ConvertToInfo(data)
		info.StockCode = R.StockCode
		infos = append(infos, info)
	}

	return infos
}

// Info ...
type Info struct {
	StockCode       string
	Date            string // 日期
	Yield           string // 殖利率(%): 每股股利／收盤價*100%
	DividendYear    string // 股利年度
	PEratio         string // 本益比: 收盤價／每股參考稅後純益
	WorthRatio      string // 股價淨值比: 收盤價／每股參考淨值
	FinancialReport string // 財報 年/季
}
