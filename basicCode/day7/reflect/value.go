package reflect

import (
	"fmt"
	"reflect"
)

/**
 *
 * @date 2022/7/21 10:21
 * @version 0.1
 * @author KongLingJie
 *
 */

func get_value() {
	iValue := reflect.ValueOf(3)
	sValue := reflect.ValueOf("abc")

	fmt.Println(iValue)
	fmt.Println(sValue)

	fmt.Println(iValue.Kind() == reflect.Int)
	fmt.Println(sValue.Kind() == reflect.String)

	iType := iValue.Type().Kind()
	sType := sValue.Type().Kind()

	fmt.Println(iType)
	fmt.Println(sType)

	// 当调用 IsNil IsZero 时 需要先判断 IsVaild() == true 才能调用,否则会报错

	var i interface{}
	x := reflect.ValueOf(i)
	fmt.Printf("x value IsValid %t \n", x.IsValid()) // interface 是否指向具体类型  FALSE

	type User struct {
		Name string
		float64
		sex int
	}
	var u1 *User = nil
	u := reflect.ValueOf(u1)
	if u.IsValid() {
		fmt.Printf("u value IsNil %t \n", u.IsNil()) // nil  初始赋值为 nil
	}

	var u2 User
	v2 := reflect.ValueOf(u2)
	if ok := v2.IsZero(); ok {
		fmt.Printf("v2 value IsZero %t \n", v2.IsZero()) // 是否是初始默认值
	}
}

func get_is_type() {
	// 以下类型,	只声明并未初始化,它们的默认值都是 nil
	// pointer,func,interface,channel,map,slice
	var val func(a, b int) int
	fmt.Println(val == nil)
	var arr []int
	fmt.Println(arr == nil)
	var i interface{}
	fmt.Println(i == nil)
	var cha chan int
	fmt.Println(cha == nil)
	var m map[int]int
	fmt.Println(m == nil)
}

type Dog struct {
	Name string
	Age  int
}

func set_value() {
	s := "hello"
	rs := reflect.ValueOf(&s)   //必须传递指针才能修改数据
	rs.Elem().SetString("Jack") // 需要先调Elem,把指针value转成非指针value
	fmt.Println(s)

	u := Dog{"rose", 1}

	rsu := reflect.ValueOf(&u)

	setU := rsu.Elem().FieldByName("Name")
	fmt.Println(setU)
	if setU.CanSet() {
		setU.SetString("tom")
	}

	setI := rsu.Elem().FieldByName("Age")

	if setI.CanInt() {
		setI.Set(reflect.ValueOf(10))
	}
	fmt.Println(u)

}

func ValueFunc() {
	//get_value()

	//get_is_type()

	set_value()
}
