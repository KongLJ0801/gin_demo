/*

   @author #Kkk
   @File Code005
   @version 0.1
   @date 2022/9/2 10:59
   @Description: 005: init() 函数是什么时候执行的?
   @Analysis: 初始化顺序 从最底层开始执行init函数,也就是先初始化不依赖任何包的那些包

*/

package main

import (
	"fmt"
	"gin_demo/InterviewCode/Code005/initFunc"
)

func init() {
	fmt.Println("main-package")
}

func main() {

	code001()

	code002()

}

func code001() {
	// 不能为 string 类型的变量赋值为 nil 也不能进行比较
	// nil 是那些类型的零值 (pointer,channel,map,slice,fuc,interface)

	//var s string = nil
	//if s == nil {
	//	s = "default"
	//}
	//fmt.Println(s)
}

func code002() {
	initFunc.TestInitFunc()
	// 1.package-initFuc-init
	// 2.main-package
	// 3.package-initFuc-TestInitFunc
}
