package distributedLock

import (
	"context"
	"fmt"
	"testing"
	"time"
)

var options = &ClientOptions{
	Addr:        "127.0.0.1:6379",
	Password:    "",
	Network:     "tcp",
	PoolSize:    100,
	MinIdleConn: 10,
	DialTimeout: 5 * time.Second,
	PoolTimeout: 5 * time.Second,
}
var Rdb = NewClient(options)
var Lock = NewRedisLock("My-Distributed-Lock", Rdb, WithBlock())

func SecondLock() {
	fmt.Println("我要开始第二加锁啦")
	err := Lock.Lock(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("第二次取锁成功")
	Lock.Unlock(context.Background())
	fmt.Println("第二次解锁成功")
}

func Test(t *testing.T) {
	fmt.Println("我要开始加锁啦")
	err := Lock.Lock(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("第一次取锁成功")
	go SecondLock()

	_ = Lock.Unlock(context.Background())
	fmt.Println("第一次解锁成功")
	time.Sleep(20 * time.Second)
}
