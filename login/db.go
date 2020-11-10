package jlogin

import (
	jredis "JFFun/database/redis"
	jsql "JFFun/database/sql"
	jtable "JFFun/database/table"
	"database/sql"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func (m *MLogin) getDB() *sql.DB {
	return jsql.GetDB(m.cfg.DataBase)
}

func (m *MLogin) getCache() *redis.Client {
	return jredis.GetDB(m.cfg.Cache)
}

func (m *MLogin) createDBTables() error {
	if err := m.createAccountTable(); err != nil {
		return err
	}

	if err := m.createGuestTable(); err != nil {
		return err
	}

	return nil
}

func (m *MLogin) createAccountTable() error {
	if _, err := m.getDB().Exec(fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s`(\n`id` CHAR(50) NOT NULL,\n`auth` INT NOT NULL,\nPRIMARY KEY (`id`)\n)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4", jtable.Account)); err != nil {
		return err
	}
	return nil
}

func (m *MLogin) createGuestTable() error {
	if _, err := m.getDB().Exec(fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s`(\n`id` CHAR(50) NOT NULL,\n`pid` CHAR(50) NOT NULL,\nPRIMARY KEY (`id`),\nUNIQUE (`pid`)\n)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4", jtable.Guest)); err != nil {
		return err
	}
	return nil
}
