package bwibbu

import "github.com/YWJSonic/ReptileService/foundation"

const (
	collectionFlagkey string = "Bwibbu_Month"
)

// Result ...
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
		info.Date = R.Date
		info.StockCode = R.StockCode
		infos = append(infos, info)
	}

	return infos
}

// ConvertToInfo ...
func ConvertToInfo(Data []interface{}) Info {
	var info Info
	if len(Data) > 4 {
		info.Date = foundation.InterfaceToString(Data[0])
		info.YieldRate = foundation.InterfaceToString(Data[1])
		if Data[2] != "" {
			info.DividendYear = foundation.InterfaceToInt(Data[2])
		} else {
			info.DividendYear = 0
		}
		info.PeRatio = foundation.InterfaceToString(Data[3])
		info.WorthRatio = foundation.InterfaceToString(Data[4])
		info.FinancialReport = foundation.InterfaceToString(Data[5])
	} else {
		info.Date = foundation.InterfaceToString(Data[0])
		info.PeRatio = foundation.InterfaceToString(Data[1])
		info.YieldRate = foundation.InterfaceToString(Data[2])
		info.WorthRatio = foundation.InterfaceToString(Data[3])
	}
	return info
}

// Info 2017 3月以後的結構(不含3月)
type Info struct {
	StockCode       string
	Date            string // 日期
	YieldRate       string // 殖利率(%): 每股股利／收盤價*100%
	DividendYear    int    // 股利年度
	PeRatio         string // 本益比: 收盤價／每股參考稅後純益
	WorthRatio      string // 股價淨值比: 收盤價／每股參考淨值
	FinancialReport string // 財報 年/季
}
