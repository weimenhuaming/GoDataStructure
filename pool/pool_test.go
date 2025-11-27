package pool

import (
	"bytes"
	"io/ioutil"
	"sync"
	"testing"
)

var pool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}
var fileName = "test_sync_pool.log"
var data = make([]byte, 10000)

func BenchmarkWriteFile(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := new(bytes.Buffer)
		buf.Reset() // Reset 缓存区，不然会连接上次调用时保存在缓存区里的内容
		buf.Write(data)
		_ = ioutil.WriteFile(fileName, buf.Bytes(), 0644)
	}
}

func BenchmarkWriteFileWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := pool.Get().(*bytes.Buffer) // 如果是第一个调用，则创建一个缓冲区

		buf.Reset() // Reset 缓存区，不然会连接上次调用时保存在缓存区里的内容
		buf.Write(data)
		_ = ioutil.WriteFile(fileName, buf.Bytes(), 0644)

		pool.Put(buf) // 将缓冲区放回 sync.Pool中
	}
}
