package main

import "fmt"

var num int = 0

func main() {
	//定义一个channel
	//c := make(chan int)
	//var ints chan int = make(chan int)

	go func() {
		defer fmt.Println("goroutine结束")

		fmt.Println("goroutine 正在运行...")
		num = 666
		//c <- 666 //将666 发送给c
	}()

	//num := <-c //从c中接受数据，并赋值给num

	fmt.Println("num = ", num)
	fmt.Println("main goroutine 结束...")
}
