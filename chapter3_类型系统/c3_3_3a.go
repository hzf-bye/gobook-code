/*
代码所在章节：3.3.2节
*/


package main

import "fmt"

type S struct {
	a int
}

func (s S) Get() int  {

	return s.a
}

func (s *S ) Set(i int)  {

	s.a = i
}

func (s *S) Print()  {

	fmt.Printf("%p, %v, %d\n", s, s, s.a)
}


func main()  {

	var s1 = S{a: 1}
	var s2 = &S{a: 2}

	//普通方法调用
	s1.Get()
	s1.Set(5)
	s2.Get()
	s2.Set(3)

	//方法表达式调用
	(S).Get(s1)

	//方法表达式调用
	f1 := S.Get
	f1(s1)

	//方法表达式调用
	f2 := (S).Get
	f2(s1)

	//如下方法表达式调用都是等价的
	(*S).Set(&s1, 1)
	f3 := (*S).Set
	f3(&s1, 1)


}
