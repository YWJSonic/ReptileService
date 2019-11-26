package bwibbu

import (
	"fmt"
	"time"

	"github.com/YWJSonic/ReptileService/foundation"
	"github.com/YWJSonic/ReptileService/handlehttp"
)

// CopyData ...
func CopyData(Num string) {
	// result, err := Get(Num)
	// if err != nil {
	// 	panic((err))
	// }
	// Infos := result.GetInfos()
	// for _, info := range Infos {
	// err = handledb.Setstockyear(info.StockCode, info.Year, info.StockCount, info.StockPrice, info.DealCount, info.TopPrice, info.TopPriceDate, info.BottomPrice, info.BottomPriceDate, info.AVGPrice)
	// if err != nil {
	// 	panic(err)
	// }
	// }
}

// Get data
// Num: 股票代號
func Get(Num, Date string) (*Result, error) {
	data := &Result{}
	result := handlehttp.HTTPGetRequest(handlehttp.ConnectPool(), fmt.Sprintf("https://www.twse.com.tw/exchangeReport/BWIBBU?response=json&date=%s&stockNo=%s&_=%d", Date, Num, time.Now().Unix()*1000), nil)
	err := foundation.ByteToStruct(result, &data)
	if err != nil {
		return nil, err
	}
	data.StockCode = Num
	data.Original = string(result)
	return data, nil
}

// ConvertToInfo ...
func ConvertToInfo(Data []interface{}) Info {
	var info Info
	info.Date = foundation.InterfaceToString(Data[0])
	info.Yield = foundation.InterfaceToString(Data[1])
	info.DividendYear = foundation.InterfaceToString(Data[2])
	info.PEratio = foundation.InterfaceToString(Data[3])
	info.WorthRatio = foundation.InterfaceToString(Data[4])
	info.FinancialReport = foundation.InterfaceToString(Data[5])
	return info
}
