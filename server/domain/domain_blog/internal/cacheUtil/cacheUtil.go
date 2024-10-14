package cacheUtil

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var Cache *cache.Cache

func InitCache() {
	c := cache.New(7*24*time.Hour, 24*time.Hour)
	Cache = c
}

func Set(key string, value interface{}, d time.Duration) {
	Cache.Set(key, value, d)
}

func SetDefault(key string, value interface{}) {
	Cache.SetDefault(key, value)
}

func Get(key string) (interface{}, bool) {
	return Cache.Get(key)
}

func Delete(key string) {
	Cache.Delete(key)
}
