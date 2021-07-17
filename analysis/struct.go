package analysis

import (
	"fmt"
	"time"

	"github.com/YWJSonic/ReptileService/MISTWSEcom/stock"
	"github.com/YWJSonic/ReptileService/analysis/analysisday"
	"github.com/YWJSonic/ReptileService/dbhandle"
	"github.com/YWJSonic/ReptileService/foundation"
	"github.com/YWJSonic/ReptileService/routineswitch"
)

// Manager ...
type Manager struct {
	PriceDetails   map[string]*analysisday.Info
	routingswitchs map[string]*routineswitch.Info
}

// ShowPriceDetail ...
func (M *Manager) ShowPriceDetail(StockCode string) {
	M.PriceDetails[StockCode].PrintPriceDeatil()
}

// SetPriceDetail ...
func (M *Manager) SetPriceDetail(data stock.MsgInfo) {
	stockCode := data.C
	M.PriceDetails[stockCode].SetPriceDeatil(&data)
}

// StopCollectionPriceDetail ...
func (M *Manager) StopCollectionPriceDetail(StockCode string) {
	if _, ok := M.routingswitchs[StockCode]; !ok {
		fmt.Println("Error StopCollectionPriceDetail ", StockCode)
		return
	}
	M.routingswitchs[StockCode].Close()

}

// CollectionPriceDetail ...
func (M *Manager) CollectionPriceDetail(StockCode string) {
	routingswitch := &routineswitch.Info{}
	M.PriceDetails[StockCode] = analysisday.NewAnalysisDay()
	M.routingswitchs[StockCode] = routingswitch

	go func(stockcode string, pricedeatil *analysisday.Info, routingswitch *routineswitch.Info) {
		// mistwsecom.Collection(stockcode, routingswitch)

		var err error
		var result *stock.UpdateInfo
		var names *stock.NamesData
		var timelimit, timestart time.Time
		var info stock.MsgInfo
		var key string
		var keys []string
		timelimit, err = time.Parse("2006-01-02 15:04:05 -0700 MST", time.Now().Format("2006-01-02")+" 14:31:00 +0800 CST")
		timestart, err = time.Parse("2006-01-02 15:04:05 -0700 MST", time.Now().Format("2006-01-02")+" 08:30:00 +0800 CST")

		names, err = stock.GetName(StockCode)
		if err != nil {
			panic(err)
		}
		Exch := names.Names[0].GetExch()

		if time.Now().Before(timestart) {
			fmt.Println(timestart.Sub(time.Now()))
			time.Sleep(timestart.Sub(time.Now()))
		}

		olddatas, err := stock.GetOldData(stockcode, timestart.Format("2006-01-02"))
		if err != nil {
			panic(err)
		}
		for _, olddata := range olddatas {
			pricedeatil.SetPriceDeatil(&olddata)
		}

		for {
			if routingswitch.IsClose() {
				break
			}
			if time.Now().After(timelimit) {
				break
			}

			result, err = stock.GetUpdateMsg(Exch)
			if err != nil {
				fmt.Println(err)
				time.Sleep(time.Millisecond * 500)
			} else {
				key = result.MsgArray[0].Date() + result.MsgArray[0].Time()
				if !foundation.IsIncludeStr(key, keys) {
					keys = append(keys, key)
					info = result.MsgArray[0]
					dbhandle.Instance.SetTransactiondetail(info.C, info.D, info.T, info.TS, info.TK0, info.TK1, info.TLong, info.CH, info.N, info.NF, info.Y, info.Z, info.IP, info.TV, info.A, info.F, info.B, info.G, info.EX, info.IT, info.MT, info.O, info.OA, info.OB, info.OT, info.OV, info.OZ, info.I, info.L, info.H, info.V, info.W, info.U, info.S, info.P, info.PS, info.PZ)
				}
				pricedeatil.SetPriceDeatil(&result.MsgArray[0])
				time.Sleep(time.Second * 3)
			}
		}
		fmt.Println("Stop Collection ", StockCode)

	}(StockCode, M.PriceDetails[StockCode], routingswitch)
}
