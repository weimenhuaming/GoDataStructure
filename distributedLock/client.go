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
	Network        string
	Addr           string
	Password       string
	DB             int           // redis数据库
	PoolSize       int           // 连接池大小
	MinIdleConns   int           // 最小空闲数
	IdleTimeout    time.Duration //
	PoolTimeout    time.Duration // 连接池最大
	MaxActiveConns int
	Wait           bool
}

type ClientOption func(*ClientOptions)

// NewClient 函数式可选模式
func NewClient(address, password string, opts ...ClientOption) *Client {
	options := &ClientOptions{
		Addr:     address,
		Password: password,
		Network:  "tcp",
	}

	for _, opt := range opts {
		opt(options)
	}

	// 初始化数据
	repairClientOptions(options)

	redisOpts := &redis.Options{
		Network:      options.Network,
		Addr:         options.Addr,
		Password:     options.Password,
		DB:           options.DB,
		PoolSize:     options.PoolSize,
		MinIdleConns: options.MinIdleConns, // 连接池中保持的最小空闲连接数
		DialTimeout:  options.IdleTimeout,  // 建立链接的超时
		PoolTimeout:  options.PoolTimeout,  // 从连接池获取连接超时
	}

	client := redis.NewClient(redisOpts)
	return &Client{
		client: client,
	}
}

func repairClientOptions(opts *ClientOptions) {
	if opts.PoolSize <= 0 {
		opts.PoolSize = 10
	}
	if opts.IdleTimeout <= 0 {
		opts.IdleTimeout = 300 * time.Second
	}
	if opts.PoolTimeout <= 0 {
		if opts.Wait {
			opts.PoolTimeout = 30 * time.Second
		} else {
			opts.PoolTimeout = 0
		}
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
