/*
代码所在章节：5.1.1节
*/

package main

import "runtime"

/**
func GOMAXPROCS(n int) int用来设置或查询可以井发执行的 goroutine数目 ，
n大于 1 表示设置 GOMAXPROCS 值 ， 否则表示查询当前的 GOMAXPROCS 值 。如下:
 */
func main() {
	//获取当前GOMAXPROCS的值
	println("GOMAXPROCS=", runtime.GOMAXPROCS(0))

	//设置GOMAXPROCS的值为2
	runtime.GOMAXPROCS(2)

	//获取当前GOMAXPROCS的值
	println("GOMAXPROCS=", runtime.GOMAXPROCS(0))
}
