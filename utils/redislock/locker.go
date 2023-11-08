package redislock

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type Locker interface {
	Lock(ctx context.Context, key string, duration time.Duration) (bool, error)
	UnLock(ctx context.Context, key string) (bool, error)
}

type RedisLocker struct {
	redisClient *redis.Client
}

func NewRedisLocker(redisClient *redis.Client) Locker {
	return &RedisLocker{redisClient: redisClient}
}

func (r *RedisLocker) Lock(ctx context.Context, key string, duration time.Duration) (bool, error) {
	return r.redisClient.SetNX(ctx, key, 1, duration).Result()
}

func (r *RedisLocker) UnLock(ctx context.Context, key string) (bool, error) {
	i, err := r.redisClient.Del(ctx, key).Result()
	return i > 0, err
}

var ErrLockFailed = errors.New("lock failed within retry times")

// WithLocker unlock whatever callback error happened
func WithLocker(ctx context.Context, locker Locker, key string, duration time.Duration, retries int, f func() error) error {
	var err error
	for ok := true; ok; ok = retries > 0 {
		ok, err = locker.Lock(ctx, key, duration)
		if err != nil {
			return err
		}
		if !ok {
			fmt.Println("lock failed")
			time.Sleep(time.Second)
			retries--
			continue
		} else {
			fmt.Println("lock success")
			err = f()
			_, _ = locker.UnLock(ctx, key)
			return err
		}
	}
	return ErrLockFailed
}

// WithLockerOnlyWhenErrorUnlock unlock only when callback error happened
func WithLockerOnlyWhenErrorUnlock(ctx context.Context, locker Locker, key string, duration time.Duration, retries int, f func() error) error {
	var err error
	for ok := true; ok; ok = retries > 0 {
		ok, err = locker.Lock(ctx, key, duration)
		if err != nil {
			return err
		}
		if !ok {
			time.Sleep(time.Second)
			retries--
			continue
		} else {
			err = f()
			if err != nil {
				_, _ = locker.UnLock(ctx, key)
				return err
			}
			return nil
		}
	}
	return ErrLockFailed
}
