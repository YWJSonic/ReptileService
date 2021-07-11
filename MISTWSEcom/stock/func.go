package stock

import (
	"fmt"
	"time"

	"github.com/YWJSonic/ReptileService/foundation"
	"github.com/YWJSonic/ReptileService/handledb"
	"github.com/YWJSonic/ReptileService/handlehttp"
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

// GetOldData date:"2006-01-02"
func GetOldData(stockcode, date string) ([]MsgInfo, error) {
	result, err := handledb.Instance.GetTransactiondetail(stockcode, date)
	if err != nil {
		return nil, err
	}
	msginfos := DBMapConvertToMsgInfos(result)
	return msginfos, nil
}

// DBMapConvertToMsgInfos ...
func DBMapConvertToMsgInfos(datas []map[string]interface{}) []MsgInfo {
	var msginfos []MsgInfo
	for _, data := range datas {
		msginfo := MsgInfo{
			TS:    foundation.InterfaceToString(data["TS"]),
			TK0:   foundation.InterfaceToString(data["TK0"]),
			TK1:   foundation.InterfaceToString(data["TK1"]),
			TLong: foundation.InterfaceToString(data["TLong"]),
			CH:    foundation.InterfaceToString(data["CH"]),
			C:     foundation.InterfaceToString(data["C"]),
			N:     foundation.InterfaceToString(data["N"]),
			NF:    foundation.InterfaceToString(data["NF"]),
			D:     foundation.InterfaceToString(data["D"]),
			Y:     foundation.InterfaceToString(data["Y"]),
			Z:     foundation.InterfaceToString(data["Z"]),
			IP:    foundation.InterfaceToString(data["IP"]),
			TV:    foundation.InterfaceToString(data["TV"]),
			A:     foundation.InterfaceToString(data["A"]),
			F:     foundation.InterfaceToString(data["F"]),
			B:     foundation.InterfaceToString(data["B"]),
			G:     foundation.InterfaceToString(data["G"]),
			EX:    foundation.InterfaceToString(data["EX"]),
			IT:    foundation.InterfaceToString(data["IT"]),
			MT:    foundation.InterfaceToString(data["MT"]),
			O:     foundation.InterfaceToString(data["O"]),
			OA:    foundation.InterfaceToString(data["OA"]),
			OB:    foundation.InterfaceToString(data["OB"]),
			OT:    foundation.InterfaceToString(data["OT"]),
			OV:    foundation.InterfaceToString(data["OV"]),
			OZ:    foundation.InterfaceToString(data["OZ"]),
			I:     foundation.InterfaceToString(data["I"]),
			L:     foundation.InterfaceToString(data["L"]),
			H:     foundation.InterfaceToString(data["H"]),
			V:     foundation.InterfaceToString(data["V"]),
			W:     foundation.InterfaceToString(data["W"]),
			U:     foundation.InterfaceToString(data["U"]),
			T:     foundation.InterfaceToString(data["T"]),
			S:     foundation.InterfaceToString(data["S"]),
			P:     foundation.InterfaceToString(data["P"]),
			PS:    foundation.InterfaceToString(data["PS"]),
			PZ:    foundation.InterfaceToString(data["PZ"])}
		msginfos = append(msginfos, msginfo)
	}
	return msginfos
}
