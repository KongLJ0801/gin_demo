/*

   @author #Kkk
   @File Code013
   @version 0.1
   @date 2022/9/5 15:40
   @Description: 013: 下面代码为什么会报错?
   @Analysis:
	不允许直接为map中的结构体成员赋值,只能为结构体整体赋值
	有一点很重要：map中的元素不是变量，因此不能寻址！！
	所以如果要修改map中的struct的成员变量的值，需要将struct修改成指针形式/定义具体的整体结构
*/

package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
}

func main() {
	// m := map[string]Student{"people": {"AAA"}}
	// m["people"].Name = "BBB"

	//code1()
	code2()
}

func code1() {
	// code1  语法限制问题
	m := map[string]Student{"people": {"AAA"}}
	// m["people"] 获取key地址,如不存在则返回这段内存的起始地址
	m["people"] = Student{"BBB"}
	x := reflect.TypeOf(m["people"])
	fmt.Println(x.Kind())
	fmt.Printf("%p,%v,%T \n", m, m["people1"], m["people1"])
}

func code2() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			return
		}
	}()
	// code2 类型限制问题
	m := map[string]*Student{"people": {"AAA"}}
	// m["people1"] 获取*Student的地址
	m["people"].Name = "BBB"

	// 不存在 会收获一个空指针异常
	m["people1"].Name = "BBB"
	// runtime error: invalid memory address or nil pointer dereference

}
