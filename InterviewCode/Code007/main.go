/*

   @author #Kkk
   @File Code007
   @version 0.1
   @date 2022/9/2 14:08
   @Description: 007: 代码执行有问题吗?
   @Analysis:
	*Kkk 类实现了接口, 但是Kkk 类没实现接口.

*/

package main

import "fmt"

type Dog interface {
	Cat(string) string
}

type Kkk struct{}

func (k *Kkk) Cat(str string) (hi string) {
	if str == "wang" {
		hi = "The dog barked for a while "
	} else {
		hi = "hi"
	}
	return
}

func main() {

	//var dog Dog = Kkk{} //***Analysis***
	//hi := "wang"
	//fmt.Println(dog.Cat(hi))

	// 解答
	code001()
}

func code001() {
	var dog Dog = &Kkk{}
	hi := "wang"
	fmt.Println(dog.Cat(hi))
}
