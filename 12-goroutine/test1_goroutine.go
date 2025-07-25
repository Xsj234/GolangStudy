package main

import (
	"fmt"
	"sync"
	"time"
)

// 子goroutine
func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new Goroutine : i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

// 主goroutine
func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	//创建一个go程 去执行newTask() 流程
	go newTask()

	go func() {
		defer wg.Done() // 任务结束时，调用Done减少一个
		time.Sleep(3 * time.Second)

		fmt.Println("goroutine end")
	}()

	wg.Wait()
	fmt.Println("main goroutine exit")

	/*
		i := 0
		for {
			i++
			fmt.Printf("main goroutine: i = %d\n", i)
			time.Sleep(1 * time.Second)
		}
	*/
}
