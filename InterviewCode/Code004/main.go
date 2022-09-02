/*

   @author #Kkk
   @File Code004
   @version 0.1
   @date 2022/9/2 10:53
   @Description: 004: 如何判断2个字符串切片是相等?
   @Analysis:
	1- 切片不能通过 == 来进行比较;
	2- 使用 reflect 包的 DeepEqual ;
    2.1- 先判断切片类型是否相等,在判断元素是否相等;


*/

package main

import (
	"fmt"
	"reflect"
)

func main() {
	sMake1 := make([]string, 1)
	sMake1 = append(sMake1, "slice")

	sMake2 := make([]string, 1)
	sMake2 = append(sMake2, "sliceName")

	if ok := reflect.DeepEqual(sMake1, sMake2); ok {
		fmt.Println("ok")
	}

	code001()
}
func code001() {
	fmt.Println([...]string{"1"} == [...]string{"1"}) // 数组比较
	// Invalid operation: []string{"1"} == []string{"1"} (the operator == is not defined on []string)
	// fmt.Println([]string{"1"} == []string{"1"}) // 且不能直接比较
	// 可以使用 reflect.DeepEqual比较
	if ok := reflect.DeepEqual([]string{"1"}, []string{"1"}); ok {
		fmt.Println("ok")
	}
}
