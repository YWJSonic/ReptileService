package localdb

import (
	"fmt"

	"github.com/YWJSonic/ReptileService/constants"
	"github.com/YWJSonic/ReptileService/dbhandle/localdb/localDBDriver"
)

type LocalDB struct {
	driver *localDBDriver.Driver
}

func Connect(setting struct{ Path string }) (db *LocalDB, err error) {
	return &LocalDB{
		localDBDriver.NewDriver(setting),
	}, nil
}

func FormatKey(datas ...interface{}) string {
	var key string
	for _, data := range datas {
		key += fmt.Sprintf("%v_", data)
	}
	if len(key) > 0 {
		key = key[:len(key)-1]
	}
	return key
}

// GetTransactiondetail ...
func (self *LocalDB) GetTransactiondetail(StockCode string, date string) ([]map[string]interface{}, error) {

	result, err := self.driver.GetLike("transactiondetail", FormatKey(StockCode, date))
	if err != nil {
		return nil, err
	}

	convertData := []map[string]interface{}{}
	for _, data := range result {
		convertData = append(convertData, data.(map[string]interface{}))
	}

	return convertData, nil
}

// SetTransactiondetail ...
func (self *LocalDB) SetTransactiondetail(C, D, T, TS, TK0, TK1, TLong, CH, N, NF, Y, Z, IP, TV, A, F, B, G, EX, IT, MT, O, OA, OB, OT, OV, OZ, I, L, H, V, W, U, S, P, PS, PZ string) error {
	type insertData struct {
		C, D, T, TS, TK0, TK1, TLong, CH, N, NF, Y, Z, IP, TV, A, F, B, G, EX, IT, MT, O, OA, OB, OT, OV, OZ, I, L, H, V, W, U, S, P, PS, PZ string
	}

	key := FormatKey(C, D, T)
	data := insertData{C, D, T, TS, TK0, TK1, TLong, CH, N, NF, Y, Z, IP, TV, A, F, B, G, EX, IT, MT, O, OA, OB, OT, OV, OZ, I, L, H, V, W, U, S, P, PS, PZ}
	_, err := self.driver.Set("transactiondetail", key, data)
	if err != nil {
		return err
	}
	return nil
}

// GetstockAlreadyDay ...
func (self *LocalDB) GetstockAlreadyDay(StockCode string) ([]map[string]interface{}, error) {

	result, err := self.driver.GetLike("stockday", StockCode)
	if err != nil {
		return nil, err
	}

	convertData := []map[string]interface{}{}
	for _, data := range result {
		convertData = append(convertData, data.(map[string]interface{}))
	}
	return convertData, nil
}

// Setstockyear ...
func (self *LocalDB) Setstockyear(data ...interface{}) error {
	type insertData struct {
		StockCode       string
		Year            int
		StockPrice      string
		StockCount      string
		DealCount       string
		TopPrice        string
		TopPriceDate    string
		BottomPrice     string
		BottomPriceDate string
		AVGPrice        string
	}

	Key := FormatKey(data[0].(string), data[1].(int))
	importData := insertData{
		StockCode:       data[0].(string),
		Year:            data[1].(int),
		StockPrice:      data[2].(string),
		StockCount:      data[3].(string),
		DealCount:       data[4].(string),
		TopPrice:        data[5].(string),
		TopPriceDate:    data[6].(string),
		BottomPrice:     data[7].(string),
		BottomPriceDate: data[8].(string),
		AVGPrice:        data[9].(string),
	}
	_, err := self.driver.Set("stockyear", Key, importData)
	if err != nil {
		return err
	}
	return nil
}

// Setstockmonth ...
func (self *LocalDB) Setstockmonth(data ...interface{}) error {
	type insertData struct {
		Stockcode       string `json:"stockcode"`
		Datayear        int    `json:"datayear"`
		Datamonth       int    `json:"datamonth"`
		Stockprice      string `json:"stockprice"`
		Stockcount      string `json:"stockcount"`
		Dealcount       string `json:"dealcount"`
		Weightsavgprice string `json:"weightsavgprice"`
		Topprice        string `json:"topprice"`
		Bottomprice     string `json:"bottomprice"`
		Turnover        string `json:"turnover"`
	}

	Key := FormatKey(data[0].(string), data[1].(int), data[2].(int))
	importData := insertData{
		Stockcode:       data[0].(string),
		Datayear:        data[1].(int),
		Datamonth:       data[2].(int),
		Stockprice:      data[3].(string),
		Stockcount:      data[4].(string),
		Dealcount:       data[5].(string),
		Weightsavgprice: data[6].(string),
		Topprice:        data[7].(string),
		Bottomprice:     data[8].(string),
		Turnover:        data[9].(string),
	}
	_, err := self.driver.Set("stockmonth", Key, importData)
	if err != nil {
		return err
	}
	return nil
}

// Setstockday ...
func (self *LocalDB) Setstockday(data ...interface{}) error {
	type insertData struct {
		Stockcode   string `json:"stockcode"`
		Datayear    int    `json:"datayear"`
		Datamonth   int    `json:"datamonth"`
		Dataday     int    `json:"dataday"`
		Stockprice  string `json:"stockprice"`
		Stockcount  string `json:"stockcount"`
		Openprice   string `json:"openprice"`
		Closeprice  string `json:"closeprice"`
		Topprice    string `json:"topprice"`
		Bottomprice string `json:"bottomprice"`
		Diffprice   string `json:"diffprice"`
		Dealcount   string `json:"dealcount"`
	}

	Key := FormatKey(data[0].(string), data[1].(int), data[2].(int), data[3].(int))
	importData := insertData{
		Stockcode:   data[0].(string),
		Datayear:    data[1].(int),
		Datamonth:   data[2].(int),
		Dataday:     data[3].(int),
		Stockprice:  data[4].(string),
		Stockcount:  data[5].(string),
		Openprice:   data[6].(string),
		Closeprice:  data[7].(string),
		Topprice:    data[8].(string),
		Bottomprice: data[9].(string),
		Diffprice:   data[10].(string),
		Dealcount:   data[11].(string),
	}
	_, err := self.driver.Set("stockday", Key, importData)
	if err != nil {
		return err
	}
	return nil
}

// Getcollectionflag ...
func (self *LocalDB) Getcollectionflag(stockCode, flag string) ([]map[string]interface{}, error) {
	Key := FormatKey(stockCode, flag)
	result, err := self.driver.GetLike("collectionflag", Key)
	if err != nil {
		if err.Error() == constants.NoData {
			return make([]map[string]interface{}, 0), nil
		}
		return nil, err
	}

	convertData := []map[string]interface{}{}
	for _, idata := range result {
		data := idata.(map[string]interface{})
		if data["StockCode"].(string) != stockCode {
			continue
		}

		convertData = append(convertData, data)
	}

	return convertData, nil
}

// Setcollectionflag ...
func (self *LocalDB) Setcollectionflag(stockCode, flag, date string) error {
	type insertData struct {
		StockCode, Flag, Date string
	}

	Key := FormatKey(stockCode, flag, date)
	importData := insertData{
		StockCode: stockCode,
		Flag:      flag,
		Date:      date,
	}
	_, err := self.driver.Set("collectionflag", Key, importData)
	if err != nil {
		return err
	}
	return nil
}

func (self *LocalDB) SetBwibbu(data ...interface{}) error {
	type insertData struct {
		StockCode       string `json:"stockcode"`
		Date            string `json:"date"`
		YieldRate       string `json:"yieldrate"`
		DividendYear    int    `json:"dividendyear"`
		PeRatio         string `json:"peratio"`
		WorthRatio      string `json:"worthratio"`
		FinancialReport string `json:"financialreport"`
	}

	Key := FormatKey(data[0].(string), data[1].(string))
	importData := insertData{
		StockCode:       data[0].(string),
		Date:            data[1].(string),
		YieldRate:       data[2].(string),
		DividendYear:    data[3].(int),
		PeRatio:         data[4].(string),
		WorthRatio:      data[5].(string),
		FinancialReport: data[6].(string),
	}
	_, err := self.driver.Set("bwibbu", Key, importData)
	if err != nil {
		return err
	}
	return nil
}

func (self *LocalDB) SetLegalperson(flag, date string, data interface{}) error {

	Key := FormatKey(flag, date)
	_, err := self.driver.Set("legalperson", Key, data)
	if err != nil {
		return err
	}
	return nil
}
