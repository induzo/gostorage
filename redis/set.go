package redis

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

// Set return bytes if key is present
func Set(rp *redis.Pool, key, val interface{}) error {
	red := rp.Get()

	defer func() {
		if err := red.Close(); err != nil {
			log.Fatalf("Get(%v) Close: %v", key, err)
		}
	}()

	_, err := red.Do("SET", key, val)
	if err != nil {
		return fmt.Errorf("Exists GET(%s): %v", key, err)
	}

	return nil
}
