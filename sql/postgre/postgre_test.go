/*
 * @Author: Vongola
 * @LastEditTime: 2021-02-08 12:49:46
 * @LastEditors: Vongola
 * @Description: file content
 * @FilePath: \JFFun\sql\postgre\postgre_test.go
 * @Date: 2021-02-08 12:40:21
 * @描述: 文件描述
 */

package postgre

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/lib/pq"
)

func Test_Connect(t *testing.T) {
	db, err := connect(context.Background(), "postgres", "123456", "localhost", "9999", "postgres")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

func connect(ctx context.Context, account string, password string, address string, port string, database string) (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", account, password, address, port, database)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return db, err
	}
	err = db.PingContext(ctx)
	return db, err
}
