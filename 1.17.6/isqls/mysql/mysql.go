package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/flower/script-go/iutils"
)

func TestISqlsMysqls() {

	/*
		谨慎 字符串替换
	*/
	var sqlBase string

	sqlBase = "SELECT %s FROM table WHERE string LIKE %string%\n"
	log.Println(
		sqlBase,
		strings.Replace(sqlBase, "%s", "count(*)", 1),
		strings.Replace(sqlBase, "%s", "*", 1),
	)

}

/*
	_ "github.com/go-sql-driver/mysql" mysql 驱动

	"database/sql" golang 官方库 标准 接口/API/头文件 无实现 具体实现为 mysql 驱动
	"github.com/jmoiron/sqlx" 则是其扩展

	https://go.dev/doc/database/execute-transactions#best_practices
	http://go-database-sql.org/index.html
	https://jmoiron.github.io/sqlx/

	----

	orm
		sql insert|delete|update|select
			insert
				INSERT INTO {} (xxx, xxx, ..., xxx) VALUES (xxx, xxx, ..., xxx)
				指定字段 & 字段数 小于 数据库实际字段数 & 未指定字段 存在 默认值 则可行

			select
				生产环境 不建议 使用 SELECT * ...
				慢 & 对象字段数 小于 数据库实际字段数 则当 unsafe 为 false 会报错
				对象字段数 大于 数据库实际字段数 则当 unsafe 为 true 也可行
				对象实例 初始化 零值

			生产环境，发版，代码版本 与 数据库版本 兼容性
				先增加 数据库字段
				再更新 代码版本
				再删除 数据库字段

				往往，新增表 或者 预留字段 字段映射(点表)
*/

type Users struct {
	UserId   string `db:"user_id"`
	Username string `db:"username"`
	Balance  int64  `db:"balance"`
	Fake     string `db:"fake"`
	// Email    string `db:"email"`
}

func TestISqlsMysqlSqlx() {
	var (
		users []Users

		sql string = "SELECT * FROM users WHERE balance > ?"

		err error
	)

	db, _ := iutils.NewMySqlSqlxDbClient()
	tx, _ := db.Beginx()
	defer tx.Rollback()

	/*
		func queryx
	*/
	rs, _ := tx.Queryx(
		sql,
		0,
	)

	defer rs.Close()

	// rs.StructScan(users)
	// [mysql] 2023/03/31 00:23:16 packets.go:446: busy buffer
	// [mysql] 2023/03/31 00:23:16 connection.go:173: bad connection
	// ----
	// expected struct but got slice
	// ! slice & array
	for rs.Next() {
		u := Users{}
		err = rs.StructScan(&u)
		if err != nil {
			return
		}
		users = append(users, u)
	}

	iutils.Log(users)

	/*
		func queryrowx
	*/
	u := Users{}

	r := tx.QueryRowx(
		sql,
		0,
	)

	r.StructScan(&u)
	iutils.Log(&u)

	/*
		func select
	*/
	users = []Users{}

	_ = tx.Select(
		&users,
		sql,
		0,
	)

	iutils.Log(users)

	/*
		func exec
		不适用 SELECT
	*/
	_, err = tx.Exec(
		sql,
		0,
	)

	_, err = tx.Exec(
		"UPDATE users SET balance=balance+?",
		7,
	)

	_, err = tx.Exec(
		"UPDATE users SET balance=balance+?",
		-8,
	)

	/*
		func mustexec
		func (tx *Tx) MustExec(query string, args ...interface{}) sql.Result
	*/

	if err = tx.Commit(); err != nil {
		return
	}

	err = db.Close()
	if err != nil {
		return
	}

	return
}

func checkMysqlExecSqlResult(format string, err error, r sql.Result, rowsAffectedMin, rowsAffectedMax int64) (int64, error) {
	if format == "" {
		format = "E CheckMysqlExecSqlResult err: %w"
	}

	if err != nil {
		err = fmt.Errorf(format, err)
		return 0, err
	}

	n, err := r.RowsAffected()
	if err != nil {
		err = fmt.Errorf(format, err)
		return 0, err
	}

	if n < rowsAffectedMin || n > rowsAffectedMax {
		err = fmt.Errorf(
			format,
			fmt.Errorf(
				"rowsAffected %v lt rowsAffectedMin %v or gt rowsAffectedMax %v",
				n,
				rowsAffectedMin,
				rowsAffectedMax,
			),
		)

		return n, err
	}

	return n, err
}

type Iterms struct {
	Id     string  `db:"id"`
	Price  float64 `db:"price"`
	Number int64   `db:"number"`
}

func TestISqlsMysqlSqlxConcurrent() {

	var executeTransaction func(db *sqlx.DB, userId string, itermId string) (err error) = func(db *sqlx.DB, userId string, itermId string) (err error) {
		tx, _ := db.Beginx()
		defer tx.Rollback()

		iterms := []Iterms{}

		err = tx.Select(
			&iterms,
			"SELECT * FROM iterms WHERE number > 0 AND id = ? ",
			itermId,
		)
		if err != nil {
			return iutils.LogError(err)
		}

		if len(iterms) < 1 {
			return iutils.LogError(
				errors.New("E iterms not enough"))
		}

		r, err := tx.Exec(
			"UPDATE users SET balance=balance-? WHERE user_id=? AND balance>=?",
			iterms[0].Price,
			userId,
			iterms[0].Price,
		)
		_, err = checkMysqlExecSqlResult("", err, r, 1, 1)
		if err != nil {
			return iutils.LogError(err)
		}

		r2, err := tx.Exec(
			"UPDATE iterms SET number=number-? WHERE id=? AND number>=?",
			1,
			iterms[0].Id,
			1,
		)
		_, err = checkMysqlExecSqlResult("", err, r2, 1, 1)
		if err != nil {
			return iutils.LogError(err)
		}

		if err = tx.Commit(); err != nil {
			return iutils.LogError(err)
		}

		// ? recycle tx

		return
	}

	// log.Println("I TestISqlsMysqlSqlxConcurrent start")

	db, _ := iutils.NewMySqlSqlxDbClient()

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			err := executeTransaction(db, "1", "2")
			if err != nil {
				log.Printf("I %d fail, err: %v\n", i, err)
				return
			}

			log.Printf("I %d succ\n", i)
		}(i)
	}

	wg.Wait()

	// log.Println("I TestISqlsMysqlSqlxConcurrent over")

	return
}
