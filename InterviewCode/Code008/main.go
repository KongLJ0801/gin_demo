/*

   @author #Kkk
   @File Code008
   @version 0.1
   @date 2022/9/2 14:17
   @Description: 008: 如何交换两个变量的值?
   @Analysis:
	Go语言中传参都是值拷贝,
	需要使用指针类型才可以交换.

*/

package main

import "fmt"

func swap(a, b int) {
	// 局部变量
	a, b = b, a
	fmt.Println(a, b) // 2,1

}

func main() {
	a, b := 1, 2
	swap(a, b)
	fmt.Println(a, b) // 1,2

	// 解答
	c, d := 3, 4
	// 传递地址
	code(&c, &d)
	fmt.Println(c, d) // 4,3

}

func code(c, d *int) {
	*c, *d = *d, *c
}
