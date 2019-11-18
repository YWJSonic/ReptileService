package handledb

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/YWJSonic/ReptileService/MISTWSEcom/stock"
	"github.com/YWJSonic/processdb"
)

// SetInit init value
func SetInit(setting *struct{ DBUser, DBPassword, DBIP, DBPORT, DBName string }) error {
	_, err := connectstockBDSQL(setting)
	if err != nil {
		return err
	}

	return nil
}

// Connect New connect
func connectstockBDSQL(setting *struct{ DBUser, DBPassword, DBIP, DBPORT, DBName string }) (db *sql.DB, err error) {
	if stockBDSQL == nil {
		stockBDSQL = new(processdb.SQLCLi)
		sqlstr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&timeout=30s", setting.DBUser, setting.DBPassword, setting.DBIP, setting.DBPORT, setting.DBName)
		db, err := sql.Open("mysql", sqlstr)

		connMaxLifetime := 59 * time.Second
		maxIdleConns := 50
		maxOpenConns := 50

		db.SetConnMaxLifetime(time.Duration(connMaxLifetime))
		db.SetMaxIdleConns(maxIdleConns)
		db.SetMaxOpenConns(maxOpenConns)
		if err != nil {
			return nil, err
		}

		stockBDSQL.DB = db
	}

	return stockBDSQL.DB, nil
}

// SetTransactiondetail ...
func SetTransactiondetail(data stock.MsgInfo) error {
	_, err := processdb.CallWrite(stockBDSQL.DB, processdb.MakeProcedureQueryStr("SetTransactiondetail", 37), data.C, data.D, data.T, data.TS, data.TK0, data.TK1, data.TLong, data.CH, data.N, data.NF, data.Y, data.Z, data.IP, data.TV, data.A, data.F, data.B, data.G, data.EX, data.IT, data.MT, data.O, data.OA, data.OB, data.OT, data.OV, data.OZ, data.I, data.L, data.H, data.V, data.W, data.U, data.S, data.P, data.PS, data.PZ)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// Setstockyear ...
func Setstockyear(data ...interface{}) error {
	Result, err := processdb.CallWrite(stockBDSQL.DB, processdb.MakeProcedureQueryStr("Setstockyear", len(data)), data...)
	if err != nil {
		fmt.Println(Result)
		return err
	}
	return nil
}

// Setstockmonth ...
func Setstockmonth(data ...interface{}) error {
	Result, err := processdb.CallWrite(stockBDSQL.DB, processdb.MakeProcedureQueryStr("Setstockmonth", len(data)), data...)
	if err != nil {
		fmt.Println(Result)
		return err
	}
	return nil
}

// Setstockday ...
func Setstockday(data ...interface{}) error {
	Result, err := processdb.CallWrite(stockBDSQL.DB, processdb.MakeProcedureQueryStr("Setstockday", len(data)), data...)
	if err != nil {
		fmt.Println(Result)
		return err
	}
	return nil
}
