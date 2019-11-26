package analysis

import (
	"github.com/YWJSonic/ReptileService/analysis/analysisday"
	"github.com/YWJSonic/ReptileService/routineswitch"
	"math"
)

// GetAnalysisManager Get
func GetAnalysisManager() *Manager {
	if analysisManager == nil {
		manager := &Manager{}
		manager.PriceDetails = make(map[string]*analysisday.Info)
		manager.routingswitchs = make(map[string]*routineswitch.Info)
		analysisManager = manager
	}

	return analysisManager
}

// PayForBuy 計算購買花費金額
func PayForBuy(stockprice float64, stockCount int64) (price, feeprice, total int64) {
	price = int64(stockprice * float64(stockCount) * 1000)
	feeprice = int64(math.Floor(float64(price) * fee))
	total = price + feeprice
	return
}

// ResultForSell 計算賣出回收金額
func ResultForSell(stockprice float64, stockCount int64) (price, feeprice, taxprice, total int64) {
	price = int64(stockprice * float64(stockCount) * 1000)
	feeprice = int64(math.Floor(float64(price) * fee))
	taxprice = int64(math.Floor(float64(price) * tax))
	total = price - feeprice - taxprice
	return
}

// ProfitMath 計算交易利潤
func ProfitMath(buyPrice, sellPrice float64) int64 {
	var Profit int64
	buyprice, buyfee, _ := PayForBuy(buyPrice, 1)
	sellprice, sellfee, selltax, _ := ResultForSell(sellPrice, 1)

	Profit = sellprice - buyprice - buyfee - sellfee - selltax
	return Profit
}
