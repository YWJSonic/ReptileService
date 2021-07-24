package fmsrfk

// URL https://www.twse.com.tw/exchangeReport/FMSRFK?response=json&date=20020101&stockNo=2409&_=1573433055138

const (
	CollectionFlagkey = "Month"
)

// Result 回傳資料
type Result struct {
	Original  string
	StockCode string          // 個股編號
	Date      string          `json:"date"`
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

// Info 個股月成交資訊
type Info struct {
	StockCode       string
	Year            int    // "年度"
	Month           int    // "月份"
	StockPrice      string // "成交金額(A)"
	StockCount      string // "成交股數(B)"
	DealCount       string // "成交筆數"
	TopPrice        string // "最高價"
	BottomPrice     string // "最低價"
	WeightsAvgPrice string // "加權(A/B)平均價"
	Turnover        string // "週轉率(%)"
}
