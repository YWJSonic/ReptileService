package legalperson

const (
	CollectionFlagkey   string = "legalperson"
	InfoDataLen                = 19
	Info20171215DataLen        = 16
	Info20141128DataLen        = 12
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
func (R *Result) GetInfos() []interface{} {
	var infos []interface{}

	for _, data := range R.Data {
		if len(data) == InfoDataLen {
			info := ConvertToInfo(data)
			info.Date = R.Date
			infos = append(infos, info)
		} else if len(data) == Info20171215DataLen {
			info := ConvertToInfo20171215(data)
			info.Date = R.Date
			infos = append(infos, info)
		} else if len(data) == Info20141128DataLen {
			info := ConvertToInfo20141128(data)
			info.Date = R.Date
			infos = append(infos, info)
		}

	}
	return infos
}

// GetInfos 將interface 轉成Info
func (R *Result) GetInfo20171215s() []InfoBefore20171215 {
	var infos []InfoBefore20171215

	for _, data := range R.Data {
		info := ConvertToInfo20171215(data)
		info.Date = R.Date
		infos = append(infos, info)
	}
	return infos
}

// GetInfos 將interface 轉成Info
func (R *Result) GetInfo20141128s() []InfoBefore20141128 {
	var infos []InfoBefore20141128

	for _, data := range R.Data {
		info := ConvertToInfo20141128(data)
		info.Date = R.Date
		infos = append(infos, info)
	}
	return infos
}

// Info 12月15以後的結構(不含15號)
type Info struct {
	Date                 string // 日期
	StockCode            string // "證券代號"
	StockName            string // "證券名稱"
	OutCountryBuy        string // "外陸資買進股數(不含外資自營商)"
	OutCountrySell       string // "外陸資賣出股數(不含外資自營商)"
	OutCountryDiff       string // "外陸資買賣超股數(不含外資自營商)" = OutCountryBuy - OutCountrySell
	OutCountryBankBuy    string // "外資自營商買進股數"
	OutCountryBankSell   string // "外資自營商賣出股數"
	OutCountryBankDiff   string // "外資自營商買賣超股數"
	BankEtfBuy           string // "投信買進股數"
	BankEtfSell          string // "投信賣出股數"
	BankEtfDiff          string // "投信買賣超股數" = BankEtfBuy - BankEtfSell
	BankStockDiff        string // "自營商買賣超股數" = BankStockSelfDiff + BankStockHedgingDiff
	BankStockSelfBuy     string // "自營商買進股數(自行買賣)"
	BankStockSelfSell    string // "自營商賣出股數(自行買賣)"
	BankStockSelfDiff    string // "自營商買賣超股數(自行買賣)" = BankStockSelfBuy - BankStockSelfSell
	BankStockHedgingBuy  string // "自營商買進股數(避險)" = BankStockHedgingBuy + OutCountryBankBuy
	BankStockHedgingSell string // "自營商賣出股數(避險)" = BankStockHedgingSell + OutCountryBankSell
	BankStockHedgingDiff string // "自營商買賣超股數(避險)" = BankStockHedgingBuy - BankStockHedgingSell
	TotalDiff            string // "三大法人買賣超股數" = OutCountryDiff + BankEtfDiff + BankStockDiff
}

// 20171215 號前使用(包含15號)
type InfoBefore20171215 struct {
	Date                 string // 日期
	StockCode            string // "證卷代號"
	StockName            string // "證卷名稱"
	OutCountryBuy        string // "外資買進股數"
	OutCountrySell       string // "外資賣出股數"
	OutCountryDiff       string // "外資買賣超股數" = OutCountryBuy - OutCountrySell
	BankEtfBuy           string // "投信買進股數"
	BankEtfSell          string // "投信賣出股數"
	BankEtfDiff          string // "投信買賣超股數" = BankEtfBuy - BankEtfSell
	BankStockDiff        string // "自營商買賣超股數" = BankStockSelfDiff + BankStockHedgingDiff
	BankStockSelfBuy     string // "自營商買進股數(自行買賣)"
	BankStockSelfSell    string // "自營商賣出股數(自行買賣)"
	BankStockSelfDiff    string // "自營商買賣超股數(自行買賣)" = BankStockSelfBuy - BankStockSelfSell
	BankStockHedgingBuy  string // "自營商買進股數(避險)"
	BankStockHedgingSell string // "自營商賣出股數(避險)"
	BankStockHedgingDiff string // "自營商買賣超股數(避險)" = BankStockHedgingBuy - BankStockHedgingSell
	TotalDiff            string // "三大法人買賣超股數" = OutCountryDiff + BankEtfDiff + BankStockDiff
}

// 20141128 號前使用(包含28號)
type InfoBefore20141128 struct {
	Date           string // 日期
	StockCode      string // "證卷代號"
	StockName      string // "證卷名稱"
	OutCountryBuy  string // "外資買進股數"
	OutCountrySell string // "外資賣出股數"
	OutCountryDiff string // "外資買賣超股數" = OutCountryBuy - OutCountrySell
	BankEtfBuy     string // "投信買進股數"
	BankEtfSell    string // "投信賣出股數"
	BankEtfDiff    string // "投信買賣超股數" = BankEtfBuy - BankEtfSell
	BankStockBuy   string // "自營商買進股數"
	BankStockSell  string // "自營商賣出股數"
	BankStockDiff  string // "自營商買賣超股數" = BankStockBuy - BankStockSell
	TotalDiff      string // "三大法人買賣超股數" = OutCountryDiff + BankEtfDiff + BankStockDiff
}
