/*
代码所在章节：6.1.1节
*/

package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string "学生姓名"
	Age  int    `a:"1111"b:"3333"` //这个不是单引号，而是~键上的符号
}

func main() {
	s := Student{}
	rt := reflect.TypeOf(s)
	////通过名字查找获取 struct 字段，struct 类型专用的方法
	fieldName, ok := rt.FieldByName("Name")

	//取tag数据
	if ok {
		fmt.Println(fieldName.Tag)
	}
	fieldAge, ok2 := rt.FieldByName("Age")

	/*可以像JSON 一样，取TAG 里的数据，注意，设置时，两个之间无逗
	号,键名无引号*/
	if ok2 {
		fmt.Println(fieldAge.Tag)
		fmt.Println(fieldAge.Tag.Get("a"))
		fmt.Println(fieldAge.Tag.Get("b"))
	}

	//返回包含包名的类型名字，对于未命名类型返回的是空
	fmt.Println("type_Name:", rt.Name())
	//返回字 段数目，struct 类型专用的方法
	fmt.Println("type_NumField:", rt.NumField())
	//返回类型的包路径，如采类型是预声明类型或未命名类型，则返回空字符串
	fmt.Println("type_PkgPath:", rt.PkgPath())
	fmt.Println("type_String:", rt.String())

	//Kind() 返回该类型的底层基础类型
	fmt.Println("type.Kind.String:", rt.Kind().String())
	fmt.Println("type.String()=", rt.String())

	//获取结构类型那个的字段名称
	for i := 0; i < rt.NumField(); i++ {
		fmt.Printf("type.Field[%d].Name:=%v \n", i, rt.Field(i).Name)
	}

	fmt.Println("=================")

	sc := make([]int, 10)
	sc = append(sc, 1, 2, 3)
	sct := reflect.TypeOf(sc)

	//获取slice元素的Type
	//返回类型的元素类型，该方法只适合 Array、 Chan, Map, Ptr, Slice 类型
	scet := sct.Elem()

	fmt.Println("slice element type.Kind()=", scet.Kind())
	fmt.Printf("slice element type.Kind()=%d\n", scet.Kind())
	fmt.Println("slice element type.String()=", scet.String())

	fmt.Println("slice element type.Name()=", scet.Name())
	//返回 一个类型的方法的个数
	fmt.Println("slice type.NumMethod()=", scet.NumMethod())
	fmt.Println("slice type.PkgPath()=", scet.PkgPath())
	fmt.Println("slice type.PkgPath()=", sct.PkgPath())

	fmt.Println("=================")

	fmt.Println("slice element type.Kind()=", sct.Kind())
	fmt.Printf("slice element type.Kind()=%d\n", sct.Kind())
	fmt.Println("slice element type.String()=", sct.String())

	fmt.Println("slice element type.Name()=", sct.Name())
	//返回 一个类型的方法的个数
	fmt.Println("slice type.NumMethod()=", sct.NumMethod())
	fmt.Println("slice type.PkgPath()=", sct.PkgPath())
	fmt.Println("slice type.PkgPath()=", sct.PkgPath())

}
