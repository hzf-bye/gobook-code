/*
代码所在章节：2.1.3节
*/

package main

import (
	"fmt"
	"sync"
)

func chvalue(a int) int {
	a = a + 1
	return a
}

func chpointer(a *int) {
	*a = *a + 1
	return
}


func main() {
	a := 10
	chvalue(a)
	fmt.Println(a)
	chpointer(&a)
	fmt.Println(a)


	wg := sync .WaitGroup{}

	si:= []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := range si {
		wg.Add(1)
		go func() {
			println(i)
			wg.Done()
		}()
	}

	wg.Wait()

}
