package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
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

// 生产者函数
func producer(id int, buffer chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // 函数退出时通知 WaitGroup

	for i := 0; i < 5; i++ { // 每个生产者生产5个数据
		data := rand.Intn(100)
		fmt.Printf("Producer %d producing: %d\n", id, data)
		buffer <- data                                               // 将数据放入缓冲区（如果缓冲区满则阻塞）
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500))) // 模拟生产耗时
	}
	fmt.Printf("Producer %d exit\n", id)
}

// 消费者函数
func consumer(id int, buffer <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() // 函数退出时通知 WaitGroup

	for data := range buffer { // 循环从 channel 读取数据（channel关闭时自动退出）
		fmt.Printf("Consumer %d consumed: %d\n", id, data)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000))) // 模拟消费耗时
	}
	fmt.Printf("Consumer %d exit\n", id)

}
