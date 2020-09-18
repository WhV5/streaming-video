//============================================================
// 描述:
// 作者: Henry
// 日期: 2019-12-03 17:21
// 版权: Henry @Since 2018
//
//============================================================
package session

import (
	"github.com/StreamingVideo/api/dbops"
	"github.com/StreamingVideo/api/defs"
	"github.com/StreamingVideo/api/utils"
	"sync"
	"time"
)

type SimpleSession struct {
	Username string
	TTL      int64
}

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func nowIntMilli() int64 {
	return time.Now().UnixNano() / 100000
}

func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}

func LoadSessionsFromDB() {
	r, err := dbops.RetrieveAllSessions()
	if err != nil {
		return
	}
	r.Range(func(k, v interface{}) bool {
		ss := v.(*defs.SimpleSession)
		sessionMap.Store(k, ss)
		return true
	})
}

func GenerateNewSessionId(un string) string {
	id, _ := utils.NewUUID()
	ct := nowIntMilli()
	ttl := ct + 30*60*1000

	ss := &defs.SimpleSession{
		Username: un,
		TTL:      ttl,
	}
	sessionMap.Store(id, ss)
	dbops.InsertSession(id, ttl, un)

	return id
}

func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	if ok {
		ct := nowIntMilli()
		if ss.(*defs.SimpleSession).TTL < ct {
			//delete expired session
			deleteExpiredSession(sid)
			return "", true
		}
		return ss.(*defs.SimpleSession).Username, false
	}
	return "", true
}
