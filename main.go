package main

import (
	"os"
	"os/exec"
	"time"

	"github.com/YWJSonic/ReptileService/analysis"
	"github.com/YWJSonic/ReptileService/handledb"
	_ "github.com/go-sql-driver/mysql"
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
	stockCode := os.Args[1]

	setting := struct{ DBUser, DBPassword, DBIP, DBPORT, DBName string }{
		DBUser:     "sony79410",
		DBPassword: "Sonic79410",
		DBIP:       "127.0.0.1",
		DBPORT:     "3306",
		DBName:     "stock",
	}
	err := handledb.SetInit(&setting)
	if err != nil {
		panic(err)
	}

	// stockcodes := []string{"2377", "2882", "2888", "3481", "4904", "8046", "2317", "3481", "2376", "2449"}
	// for _, stockcode := range stockcodes {
	// 	twsecom.Collection(stockcode)
	// }

	// routingswitch := &routineswitch.Info{}
	// mistwsecom.Collection(stockCode, routingswitch)

	analysis.GetAnalysisManager().CollectionPriceDetail(stockCode)
	for {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
		analysis.GetAnalysisManager().ShowPriceDetail(stockCode)
		time.Sleep(time.Second * 5)
	}

	// analysis.GetAnalysisManager().StopCollectionPriceDetail("2409")

	// value1, _ := strconv.ParseFloat(os.Args[1], 64)
	// value2, _ := strconv.ParseFloat(os.Args[2], 64)

	// fmt.Println(analysis.ProfitMath(value1, value2))
	// fmt.Println("---------------------------------------------")
	// time.Sleep(time.Hour)
}
