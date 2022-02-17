package main

import (
	"sync"
)

func main() {

	wg := sync .WaitGroup{}

	si:= []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := range si {
		wg.Add(1)
		go func() {
			/**
			一直打印9，原因：
			(1) for range 下的迭代变量 1 的值是共用的。
			(2) main 函数所在的 goroutine 和后续启动的 goroutines 存在竞争关系 。
			 */
			println(i)
			wg.Done()
		}()
	}

	wg.Wait()

	for i := range si {
		wg.Add(1)
		go func(a int) {

			//这里有一个实参到形参的位拷贝，问题得以解决。
			println(a)
			wg.Done()
		}(i)
	}

	wg.Wait()

}