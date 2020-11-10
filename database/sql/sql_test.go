package jsql

import (
	"testing"
)

func Test_Duple(t *testing.T) {
	ConnectDB("JFFun", "vongola", "123", "localhost:3306")
	db := GetDB("JFFun")
	tx, err := db.Begin()
	if err != nil {
		return
	}

	_, err = tx.Exec("insert into `guest` (`id`,`pid`) values (?,?)", "guest", "1111", "ffae")
	if err != nil {
		tx.Rollback()
		return
	}

	_, err = tx.Exec("insert into `account` (`id`) values (?)", "b20a3bd2-e93d-46cf-8f26-3c389c1fffae")
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
}
