package main

import (
	twsecom "github.com/YWJSonic/ReptileService/TWSEcom"
	"github.com/YWJSonic/ReptileService/dbhandle"
	"github.com/YWJSonic/ReptileService/httphandle"
)

// URL https://www.twse.com.tw/exchangeReport/FMNPTK?response=json&stockNo=2409&_=1573430069096
// URL https://www.twse.com.tw/exchangeReport/FMSRFK?response=json&date=20020101&stockNo=2409&_=1573433055138

// URL https://www.twse.com.tw/exchangeReport/STOCK_DAY?response=json&date=20190801&stockNo=2409&_=1573440324018

// URL https://www.twse.com.tw/exchangeReport/MI_INDEX?response=json&date=20191108&type=ALL&_=1573441130523
// URL https://www.twse.com.tw/exchangeReport/MI_INDEX?response=json&date=20191107&type=ALL&_=1573441130533

// URL https://www.twse.com.tw/exchangeReport/MI_INDEX?response=json&date=20191107&type=MS&_=1573441130592

// URL　https://mis.twse.com.tw/stock/api/getStockNames.jsp?n=2409&_=1573452768792

// URL https://www.twse.com.tw/exchangeReport/BWIBBU?response=json&date=20190101&stockNo=2409&_=1574393051372
func main() {
	// stockCode := os.Args[1]

	// setting := struct{ DBUser, DBPassword, DBIP, DBPORT, DBName string }{
	// 	DBUser:     "sony79410",
	// 	DBPassword: "Sonic79410",
	// 	DBIP:       "127.0.0.1",
	// 	DBPORT:     "3306",
	// 	DBName:     "stock",
	// }
	dbHandle := dbhandle.NewDBHandle()
	dbHandle.ConnectLocalDB(struct{ Path string }{Path: "./DB"})
	dbhandle.Instance = dbHandle

	httphandle.Instans = httphandle.NewHttpHandle()

	stockcodes := []string{"2356", "2412", "2449", "2834", "5283", "2002", "2881", "2888", "6505"}
	for _, stockcode := range stockcodes {
		twsecom.Collection(stockcode) //stockCode)
	}
}
