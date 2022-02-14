/*
代码所在章节：4.1.3节
*/

package main

import "fmt"

type Printer interface {
	Print(i int)
}

type Map map[string]string

func (m Map) Print() {
	for _, key := range m {
		fmt.Println(key)
	}
}

type S struct{}

func (s S) Print(i int) {
	println("print")
}

func main() {
	var i Printer
	//	i.Print() // panic: runtime error: invalid memory address or nil pointer dereference

	//必须初始化
	i = S{}
	i.Print(1)

	mp := make(Map, 10)
	mp["hi"] = "tata"
	mp.Print()

}
