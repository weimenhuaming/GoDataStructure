package distributedLock

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

// Client Redis 封装一个新的客户端（将SetNEX和Eval暴露出来）
type Client struct {
	client *redis.Client
}

type ClientOptions struct {
	Network     string
	Addr        string
	Password    string
	DB          int           // redis数据库
	PoolSize    int           // 连接池提供的连接数量
	MinIdleConn int           // 最小空闲数
	DialTimeout time.Duration // 建立链接的超时
	PoolTimeout time.Duration // 从连接池获取连接的最大等待时间
}

// NewClient 函数式可选模式
func NewClient(options *ClientOptions) *Client {
	// 基础的配置
	redisOpts := &redis.Options{
		Network:      options.Network,
		Addr:         options.Addr,
		Password:     options.Password,
		DB:           0,
		PoolSize:     options.PoolSize,
		MinIdleConns: options.MinIdleConn, // 连接池中保持的最小空闲连接数
		DialTimeout:  options.DialTimeout, // 建立链接的超时
		PoolTimeout:  options.PoolTimeout, // 从连接池获取连接超时
	}
	// 获取服务端
	client := redis.NewClient(redisOpts)
	return &Client{
		client: client,
	}
}

// SetNEX 在键不存在时设置值并指定过期时间（秒）
func (c *Client) SetNEX(ctx context.Context, key, value string, expireSeconds int64) (bool, error) {
	return c.client.SetNX(ctx, key, value, time.Duration(expireSeconds)*time.Second).Result()
}

// Eval 执行Lua脚本
func (c *Client) Eval(ctx context.Context, script string, keyCount int, keysAndArgs []interface{}) (interface{}, error) {
	keys := make([]string, 0, keyCount)
	args := make([]interface{}, 0, len(keysAndArgs)-keyCount)

	for i := 0; i < keyCount; i++ {
		if i >= len(keysAndArgs) {
			return nil, errors.New("insufficient keys in keysAndArgs")
		}
		key, ok := keysAndArgs[i].(string)
		if !ok {
			return nil, errors.New("key is not a string")
		}
		keys = append(keys, key)
	}

	if len(keysAndArgs) > keyCount {
		args = keysAndArgs[keyCount:]
	}

	return c.client.Eval(ctx, script, keys, args...).Result()
}
