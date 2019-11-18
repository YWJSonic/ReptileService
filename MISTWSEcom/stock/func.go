package stock

import (
	"fmt"
	"time"

	foundation "github.com/YWJSonic/ReptileService/foundation"
	handlehttp "github.com/YWJSonic/ReptileService/handlehttp"
)

// URL https://mis.twse.com.tw/stock/api/getStockInfo.jsp?ex_ch=tse_2409.tw&json=1&delay=0&_=1573451359721

// GetName ...
func GetName(Num string) (*NamesData, error) {
	info := &NamesData{}
	result := handlehttp.HTTPGetRequest(handlehttp.ConnectPool(), fmt.Sprintf("https://mis.twse.com.tw/stock/api/getStockNames.jsp?n=%s&_=%d", Num, time.Now().Unix()*1000), nil)
	err := foundation.ByteToStruct(result, &info)
	if err != nil {
		return nil, err
	}
	return info, nil

}

// GetUpdateMsg ...
func GetUpdateMsg(exch string) (*UpdateInfo, error) {
	info := &UpdateInfo{}
	result := handlehttp.HTTPGetRequest(handlehttp.ConnectPool(), fmt.Sprintf("https://mis.twse.com.tw/stock/api/getStockInfo.jsp?ex_ch=%s&json=1&delay=0&_=%d", exch, time.Now().Unix()*1000), nil)
	err := foundation.ByteToStruct(result, &info)
	if err != nil {
		return nil, err
	}
	info.Original = string(result)
	return info, nil
}
