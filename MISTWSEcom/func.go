package mistwsecom

import (
	"fmt"
	"time"

	"github.com/YWJSonic/ReptileService/MISTWSEcom/stock"
	"github.com/YWJSonic/ReptileService/dbhandle"
	"github.com/YWJSonic/ReptileService/foundation"
	"github.com/YWJSonic/ReptileService/routineswitch"
)

// Collection ...
func Collection(StockCode string, routing *routineswitch.Info) {

	var err error
	var result *stock.UpdateInfo
	var names *stock.NamesData
	var keys []string
	var key string
	var timelimit, timestart time.Time
	var info stock.MsgInfo

	timelimit, err = time.Parse("2006-01-02 15:04:05 -0700 MST", time.Now().Format("2006-01-02")+" 14:31:00 +0800 CST")
	timestart, err = time.Parse("2006-01-02 15:04:05 -0700 MST", time.Now().Format("2006-01-02")+" 08:30:00 +0800 CST")
	timestart.Zone()
	if err != nil {
		fmt.Println(err)
	}

	for {
		names, err = stock.GetName(StockCode)
		if err != nil {
			fmt.Println(err)
			time.Sleep(time.Second)
		} else {
			break
		}
	}

	stockcode := names.Names[0].GetExch()
	if time.Now().Before(timestart) {
		fmt.Println(timestart.Sub(time.Now()))
		time.Sleep(timestart.Sub(time.Now()))
	}

	for {
		if routing.IsClose() {
			fmt.Println("Afteroutingr IsClose ")
			break
		}
		if time.Now().After(timelimit) {
			fmt.Println("After: ", timelimit)
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
				info = result.MsgArray[0]
				dbhandle.Instance.SetTransactiondetail(info.C, info.D, info.T, info.TS, info.TK0, info.TK1, info.TLong, info.CH, info.N, info.NF, info.Y, info.Z, info.IP, info.TV, info.A, info.F, info.B, info.G, info.EX, info.IT, info.MT, info.O, info.OA, info.OB, info.OT, info.OV, info.OZ, info.I, info.L, info.H, info.V, info.W, info.U, info.S, info.P, info.PS, info.PZ)
			}
			time.Sleep(time.Second * 5)
		}
	}
}
