package _break

import "fmt"

/**
 *
 * @date 2022/7/18 15:29
 * @version 0.1
 * @author KongLingJie
 *
 */

func break_for() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < len(arr); i++ {
		//if i > 3 {
		//	break  跳出循环,后面循环不在执行.
		//}
		if i%3 == 0 {
			continue // 跳出当前循环,后续代码不在执行 执行下次循环
		}
		fmt.Printf("for is i =%d,", i)
	}
}

func BreakMain() {
	break_for()
}
