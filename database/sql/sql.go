package jsql

import (
	jlog "JFFun/log"
	jtag "JFFun/log/tag"
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

var dbs = make(map[string]*sql.DB)

//GetDB 获取DB
func GetDB(db string) *sql.DB {
	return dbs[db]
}

//ConnectDB 连接数据库
func ConnectDB(db string, usr string, password string, addr string) error {
	cfg := mysql.NewConfig()
	cfg.DBName = db
	cfg.User = usr
	cfg.Passwd = password
	cfg.Net = "tcp"
	cfg.Addr = addr

	database, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		jlog.Error(jtag.DataBase, err)
		return err
	}
	if err = database.Ping(); err != nil {
		jlog.Error(jtag.DataBase, err)
		return err
	}

	dbs[db] = database
	return nil
}

//Close 关闭所有数据库
func Close() {
	for _, v := range dbs {
		v.Close()
	}
}
