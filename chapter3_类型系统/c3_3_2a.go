/*
代码所在章节：3.3.2节
*/


package main

import "fmt"

type T struct {
	a int
}

func (t T) Get() int  {

	return t.a
}

func (t *T ) Set(i int)  {

	t.a = i
}

func (t *T) Print1()  {

	fmt.Printf("%p, %v, %d\n", t, t, t.a)
}

func (t T) Print2()  {

	fmt.Printf("%p, %v, %d\n", &t, &t, t.a)
}

func main()  {

	var t = &T{}
	f := t.Set

	//方法值调用
	f(2)
	t.Print1()
	t.Print2()

	f(3)
	t.Print1()
	t.Print2()

}
