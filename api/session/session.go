package session

import (
	"sync"
	"time"
	"video-server/api/model"
	"video-server/api/utils"
)

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

//返回现在的时间
func NowInMilli() int64 {
	return time.Now().UnixNano() / 1000000
}

//根据session_id删除session
func DeleteExpiredSession(sid string) {
	//首先从sessionMap中删除
	sessionMap.Delete(sid)
	//然后删除数据库中的
	DeleteSession(sid)
}

//加载数据库中的session信息
func LoadSessionFromDB() {
	r, err := RetrieveAllSessions() //从数据库中获取
	if err != nil {
		return
	}

	r.Range(func(key, value interface{}) bool {
		ss := value.(*model.SimpleSession) //使用断言
		sessionMap.Store(key, ss)
		return true
	})
}

//生成新的sessionId
func GenerateNewSessionId(un string) string {
	id, _ := utils.NewUUID()
	ct := NowInMilli()
	ttl := ct + 30*60*1000 // Severside session valid time: 30 min

	ss := &model.SimpleSession{Username: un, TTL: ttl}
	sessionMap.Store(id, ss)
	InsertSession(id, ttl, un) //插入到数据库
	return id
}

//判断session是否过期
func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	if ok {
		ct := NowInMilli()
		if ss.(*model.SimpleSession).TTL < ct {
			//从数据库中删除过期的sid
			DeleteExpiredSession(sid)
			return "", true
		}
		return ss.(*model.SimpleSession).Username, false
	}
	return "", true
}
