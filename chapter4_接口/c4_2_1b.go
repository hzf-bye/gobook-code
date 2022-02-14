/*
代码所在章节：4.2.1节
*/

package main

import "fmt"

/*
语义分析 :
(1) TypeName是具体类型名，此时如果接口 i绑定的实例类型就是具体类型 TypeName, 则 ok 为 tiue， 变量。的类型就是 TypeName，变量 。的值就是接口绑定的 实例值的副本(当然 实例可能是指针值，那就是指针值的副本)。
( 2 )TypeName是接口类型名， 此时如果接口 i绑定的实例 的类型满足接口类型 TypeName, 则ok为tiue，变量。的类型就是接口类型TypeName, o底层绑定的具体类型实例是i绑定的实 例的副 本(当然实例可能是 指针值，那就是指针值的副 本)。
(3)如果上述两个都不满足，则ok为false， 变量。是TypeName类型的“零值”，此种条 件分支下程序逻辑不应该再去引用。，因为此时的。没有意义 。
 */
type Inter interface {
	Ping()
	Pang()
}

type Anter interface {
	Inter
	String()
}

type St struct {
	Name string
}

func (St) Ping() {
	println("ping")
}
func (*St) Pang() {
	println("pang")
}

func main() {
	st := &St{"andes"}
	var i interface{} = st

	//判断i绑定的实例是否实现了接口类型Inter
	if o, ok := i.(Inter); ok {
		o.Ping()
		o.Pang()
	}

	if p, ok := i.(Anter); ok {
		//i没有实现接口Anter，所以程序不会执行到这里
		p.String()
	}

	//判断i绑定的实例是否就是具体类型St
	if s, ok := i.(*St); ok {
		fmt.Printf("%s", s.Name)
	}

}
