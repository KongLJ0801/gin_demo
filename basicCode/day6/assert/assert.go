package assert

import "fmt"

/**
 *
 * @date 2022/7/19 16:52
 * @version 0.1
 * @author KongLingJie
 *
 */

func assert(i interface{}) {
	switch i := i.(type) {
	case int:
		fmt.Printf("%d \n", i)
	case float64:
		fmt.Printf("%.3f \n", i)
	case uint16, string, byte:
		fmt.Printf("%T %v \n", i, i)
	}
}

func AssertFunc() {
	var i interface{}
	var a int
	var b float64
	var c byte

	i = a
	assert(i)
	i = b
	assert(i)
	i = c
	assert(i)
	fmt.Printf("%v %T", nil, nil)
}
