//go:build redis
// +build redis

package ss

import (
	"github.com/go-redis/redis"
	"github.com/wwqk4444/ccsexyz_utils"
)

type redisFilterImpl struct {
	*redis.Client
}

func (r *redisFilterImpl) TestAndAdd(k []byte) bool {
	if r == nil || r.Client == nil {
		return false
	}
	_, err := r.Client.GetSet(utils.SliceToString(k), "true").Result()
	return err == nil
}

func newRedisFilter(addr, pass string, db int) bytesFilter {
	return &redisFilterImpl{
		Client: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: pass,
			DB:       db,
		}),
	}
}
