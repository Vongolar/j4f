package jgate

import (
	jsql "JFFun/database/sql"
	"database/sql"
)

func (m *MGate) getDB() *sql.DB {
	return jsql.GetDB(m.cfg.DataBase)
}
