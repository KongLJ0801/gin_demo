package main

/**
 *
 * @date 2022/7/13 16:47
 * @version 0.1
 * @author KongLingJie
 *
 */

import (
	"fmt"
	"gin_demo/basicCode/day1/count"
)

/*
	方法名/全局变量名(函数外部定义)
		大写 : 其他package可访问
		小写 : 仅本package可访问

	全局 变量使用 var / const 来定义
		同包/不同文件 也可访问
	局部 变量使用 := 来定义
*/

func nums() {
	a := count.Add(1, 2) // 导包
	b := count.Cut(1, 2)
	c := 1 * 2 // 加减乘除  取余
	d := 1 / 2
	e := 1 % 2

	fmt.Println("Add=", a)
	fmt.Println("Cut=", b)
	fmt.Println("c=", c)
	fmt.Println("d=", d)
	fmt.Println("e=", e)

}

func rel() { //  与 或 非  运算符

	a := 1 == 2
	b := 1 != 2
	c := 1 < 2
	d := 1 > 2
	e := 1 <= 2

	f := a && b || !c

	fmt.Println("a=", a)
	fmt.Println("b=", b)
	fmt.Println("c=", c)
	fmt.Println("d=", d)
	fmt.Println("e=", e)
	fmt.Println("f=", f)
}

func cst() {
	const PI float64 = 3.14 // 显示定义 需要初始值
	const (
		A = 100 // 默认值,后面值不修改默认同上
		B
	)
	const (
		C = iota // 占位 递增 0 1 2 3
		D
		_ = 10 // 占位无法使用  但会修改后面的值
		E
		F = iota
		G
	)
	fmt.Println(PI, A, B, C, D, E, F, G)
}

func main() {
	cst()
}
