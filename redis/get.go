package redis

import (
	"errors"
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

var ErrNotFound = errors.New("key %s not found")

// Get return bytes if key is present
func Get(rp *redis.Pool, key string) (interface{}, error) {
	red := rp.Get()

	defer func() {
		if err := red.Close(); err != nil {
			log.Fatalf("Get(%v) Close: %v", key, err)
		}
	}()

	rawBytes, err := red.Do("GET", key)
	if err != nil {
		return nil, fmt.Errorf("Exists GET(%s): %v", key, err)
	}

	if rawBytes == nil {
		return nil, ErrNotFound
	}

	return rawBytes, nil
}
