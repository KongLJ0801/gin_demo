package main

import (
	"fmt"
	_goto "gin_demo/basicCode/day4/goto"
)

/**
 *
 * @date 2022/7/18 10:55
 * @version 0.1
 * @author KongLingJie
 *	流程控制
 */
func isIf() {
	a := 1
	b := 2
	if !(a > b) {
		fmt.Println("true")
	}

	if c, d, e := 1, 2, 3; c < d && (c < e || c > 0) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	color := "red"

	if color == "red" {
		fmt.Println("red")
	} else if color == "green" {
		fmt.Println("green")
	} else {
		fmt.Println("false")
	}

	m := make(map[int]string, 10)
	m[0] = "abc"
	if value, exists := m[3]; exists {
		fmt.Println(value)
	} else {
		fmt.Println(exists)
	}
}

func isSwitch() {
	color := "red"
	switch color {
	case "red":
		fmt.Println("red")
		//使用 fallthrough 会强制执行后面的 case 语句，fallthrough 不会判断下一条 case 的表达式结果是否为 true。
		fallthrough
	case "green":
		fmt.Println("green")
	default:
		fmt.Println("no")
	}
}
func isSwitchType(num interface{}) {
	switch val := num.(type) {
	case int:
		fmt.Printf("num is int = %d \n", val)
	case string:
		fmt.Printf("num is string = %s \n", val)
	case bool, float64:
		fmt.Printf("num is bool = %v \n", val)
	default:
		fmt.Println("其他类型")
	}
	//switch num.(type) {
	//case int:
	//	val := num.(int)
	//	fmt.Printf("num is int = %d \n", val)
	//case string:
	//	val := num.(string)
	//
	//	fmt.Printf("num is string = %s \n", val)
	//case bool, float64:
	//	val := num.(string) / val := num.(bool)
	//	fmt.Printf("num is bool = %v \n", val)
	//default:
	//	fmt.Println("其他类型")
	//}

}

func isSelect() {

	//val := "tom"

}

func main() {
	//isIf()
	//isSwitch()
	//isSwitchType(false)
	//isSelect()

	//loop.LoopMain()
	//_break.BreakMain()

	_goto.GotoMain()
}
