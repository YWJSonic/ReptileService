package mysql

import (
	"database/sql"
	"errors"
)

type Mysql struct {
}

func Connect(setting *struct{ DBUser, DBPassword, DBIP, DBPORT, DBName string }) (db *sql.DB, err error) {
	// if stockBDSQL == nil {
	// 	stockBDSQL = new(processdb.SQLCLi)
	// 	sqlstr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&timeout=30s", setting.DBUser, setting.DBPassword, setting.DBIP, setting.DBPORT, setting.DBName)
	// 	db, err := sql.Open("mysql", sqlstr)

	// 	connMaxLifetime := 59 * time.Second
	// 	maxIdleConns := 50
	// 	maxOpenConns := 50

	// 	db.SetConnMaxLifetime(time.Duration(connMaxLifetime))
	// 	db.SetMaxIdleConns(maxIdleConns)
	// 	db.SetMaxOpenConns(maxOpenConns)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	stockBDSQL.DB = db
	// }

	// return stockBDSQL.DB, nil
	return nil, errors.New("db fail")
}

// GetTransactiondetail ...
func (self *Mysql) GetTransactiondetail(StockCode string, date string) ([]map[string]interface{}, error) {
	// result, err := processdb.CallReadOutMap(stockBDSQL.DB, "GetTransactiondetail", StockCode, date)
	// if err != nil {
	// 	return nil, err
	// }
	// return result, nil
	return nil, nil
}

// SetTransactiondetail ...
func (self *Mysql) SetTransactiondetail(C, D, T, TS, TK0, TK1, TLong, CH, N, NF, Y, Z, IP, TV, A, F, B, G, EX, IT, MT, O, OA, OB, OT, OV, OZ, I, L, H, V, W, U, S, P, PS, PZ string) error {
	// _, err := processdb.CallWrite(stockBDSQL.DB, processdb.MakeProcedureQueryStr("SetTransactiondetail", 37), C, D, T, TS, TK0, TK1, TLong, CH, N, NF, Y, Z, IP, TV, A, F, B, G, EX, IT, MT, O, OA, OB, OT, OV, OZ, I, L, H, V, W, U, S, P, PS, PZ)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }
	return nil
}

// GetstockAlreadyDay ...
func (self *Mysql) GetstockAlreadyDay(StockCode string) ([]map[string]interface{}, error) {
	// result, err := processdb.CallReadOutMap(stockBDSQL.DB, "GetstockAlreadyDay", StockCode)
	// if err != nil {
	// 	return nil, err
	// }
	// return result, nil
	return nil, nil
}

// Setstockyear ...
func (self *Mysql) Setstockyear(data ...interface{}) error {
	// Result, err := processdb.CallWrite(stockBDSQL.DB, processdb.MakeProcedureQueryStr("Setstockyear", len(data)), data...)
	// if err != nil {
	// 	fmt.Println(Result)
	// 	return err
	// }
	return nil
}

// Setstockmonth ...
func (self *Mysql) Setstockmonth(data ...interface{}) error {
	// Result, err := processdb.CallWrite(stockBDSQL.DB, processdb.MakeProcedureQueryStr("Setstockmonth", len(data)), data...)
	// if err != nil {
	// 	fmt.Println(Result)
	// 	return err
	// }
	return nil
}

// Setstockday ...
func (self *Mysql) Setstockday(data ...interface{}) error {
	// Result, err := processdb.CallWrite(stockBDSQL.DB, processdb.MakeProcedureQueryStr("Setstockday", len(data)), data...)
	// if err != nil {
	// 	fmt.Println(Result)
	// 	return err
	// }
	return nil
}

// Getcollectionflag ...
func (self *Mysql) Getcollectionflag(StockCode, Flag string) ([]map[string]interface{}, error) {
	// result, err := processdb.CallReadOutMap(stockBDSQL.DB, "Getcollectionflag", StockCode, Flag)
	// if err != nil {
	// 	return nil, err
	// }
	// return result, nil
	return nil, nil
}

// Setcollectionflag ...
func (self *Mysql) Setcollectionflag(StockCode, Flag, Date string) error {
	// Result, err := processdb.CallWrite(stockBDSQL.DB, processdb.MakeProcedureQueryStr("Setcollectionflag", 3), StockCode, Flag, Date)
	// if err != nil {
	// 	fmt.Println(Result)
	// 	return err
	// }
	return nil
}
