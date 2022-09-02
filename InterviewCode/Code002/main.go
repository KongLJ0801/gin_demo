/*

   @author #Kkk
   @File Code002
   @version 0.1
   @date 2022/9/2 10:22
   @Description: 002: 读写未初始化的channel会怎样?
   @Analysis: 对未初始化的channel进行(读,写,关闭)都会发生阻塞.

*/

package main

import "fmt"

func main() {
	// 写
	chWrite()
	// 读
	chRead()
	// 关闭
	chClose()

	// channel 初始化
	code001()
}

func chWrite() {
	var ch chan int                   // 只声明，并没有初始化
	fmt.Printf("chWrite is %v\n", ch) // ch is <nil>
	ch <- 1                           // fatal error: all goroutines are asleep - deadlock!
}

func chRead() {
	var ch chan int // 只声明，并没有初始化
	r := <-ch
	fmt.Printf("r is %v\n", r) // fatal error: all goroutines are asleep - deadlock!
}

func chClose() {
	var ch chan int
	close(ch) // panic: close of nil channel
}

func code001() {
	ch := make(chan int)        // 初始化
	defer close(ch)             // 关闭 (defer:函数即将返回时执行,先进后出)
	go func() { ch <- 3 + 4 }() // 写
	r := <-ch                   // 读
	fmt.Println(r)
}
