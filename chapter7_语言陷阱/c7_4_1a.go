package main

import (
	"fmt"
)

func f(a [3]int)  {
	a[2] = 10
	fmt.Printf("%p, %v\n", &a, a) //0xc0000180f0, [1 2 10]
}

func main()  {

	a := [3]int{1,2,3}

	//直接赋值是值拷贝
	b := a
	//修改 a元素值并不影响 b
	a[2] = 4

	fmt.Printf("%p, %v\n", &a, a) //0xc000018090, [1 2 4]
	fmt.Printf ("%p, %v\n",&b, b) //0xc0000180a8, [1 2 3]

	//数纽作为函数参数仍然是值拷贝
	f(a)

	c := struct {
		s [3]int
	}{
		s : a,
	}

	//结构体是值拷贝，内部的数组也是值拷贝
	d := c
	//修改 c 中的数纽元素值并不影响 a
	c.s[2] = 30
	//修改 d 中的数纽元素值并不影响 a
	d.s[2] = 20
	fmt.Printf ("%p, %v\n",&a, a) //0xc000018090, [1 2 4]
	fmt.Printf ("%p, %v\n",&c, c) //0xc000018120, [1 2 30]
	fmt.Printf ("%p, %v\n",&d, d) //0xc000018138, [1 2 20]

	s1 := s {
		1,
	}
	map1 := map[string]int{"hzf":1, "mamba":2}
	f2(&s1, map1)
	println(s1.a)
	fmt.Println(map1)
}

type s struct {
	a int
}
func f2(s1 *s, map1 map[string]int)  {

	s1.a = 200
	map1["hzf"] = 200
}