//============================================================
// 描述:
// 作者: Henry
// 日期: 2019-12-04 16:27
// 版权: Henry @Since 2018
//
//============================================================
package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("mysql", "root:sa@123@tcp(localhost:3306)/video_server?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}
