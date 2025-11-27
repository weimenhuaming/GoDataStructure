package producer_and_consumer

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func Test(t *testing.T) {
	// 初始化随机种子
	rand.Seed(time.Now().UnixNano())

	// 创建一个带缓冲的 channel 作为队列（容量为10）
	buffer := make(chan int, 10)

	// 使用 WaitGroup 等待所有 goroutine 完成
	var wg sync.WaitGroup

	// 启动 3 个生产者
	wg.Add(3)
	for i := 1; i <= 3; i++ {
		go producer(i, buffer, &wg)
	}

	// 启动 2 个消费者
	wg.Add(2)
	for i := 1; i <= 2; i++ {
		go consumer(i, buffer, &wg)
	}

	// 等待所有生产者和消费者完成
	wg.Wait()
	close(buffer) // 关闭 channel（实际在本例中生产者会自动退出）
	fmt.Println("All tasks completed!")
}
