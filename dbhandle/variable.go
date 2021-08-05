package dbhandle

type IDBHandle interface {
	// 詳細交易資訊
	GetTransactiondetail(StockCode string, date string) ([]map[string]interface{}, error)
	SetTransactiondetail(C, D, T, TS, TK0, TK1, TLong, CH, N, NF, Y, Z, IP, TV, A, F, B, G, EX, IT, MT, O, OA, OB, OT, OV, OZ, I, L, H, V, W, U, S, P, PS, PZ string) error

	// GetstockAlreadyDay(StockCode string) ([]map[string]interface{}, error)

	// 歷史個股成交資訊
	Setstockyear(data ...interface{}) error
	Setstockmonth(data ...interface{}) error
	Setstockday(data ...interface{}) error

	// 個股日本益比、殖利率及股價淨值比
	SetBwibbu(data ...interface{}) error

	// 三大法人買賣金額統計表
	SetLegalperson(flag, date string, data interface{}) error

	// 已取得資料紀錄
	Getcollectionflag(StockCode, Flag string) ([]map[string]interface{}, error)
	Setcollectionflag(StockCode, Flag, Date string) error
	// Set
}

// var stockBDSQL *processdb.SQLCLi
type DBHandle struct {
	IDBHandle
}

var Instance IDBHandle
