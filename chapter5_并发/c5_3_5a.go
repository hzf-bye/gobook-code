/*
代码所在章节：5.3.5节
*/

package main

import (
	"context"
	"fmt"
	"time"
)

//define a new type include a Context Field
type otherContext struct {
	context.Context
}

func main() {

	//Construct a *cancelCtx type object
	//使用 context.Background ()构建一个 WithCancel 类型的上下文
	ctxa , cancel := context.WithCancel(context.Background())
	/*
	ctxa 内部状态一一>ctxa=&cancelCtx {
				Context: new(emptyCtx)
				}

	*/
	//work 模拟运行并检测前端的退出通知
	go work(ctxa, "work1")

	//Construct a *timerCtx type object wrapped by *cancelCtx
	//使用 WithDeadline 包装前面的上下文对象 ctxa
	tm := time.Now().Add(3 * time.Second)
	ctxb, _ := context.WithDeadline(ctxa, tm)
	/*
	ctxb 内部状态---〉 ctxb=&timerCtx{
				cancelCtx: ctxa
				dataline :tm
		}
		同时触发 ctxa，在 children 中维护 ctxb 作为子节点
	*/

	go work(ctxb, "work2")

	//使用 WithValue 包装前面的上下文对象 ctxb
	oc := otherContext{ctxb}
	//Construct a *cancelCtx type object wrapped by oc
	ctxc := context.WithValue(oc, "key", "god andes,pass from main ")

	/*
				ctxc -一 > ctxc=&cancelCtx {
								Context: oc
							}
		同时通过 oc.Context 找到 ctxb，通过 ctxb.cancelCtx 找到 ctxa，在 ctxa 的 children 字段中维护 ctxc 作为其子节点
	*/
	go workWithValue(ctxc, "work3")

	//故意“sleep” 10 秒， 让 work2、 work3 超时退出
	time.Sleep(10 * time.Second)
	//显式调用 work1 的 cancel 方法通知其退出
	cancel()
	//等待 work1 打印退出信息
	time.Sleep(5 * time.Second)
	fmt.Println("main stop")
}

//do something
func work(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s get msg to cancel\n", name)
			return
		default:
			fmt.Printf("%s is running \n", name)
			time.Sleep(1 * time.Second)
		}
	}
}

//do something and pass values by context
//等待前端的退出通知，并试图获取 Context 传递的数据
func workWithValue(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s get msg to cancel\n", name)
			return
		default:
			value := ctx.Value("key").(string)
			fmt.Printf("%s is running value=%s \n", name, value)
			time.Sleep(1 * time.Second)
		}
	}
}
