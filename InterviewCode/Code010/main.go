/*

   @author #Kkk
   @File Code010
   @version 0.1
   @date 2022/9/2 14:43
   @Description: 010: struct{}有什么用处?
   @Analysis:
	struct{}和 type T struct{} 为同一内存地址
	内存大小 size:0   (unsafe.Pointer(&zerobase) 8字节).
	结构体变量的值无关紧要.
		无需传递数据的channel
		map的value设置为struct{}
		只包含方法,不包含任何字段的结构体

*/

package main

import "fmt"

type Kkk struct{}

func main() {
	k1 := struct{}{}
	k2 := struct{}{}
	kkk := Kkk{}
	fmt.Printf("&k1=%p, &k2=%p,&kkk=%p ", &k1, &k2, &kkk)
}
