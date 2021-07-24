package bwibbu

import (
	"fmt"

	"github.com/YWJSonic/ReptileService/dbhandle"
	"github.com/YWJSonic/ReptileService/foundation"
	"github.com/YWJSonic/ReptileService/httphandle"
)

// 個股日本益比、殖利率及股價淨值比
// https://www.twse.com.tw/exchangeReport/BWIBBU?response=json&date=20210718&stockNo=2449&_=1626579187220

// GetAlreadyDate ...
func GetAlreadyDate(StockCode string) ([]map[string]interface{}, error) {
	return dbhandle.Instance.Getcollectionflag(StockCode, CollectionFlagkey)
}

// CopyData ...
func CopyData(stockCode, date string, cacheTime int64) error {
	result, err := Get(stockCode, date, cacheTime)
	if err != nil {
		return err
	}

	Infos := result.GetInfos()
	for _, info := range Infos {
		err = dbhandle.Instance.SetBwibbu(info.StockCode, info.Date, info.YieldRate, info.DividendYear, info.PeRatio, info.WorthRatio, info.FinancialReport)
		if err != nil {
			return err
		}

	}

	err = dbhandle.Instance.Setcollectionflag(stockCode, CollectionFlagkey, date)
	if err != nil {
		return err
	}
	return nil
}

// Get data
// Num: 股票代號
func Get(stockCode, date string, cacheTime int64) (*Result, error) {
	data := &Result{}
	result := httphandle.Instans.HTTPGetRequest(fmt.Sprintf("https://www.twse.com.tw/exchangeReport/BWIBBU?response=json&date=%s&stockNo=%s&_=%d", date, stockCode, cacheTime), nil)
	err := foundation.ByteToStruct(result, &data)
	if err != nil {
		return nil, err
	}
	data.StockCode = stockCode
	data.Original = string(result)
	return data, nil
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
