package reflect

import (
	"fmt"
	"reflect"
)

/**
 *
 * @date 2022/7/21 13:56
 * @version 0.1
 * @author KongLingJie
 *
 */

type Cat struct {
	Name string
	age  int
}

// 反射 创建结构体 struct
func newStruct() {

	typeC := reflect.TypeOf(Cat{"rose", 1})

	nCat := reflect.New(typeC)

	v := nCat.Elem().FieldByName("Name")

	if v.CanSet() {
		v.SetString("tom")
		fmt.Println(v)
	}

	if val, ok := nCat.Interface().(*Cat); ok {
		fmt.Println(val)
	} else {
		fmt.Println("----")
	}
}

// 反射 创建切片 slice
func newSlice() {
	var sliceCat []Cat
	typeC := reflect.TypeOf(sliceCat)
	valueC := reflect.MakeSlice(typeC, 1, 20)
	x := valueC.Len()
	for i := 0; i < x; i++ {
		valueC.Index(i).Set(reflect.ValueOf(Cat{"rose", 2}))
	}
	if val, ok := valueC.Interface().([]Cat); ok {
		fmt.Println(val, ok)

		valueC2 := reflect.ValueOf(&val)

		valueC2.Elem().SetCap(10) // 原始内存大小的越界问题
		valueC2.Elem().SetLen(5)
		fmt.Println(valueC2.Elem().Len())
		fmt.Println(valueC2.Elem().Cap())
		valueC2.Elem().Index(0).FieldByName("Name").SetString("tom")
		valueC2.Elem().Index(2).Set(reflect.ValueOf(Cat{"jack", 21}))
		valueC2.Elem().Index(3).Set(reflect.ValueOf(Cat{"rose", 22}))

		v := valueC2.Elem().Interface().([]Cat)
		fmt.Println(v, val)
	}

}

// 反射 创建map
func newMap() {
	var m map[int]*Cat
	typeM := reflect.TypeOf(m)
	valueC := reflect.MakeMap(typeM)
	//valueC := reflect.MakeMapWithSize(typeM, 10) // 指定容量

	key := reflect.ValueOf(2)
	c := &Cat{"tom", 1}
	val := reflect.ValueOf(c)

	valueC.SetMapIndex(key, val)

	if v, ok := valueC.Interface().(map[int]*Cat); ok {
		for i, cat := range v {
			fmt.Println(i, cat)
		}

	}

}

func Add(a, b int) int {
	return a + b
}

// 反射 调用func
func call_func() {
	func_value := reflect.ValueOf(Add)
	c := func_value.Call([]reflect.Value{reflect.ValueOf(3), reflect.ValueOf(5)})
	if r, ok := c[0].Interface().(int); ok {
		fmt.Println(r)
	}
}

func NewReflect() {
	//newStruct()

	//newSlice()

	//newMap()

	call_func()
}
