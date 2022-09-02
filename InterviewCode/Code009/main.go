/*

   @author #Kkk
   @File Code009
   @version 0.1
   @date 2022/9/2 14:37
   @Description: 009: 读写已关闭的channel会怎么样?
   @Analysis:
	1- 有缓存,有数据  true
	2- 无缓存,无数据  false

*/

package main

import "fmt"

func main() {
	// 写已经关闭的channel
	chCloseWrite() // panic

	// 读已经关闭的channel
	chCloseRead()
}
func chCloseWrite() {
	ch := make(chan int, 2)
	close(ch)
	ch <- 1
}
func chCloseRead() {
	ch := make(chan int, 2)
	ch <- 1
	close(ch)
	num, ok := <-ch
	fmt.Println(num, ok) // 1 , true
	num, ok = <-ch
	fmt.Println(num, ok) // 0 , false
}
