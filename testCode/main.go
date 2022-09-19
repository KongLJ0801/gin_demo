/*

   @author #Kkk
   @File testCode
   @Description:
   @version 0.1
   @date 2022/8/24 19:36

*/

package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type Dictionay[K comparable, V any] map[K]V

func main() {
	dict := Dictionay[string, float64]{"string": 1}
	fmt.Printf("dict: %#v \n", dict)
	//funcName()
	//funcName1()
	//funcName2()
	//funcName3()

	x := []int64{23, -48, 19, 10, -10, 90}

	s := FindMinElement(x)
	fmt.Println(s)

}

type intMinType interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

func FindMinElement[T intMinType](arr []T) T {
	minNum := arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] < minNum {
			minNum = arr[i]
		}
	}
	return minNum
}

func funcName3() {
	x := "[qti/kona/kona:10/VC.20220815.022/eng.qianxi.20220815.124113:userdebug/test-keys]"
	println(strings.Trim(x, "[]"))
}

func funcName2() {
	result := NewResult("xxxx", nil)
	result.Version = 1
	xx, _ := json.Marshal(result)
	fmt.Println(string(xx))
}

type Result struct {
	// 状态码
	Code int `json:"code"`
	// 返回描述
	Message string `json:"message"`
	// 消息标识
	CallBackKey string `json:"callBackKey"`
	// 消息体
	Content interface{} `json:"content,omitempty"`
	// 版本号
	Version interface{} `json:"version,omitempty"`
}

func NewResult(callBackKey string, content interface{}) *Result {
	result := &Result{
		Code:        0,
		Message:     "success",
		CallBackKey: callBackKey,
		Content:     content,
	}
	return result
}

func funcName1() {
	for i := 0; i < 9; i++ {
		num := 28 + i
		s := fmt.Sprintf("cat /sys/class/thermal/thermal_zone%d/temp | awk '{print$1}'", num)
		fmt.Println(s)
	}
}

func funcName() {
	var pi float64
	b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
	buf := bytes.NewBuffer(b)
	err := binary.Read(buf, binary.LittleEndian, &pi)
	if err != nil {
		log.Fatalln("binary.Read failed:", err)
	}
	fmt.Println(pi)

	var reader *bufio.Reader

	earl, _ := reader.Peek(4)
	fmt.Println(earl)
}
