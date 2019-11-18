package analysisday

import (
	"fmt"

	"github.com/YWJSonic/ReptileService/foundation"
)

// Info pricedetail info
type Info struct {
	PriceDeatil    map[string]int64
	PriceDetailLog []string
}

// PrintPriceDeatil ...
func (I *Info) PrintPriceDeatil() {
	for key, value := range I.PriceDeatil {
		fmt.Println(key, value)
	}
}

// SetPriceDeatil ...
func (I *Info) SetPriceDeatil(info GetPriceDeatil) {

	if !foundation.IsIncludeStr(info.Time(), I.PriceDetailLog) {
		I.PriceDeatil[info.Price()] = info.Count()
		I.PriceDetailLog = append(I.PriceDetailLog, info.Time())
	}

}
