package redislock

import (
	"context"
	"github.com/redis/go-redis/v9"
	"golang.org/x/sync/errgroup"
	"log"
	"testing"
	"time"
)

func TestNewRedisLocker(t *testing.T) {
	locker := NewRedisLocker(redis.NewClient(&redis.Options{
		Addr: "localhost:6380",
	}))
	ctx := context.Background()
	ok, err := locker.Lock(ctx, "test", time.Second*10)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("lock failed")
	}
	willFail, err := locker.Lock(ctx, "test", time.Second*10)
	if err != nil {
		t.Fatal(err)
	}
	if willFail {
		t.Fatal("should lock failed")
	}
	ok, err = locker.UnLock(ctx, "test")
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("unlock failed")
	}
	ok, err = locker.Lock(ctx, "test", time.Second*10)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("lock failed")
	}
	_, _ = locker.UnLock(ctx, "test")

}

func TestWithLocker(t *testing.T) {
	locker := NewRedisLocker(redis.NewClient(&redis.Options{
		Addr: "localhost:6380",
	}))
	var wg errgroup.Group
	wg.Go(func() error {
		err := WithLocker(context.Background(), locker, "test", time.Second*10, 3, func() error {
			log.Println("the first success")
			time.Sleep(time.Second * 5)
			return nil
		})
		if err != nil {
			t.Log(err)
		}
		return nil
	})
	wg.Go(func() error {
		err := WithLocker(context.Background(), locker, "test", time.Second*10, 3, func() error {
			log.Println("the twice success")
			time.Sleep(time.Second * 5)
			return nil
		})
		if err != nil {
			t.Log(err)
		}
		return nil
	})
	err := wg.Wait()
	if err != nil {
		t.Log(err)
	}
}
