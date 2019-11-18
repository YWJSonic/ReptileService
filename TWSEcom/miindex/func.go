package miindex

import (
	"fmt"
	"time"

	foundation "github.com/YWJSonic/ReptileService/foundation"
	handlehttp "github.com/YWJSonic/ReptileService/handlehttp"
)

// SelectType ...
var SelectType = map[string]string{
	"全部":        "ALL",
	"大盤統計資訊":    "MS",
	"收盤指數資訊":    "IND",
	"委託及成交統計資訊": "MS2",
	"全部(不含權證、牛熊證、可展延牛熊證)": "ALLBUT0999",
	"封閉式基金":       "49",
	"ETF":         "0099P",
	"ETN":         "29999",
	"受益證券":        "019919T",
	"認購權證(不含牛證)":  "999",
	"認售權證(不含熊證)":  "0999P",
	"牛證(不含可展延牛證)": "0999C",
	"熊證(不含可展延熊證)": "0999B",
	"可展延牛證":       "0999X",
	"可展延熊證":       "0999Y",
	"附認股權特別股":     "0999GA",
	"附認股權公司債":     "0999GD",
	"認股權憑證":       "0999G9",
	"可轉換公司債":      "CB",
	"水泥工業":        "1",
	"食品工業":        "2",
	"塑膠工業":        "3",
	"紡織纖維":        "4",
	"電機機械":        "5",
	"電器電纜":        "6",
	"化學生技醫療":      "7",
	"玻璃陶瓷":        "8",
	"造紙工業":        "9",
	"鋼鐵工業":        "10",
	"橡膠工業":        "11",
	"汽車工業":        "12",
	"電子工業":        "13",
	"建材營造":        "14",
	"航運業":         "15",
	"觀光事業":        "16",
	"金融保險":        "17",
	"貿易百貨":        "18",
	"綜合":          "19",
	"其他":          "20",
	"化學工業":        "21",
	"生技醫療業":       "22",
	"油電燃氣業":       "23",
	"半導體業":        "24",
	"電腦及週邊設備業":    "25",
	"光電業":         "26",
	"通信網路業":       "27",
	"電子零組件業":      "28",
	"電子通路業":       "29",
	"資訊服務業":       "30",
	"其他電子業":       "31",
	"存託憑證":        "9299"}

// SelectTypeIndex ...
var SelectTypeIndex = []string{
	"全部",
	"大盤統計資訊",
	"收盤指數資訊",
	"委託及成交統計資訊",
	"全部(不含權證、牛熊證、可展延牛熊證)",
	"封閉式基金",
	"ETF",
	"ETN",
	"受益證券",
	"認購權證(不含牛證)",
	"認售權證(不含熊證)",
	"牛證(不含可展延牛證)",
	"熊證(不含可展延熊證)",
	"可展延牛證",
	"可展延熊證",
	"附認股權特別股",
	"附認股權公司債",
	"認股權憑證",
	"可轉換公司債",
	"水泥工業",
	"食品工業",
	"塑膠工業",
	"紡織纖維",
	"電機機械",
	"電器電纜",
	"化學生技醫療",
	"玻璃陶瓷",
	"造紙工業",
	"鋼鐵工業",
	"橡膠工業",
	"汽車工業",
	"電子工業",
	"建材營造",
	"航運業",
	"觀光事業",
	"金融保險",
	"貿易百貨",
	"綜合",
	"其他",
	"化學工業",
	"生技醫療業",
	"油電燃氣業",
	"半導體業",
	"電腦及週邊設備業",
	"光電業",
	"通信網路業",
	"電子零組件業",
	"電子通路業",
	"資訊服務業",
	"其他電子業",
	"存託憑證"}

// GetBySelectTypeIndex ...
func GetBySelectTypeIndex(selectTypeIndex int, Date string) (*Info, error) {
	return GetBySelectType(SelectTypeIndex[selectTypeIndex], Date)
}

// GetBySelectType ...
func GetBySelectType(selectType string, Date string) (*Info, error) {
	info := &Info{}
	ip := fmt.Sprintf("https://www.twse.com.tw/exchangeReport/MI_INDEX?response=json&date=%s&type=%s&_=%d", Date, SelectType[selectType], time.Now().Unix()*1000)
	result := handlehttp.HTTPGetRequest(handlehttp.ConnectPool(), ip, nil)
	// fmt.Println(string(result))
	err := foundation.ByteToStruct(result, &info)
	if err != nil {
		return nil, err
	}
	return info, nil
}
