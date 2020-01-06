package fmsrfk

import (
	"errors"
	"fmt"
	"time"

	foundation "github.com/YWJSonic/ReptileService/foundation"
	"github.com/YWJSonic/ReptileService/handledb"
	handlehttp "github.com/YWJSonic/ReptileService/handlehttp"
)

// CopyData ...
func CopyData(Num string, Date string) error {
	result, err := Get(Num, Date)
	if err != nil {
		return err
	}
	Infos := result.GetInfos()
	for _, info := range Infos {
		err = handledb.Setstockmonth(info.StockCode, info.Year, info.Month, info.StockPrice, info.StockCount, info.DealCount, info.WeightsAvgPrice, info.TopPrice, info.BottomPrice, info.Turnover)
		if err != nil {
			return err
		}
	}
	err = handledb.Setcollectionflag(Num, "Month", Date[0:len(Date)-2])
	if err != nil {
		return err
	}
	return nil
}

// Get data
// Num: 股票代號
// Date: 查詢日期 20020101
func Get(Num string, Date string) (*Result, error) {
	data := &Result{}
	result := handlehttp.HTTPGetRequest(handlehttp.ConnectPool(), fmt.Sprintf("https://www.twse.com.tw/exchangeReport/FMSRFK?response=json&date=%s&stockNo=%s&_=%d", Date, Num, time.Now().Unix()*1000), nil)
	err := foundation.ByteToStruct(result, &data)
	if err != nil {
		return nil, err
	}
	data.StockCode = Num
	data.Original = string(result)
	if data.Stat != "OK" {
		return nil, errors.New(data.Stat)
	}
	return data, nil
}

// ConvertToInfo ...
func ConvertToInfo(Data []interface{}) Info {
	var info Info

	info.Year = foundation.InterfaceToInt(Data[0])
	info.Month = foundation.InterfaceToInt(Data[1])
	info.TopPrice = foundation.InterfaceToString(Data[2])
	info.BottomPrice = foundation.InterfaceToString(Data[3])
	info.WeightsAvgPrice = foundation.InterfaceToString(Data[4])
	info.DealCount = foundation.InterfaceToString(Data[5])
	info.StockPrice = foundation.InterfaceToString(Data[6])
	info.StockCount = foundation.InterfaceToString(Data[7])
	info.Turnover = foundation.InterfaceToString(Data[8])
	return info
}

// GetAlreadyDate ...
func GetAlreadyDate(StockCode string) ([]map[string]interface{}, error) {
	return handledb.Getcollectionflag(StockCode, "Month")
}
