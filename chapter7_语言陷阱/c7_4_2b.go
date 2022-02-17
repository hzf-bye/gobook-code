package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main()  {

	a := []int{0, 1, 2, 3, 4, 5, 6}
	b := a[0:4]
	as := (*reflect.SliceHeader) (unsafe.Pointer(&a))
	bs := (*reflect.SliceHeader) (unsafe.Pointer(&b))

	//a、 b共享底层数组
	fmt.Printf ("a=%v, len=%d, cap=%d, type=%d\n", a, len(a), cap(a), as.Data)
	fmt.Printf ("b=%v, len=%d, cap=%d, type=%d\n", b, len(b), cap(b), bs.Data)

	b = append(b, 10, 11, 12)

	//a, b 继续共享底层数组，修改 b 会影响共亭的底层数组
	fmt.Printf ("a=%v, len=%d, cap=%d\n", a, len(a), cap(a))
	fmt.Printf ("b=%v, len=%d, cap=%d\n", b, len(b), cap(b))

	//len(b)=7 ，底层数组容量是7， 此时需要重新分自己数组，并将原来数组复制 到 新数组
	b = append(b , 13 , 14)

	as = (*reflect.SliceHeader) (unsafe.Pointer(&a))
	bs = (*reflect.SliceHeader) (unsafe.Pointer(&b))
	//可以看到 a 和 b 指向底层数组的指针已经不同了
	fmt.Printf ("a=%v, len=%d, cap=%d, type=%d\n", a, len(a), cap(a), as.Data)
	fmt.Printf ("b=%v, len=%d, cap=%d, type=%d\n", b, len(b), cap(b), bs.Data)

	/*
	问题总结:多个切片共享一个底层数组，其中一个切片的 append操作可能引发如下两种 情况。
	(1) append 追加的元素没有超过底层数组的容量，此种 append 操作会直接操作共享的底 层数组，
	如果其他切片有引用数组被覆盖的元素，则会导致其他切片的值也隐式地发生变化 。
	(2) append 追加的元素加上原来的元素如果超出底层数组的容量，则此种 append 操作会 重新申请新数组，并将原来数组值复制到新数组。
	由于有这种二义性，所以在使用切片的过程中应该尽量避免多个切面共享底层数组，可以使用copy进行显式的复制。
	 */
}
