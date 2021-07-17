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

// 網站限制
// 1分鐘 內請求 25次 以下
// 建議每 4~5秒 一次
// 被鎖定後 1小時 才會解鎖

// YearSlipDupliy 執行續分割年分
var YearSlipDupliy = 5

// Collection ...
func Collection(StockCode string) {
	var YearSlice []int
	thisYear := time.Now().Year()
	NextYear := thisYear
	LastYear := thisYear - 11

	for NextYear > LastYear {
		YearSlice = append(YearSlice, NextYear)
		NextYear = NextYear - YearSlipDupliy
	}

	collectionflag := []map[string]interface{}{}
	collectionflag, err := stockday.GetAlreadyDate(StockCode)
	if err != nil {
		fmt.Println(err)
		return
	}

	for index, count := 0, len(YearSlice); index < count; index++ {
		if index == count-1 {
			DayCollection(StockCode, YearSlice[index], LastYear, collectionflag)
		} else {
			DayCollection(StockCode, YearSlice[index], YearSlice[index+1], collectionflag)
		}
	}

	MonthCollection(StockCode, thisYear, LastYear, collectionflag)

	YearCollection(StockCode, collectionflag)
}

// DayCollection ...
func DayCollection(StockCode string, StartYear, EndYear int, collectionflag []map[string]interface{}) {
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

			err = stockday.CopyData(StockCode, date, time.Now().Unix()*1000)
			fmt.Printf("Collection %s daily day stock\n", date)
			time.Sleep(time.Second * 1)

			if err != nil {
				fmt.Println("YearCollection Error:", date, err)
				return
			}

			month--
		}
	}
	fmt.Printf("Stock %s %d ~ %d daily data finish!!!\n", StockCode, StartYear, EndYear)
}

// MonthCollection ...
func MonthCollection(StockCode string, StartYear, LastYear int, collectionflag []map[string]interface{}) {
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
			fmt.Println("MonthCollection Error:", date, err)
			return
		}
	}
	fmt.Println("Stock " + StockCode + " monthly data finish!!!")
}

// YearCollection ...
func YearCollection(StockCode string, collectionflag []map[string]interface{}) {
	// Stock Year Data
	var err error

	date := strconv.Itoa(time.Now().Year())
	if IsInCollectionFlag(date, "Year", collectionflag) {
		fmt.Printf("Year %s IsSkip!\n", date)
		return
	}

	err = fmnptk.CopyData(StockCode)
	if err != nil {
		fmt.Println("YearCollection Error:", date, err)
		return
	}
	fmt.Println("Stock " + StockCode + " yearly data finish!!!")
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
