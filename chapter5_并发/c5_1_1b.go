/*
代码所在章节：5.1.1节
*/

package main

import (
	"runtime"
	"time"
)

/**
goroutine 有如下特性:
• go 的执行是非阻塞的，不会等待 。
• go 后面的函数的返回值会被忽略 。
• 调度器不能保证 多个 goroutine 的执行次序 。
• 没有父子 goroutine 的概念，所有的 goroutine 是平等地被调度和执行的 。
• Go 程序在执行时会单独为 main 函数创建一个 goroutine，遇到其他 go 关键字时再去创 建其他的 goroutine。
• Go 没有暴露 goroutine id 给用户，所以不能在 一个 goroutine 里面显式地操作另 一个 goroutine， 不过 runtime 包提供了 一些函数访 问和设置 goroutine 的相关信息 。
 */
func sum() int {
	sum := 0
	for i := 0; i < 10000; i++ {
		sum += i
	}
	println(sum)
	time.Sleep(1 * time.Second)
	return sum
}

func main() {
	go sum()
	//NumGoroutine可以返回当前程序的goroutine数目
	println("NumGoroutine=", runtime.NumGoroutine())

	//main goroutine故意sleep5秒,防止其提前退出
	time.Sleep(5 * time.Second)
}
