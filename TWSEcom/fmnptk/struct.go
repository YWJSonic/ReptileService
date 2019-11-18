package fmnptk

// URL https://www.twse.com.tw/exchangeReport/FMNPTK?response=json&stockNo=2409&_=1573430069096

// Result 回傳資料
type Result struct {
	Original  string
	StockCode string          // 個股編號
	Stat      string          `json:"stat"`
	Title     string          `json:"title"`
	Fields    []string        `json:"fields"`
	Fields2   []string        `json:"fields2"`
	Data      [][]interface{} `json:"data"`
	Data2     [][]interface{} `json:"data2"`
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

// Info 個股年成交資訊
type Info struct {
	StockCode       string // 個股編號
	Year            int    // 年度
	StockPrice      string // 成交金額
	StockCount      string // 成交股數
	DealCount       string // 成交筆數
	TopPrice        string // 最高價
	TopPriceDate    string // 日期
	BottomPrice     string // 最低價
	BottomPriceDate string // 日期
	AVGPrice        string // 收盤平均價
}
