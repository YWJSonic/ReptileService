package analysis

import (
	"fmt"
	"time"

	"github.com/YWJSonic/ReptileService/MISTWSEcom/stock"
	"github.com/YWJSonic/ReptileService/analysis/analysisday"
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

		names, err = stock.GetName(StockCode)
		if err != nil {
			panic(err)
		}
		Exch := names.Names[0].GetExch()
		for {
			if routingswitch.IsClose() {
				break
			}
			result, err = stock.GetUpdateMsg(Exch)
			pricedeatil.SetPriceDeatil(&result.MsgArray[0])
			if err != nil {
				fmt.Println(err)
				time.Sleep(time.Second)
			} else {
				time.Sleep(time.Second * 5)
			}
		}
		fmt.Println("Stop Collection ", StockCode)

	}(StockCode, M.PriceDetails[StockCode], routingswitch)
}
