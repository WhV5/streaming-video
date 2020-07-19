//============================================================
// 描述:
// 作者: Henry
// 日期: 2019-12-04 16:27
// 版权: Henry @Since 2018
//
//============================================================
package dbops

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func AddVideoDeletionRecord(vid string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO video_del_rec(video_id) VALUES (?)")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(vid)
	if err != nil {
		log.Printf("AddVideoDeletetionRecord error : %v", err)
		return err
	}

	defer stmtIns.Close()
	return nil
}
