package iutils

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	user     = "root"
	password = "root"
	addr     = "127.0.0.1"
	dbStr    = "test"
)

func NewMySqlSqlxDbClient() (cli *sqlx.DB, err error) {
	path := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8", user, password, addr, dbStr,
	)

	// *sqlx.DB 线程安全 & 并发安全
	cli, err = sqlx.Open("mysql", path)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// defer cli.Close()

	/*
		unsafe
	*/
	cli = cli.Unsafe()

	// ping
	err = cli.Ping()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return
}
