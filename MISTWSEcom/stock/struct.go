package stock

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/YWJSonic/ReptileService/foundation"
)

// URL https://mis.twse.com.tw/stock/api/getStockInfo.jsp?ex_ch=tse_2409.tw&json=1&delay=0&_=1573451359721

// UpdateInfo ...
type UpdateInfo struct {
	Original  string
	MsgArray  []MsgInfo   `json:"msgArray"`
	UserDelay int64       `json:"userDelay"`
	RTmessage string      `json:"rtmessage"`
	Referer   string      `json:"referer"`
	QueryTime interface{} `json:"queryTime"`
	RTcode    string      `json:"rtcode"`
}

// NamesData ...
type NamesData struct {
	Names     []NameInfo `json:"datas"`
	RTmessage string     `json:"rtmessage"`
	RTcode    string     `json:"rtcode"`
}

// NameInfo ...
type NameInfo struct {
	C   string `json:"c"`
	N   string `json:"n"`
	Key string `json:"key"`
}

// GetExch 取得刷新資料用的ex_ch參數
func (N *NameInfo) GetExch() string {
	data := strings.Split(N.Key, "_")
	return data[0] + "_" + data[1]
}

// MsgInfo 當日每5秒交易資訊
type MsgInfo struct {
	Original string
	TS       string `json:"ts"`
	TK0      string `json:"tk0"`
	TK1      string `json:"tk1"`
	TLong    string `json:"tlong"`
	CH       string `json:"ch"` // 個股索引代號
	C        string `json:"c"`  // 個股編號
	N        string `json:"n"`  // 個股名稱
	NF       string `json:"nf"` // 個股全名
	D        string `json:"d"`  // 日期
	Y        string `json:"y"`  // 開盤價
	Z        string `json:"z"`  // 最近成交價
	IP       string `json:"ip"` // 0=跌 1=漲
	TV       string `json:"tv"` // 當盤成交量
	A        string `json:"a"`  // 賣出價[低 中 高]
	F        string `json:"f"`  // 賣出量 [低 中 高]
	B        string `json:"b"`  // 買進價[高 中 低]
	G        string `json:"g"`  // 買進量 [高 中 低]
	EX       string `json:"ex"`
	IT       string `json:"it"`
	MT       string `json:"mt"`
	O        string `json:"o"`  // 開盤價
	OA       string `json:"oa"` // 揭示賣價
	OB       string `json:"ob"` // 揭示買價
	OT       string `json:"ot"` // 揭示時間
	OV       string `json:"ov"` // 參考成交量
	OZ       string `json:"oz"` // 參考成交價
	I        string `json:"i"`
	L        string `json:"l"` // 本日最低價
	H        string `json:"h"` // 本日最高價
	V        string `json:"v"` // 本日累積成交量
	W        string `json:"w"` // 本日跌停價
	U        string `json:"u"` // 本日漲停價
	T        string `json:"t"` // 時間
	S        string `json:"s"`
	P        string `json:"p"`
	PS       string `json:"ps"` // 最近成交量
	PZ       string `json:"pz"` // 最近成交價
}

// 每日資料分析介面

// StockCode ...
func (M *MsgInfo) StockCode() string {
	return M.C
}

// StockNmae ...
func (M *MsgInfo) StockNmae() string {
	return M.N
}

// Price ...
func (M *MsgInfo) Price() string {
	return M.Z
}

// Count ...
func (M *MsgInfo) Count() int64 {
	var count int64
	var err error
	if count, err = strconv.ParseInt(M.S, 10, 64); err != nil {
		return 0
	}
	return count
}

// Time data create time
func (M *MsgInfo) Time() string {
	return M.T
}

// Date ...
func (M *MsgInfo) Date() string {
	return M.D
}

// OpenPrice ...
func (M *MsgInfo) OpenPrice() string {
	return M.O
}

// YesterdayPrice ...
func (M *MsgInfo) YesterdayPrice() string {
	return M.Y
}

//////////////////////////

// Print ...
func (M *MsgInfo) Print() {
	var Diff, DiffPre string
	var BuyTop5Money, BuyTop5Count, SellTop5Money, SellTop5Count []string
	var err error

	if Diff, err = M.GetDiffStr(); err != nil {
		fmt.Println(err)
		return
	}
	if DiffPre, err = M.GetDiffPreStr(); err != nil {
		fmt.Println(err)
		return
	}
	if BuyTop5Money, err = M.GetBuyMoneyTop5Str(); err != nil {
		fmt.Println(err)
		return
	}
	if BuyTop5Count, err = M.GetBuyCountTop5Str(); err != nil {
		fmt.Println(err)
		return
	}
	if SellTop5Money, err = M.GetSellMoneyBottom5Str(); err != nil {
		fmt.Println(err)
		return
	}
	if SellTop5Count, err = M.GetSellCountBottom5Str(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s %s %s %s\n", M.N, M.Z, Diff, DiffPre)
	fmt.Printf("%s %s\n", foundation.ReverseArray(BuyTop5Money), SellTop5Money)
	fmt.Printf("%s %s\n", foundation.ReverseArray(BuyTop5Count), SellTop5Count)
	fmt.Println("------------------------------------------------------")
}

// GetDiffStr 漲跌價差
func (M *MsgInfo) GetDiffStr() (string, error) {
	value, err := M.GetDiff()
	if err != nil {
		return "-", nil
	}
	result := strconv.FormatFloat(value, 'f', 2, 64)
	return result, nil
}

// GetDiff 漲跌價差
func (M *MsgInfo) GetDiff() (float64, error) {
	z, err1 := strconv.ParseFloat(M.Z, 64)
	y, err2 := strconv.ParseFloat(M.Y, 64)

	if err1 != nil {
		return z, err1
	}
	if err2 != nil {
		return y, err1
	}
	result := z - y
	return result, nil
}

// GetDiffPreStr 漲跌價差百分比
func (M *MsgInfo) GetDiffPreStr() (string, error) {

	value, err := M.GetDiffPre()
	if err != nil {
		return "-", nil
	}
	result := strconv.FormatFloat(value, 'f', 2, 64) + "%"

	return result, nil
}

// GetDiffPre 漲跌價差百分比
func (M *MsgInfo) GetDiffPre() (float64, error) {
	z, err1 := strconv.ParseFloat(M.Z, 64)
	y, err2 := strconv.ParseFloat(M.Y, 64)

	if err1 != nil {
		return 0, err1
	}
	if err2 != nil {
		return 0, err1
	}
	result := (z - y) / y * 100

	return result, nil
}

// GetSellMoneyBottom5Str 前五賣出價格
func (M *MsgInfo) GetSellMoneyBottom5Str() ([]string, error) {
	selltop5 := strings.Split(M.A, "_")
	return selltop5, nil
}

// GetSellMoneyBottom5 前五賣出價格
func (M *MsgInfo) GetSellMoneyBottom5() ([]float64, error) {
	values, err := M.GetSellMoneyBottom5Str()
	if err != nil {
		return nil, err
	}
	var result []float64

	for _, item := range values {
		value, errMsg := strconv.ParseFloat(item, 64)
		if errMsg != nil {
			fmt.Println("Error GetSellMoneyBottom5")
			continue
		}
		result = append(result, value)
	}
	return result, nil
}

// GetBuyMoneyTop5Str 前五買進價格
func (M *MsgInfo) GetBuyMoneyTop5Str() ([]string, error) {
	selltop5 := strings.Split(M.B, "_")
	return selltop5, nil
}

// GetBuyMoneyTop5 前五買進價格
func (M *MsgInfo) GetBuyMoneyTop5() ([]float64, error) {
	values, err := M.GetBuyMoneyTop5Str()
	if err != nil {
		return nil, err
	}
	var result []float64

	for _, item := range values {
		value, errMsg := strconv.ParseFloat(item, 64)
		if errMsg != nil {
			fmt.Println("Error GetSellMoneyBottom5")
			continue
		}
		result = append(result, value)
	}
	return result, nil
}

// GetSellCountBottom5Str 前五賣出數量
func (M *MsgInfo) GetSellCountBottom5Str() ([]string, error) {
	selltop5 := strings.Split(M.F, "_")
	return selltop5, nil
}

// GetSellCountBottom5 前五賣出數量
func (M *MsgInfo) GetSellCountBottom5() ([]float64, error) {
	values, err := M.GetSellCountBottom5Str()
	if err != nil {
		return nil, err
	}
	var result []float64

	for _, item := range values {
		value, errMsg := strconv.ParseFloat(item, 64)
		if errMsg != nil {
			fmt.Println("Error GetSellMoneyBottom5")
			continue
		}
		result = append(result, value)
	}
	return result, nil
}

// GetBuyCountTop5Str 前五買進數量
func (M *MsgInfo) GetBuyCountTop5Str() ([]string, error) {
	selltop5 := strings.Split(M.G, "_")
	return selltop5, nil
}

// GetBuyCountTop5 前五買進數量
func (M *MsgInfo) GetBuyCountTop5() ([]float64, error) {
	values, err := M.GetBuyCountTop5Str()
	if err != nil {
		return nil, err
	}
	var result []float64

	for _, item := range values {
		value, errMsg := strconv.ParseFloat(item, 64)
		if errMsg != nil {
			fmt.Println("Error GetSellMoneyBottom5")
			continue
		}
		result = append(result, value)
	}
	return result, nil
}

// IsPriceUp 本日漲跌
func (M *MsgInfo) IsPriceUp() bool {

	if value, err := M.GetDiff(); err != nil {
		return false
	} else if value > 0 {
		return true
	}

	return false
}
