package twsecom

import (
	"fmt"
	"time"

	"github.com/YWJSonic/ReptileService/TWSEcom/fmnptk"
	"github.com/YWJSonic/ReptileService/TWSEcom/fmsrfk"
	"github.com/YWJSonic/ReptileService/TWSEcom/stockday"
)

// Collection ...
func Collection(StockCode string) {
	RunCount := make(chan int)

	go func(StockCode string) {
		// Stock Day Data
		var date string
		for year := time.Now().Year(); year > 2015; year-- {
			for month := 12; month > 0; month-- {
				date = fmt.Sprintf("%d%02d01", year, month)
				stockday.CopyData(StockCode, date)
				fmt.Printf("Collection %s\n daily day stock", date)
				time.Sleep(time.Second * 5)
			}
		}
		fmt.Println("Stock " + StockCode + " day data finish!!!")
		RunCount <- 1
	}(StockCode)

	go func(StockCode string) {
		// Stock Day Data
		var date string
		for year := 2015; year > 2009; year-- {
			for month := 12; month > 0; month-- {
				date = fmt.Sprintf("%d%02d01", year, month)
				stockday.CopyData(StockCode, date)
				fmt.Printf("Collection %s\n daily day stock", date)
				time.Sleep(time.Second * 5)
			}
		}
		fmt.Println("Stock " + StockCode + " day data finish!!!")
		RunCount <- 1
	}(StockCode)

	go func(StockCode string) {
		// Stock Month Data
		var date string
		for year := time.Now().Year(); year >= 2010; year-- {
			date = fmt.Sprintf("%d0101", year)
			fmsrfk.CopyData(StockCode, date)
			fmt.Printf("Collection %s\n daily day stock", date)
			time.Sleep(time.Second)
		}
		fmt.Println("Stock " + StockCode + " month data finish!!!")
		RunCount <- 1
	}(StockCode)

	go func(StockCode string) {
		// Stock Year Data
		fmnptk.CopyData(StockCode)
		fmt.Println("Stock " + StockCode + " year data finish!!!")
		RunCount <- 1
	}(StockCode)

	<-RunCount
	<-RunCount
	<-RunCount
	<-RunCount
}
