package distributedLock

import (
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	goredislib "github.com/redis/go-redis/v9"
)

func test() {
	// 创建多个独立Redis实例的连接池
	var pools []redis.Pool
	// 对应的地址
	addrs := []string{
		"localhost:6379",
		"localhost:6380",
		"localhost:6381",
		"localhost:6382",
		"localhost:6383", // 推荐使用奇数个节点（如5个）
	}

	for _, addr := range addrs {
		client := goredislib.NewClient(&goredislib.Options{
			Addr: addr, // 每个地址对应独立Redis实例
		})
		// 将 go-redis 客户端转换为 redsync 需要的 Pool 接口
		pool := goredis.NewPool(client)
		pools = append(pools, pool)
	}

	// 创建 redsync 实例（需传入多个 Pool）
	rs := redsync.New(pools...)

	// 创建互斥锁
	mutex := rs.NewMutex("my-global-mutex",
		redsync.WithExpiry(10*time.Second), // 建议设置超时
	)

	// 获取锁
	if err := mutex.Lock(); err != nil {
		panic("获取锁失败: " + err.Error())
	}

	// 业务逻辑...

	// 释放锁
	if ok, err := mutex.Unlock(); !ok || err != nil {
		panic("释放锁失败: " + err.Error())
	}
}
