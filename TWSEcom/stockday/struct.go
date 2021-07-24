package stockday

// URL https://www.twse.com.tw/exchangeReport/STOCK_DAY?response=json&date=20190801&stockNo=2409&_=1573440324018

const (
	CollectionFlagkey = "Day"
)

// Result 個股日交資訊
type Result struct {
	Oringial  string
	StockCode string
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

// Info 個股日成交資訊
type Info struct {
	StockCode   string
	Year        int    // 年
	Month       int    // 月
	Day         int    // 日
	StockPrice  string // 成交金額
	StockCount  string // 成交股數
	DealCount   string // 成交筆數
	OpenPrice   string // 開盤價
	ClosePrice  string // 收盤價
	TopPrice    string // 最高價
	BottomPrice string // 最低價
	DiffPrice   string // 漲跌價差
}
