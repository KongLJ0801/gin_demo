/*

   @author #Kkk
   @File Code003
   @version 0.1
   @date 2022/9/2 10:24
   @Description: 003: new 和 make 的区别?
   @Analysis:
	1- new 分配内存,不会初始化内存,内存置零. 返回指针,新内存.
	2- make 用于(slice,map,channel)创建,它会初始化T类型,不会简单的置为零值. 引用类型, 使用前必须初始化.
    3- 二者都是用来做内存分配的.

*/

package main

import "fmt"

func main() {
	initNew()
	initMake()
}

func initNew() {
	s := new([]int)
	fmt.Printf("%T \n", s)        // *[]int
	fmt.Printf("%v \n", s)        //  &[]
	fmt.Println(len(*s), cap(*s)) // 0 ,0
	// (*s)[0]=1    s 只是个指针,new 并未分配slice的底层数组,所以为slice的元素赋值发生 panic
}
func initMake() {

	s := make([]int, 5)
	fmt.Printf("%T \n", s)      // []int
	fmt.Printf("%v \n", s)      // [0 0 0 0 0]
	fmt.Println(len(s), cap(s)) // 5,5

	// make -> address -> data=nil,len=0,cap=0
	// s -> address -> data=[0 0 0 0 0],len=5,cap=5

}
