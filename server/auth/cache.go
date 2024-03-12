package auth

import (
	"sync"
	"time"
)

// UserCache 缓存
type UserCache struct {
	LifeCircle int64
	UserInfo   UserInfo
	TouchTime  int64
}

// TokenCache 缓存用户与token的映射关系，后期需要定期清理功能
var tokenCache = make(map[string]*UserCache, 0)
var lock = &sync.Mutex{}

func Put(key string, userInfo *UserCache) {
	lock.Lock()
	defer lock.Unlock()
	tokenCache[key] = userInfo
}

func Get(token string) UserInfo {
	lock.Lock()
	defer lock.Unlock()
	cache, ok := tokenCache[token]
	if ok {
		if cache.LifeCircle <= 0 {
			return cache.UserInfo
		}
		ctime := time.Now().UnixNano() / int64(time.Millisecond)
		//换成失效判断
		if ctime-cache.TouchTime > cache.LifeCircle {
			delete(tokenCache, token)
			return nil
		}
		cache.TouchTime = ctime
		return cache.UserInfo
	}
	return nil
}
func TouchUserInfoTime(token string) {
	lock.Lock()
	defer lock.Unlock()
	tokenCache[token].TouchTime = time.Now().UnixNano() / int64(time.Millisecond)
}
