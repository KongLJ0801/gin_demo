package _type

import (
	"errors"
	"fmt"
	"math"
	"strings"
	"unsafe"
)

/**
 *
 * @date 2022/7/14 14:16
 * @version 0.1
 * @author KongLingJie
 *
 */

func TypeMain() {
	//default_value()
	//scope()
	//errs()

	//---------------------
	//var a Mss
	//
	//fmt.Printf("%T", a)
	//
	//var user User
	//user = User{"tom", 1}
	//user.getUser()
	//
	//var mss ms
	//mss = ms{}
	//mss.say()
	// ---------------------

	//stringArr()
	//stringTmp()
	//stringFunc()
	//stringJoin()

	typeOr()

}

func default_value() {
	var a int
	var b byte
	var c float32
	var d float64
	var e bool
	var f string
	var g rune
	var arr []int

	fmt.Printf("default_value of int %d \n", a)
	fmt.Printf("default_value of byte %d \n", b)
	fmt.Printf("default_value of float32 %f \n", c)
	fmt.Printf("default_value of float64 %f \n", d)
	fmt.Printf("default_value of bool %t \n", e)
	fmt.Printf("default_value of string [%s] \n", f)
	fmt.Printf("default_value of rune %d [%c]\n", g, g)
	fmt.Printf("default_value of slice %v , arr is nil %t \n", arr, arr == nil)

}

func scope() {
	//var a byte = 255
	//var b int8 = 127
	//var c uint = 255
	//var d float32 = 2.4578
	//var e bool
	//var f string
	//var g rune
	//var arr []int

	//fmt.Printf("default_value of int %d \n", a)
	//fmt.Printf("default_value of byte %d \n", b)
	//fmt.Printf("default_value of uint %d \n", c)
	//fmt.Printf("default_value of float64 %.2f,%.3f,%g \n", d, d, d)
	//fmt.Printf("default_value of bool %t \n", e)
	//fmt.Printf("default_value of string [%s] \n", f)
	//fmt.Printf("default_value of rune %d [%c]\n", g, g)
	//fmt.Printf("default_value of slice %v , arr is nil %t \n", arr, arr == nil)

	//var m complex64
	//m = complex(2, 6)
	//fmt.Printf("%T %T \n", real(m), imag(m)) // %t bool  %T 类型
	//fmt.Printf("%f,%f \n", real(m), imag(m))
	//fmt.Printf("%T \n", m)

	//// var o bool / var o = true    // 默认值是 FALSE
	//o := true //声明并赋值
	//o = false //  赋值
	//fmt.Printf("o is value : %t,%T\n", o, o)

	var a int
	var pointer unsafe.Pointer = unsafe.Pointer(&a)
	var p uintptr = uintptr(pointer)
	var ptr *int = &a // *int 指向整形的指针类型

	fmt.Printf("pointer %p, p %d %x,ptr %p", pointer, p, p, ptr)

	//var b = 100 //  int
	//var b byte = 100 // byte 等价于 uint8; rune 等价于 int32

	//var r rune = '中'
	//fmt.Printf(" %T %d %c\n", r, r, r)
}

func errs() {
	var e error
	e = errors.New("divide by zero")
	fmt.Printf("%v \n", e)
	fmt.Printf("%+v \n", e)
	fmt.Printf("%#v \n", e)

	type user struct {
		Name string
		Age  int
	}

	jack := user{
		"杰克",
		1,
	}
	fmt.Printf("%v \n", jack)  //value
	fmt.Printf("%+v \n", jack) //object
	fmt.Printf("%#v \n", jack) // pkg.struct.object

}

type User struct {
	Name string
	Age  int
}

func (u User) getUser() {
	fmt.Printf("my name is %s \n", u.Name)
}

//	自定义类型
type ms map[string]int

func (m ms) say() {
	fmt.Printf("ms： %d \n", m["hole"])
}

//  类型别名

type Mss = map[string]int

func stringArr() {
	name := "孔令杰"
	arr := []rune(name)

	for _, b := range arr {
		fmt.Printf("%d \n", b)
	}
	fmt.Printf("name is length %d ,arr is length %d \n", len(name), len(arr))

	for _, s := range name {
		fmt.Printf("%c", s)

	}
}

func stringTmp() {
	s1 := "hello \n world"

	s2 := `
	hello world
`
	fmt.Printf("s1  is value %s \n,s2 is value %s", s1, s2)
}

func stringFunc() {
	s := "hello world , tom"

	arr := strings.Split(s, ",") // 字符串分割

	for _, s2 := range arr {
		fmt.Printf("%s \n", s2)
	}
	// bool
	fmt.Printf("%t \n", strings.Contains(s, "tom"))  // 是否包含字符/字符串
	fmt.Printf("%t \n", strings.HasSuffix(s, "tom")) // 是否以某个字符串结尾
	fmt.Printf("%t \n", strings.HasPrefix(s, "tom")) // 是否以某个字符串开始

	// index
	fmt.Printf("%d \n", strings.Index(s, "o"))     // 字符匹配 从头开始匹配 返回 index
	fmt.Printf("%d \n", strings.LastIndex(s, "o")) // 字符匹配 从尾开始匹配 返回 index
	fmt.Printf("%d \n", strings.IndexAny(s, "oe")) // 字符匹配 从头开始匹配 (同时检索 'o/e',包含就直接返回) index

}

func stringJoin() {
	s1 := "aaa"
	s2 := "bbb"
	s3 := "ccc"

	S1 := s1 + " " + s2 + " " + s3
	S2 := fmt.Sprintf("%s %s %s", s1, s2, s3)

	// 性能极高
	S3 := strings.Join([]string{s1, s2, s3}, "-")

	// 性能极高
	sb := strings.Builder{}

	sb.WriteString(s1)
	sb.WriteString("-")
	sb.WriteString(s2)
	sb.WriteString("-")
	sb.WriteString(s3)

	S4 := sb.String()

	fmt.Printf(" %s \n %s \n %s \n %s \n", S1, S2, S3, S4)
}

func typeOr() { // 强制类型转换

	// 低转高没问题, 高转低存在精度丢失  (常量也不能强制转换)

	var ua uint64 = math.MaxUint64
	fmt.Printf("ua=%d\n", ua)

	i8 := uint8(ua) //无符号
	var ub uint64 = uint64(i8)
	fmt.Printf("i8=%d\n", i8)
	fmt.Printf("ub=%d\n", ub)
}
