/*

   @author #Kkk
   @File code_error
   @Description: 异常处理 -> panic/recover
   @version 0.1
   @date 2022/8/8 17:14

*/

package main

import (
	"errors"
	"fmt"
)

func main() {
	receivePanic()
}

func receivePanic() {

	defer coverPanic()
	//panic("str")
	//panic(1)
	panic(errors.New("error "))
}
func coverPanic() {
	message := recover()
	switch message.(type) {
	case string:
		fmt.Println("string panic message : ", message)
	case error:
		fmt.Println("error panic message : ", message)
	default:
		fmt.Println("panic message : ", message)
	}
}
