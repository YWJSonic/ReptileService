package stockday

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	foundation "github.com/YWJSonic/ReptileService/foundation"
	"github.com/YWJSonic/ReptileService/handledb"
	handlehttp "github.com/YWJSonic/ReptileService/handlehttp"
)

// URL https://www.twse.com.tw/exchangeReport/STOCK_DAY?response=json&date=20190801&stockNo=2409&_=1573440324018

// CopyData ...
func CopyData(Num string, Date string) error {
	result, err := Get(Num, Date)
	if err != nil {
		return err
	}
	Infos := result.GetInfos()
	if Infos == nil {
		return errors.New("empty info")
	}
	for _, info := range Infos {
		err = handledb.Setstockday(info.StockCode, info.Year, info.Month, info.Day, info.StockPrice, info.StockCount, info.OpenPrice, info.ClosePrice, info.TopPrice, info.BottomPrice, info.DiffPrice, info.DealCount)
		if err != nil {
			return err
		}
	}
	err = handledb.Setcollectionflag(Num, "Day", Date[0:len(Date)-2])
	if err != nil {
		return err
	}
	return nil
}

// Get data
// Num: 股票代號
// Date: 查詢日期 20020101
// 查詢範圍: 20020101~20020131
func Get(Num string, Date string) (*Result, error) {
	data := &Result{}
	result := handlehttp.HTTPGetRequest(handlehttp.ConnectPool(), fmt.Sprintf("https://www.twse.com.tw/exchangeReport/STOCK_DAY?response=json&date=%s&stockNo=%s&_=%d", Date, Num, time.Now().Unix()*1000), nil)
	err := foundation.ByteToStruct(result, &data)
	if err != nil {
		return nil, err
	}
	data.Oringial = string(result)
	data.StockCode = Num
	if data.Stat != "OK" {
		return nil, errors.New(data.Stat)
	}
	return data, nil
}

// ConvertToInfo ...
func ConvertToInfo(Data []interface{}) Info {
	var info Info
	var year, month, day int
	var err error

	DateStr := strings.ReplaceAll(foundation.InterfaceToString(Data[0]), " ", "")
	DateSplit := strings.Split(DateStr, "/")

	if year, err = strconv.Atoi(DateSplit[0]); err != nil {
		fmt.Println(Data)
		panic(err)
	}
	if month, err = strconv.Atoi(DateSplit[1]); err != nil {
		fmt.Println(Data)
		panic(err)
	}
	if day, err = strconv.Atoi(DateSplit[2]); err != nil {
		fmt.Println(Data)
		panic(err)
	}

	info.Year = year
	info.Month = month
	info.Day = day
	info.StockCount = foundation.InterfaceToString(Data[1])
	info.StockPrice = foundation.InterfaceToString(Data[2])
	info.OpenPrice = foundation.InterfaceToString(Data[3])
	info.TopPrice = foundation.InterfaceToString(Data[4])
	info.BottomPrice = foundation.InterfaceToString(Data[5])
	info.ClosePrice = foundation.InterfaceToString(Data[6])
	info.DiffPrice = foundation.InterfaceToString(Data[7])
	info.DealCount = foundation.InterfaceToString(Data[8])
	return info
}

// GetAlreadyDate ...
func GetAlreadyDate(StockCode string) ([]map[string]interface{}, error) {
	return handledb.Getcollectionflag(StockCode, "Day")
}
