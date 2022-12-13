package cache

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/encse/altnet/lib/config"
	log "github.com/sirupsen/logrus"
)

func Cached[T any](key string, expiration time.Duration, f func() (T, error)) (T, error) {
	res, err := Get[T](key)
	if err != nil {
		res, err = f()
		if err != nil {
			return res, err
		}
		return res, Set(key, res, expiration)
	} else {
		log.Print(fmt.Sprintf("Return %s from cache", key))
	}
	return res, nil
}

func Get[T any](key string) (T, error) {
	var result T

	mc := memcache.New(config.Get().Memcached.Server)
	item, err := mc.Get(key)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(item.Value, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func Set[T any](key string, value T, expiration time.Duration) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	item := &memcache.Item{
		Key:        key,
		Value:      bytes,
		Expiration: int32(expiration.Seconds()),
	}
	mc := memcache.New(config.Get().Memcached.Server)
	return mc.Set(item)
}
