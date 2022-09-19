/*

   @author #Kkk
   @File Code011
   @version 0.1
   @date 2022/9/5 10:15
   @Description: 011: 输出什么? (slice|append)
   @Analysis:
	sMake/sMake1 : 局部变量拷贝,共用一个底层数组
	sMake[:10] 基于同一底层数组创建一个新的切片

	sMake : []int,[], 地址
	sMake1 : []int,[10,100,30], 地址
	sMake[:10] : []int,[2,100,30,0,0,0,0,0,0,0], 地址

*/

package main

import "fmt"

func main() {
	// 切片
	sMake := make([]int, 0, 10)
	// 传参:值拷贝,(slice:地址 指针)
	var appendFunc = func(sMake1 []int) {
		// 局部  sMake1.len=3
		sMake1 = append(sMake1, 10, 20, 30)
		sMake1[1] = 100
		fmt.Printf("%T,%v,%p \n", sMake1, sMake1, sMake1) // []int,[10 20 30],地址
	}
	// 修改 : sMake[0] = 1  (panic, 因为sMake.len=0)
	fmt.Printf("%T,%v,%p \n", sMake, sMake, sMake) //  []int,[],地址
	appendFunc(sMake)
	fmt.Printf("%T,%v,%p \n", sMake, sMake, sMake) //  []int,[],地址
	// sMake[:10] 基于同一底层数组创建一个新的切片 sMake[start:end:cap]    len: [start:end]
	sSon := sMake[:10]
	sSon[0] = 2
	fmt.Printf("%T,%v,%p \n", sMake[:10], sMake[:10], sMake[:10]) // []int,[10 20 30 0 0 0 0 0 0 0],地址
	fmt.Printf("%T,%v,%p \n", sSon, sSon, sSon)                   // []int,[10 20 30 0 0 0 0 0 0 0],地址

}
