package cache

import (
	"github.com/gomodule/redigo/redis"
	"sync"
	"time"
)

var (
	once          sync.Once
	cacheConnPool *redis.Pool
)

var GetCacheConnection = func() redis.Conn {
	once.Do(func() {
		cacheConnPool = &redis.Pool{
			Dial: func() (conn redis.Conn, err error) {
				return redis.Dial("tcp", "localhost:6379")
			},
			MaxIdle:     10,
			IdleTimeout: 240 * time.Second,
		}
	})
	return cacheConnPool.Get()
}
