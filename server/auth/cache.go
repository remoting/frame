package auth

import "time"

// UserCache 缓存
type UserCache struct {
	LifeCircle int64
	UserInfo   UserInfo
	TouchTime  int64
}

// TokenCache 缓存用户与token的映射关系，后期需要定期清理功能
var TokenCache map[string]*UserCache = make(map[string]*UserCache, 0)

func Put(key string, userInfo *UserCache) {
	TokenCache[key] = userInfo
}

func Get(token string) UserInfo {
	cache, ok := TokenCache[token]
	if ok {
		if cache.LifeCircle <= 0 {
			return cache.UserInfo
		}
		ctime := time.Now().UnixNano() / int64(time.Millisecond)
		//换成失效判断
		if ctime-cache.TouchTime > cache.LifeCircle {
			delete(TokenCache, token)
			return nil
		}
		cache.TouchTime = ctime
		return cache.UserInfo
	}
	return nil
}
func TouchUserInfoTime(token string) {
	TokenCache[token].TouchTime = time.Now().UnixNano() / int64(time.Millisecond)
}
