package closure

import "fmt"

/**
 *
 * @date 2022/7/19 10:40
 * @version 0.1
 * @author KongLingJie
 *
 */

func sub() func() {
	i := 10
	fmt.Printf("i addr = %p \n", &i)

	return func() {

		fmt.Printf("func addr = %p \n", &i)

		i--

		fmt.Println(i)
	}
}
func add(base int) func(int) int {
	fmt.Printf("base is addr %p \n", &base)
	return func(i int) int {
		base += i
		return base
	}
}

func ClosureFunc() {
	f := sub()
	f()
	f()

	fmt.Println()
	tmp1 := add(10)
	fmt.Println(tmp1(1), tmp1(2)) // 11, 13  同一内存地址 闭包 内存不会销毁
	tmp2 := add(10)
	fmt.Println(tmp2(1), tmp2(2))
}
