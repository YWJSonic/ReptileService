package analysisday

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/YWJSonic/ReptileService/foundation"
	"github.com/YWJSonic/ReptileService/foundation/sortinterface"
)

// Info pricedetail info
type Info struct {
	isSamePrice    bool
	openPrice      string
	newPrice       string
	newCount       int64
	samePriceCount int64
	PriceDeatil    map[string]int64
	PriceDetailLog []string

	countDeatilTop int
	countLogLimit  int64 // value = count * price * 1000
	CountDeatil    map[string]*CountDeatilInfo
}

// PrintPriceDeatil ...
func (I *Info) PrintPriceDeatil() {
	var Keys []string

	Keys = foundation.MapSIKeys(I.PriceDeatil)
	sort.Sort(sort.Reverse(sort.StringSlice(Keys)))
	// sort.Strings(sort.Reverse())

	fmt.Printf("------ 價格    累積數量    現購數量    累積數量 ---\n")
	for index, key := range Keys {
		if key == I.newPrice {
			fmt.Printf("現     %s    %d        %d        %d\n", key, I.PriceDeatil[key], I.newCount, I.samePriceCount)
		} else if key == I.openPrice {
			fmt.Printf("開     %s    %d\n", key, I.PriceDeatil[key])
		} else if index == 0 {
			fmt.Printf("高     %s    %d\n", key, I.PriceDeatil[key])
		} else if len(Keys) == (index + 1) {
			fmt.Printf("低     %s    %d\n", key, I.PriceDeatil[key])
		} else {
			fmt.Printf("       %s    %d\n", key, I.PriceDeatil[key])
		}
	}

	var CountDeatilKeys []int64

	Keys = []string{}
	for key := range I.CountDeatil {
		Keys = append(Keys, key)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(Keys)))
	fmt.Printf("------ 價格    數量/次數 ---\n")
	for _, key := range Keys {

		CountDeatilKeys = foundation.MapIIKeys(I.CountDeatil[key].CountData)
		sort.Sort(sort.Reverse(sortinterface.Int64Slice(CountDeatilKeys)))

		fmt.Printf("       %s", key)
		for index, CountDeatilKey := range CountDeatilKeys {
			if index >= I.countDeatilTop {
				break
			}
			fmt.Printf("    %d/%d", CountDeatilKey, I.CountDeatil[key].CountData[CountDeatilKey])
		}
		fmt.Println("")
	}
}

// SetPriceDeatil ...
func (I *Info) SetPriceDeatil(info GetPriceDeatil) {
	if info.Price() == "" || info.Count() == 0 {
		return
	}

	if !foundation.IsIncludeStr(info.Time(), I.PriceDetailLog) {
		I.isSamePrice = I.newPrice == info.Price()
		I.openPrice = info.OpenPrice()
		I.newPrice = info.Price()
		I.newCount = info.Count()

		if I.isSamePrice {
			I.samePriceCount += info.Count()
		} else {
			I.samePriceCount = info.Count()
		}
		I.PriceDeatil[I.newPrice] += I.newCount
		I.PriceDetailLog = append(I.PriceDetailLog, info.Time())

		if value, ok := strconv.ParseFloat(I.newPrice, 64); ok == nil &&
			(value*float64(I.newCount)*1000) >= float64(I.countLogLimit) {

			if _, ok := I.CountDeatil[I.newPrice]; !ok {
				I.CountDeatil[I.newPrice] = &CountDeatilInfo{}
			}
			I.CountDeatil[I.newPrice].SetCountData(I.newCount)
		}

	}
}

// CountDeatilInfo ...
type CountDeatilInfo struct {
	CountData map[int64]int64
}

// SetCountData ...
func (CDI *CountDeatilInfo) SetCountData(count int64) {
	if CDI.CountData != nil {
		CDI.CountData[count]++
	} else {
		CDI.CountData = map[int64]int64{count: 1}
	}
}
