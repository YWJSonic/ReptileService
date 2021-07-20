package t86

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
	return dbhandle.Instance.Getcollectionflag("", collectionFlagkey)
}

// CopyData ...
func CopyData(stockCode []string, date string, cacheTime int64) error {
	// result, err := Get(date, cacheTime)
	// if err != nil {
	// 	return err
	// }

	// Infos := result.GetInfos()
	// for _, info := range Infos {
	// dbhandle.Instance.
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	// err = dbhandle.Instance.Setcollectionflag("", collectionFlagkey, date)

	// if err != nil {
	// 	return err
	// }
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
