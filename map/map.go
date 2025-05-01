package main

import "sync"

type MyConcurrentMap struct {
	mp map[string]interface{}
	mu sync.RWMutex
}

func NewMyConcurrentMap() *MyConcurrentMap {
	return &MyConcurrentMap{
		mp: make(map[string]interface{}),
	}
}

// Get 获取键值，存在返回值和true，否则返回nil和false
func (m *MyConcurrentMap) Get(key string) (interface{}, bool) {
	m.mu.RLock() // 读锁，允许多个读
	defer m.mu.RUnlock()
	val, ok := m.mp[key]
	return val, ok
}

func (m *MyConcurrentMap) Set(key string, val interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.mp[key] = val
}

func (m *MyConcurrentMap) Del(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.mp, key)
}

func (m *MyConcurrentMap) Range(f func(key string, value interface{}) bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for k, v := range m.mp {
		if !f(k, v) {
			break
		}
	}
}
