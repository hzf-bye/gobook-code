package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
make([]int,O)与 var a []int创建的切片是有区别的。前者的切片指针有分配，后者的内部指针为0。 示例如下:
 */
func main()  {

	var a []int

	b := make([]int, 0)

	if a == nil {
		println("a is nil")
	} else {
		println("a is not nil")
	}

	if b == nil {
		println("b is nil")
	} else {
		println("b is not nil")
	}

	//使用反射 中 的 SliceHeader 来获取切片 运行 时 的数据结构
	as := (*reflect.SliceHeader) (unsafe.Pointer(&a))
	bs := (*reflect.SliceHeader) (unsafe.Pointer (&b))

	fmt.Printf ("len=%d, cap=%d, type=%d\n", len (a), cap (a) , as.Data)
	fmt.Printf ("len=%d, cap=%d, type=%d\n", len (b), cap (b) , bs.Data)
}