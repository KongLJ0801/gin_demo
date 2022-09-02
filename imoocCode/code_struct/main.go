/*

   @author #Kkk
   @File code_struct
   @Description:
   @version 0.1
   @date 2022/8/9 9:27

*/

package main

import (
	"encoding/json"
	"fmt"
	"gin_demo/imoocCode/code_struct/model"
)

func main() {
	initStruct()
	setStruct()
}

// 结构体方法
func setStruct() {
	var dog = model.Dog{}
	dog.SetDogName("iii")
	name := dog.GetDogName()
	fmt.Println(dog, name)
}

// 初始化 结构体
func initStruct() {
	//var dog = model.Dog{
	//	Name: "RED",
	//	Age:  1,
	//}

	//dog := model.Dog{
	//	Name: "RED",
	//	Age:  1,
	//}

	dog := new(model.Dog)
	dog.Name = "red"
	dog.Age = 1
	fmt.Println(dog)

	// 工厂函数
	x := model.NewDog("red", 1, "ooo")
	x.Name = "xxxx"
	fmt.Println(*x, x.Color.TypeName, any(json.Marshal(x)))
	fmt.Printf("%#v", x)
}
