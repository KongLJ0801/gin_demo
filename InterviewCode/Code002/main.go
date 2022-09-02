/*

   @author #Kkk
   @File Code002
   @version 0.1
   @date 2022/9/2 10:22
   @Description: 002: 读写未初始化的channel会怎样?
   @Analysis: 对未初始化的channel进行(读,写,关闭)都会发生阻塞

*/

package main

import "fmt"

func main() {
	var ch chan int              // 只声明，并没有初始化
	fmt.Printf("ch is %v\n", ch) // ch is <nil>
	ch <- 1                      // fatal error: all goroutines are asleep - deadlock!
}
