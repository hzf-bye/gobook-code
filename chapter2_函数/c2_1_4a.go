/*
代码所在章节：2.1.4节
*/

package main

import "fmt"

func add1(a, b int) (sum int) {
	sum = a + b
	return
}

/*
func add2(a, b int) int {
	return a + b
}


func add(a, b int) (sum int) {
	anonymous := func(x, y int) int {
		return x + y
	}
	return anonymous(a, b)
}
*/

func sum(arr ...int) (sum int) {
	for _,v := range arr {
		sum += v
	}
	return
}

func suma(arr ...int) (sum int) {
	for _,v := range arr {
		sum += v
	}
	return
}

func sumb(arr []int) (sum int) {
	for _,v := range arr {
		sum += v
	}
	return
}

type sumFunc func(...int) int

func do1(f sumFunc, a, b int) int {
	return f(a, b)
}

func main() {
	//切片可以作为实参传递给不定参数，切片名后面需要加"..."
	slice := []int{1, 2, 3, 4}
	fmt.Println(sum(slice...))

	//数组不可以作为实参传递给不定参数的函数
	//array := [...]int{1, 2, 3, 4}
	//sum(array)


	fmt.Println(do1(suma, 10, 2))

	//形参为不定参数的函数与形参为切片的函数类型不相同
	fmt.Printf("%T\n", suma) // func(...int) int
	fmt.Printf("%T\n", sumb) // func([]int) int

}
