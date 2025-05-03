package distributedLock

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {
	options := &ClientOptions{
		Addr:        "127.0.0.1:6379",
		Password:    "",
		Network:     "tcp",
		PoolSize:    100,
		MinIdleConn: 10,
		DialTimeout: 5 * time.Second,
		PoolTimeout: 5 * time.Second,
	}
	Rdb := NewClient(options)
	Lock := NewRedisLock("My-Distributed-Lock", Rdb, WithBlock())
	err := Lock.Lock(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	//_ = Lock.Unlock(context.Background())
}
