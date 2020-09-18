//============================================================
// 描述:
// 作者: Henry
// 日期: 2019-12-03 17:29
// 版权: Henry @Since 2018
//
//============================================================
package dbops

import (
	"database/sql"
	"github.com/StreamingVideo/api/defs"
	"log"
	"strconv"
	"sync"
)

func InsertSession(sid string, ttl int64, uName string) error {
	ttlStr := strconv.FormatInt(ttl, 10)
	stmtIns, err := dbConn.Prepare("INSERT INTO sessions (session_id,TTL,login_name) VALUES (?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(sid, ttlStr, uName)
	if err != nil {
		return err
	}

	defer stmtIns.Close()
	return nil
}

func RetrieveSession(sid string) (*defs.SimpleSession, error) {
	ss := &defs.SimpleSession{}

	stmtOut, err := dbConn.Prepare("SELECT TTL,login_name FROM sessions WHERE session_id= ?")
	if err != nil {
		return nil, err
	}
	var ttl string
	var uName string
	err = stmtOut.QueryRow(sid).Scan(&ttl, &uName)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if res, err := strconv.ParseInt(ttl, 10, 64); err == nil {
		ss.TTL = res
		ss.Username = uName
	} else {
		return nil, err
	}
	defer stmtOut.Close()
	return ss, nil
}

func RetrieveAllSessions() (*sync.Map, error) {
	m := &sync.Map{}
	stmtOut, err := dbConn.Prepare("SELECT * FROM sessions")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	rows, err := stmtOut.Query()
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	for rows.Next() {
		var id string
		var ttlStr string
		var loginName string
		if err := rows.Scan(&id, &ttlStr, &loginName); err != nil {
			log.Printf("retrieve sessions error : %s", err)
			break
		}
		if ttl, err1 := strconv.ParseInt(ttlStr, 10, 64); err1 == nil {
			ss := &defs.SimpleSession{
				Username: loginName,
				TTL:      ttl,
			}
			m.Store(id, ss)
			log.Printf("session id : %s,ttl:%d", id, ss.TTL)
		}
	}
	return m, nil
}

func DeleteSession(sid string) error {
	stmtOut, err := dbConn.Prepare("DELETE FROM sessions WHERE session_id = ?")
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	if _, err := stmtOut.Query(sid); err != nil {
		return err
	}
	return nil
}
