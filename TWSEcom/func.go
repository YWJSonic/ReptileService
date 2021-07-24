package twsecom

import (
	"fmt"
	"strconv"
	"time"

	"github.com/YWJSonic/ReptileService/TWSEcom/bwibbu"
	"github.com/YWJSonic/ReptileService/TWSEcom/fmnptk"
	"github.com/YWJSonic/ReptileService/TWSEcom/fmsrfk"
	"github.com/YWJSonic/ReptileService/TWSEcom/legalperson"
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
	thisYear := time.Now().Year()
	LastYear := thisYear - 11

	BwibbuCollection(StockCode, thisYear, LastYear)

	DayCollection(StockCode, thisYear, LastYear)

	MonthCollection(StockCode, thisYear, LastYear)

	YearCollection(StockCode)
}

// DayCollection ...
func DayCollection(StockCode string, StartYear, EndYear int) {
	// Stock Day Data
	var month int
	var err error
	var collectionflag []map[string]interface{}
	var cacheTime int64 = time.Now().Unix() * 1000

	if collectionflag, err = stockday.GetAlreadyDate(StockCode); err != nil {
		fmt.Println("DayCollection Error:", err)
		return
	}

	dates := MonthSlice(StartYear, EndYear, true)
	for _, date := range dates {
		if IsInCollectionFlag(date[0:len(date)-2], stockday.CollectionFlagkey, collectionflag) { //} && !(now.Year() == year && now.Month() == time.Month(month)) {
			fmt.Printf("Day %s IsSkip!\n", date[0:len(date)-2])
			month--
			continue
		}

		err = stockday.CopyData(StockCode, date, cacheTime)
		fmt.Printf("Collection %s daily day stock\n", date)
		time.Sleep(time.Second * 1)

		if err != nil {
			fmt.Println("YearCollection Error:", date, err)
			return
		}

	}
	fmt.Printf("Stock %s %d ~ %d daily data finish!!!\n", StockCode, StartYear, EndYear)
}

// MonthCollection ...
func MonthCollection(StockCode string, StartYear, LastYear int) {
	// Stock Day Data
	var date string
	var err error
	var cacheTime int64 = time.Now().Unix() * 1000
	var collectionflag []map[string]interface{}

	if collectionflag, err = fmsrfk.GetAlreadyDate(StockCode); err != nil {
		fmt.Println("MonthCollection Error:", err)
		return
	}

	for year := StartYear; year > LastYear; year-- {
		if foundation.IsAfterNowTime(year, 1, 1) {
			continue
		}

		date = fmt.Sprintf("%d0101", year)

		if IsInCollectionFlag(date[0:len(date)-2], fmsrfk.CollectionFlagkey, collectionflag) {
			fmt.Printf("Month %s IsSkip!\n", date[0:len(date)-2])
			continue
		}

		err = fmsrfk.CopyData(StockCode, date, cacheTime)
		fmt.Printf("Collection %s daily Month stock\n", date)

		if err != nil {
			fmt.Println("MonthCollection Error:", date, err)
			return
		}
	}
	fmt.Println("Stock " + StockCode + " monthly data finish!!!")
}

// YearCollection ...
func YearCollection(StockCode string) {
	// Stock Year Data
	var err error
	var cacheTime int64 = time.Now().Unix() * 1000
	var collectionflag []map[string]interface{}

	if collectionflag, err = fmnptk.GetAlreadyDate(StockCode); err != nil {
		fmt.Println("YearCollection Error:", err)
		return
	}

	date := strconv.Itoa(time.Now().Year())
	if IsInCollectionFlag(date, fmnptk.CollectionFlagkey, collectionflag) {
		fmt.Printf("Year %s IsSkip!\n", date)
		return
	}

	err = fmnptk.CopyData(StockCode, cacheTime)
	if err != nil {
		fmt.Println("Year CopyData Error:", date, err)
		return
	}
	fmt.Println("Stock " + StockCode + " yearly data finish!!!")
}

func BwibbuCollection(StockCode string, StartYear, LastYear int) {
	var err error
	var cacheTime int64 = time.Now().Unix() * 1000
	var collectionflag []map[string]interface{}

	if collectionflag, err = bwibbu.GetAlreadyDate(StockCode); err != nil {
		fmt.Println("BwibbuCollection Error:", err)
		return
	}

	dates := MonthSlice(StartYear, LastYear, true)
	for _, date := range dates {
		if IsInCollectionFlag(date, bwibbu.CollectionFlagkey, collectionflag) {
			fmt.Printf("%v Bwibbu_Month %s IsSkip!\n", StockCode, date)
			continue
		}

		err = bwibbu.CopyData(StockCode, date, cacheTime)
		if err != nil {
			fmt.Println("Bwibbu_Month CopyData Error:", date, err)
			return
		}
	}
	fmt.Println("Stock " + StockCode + " Bwibbu_Month data finish!!!")
}

func LegalPersonCollection(StartYear, LastYear int) {
	var err error
	var cacheTime int64 = time.Now().Unix() * 1000
	var collectionflag []map[string]interface{}

	if collectionflag, err = legalperson.GetAlreadyDate(); err != nil {
		fmt.Println("LegalPersonCollection Error:", err)
		return
	}

	dates := DaySlice(StartYear, LastYear, true)
	for _, date := range dates {
		if IsInCollectionFlag(date, legalperson.CollectionFlagkey, collectionflag) {
			fmt.Printf("LegalPerson %v IsSkip!\n", date)
			continue
		}

		err = legalperson.CopyData(date, cacheTime)
		if err != nil {
			fmt.Println("LegalPerson CopyData Error:", date, err)
			return
		}
		fmt.Println("Stock " + date + " LegalPerson data finish!!!")
	}
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

func DaySlice(startYear, endYear int, isSkipAfterDay bool) []string {
	var result []string
	if endYear > startYear {
		return []string{}
	}
	currentTime := time.Date(startYear+1, 1, 1, 0, 0, 0, 0, time.UTC)
	for isNext := true; isNext; {
		currentTime = currentTime.AddDate(0, 0, -1)
		year, month, day := currentTime.Date()
		if isSkipAfterDay && currentTime.After(time.Now()) {
			continue
		}
		if year < endYear {
			isNext = false
			continue
		}

		result = append(result, fmt.Sprintf("%d%02d%02d", year, month, day))

	}
	return result
}

func MonthSlice(startYear, endYear int, isSkipAfterDay bool) []string {
	var result []string
	var month int = 12
	if endYear > startYear {
		return []string{}
	}
	for year := startYear; year > endYear; year-- {
		for month = 12; month > 0; month-- {
			if isSkipAfterDay && foundation.IsAfterNowTime(year, month, 1) {
				continue
			}
			result = append(result, fmt.Sprintf("%d%02d01", year, month))
		}
	}
	return result
}

func YearSlice(startYear, endYear int) []string {
	var result []string
	if endYear > startYear {
		return []string{}
	}
	for year := startYear; year > endYear; year-- {
		if foundation.IsAfterNowTime(year, 1, 1) {
			continue
		}
		result = append(result, fmt.Sprintf("%d0101", year))
	}
	return result
}
