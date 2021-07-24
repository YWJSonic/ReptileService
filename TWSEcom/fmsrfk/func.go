package fmsrfk

import (
	"errors"
	"fmt"
	"time"

	"github.com/YWJSonic/ReptileService/dbhandle"
	foundation "github.com/YWJSonic/ReptileService/foundation"
	"github.com/YWJSonic/ReptileService/httphandle"
)

// GetAlreadyDate ...
func GetAlreadyDate(StockCode string) ([]map[string]interface{}, error) {
	return dbhandle.Instance.Getcollectionflag(StockCode, "Month")
}

// CopyData ...
func CopyData(stockCode string, date string, cacheTime int64) error {
	result, err := Get(stockCode, date, cacheTime)
	if err != nil {
		return err
	}
	Infos := result.GetInfos()
	for _, info := range Infos {
		err = dbhandle.Instance.Setstockmonth(info.StockCode, info.Year, info.Month, info.StockPrice, info.StockCount, info.DealCount, info.WeightsAvgPrice, info.TopPrice, info.BottomPrice, info.Turnover)
		if err != nil {
			return err
		}
	}

	// 同一年份需重新確認
	if time.Now().Format("2006") != date[0:len(date)-2] {
		err = dbhandle.Instance.Setcollectionflag(stockCode, "Month", date[0:len(date)-2])
		if err != nil {
			return err
		}
	}
	return nil
}

// Get data
// Num: 股票代號
// Date: 查詢日期 20020101
func Get(stockCode string, date string, cacheTime int64) (*Result, error) {
	data := &Result{}
	result := httphandle.Instans.HTTPGetRequest(fmt.Sprintf("https://www.twse.com.tw/exchangeReport/FMSRFK?response=json&date=%s&stockNo=%s&_=%d", date, stockCode, cacheTime), nil)
	err := foundation.ByteToStruct(result, &data)
	if err != nil {
		return nil, err
	}
	data.StockCode = stockCode
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
