/*
代码所在章节：2.4.1节
*/

package main

var (
	a1 = 0
)

func fa3() func(i int) int {
	return func(i int) int {
		println(&a1, a)
		a1 = a1 + i
		return a1
	}
}
func main() {
	f := fa3() //f引用的外部的闭包环境包括全局变量a
	g := fa3() //f引用的外部的闭包环境包括全局变量a
	// 此时f,g引用的闭包环境中的a的值是同一个
	println(f(1)) //1
	println(g(1)) //2
	println(g(1)) //3
	println(g(1)) //4
}
