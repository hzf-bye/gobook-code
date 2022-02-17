/*
代码所在章节：5.1.2节
*/

package main

import (
	"fmt"
	"runtime"
	//	"time"
)

func main() {
	c := make(chan struct{})
	ci := make(chan int, 100)
	go func(i chan struct{}, j chan int) {
		for i := 0; i < 10; i++ {
			ci <- i
		}
		close(ci)
		//写通道
		c <- struct{}{}
	}(c, ci)

	//NumGoroutine可以返回当前程序的goroutine数目
	println("NumGoroutine=", runtime.NumGoroutine())

	//读通道c，通过通道进行同步等待
	command, ok := <-c

	println("===", ok, "===")
	fmt.Println(command)
	fmt.Printf("%T\n",command)

	//读通道c，通过通道进行同步等待
	command1, ok1 := <- ci

	println("===", ok1, "===")
	fmt.Printf("%T\n",command1)
	fmt.Printf("%v\n",command1)

	command2, ok2 := <- ci
	println("===", ok2, "===")
	fmt.Printf("%T\n",command2)
	fmt.Printf("%v\n",command2)

	//此时ci通道已经关闭，匿名函数启动的goroutine已经退出
	println("NumGoroutine=", runtime.NumGoroutine())

	//但是通道ci还是可以继续读取
	for v := range ci {
		println(v)
	}
}
