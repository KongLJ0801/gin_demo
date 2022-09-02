/*

   @author #Kkk
   @File Code001
   @version 0.1
   @date 2022/9/2 10:12
   @Description: 001:  = 和 := 的区别?
   @Analysis:
	1- = 赋值.
	 = 使用必须有var声明 / 重新赋值.
	2- := 声明变量并赋值.
	 := 是声明并赋值,系统会自动推断类型,不需要使用var关键字.

*/

package main

import (
	"errors"
	"fmt"
)

func main() {
	var a int // 初始值 0
	fmt.Println(a)
	a = 1
	b := 10
	fmt.Println(a, b)

	code001()
}

func code001() {
	// 尽管两个语句中都出现了 err，但这种重复仍然是合法的：err 在第一条语句中被声明，但在第二条语句中只是被再次赋值罢了
	a := 1
	err := testErr(a)
	fmt.Println(err) // err msg
	a = 11
	err = testErr(a)
	fmt.Println(err) // nil

}
func testErr(a int) error {
	if a < 10 {
		return errors.New("err msg")
	}
	return nil
}
