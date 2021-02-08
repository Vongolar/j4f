/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-08 12:38:33
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\modules\account\account_test.go
 * @Date: 2021-02-08 10:51:14
 * @描述: 文件描述
 */
package account

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/lib/pq"
)

func Test_ConnectSql(t *testing.T) {
	connStr := "postgres://postgres:123456@localhost:9999/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	res, err := db.Exec(`DROP TABLE IF EXISTS account;`)
	if err != nil {
		log.Fatal("create err : ", err)
	} else {
		log.Fatal("create res : ", res)
	}
}
