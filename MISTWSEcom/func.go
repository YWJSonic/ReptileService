package mistwsecom

import (
	"fmt"
	"time"

	"github.com/YWJSonic/ReptileService/MISTWSEcom/stock"
	"github.com/YWJSonic/ReptileService/foundation"
	"github.com/YWJSonic/ReptileService/handledb"
	"github.com/YWJSonic/ReptileService/routineswitch"
)

// Collection ...
func Collection(StockCode string, routing *routineswitch.Info) {

	var err error
	var result *stock.UpdateInfo
	var names *stock.NamesData
	var keys []string
	var key string
	timelimit, err := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02")+" 14:31:00")
	if err != nil {
		fmt.Println(err)
	}

	names, err = stock.GetName(StockCode)
	if err != nil {
		panic(err)
	}
	stockcode := names.Names[0].GetExch()
	for {
		if routing.IsClose() {
			break
		}
		result, err = stock.GetUpdateMsg(stockcode)
		if err != nil {
			fmt.Println(err)
			time.Sleep(time.Second)
		} else {
			fmt.Println(result.MsgArray)
			key = result.MsgArray[0].Date() + result.MsgArray[0].Time()
			if !foundation.IsIncludeStr(key, keys) {
				keys = append(keys, key)
				handledb.SetTransactiondetail(result.MsgArray[0])
			}
			time.Sleep(time.Second * 5)

			if time.Now().After(timelimit) {
				break
			}
		}

	}
}
