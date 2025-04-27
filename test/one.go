package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	counter int64 //计数器
	wg      sync.WaitGroup
)

func foo() int64 {
	atomic.AddInt64(&counter, 1) // 原子地增加计数器的值，每次加1
	return counter
}

func main() {
	wg.Add(100)
	for i := 0; i < 100; i++ { // 启动100个goroutine来增加计数器的值
		go func() {
			defer wg.Done()
			foo()
		}()
	}
	wg.Wait()
	fmt.Println("计数器的值:", atomic.LoadInt64(&counter)) // 原子地读取计数器的值
}
