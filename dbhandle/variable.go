package dbhandle

type IDBHandle interface {
	GetTransactiondetail(StockCode string, date string) ([]map[string]interface{}, error)
	SetTransactiondetail(C, D, T, TS, TK0, TK1, TLong, CH, N, NF, Y, Z, IP, TV, A, F, B, G, EX, IT, MT, O, OA, OB, OT, OV, OZ, I, L, H, V, W, U, S, P, PS, PZ string) error
	GetstockAlreadyDay(StockCode string) ([]map[string]interface{}, error)
	Setstockyear(data ...interface{}) error
	Setstockmonth(data ...interface{}) error
	Setstockday(data ...interface{}) error
	Getcollectionflag(StockCode, Flag string) ([]map[string]interface{}, error)
	Setcollectionflag(StockCode, Flag, Date string) error
}

// var stockBDSQL *processdb.SQLCLi
type DBHandle struct {
	IDBHandle
}

var Instance IDBHandle
