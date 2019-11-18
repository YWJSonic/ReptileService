package main

import (
	"fmt"
	"time"

	mistwsecom "github.com/YWJSonic/ReptileService/MISTWSEcom"
	"github.com/YWJSonic/ReptileService/handledb"
	"github.com/YWJSonic/ReptileService/routineswitch"
	_ "github.com/go-sql-driver/mysql"
)

// URL https://www.twse.com.tw/exchangeReport/FMNPTK?response=json&stockNo=2409&_=1573430069096
// URL https://www.twse.com.tw/exchangeReport/FMSRFK?response=json&date=20020101&stockNo=2409&_=1573433055138

// URL https://www.twse.com.tw/exchangeReport/STOCK_DAY?response=json&date=20190801&stockNo=2409&_=1573440324018

// URL https://www.twse.com.tw/exchangeReport/MI_INDEX?response=json&date=20191108&type=ALL&_=1573441130523
// URL https://www.twse.com.tw/exchangeReport/MI_INDEX?response=json&date=20191107&type=ALL&_=1573441130533

// URL https://www.twse.com.tw/exchangeReport/MI_INDEX?response=json&date=20191107&type=MS&_=1573441130592

// URL　https://mis.twse.com.tw/stock/api/getStockNames.jsp?n=2409&_=1573452768792
func main() {
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

	routingswitch := &routineswitch.Info{}
	go func() {
		mistwsecom.Collection("2409", routingswitch)
	}()

	// analysis.GetAnalysisManager().CollectionPriceDetail("2409")
	// for index := 50; index > 0; index-- {

	// 	cmd := exec.Command("cmd", "/c", "cls")
	// 	cmd.Stdout = os.Stdout
	// 	cmd.Run()
	// 	analysis.GetAnalysisManager().ShowPriceDetail("2409")
	// 	time.Sleep(time.Second * 5)
	// }
	// analysis.GetAnalysisManager().StopCollectionPriceDetail("2409")
	fmt.Println("---------------------------------------------")
	time.Sleep(time.Hour)
}
