package caching

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"url-shortner/config/cache"
)

var GetURLFromCache = func(shortURLID string) (string, error) {
	resp, err := redis.String(cache.GetCacheConnection().Do("GET", shortURLID))
	if err != nil {
		log.Printf("%v error happened while retrieving from cache", err.Error())
	}
	return resp, err
}

var SetURLIntoCache = func(key string, value interface{}) {
	_, err := cache.GetCacheConnection().Do("SET", key, value)
	if err != nil {
		log.Printf("%v error happened while storing in cache", err.Error())
	}
	_, errExpiry := cache.GetCacheConnection().Do("EXPIRE", key, 300)
	if errExpiry != nil {
		log.Printf("%v error happened in setting expiry", errExpiry.Error())
	}
}
