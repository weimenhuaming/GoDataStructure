package distributedLock

import (
	"context"
	"errors"
	"time"
)

// 红锁中每个节点默认的处理超时时间为 50 ms
const DefaultSingleLockTimeout = 50 * time.Millisecond

type RedLock struct {
	locks          []*RedisLock // 内置对所有节点的锁
	RedLockOptions              // 红锁的配置
}

type RedLockOption func(*RedLockOptions)

type RedLockOptions struct {
	singleNodesTimeout time.Duration
	expireDuration     time.Duration
}

func WithSingleNodesTimeout(singleNodesTimeout time.Duration) RedLockOption {
	return func(o *RedLockOptions) {
		o.singleNodesTimeout = singleNodesTimeout
	}
}
func WithRedLockExpireDuration(expireDuration time.Duration) RedLockOption {
	return func(o *RedLockOptions) {
		o.expireDuration = expireDuration
	}
}
func repairRedLock(o *RedLockOptions) {
	if o.singleNodesTimeout <= 0 {
		o.singleNodesTimeout = DefaultSingleLockTimeout
	}
}
func NewRedLock(key string, confs []*ClientOptions, opts ...RedLockOption) (*RedLock, error) {
	// 3 个节点以上，红锁才有意义
	if len(confs) < 3 {
		return nil, errors.New("can not use redLock less than 3 nodes")
	}

	r := RedLock{}
	for _, opt := range opts {
		opt(&r.RedLockOptions)
	}

	repairRedLock(&r.RedLockOptions)
	if r.expireDuration > 0 && time.Duration(len(confs))*r.singleNodesTimeout*10 > r.expireDuration {
		// 要求所有节点累计的超时阈值要小于分布式锁过期时间的十分之一
		return nil, errors.New("expire thresholds of single node is too long")
	}

	r.locks = make([]*RedisLock, 0, len(confs))
	for _, conf := range confs {
		client := NewClient(conf)
		r.locks = append(r.locks, NewRedisLock(key, client, WithExpireSeconds(int64(r.expireDuration.Seconds()))))
	}

	return &r, nil
}

// 加锁
func (r *RedLock) Lock(ctx context.Context) error {
	var successCnt int
	for _, lock := range r.locks {
		startTime := time.Now()
		// 保证请求耗时在指定阈值以内
		_ctx, cancel := context.WithTimeout(ctx, r.singleNodesTimeout)
		defer cancel()
		err := lock.Lock(_ctx)
		cost := time.Since(startTime)
		if err == nil && cost <= r.singleNodesTimeout {
			successCnt++
		}
	}

	if successCnt < len(r.locks)>>1+1 {
		// 倘若加锁失败，则进行解锁操作
		_ = r.Unlock(ctx)
		return errors.New("lock failed")
	}

	return nil
}

// 解锁时，对所有节点广播解锁
func (r *RedLock) Unlock(ctx context.Context) error {
	var err error
	for _, lock := range r.locks {
		if _err := lock.Unlock(ctx); _err != nil {
			if err == nil {
				err = _err
			}
		}
	}
	return err
}
