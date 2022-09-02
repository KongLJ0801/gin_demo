/*

   @author #Kkk
   @File code_interface
   @Description:
   @version 0.1
   @date 2022/8/9 13:57

*/

package main

import "gin_demo/imoocCode/code_interface/model"

/**
怎么定义接口?
公共方法组合起来以封装特定功能的集合,代码中通过关键字 interface来定义接口
实现接口?
隐式的实现了  不需要 implements,来实现; 只需要保持方法名,出参,入参,一致
面向对象的特性?
多态,接口的多态是多种不同的实现方式
	1.接口定义变量,
	2.创建的结构体赋值给前面的变量
	3.通过变量调用结构体中实现接口的方法
*/

func main() {
	// 1. 多态 (多种实现)
	//var b model.Info
	//b = new(model.UserInfo)
	//b.GetInfo()
	// 	接口定义变量
	user := new(model.UserInfo)
	dog := new(model.DogInfo)
	action(user)
	action(dog)

}

func action(b model.Info) string {
	b.GetInfo()
	return ""
}
