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
func GetAlreadyDate() ([]map[string]interface{}, error) {
	return dbhandle.Instance.Getcollectionflag("All", CollectionFlagkey)
}

// CopyData ...
func CopyData(date string, cacheTime int64) error {
	result, err := Get(date, cacheTime)
	if err != nil {
		return err
	}

	Infos := result.GetInfos()
	for _, info := range Infos {
		err = dbhandle.Instance.SetLegalperson(info.Date, info.StockCode, info.StockName, info.OutCountryBuy, info.OutCountrySell, info.OutCountryDiff, info.BankEtfBuy, info.BankEtfSell, info.BankEtfDiff, info.BankStockBuy, info.BankStockSell, info.BankStockDiff, info.TotalDiff)
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

// Get data
// Num: 股票代號
func Get(date string, cacheTime int64) (*Result, error) {
	data := &Result{}
	result := httphandle.Instans.HTTPGetRequest(fmt.Sprintf("https://www.twse.com.tw/fund/T86?response=json&date=%s&selectType=ALL&_=%d", date, cacheTime), nil)

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
	info.BankEtfBuy = foundation.InterfaceToString(Data[5])
	info.BankEtfSell = foundation.InterfaceToString(Data[6])
	info.BankEtfDiff = foundation.InterfaceToString(Data[7])
	info.BankStockBuy = foundation.InterfaceToString(Data[8])
	info.BankStockSell = foundation.InterfaceToString(Data[9])
	info.BankStockDiff = foundation.InterfaceToString(Data[10])
	info.TotalDiff = foundation.InterfaceToString(Data[11])
	return info
}
