package localdb

import (
	"fmt"

	"github.com/YWJSonic/ReptileService/dbhandle/localdb/localDBDriver"
)

type LocalDB struct {
	driver localDBDriver.Driver
}

func Connect(setting struct{ Path string }) (db *LocalDB, err error) {
	return &LocalDB{
		// path: setting.Path,
	}, nil
}

// GetTransactiondetail ...
func (self *LocalDB) GetTransactiondetail(StockCode string, date string) ([]map[string]interface{}, error) {

	result, err := self.driver.GetLike("transactiondetail", fmt.Sprintf("%v%v", StockCode, date))
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

	key := fmt.Sprintf("%v%v%v", C, D, T)
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

	Key := fmt.Sprintf("%v%v", data[0].(string), data[1].(int))
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
	_, err := self.driver.Set("stockday", Key, importData)
	if err != nil {
		return err
	}
	return nil
}

// Setstockmonth ...
func (self *LocalDB) Setstockmonth(data ...interface{}) error {
	// Result, err := processdb.CallWrite(stockBDSQL.DB, processdb.MakeProcedureQueryStr("Setstockmonth", len(data)), data...)
	// if err != nil {
	// 	fmt.Println(Result)
	// 	return err
	// }
	return nil
}

// Setstockday ...
func (self *LocalDB) Setstockday(data ...interface{}) error {
	// Result, err := processdb.CallWrite(stockBDSQL.DB, processdb.MakeProcedureQueryStr("Setstockday", len(data)), data...)
	// if err != nil {
	// 	fmt.Println(Result)
	// 	return err
	// }
	return nil
}

// Getcollectionflag ...
func (self *LocalDB) Getcollectionflag(stockCode, flag string) ([]map[string]interface{}, error) {
	Key := fmt.Sprintf("%v_%v", stockCode, flag)
	result, err := self.driver.GetLike("collectionflag", Key)
	if err != nil {
		return nil, err
	}

	convertData := []map[string]interface{}{}
	for _, data := range result {
		convertData = append(convertData, data.(map[string]interface{}))
	}

	return convertData, nil
}

// Setcollectionflag ...
func (self *LocalDB) Setcollectionflag(stockCode, flag, date string) error {
	type insertData struct {
		StockCode, Flag, Date string
	}

	Key := fmt.Sprintf("%v_%v_%v", stockCode, flag, date)
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
