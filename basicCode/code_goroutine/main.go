/*

   @author #Kkk
   @File code_goroutine
   @Description:
   @version 0.1
   @date 2022/8/5 14:01

*/

package main

import (
	"fmt"
	"gin_demo/basicCode/code_goroutine/context_go"
)

func action() {
	fmt.Println("Test goroutine")
}
func main() {

	//// 协程
	//go action()
	//// 主协程执行完了,销毁 子协程还没执行,所以需要沉睡2S
	//time.Sleep(1 * time.Second)
	//
	//wait.Wait()

	context_go.ContextGo()
}
