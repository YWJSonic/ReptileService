package legalperson

import (
	"fmt"

	"github.com/YWJSonic/ReptileService/dbhandle"
	"github.com/YWJSonic/ReptileService/foundation"
	"github.com/YWJSonic/ReptileService/httphandle"
)

// 三大法人買賣金額統計表
// 日表 https://www.twse.com.tw/fund/T86?response=json&date=20210707&selectType=ALL&_=1626618069309
// 月表 https://www.twse.com.tw/fund/TWT47U?response=json&date=20210701&selectType=ALL&_=1626618193051

// GetAlreadyDate ...
func GetAlreadyDate(legalpersonType string) ([]map[string]interface{}, error) {
	return dbhandle.Instance.Getcollectionflag("All"+legalpersonType, CollectionFlagkey)
}

// CopyData ...
func CopyDayData(legalpersonType, date string, cacheTime int64) error {
	result, err := GetDay(date, cacheTime)
	if err != nil {
		return err
	}

	Infos := result.GetInfos()
	if Infos != nil {
		err = dbhandle.Instance.SetLegalperson(CollectionFlagkey, date, Infos)
		if err != nil {
			return err
		}
	}

	err = dbhandle.Instance.Setcollectionflag("All", CollectionFlagkey, date)
	if err != nil {
		return err
	}

	return nil
}

// CopyData ...
func CopyMonthData(date string, cacheTime int64) error {
	result, err := GetDay(date, cacheTime)
	if err != nil {
		return err
	}

	Infos := result.GetInfos()
	if Infos != nil {
		err = dbhandle.Instance.SetLegalperson(CollectionFlagkey+"_Month", date, Infos)
		if err != nil {
			return err
		}
	}

	err = dbhandle.Instance.Setcollectionflag("All_Month", CollectionFlagkey, date)
	if err != nil {
		return err
	}

	return nil
}

// GetDay data
// Num: 股票代號
func GetDay(date string, cacheTime int64) (*Result, error) {
	data := &Result{}

	result := httphandle.Instans.HTTPGetRequest(fmt.Sprintf("https://www.twse.com.tw/fund/T86?response=json&date=%s&selectType=ALL&_=%d", date, cacheTime), nil)
	err := foundation.ByteToStruct(result, &data)
	if err != nil {
		return nil, err
	}
	data.Original = string(result)
	return data, nil
}

// GetMonth data
// Num: 股票代號
func GetMonth(date string, cacheTime int64) (*Result, error) {
	data := &Result{}

	result := httphandle.Instans.HTTPGetRequest(fmt.Sprintf("https://www.twse.com.tw/fund/TWT47U?response=json&date=%s&selectType=ALL&_=%d", date, cacheTime), nil)
	err := foundation.ByteToStruct(result, &data)
	if err != nil {
		return nil, err
	}
	data.Original = string(result)
	return data, nil
}

// ConvertToInfo ...
func ConvertToInfo(Data []interface{}) Info {
	var info Info
	info.StockCode = foundation.InterfaceToString(Data[0])
	info.StockName = foundation.InterfaceToString(Data[1])
	info.OutCountryBuy = foundation.InterfaceToString(Data[2])
	info.OutCountrySell = foundation.InterfaceToString(Data[3])
	info.OutCountryDiff = foundation.InterfaceToString(Data[4])
	info.OutCountryBankBuy = foundation.InterfaceToString(Data[5])
	info.OutCountryBankSell = foundation.InterfaceToString(Data[6])
	info.OutCountryBankDiff = foundation.InterfaceToString(Data[7])
	info.BankEtfBuy = foundation.InterfaceToString(Data[8])
	info.BankEtfSell = foundation.InterfaceToString(Data[9])
	info.BankEtfDiff = foundation.InterfaceToString(Data[10])
	info.BankStockDiff = foundation.InterfaceToString(Data[11])
	info.BankStockSelfBuy = foundation.InterfaceToString(Data[12])
	info.BankStockSelfSell = foundation.InterfaceToString(Data[13])
	info.BankStockSelfDiff = foundation.InterfaceToString(Data[14])
	info.BankStockHedgingBuy = foundation.InterfaceToString(Data[15])
	info.BankStockHedgingSell = foundation.InterfaceToString(Data[16])
	info.BankStockHedgingDiff = foundation.InterfaceToString(Data[17])
	info.TotalDiff = foundation.InterfaceToString(Data[18])
	return info
}

// ConvertToInfo ...
func ConvertToInfo20171215(data []interface{}) InfoBefore20171215 {
	var info InfoBefore20171215
	info.StockCode = foundation.InterfaceToString(data[0])
	info.StockName = foundation.InterfaceToString(data[1])
	info.OutCountryBuy = foundation.InterfaceToString(data[2])
	info.OutCountrySell = foundation.InterfaceToString(data[3])
	info.OutCountryDiff = foundation.InterfaceToString(data[4])
	info.BankEtfBuy = foundation.InterfaceToString(data[5])
	info.BankEtfSell = foundation.InterfaceToString(data[6])
	info.BankEtfDiff = foundation.InterfaceToString(data[7])
	info.BankStockDiff = foundation.InterfaceToString(data[8])
	info.BankStockSelfBuy = foundation.InterfaceToString(data[9])
	info.BankStockSelfSell = foundation.InterfaceToString(data[10])
	info.BankStockSelfDiff = foundation.InterfaceToString(data[11])
	info.BankStockHedgingBuy = foundation.InterfaceToString(data[12])
	info.BankStockHedgingSell = foundation.InterfaceToString(data[13])
	info.BankStockHedgingDiff = foundation.InterfaceToString(data[14])
	info.TotalDiff = foundation.InterfaceToString(data[15])

	return info
}

// ConvertToInfo ...
func ConvertToInfo20141128(data []interface{}) InfoBefore20141128 {
	var info InfoBefore20141128
	info.StockCode = foundation.InterfaceToString(data[0])
	info.StockName = foundation.InterfaceToString(data[1])
	info.OutCountryBuy = foundation.InterfaceToString(data[2])
	info.OutCountrySell = foundation.InterfaceToString(data[3])
	info.OutCountryDiff = foundation.InterfaceToString(data[4])
	info.BankEtfBuy = foundation.InterfaceToString(data[5])
	info.BankEtfSell = foundation.InterfaceToString(data[6])
	info.BankEtfDiff = foundation.InterfaceToString(data[7])
	info.BankStockBuy = foundation.InterfaceToString(data[8])
	info.BankStockSell = foundation.InterfaceToString(data[9])
	info.BankStockDiff = foundation.InterfaceToString(data[10])
	info.TotalDiff = foundation.InterfaceToString(data[11])
	return info
}
