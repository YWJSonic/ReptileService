package twsecom

import (
	"fmt"
	"strconv"
	"time"

	"github.com/YWJSonic/ReptileService/TWSEcom/fmnptk"
	"github.com/YWJSonic/ReptileService/TWSEcom/fmsrfk"
	"github.com/YWJSonic/ReptileService/TWSEcom/stockday"
	"github.com/YWJSonic/ReptileService/foundation"
)

// YearSlipDupliy 執行續分割年分
var YearSlipDupliy = 5

// Collection ...
func Collection(StockCode string) {
	RunCount := make(chan int)

	var YearSlice []int
	thisYear := time.Now().Year()
	NextYear := thisYear
	LastYear := thisYear - 11

	for NextYear > LastYear {
		YearSlice = append(YearSlice, NextYear)
		NextYear = NextYear - YearSlipDupliy
	}

	collectionflag, err := stockday.GetAlreadyDate(StockCode)
	if err != nil {
		fmt.Println(err)
		return
	}

	for index, count := 0, len(YearSlice); index < count; index++ {
		if index == count-1 {
			go DayCollection(StockCode, YearSlice[index], LastYear, collectionflag, RunCount)
		} else {
			go DayCollection(StockCode, YearSlice[index], YearSlice[index+1], collectionflag, RunCount)
		}
	}

	go MonthCollection(StockCode, thisYear, LastYear, collectionflag, RunCount)

	go YearCollection(StockCode, collectionflag, RunCount)

	for range YearSlice {
		<-RunCount
		fmt.Println("finish -------------111")
	}
	<-RunCount
	fmt.Println("finish -------------222")
	<-RunCount
	fmt.Println("finish -------------333")
}

// DayCollection ...
func DayCollection(StockCode string, StartYear, EndYear int, collectionflag []map[string]interface{}, chint chan int) {
	// Stock Day Data
	var date string
	var month int
	var err error
	//now := time.Now()

	for year := StartYear; year > EndYear; year-- {
		month = 12
		for month > 0 {
			if foundation.IsAfterNowTime(year, month, 1) {
				month--
				continue
			}

			date = fmt.Sprintf("%d%02d01", year, month)

			if IsInCollectionFlag(date[0:len(date)-2], "Day", collectionflag) { //} && !(now.Year() == year && now.Month() == time.Month(month)) {
				fmt.Printf("Day %s IsSkip!\n", date[0:len(date)-2])
				month--
				continue
			}

			err = stockday.CopyData(StockCode, date)
			fmt.Printf("Collection %s daily day stock\n", date)
			time.Sleep(time.Second * 1)

			if err != nil {
				fmt.Println(date, err)
			}

			month--
		}
	}
	fmt.Printf("Stock %s %d ~ %d daily data finish!!!", StockCode, StartYear, EndYear)
	chint <- 1
}

// MonthCollection ...
func MonthCollection(StockCode string, StartYear, LastYear int, collectionflag []map[string]interface{}, chint chan int) {
	// Stock Day Data
	var date string
	var err error

	for year := StartYear; year > LastYear; year-- {
		if foundation.IsAfterNowTime(year, 1, 1) {
			continue
		}

		date = fmt.Sprintf("%d0101", year)

		if IsInCollectionFlag(date[0:len(date)-2], "Month", collectionflag) {
			fmt.Printf("Month %s IsSkip!\n", date[0:len(date)-2])
			continue
		}

		err = fmsrfk.CopyData(StockCode, date)
		fmt.Printf("Collection %s daily Month stock\n", date)
		time.Sleep(time.Second * 5)

		if err != nil {
			fmt.Println(date, err)
		}
	}
	fmt.Println("Stock " + StockCode + " monthly data finish!!!")
	chint <- 1
}

// YearCollection ...
func YearCollection(StockCode string, collectionflag []map[string]interface{}, chint chan int) {
	// Stock Year Data
	var err error

	date := strconv.Itoa(time.Now().Year())
	if IsInCollectionFlag(date, "Year", collectionflag) {
		fmt.Printf("Year %s IsSkip!\n", date)
		chint <- 1
		return
	}

	err = fmnptk.CopyData(StockCode)
	if err != nil {
		fmt.Println(err)
		chint <- 1
		return
	}
	fmt.Println("Stock " + StockCode + " yearly data finish!!!")
	chint <- 1
}

// MonthDayCount ...
func MonthDayCount(year, month int) int {
	datetime := time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC)
	fmt.Println(datetime)
	return datetime.Day()
}

// ADToROC AD year to ROC year
func ADToROC(year int) int {
	return year - 1911
}

// IsInCollectionFlag ...
func IsInCollectionFlag(Date string, Flag string, CollectionFlags []map[string]interface{}) bool {
	for _, CollectionFlag := range CollectionFlags {
		if foundation.InterfaceToString(CollectionFlag["Date"]) == Date && foundation.InterfaceToString(CollectionFlag["Flag"]) == Flag {
			return true
		}
	}
	return false
}
