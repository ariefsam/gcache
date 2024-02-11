package gcache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var c = cache.New(5*time.Minute, 10*time.Minute)

func Cache[T any](key string, function func() (ret T, err error)) (ret T, err error) {
	if x, found := c.Get(key); found {
		ret = x.(T)
		return
	}
	ret, err = function()
	if err != nil {
		return
	}

	c.Set(key, ret, cache.DefaultExpiration)
	return
}

func SetDefaultCache(duration time.Duration) {
	c = cache.New(duration, duration)
}
