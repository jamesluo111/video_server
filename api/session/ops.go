package session

import (
	"sync"
	"time"
	"video_server/api/dbops"
	"video_server/api/defs"
	"video_server/api/utils"
)

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func deleteSession(sid string) {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}

func LoadSessionsFromDB() {
	sessionMap, err := dbops.RetrieveAllSession()
	if err != nil {
		return
	}
	sessionMap.Range(func(key, value interface{}) bool {
		ss := value.(*defs.SimpleSession)
		sessionMap.Store(key, ss)
		return true
	})
}

// 新建一个session并保存到缓存map中和数据库中

func GenerateNewSession(un string) (string, error) {
	sessionId := utils.NewUuid()
	nowTime := time.Now().UnixMilli()
	ttl := nowTime + 30*60*1000
	m := &defs.SimpleSession{UserName: un, TTL: ttl}
	//保存到缓存map中
	sessionMap.Store(sessionId, m)
	//保存到数据库中
	err := dbops.InsertSession(sessionId, ttl, un)
	if err != nil {
		return "", err
	}
	return sessionId, nil
}

//判断session是否过期并取出或删除

func IsSessionExpired(sid string) (string, bool) {
	//从缓存map中去除session
	m, ok := sessionMap.Load(sid)
	if ok {
		nowTime := time.Now().UnixMilli()
		if nowTime > m.(*defs.SimpleSession).TTL {
			deleteSession(sid)
			return "", true
		}
		return m.(*defs.SimpleSession).UserName, false
	}
	return "", true
}
