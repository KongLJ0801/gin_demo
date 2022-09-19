/*

   @author #Kkk
   @File Code014
   @version 0.1
   @date 2022/9/5 16:20
   @Description: 014: 代码有什么问题?
   @Analysis:
	代码存在内存泄露问题.
	channel 无缓冲
	只有一个goroutine写入成功
	其他三个阻塞在channel,且无法被回收

*/

package main

import "fmt"

type query func(string) string

func exec(name string, vs ...query) string {

	//ch := make(chan string)
	ch := make(chan string, 5)
	fn := func(i int) {
		ch <- vs[i](name)
	}
	for i, _ := range vs {
		go fn(i)
	}
	//for ele := range ch {
	//	fmt.Println(ele)
	//}
	return <-ch

}

func main() {
	ret := exec("111", func(n string) string {
		return n + "func1"
	}, func(n string) string {
		return n + "func2"
	}, func(n string) string {
		return n + "func3"
	}, func(n string) string {
		return n + "func4"
	}, func(n string) string {
		return n + "func5"
	})

	fmt.Println(ret)
}
