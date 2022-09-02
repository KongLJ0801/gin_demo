package _defer

import (
	"fmt"
)

/**
 *
 * @date 2022/7/19 11:00
 * @version 0.1
 * @author KongLingJie
 *
 */
func basic() {
	// defer 延迟输出,后进先出,逆序
	fmt.Println("A")

	defer fmt.Printf("1")

	fmt.Println("B")

	defer fmt.Printf("2")

	fmt.Println("C")
}

func deferExeTime() (i int) {
	i = 9 + 1
	defer func() { //返回值 赋值后 在去调用defer 里面的func
		fmt.Println(i)
	}()
	defer func(i int) {
		fmt.Println(i)
	}(i)                              // 注册defer时 参数初始值
	defer fmt.Printf("i -> %d \n", i) // defer 后面如果是表达式 直接赋值(变量在注册defer时 被拷贝或者计算)
	return 5
}

func deferPanic() {
	defer fmt.Println("11111111111")

	n := 0
	defer func() {
		/**
		2/0 出现异常 会发生panic,
		后续代码不执行,会暂时搁置 defer
		但不会影响外部defer/main协程后续也不会执行
		当把其他携程执行完后再来执行这个 panic
		*/
		fmt.Println(2 / n)

		defer fmt.Printf("22222222222222") // 不会执行
	}()

	defer fmt.Println("33333333333333")
}

func DeferFunc() {
	//basic()

	//panic(0) // os.Exit(2) go的整个进程都被杀掉了

	//deferExeTime()

	deferPanic()

}
