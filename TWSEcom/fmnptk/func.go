package fmnptk

import (
	"fmt"
	"time"

	"github.com/YWJSonic/ReptileService/dbhandle"
	foundation "github.com/YWJSonic/ReptileService/foundation"
	"github.com/YWJSonic/ReptileService/httphandle"
)

// URL https://www.twse.com.tw/exchangeReport/FMNPTK?response=json&stockNo=2409&_=1573430069096

// CopyData ...
func CopyData(Num string) error {
	result, err := Get(Num)
	if err != nil {
		return err
	}
	Infos := result.GetInfos()
	for _, info := range Infos {
		err = dbhandle.Instance.Setstockyear(info.StockCode, info.Year, info.StockCount, info.StockPrice, info.DealCount, info.TopPrice, info.TopPriceDate, info.BottomPrice, info.BottomPriceDate, info.AVGPrice)
		if err != nil {
			return err
		}
	}

	err = dbhandle.Instance.Setcollectionflag(Num, "Year", fmt.Sprint(time.Now().Year()))
	if err != nil {
		return err
	}
	return nil
}

// Get data
// Num: 股票代號
func Get(Num string) (*Result, error) {
	data := &Result{}
	result := httphandle.Instans.HTTPGetRequest(fmt.Sprintf("https://www.twse.com.tw/exchangeReport/FMNPTK?response=json&stockNo=%s&_=%d", Num, time.Now().Unix()*1000), nil)
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
	info.Year = foundation.InterfaceToInt(Data[0])
	info.StockPrice = foundation.InterfaceToString(Data[1])
	info.StockCount = foundation.InterfaceToString(Data[2])
	info.DealCount = foundation.InterfaceToString(Data[3])
	info.TopPrice = foundation.InterfaceToString(Data[4])
	info.TopPriceDate = foundation.InterfaceToString(Data[5])
	info.BottomPrice = foundation.InterfaceToString(Data[6])
	info.BottomPriceDate = foundation.InterfaceToString(Data[7])
	info.AVGPrice = foundation.InterfaceToString(Data[8])
	return info
}

// GetAlreadyDate ...
func GetAlreadyDate(StockCode string) ([]map[string]interface{}, error) {
	return dbhandle.Instance.Getcollectionflag(StockCode, "Year")
}
