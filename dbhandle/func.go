package dbhandle

import (
	"errors"

	"github.com/YWJSonic/ReptileService/dbhandle/localdb"
)

// DB 整合工具
func NewDBHandle() *DBHandle {
	return &DBHandle{}
}

// 連接DB
func (self *DBHandle) ConnectMysql(setting struct{ DBUser, DBPassword, DBIP, DBPORT, DBName string }) error {
	return errors.New("db connect fail")
}

func (self *DBHandle) ConnectLocalDB(setting struct{ Path string }) error {
	localDB, err := localdb.Connect(setting)
	if err != nil {
		return err
	}

	self.IDBHandle = localDB
	return nil
}
