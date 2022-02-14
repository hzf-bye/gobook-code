/*
代码所在章节：2.5.1节
*/

package main

import "fmt"

/**
可以有连续多个panic被抛出，连续多个panic的场景只有出现在延迟调用(defer)里面，否则不会出现多个panic被抛出的场景，
但只有最后一个panic能被捕获。
 */
func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	defer func() {
		panic("first defer panic")
	}()

	defer func() {
		panic("second defer panic")
	}()

	panic("main body panic")
}
