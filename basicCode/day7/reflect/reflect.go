package reflect

import (
	"fmt"
	"reflect"
)

/**
 *
 * @date 2022/7/20 11:18
 * @version 0.1
 * @author KongLingJie
 *
 */

type User struct {
	Name string `json:"myName"`
	age  int64
	uint8
}

type Student struct {
	sex int
	User
}

func (*Student) Examine(a int, b float64) (byte, string) {
	return 3, "ABC"
}

func get_type() {
	typeI := reflect.TypeOf(1)
	typeS := reflect.TypeOf("abc")
	fmt.Println(typeI)
	fmt.Println(typeI.Name())
	fmt.Println(typeI.Kind())
	fmt.Println(typeI.String())
	fmt.Println(typeS)
}

func get_struct() {
	//u := &User{"tom", 12, 1}
	//typeU := reflect.TypeOf(u)
	//fmt.Println(typeU) // 指针解析
	//// Elem 必须是指针类型 才可以使用 elem (在编译期间不会报错,只有在运行时才会报错)
	//
	//fmt.Println(typeU.Elem().Kind()) //结构体 类型
	//fmt.Println(typeU.Elem().Name()) //结构体 名称
	//fmt.Println(typeU.Elem())        // 包名
	//fmt.Println(typeU.Kind())        // 指针类型
	//
	//fmt.Println(typeU.Kind() == reflect.Ptr)           // 判断类型
	//fmt.Println(typeU.Elem().Kind() == reflect.Struct) // 判断类型

	u := User{"tom", 12, 1}
	typeU := reflect.TypeOf(u)

	fmt.Println(typeU.Size())    // 结构体内存空间
	fmt.Println(typeU.PkgPath()) // 包路径
}

func get_field() {

	u := User{"tom", 12, 1}
	typeU := reflect.TypeOf(u)

	fieldNum := typeU.NumField()

	for i := 0; i < fieldNum; i++ {
		field := typeU.Field(i)
		fmt.Println(field.Name)            // 结构体 属性名
		fmt.Println(field.Offset)          // 结构体 字节数
		fmt.Println(field.Anonymous)       // 结构体  属性是否是 匿名
		fmt.Println(field.Type)            // 结构体 属性类型
		fmt.Println(field.IsExported())    // 属性是否可导出
		fmt.Println(field.Tag.Get("json")) // 属性 tag 别名
		fmt.Println()
	}

	if v, ok := typeU.FieldByName("sex"); ok {
		fmt.Println(v)
	} else {
		fmt.Printf("%s不存在sex成员变量", typeU.Name())
	}

	s := Student{1, User{Name: "rose", age: 12, uint8: 1}}

	typeS := reflect.TypeOf(s)

	v := typeS.FieldByIndex([]int{2, 2}) // 获取内层结构体嵌套 属性

	fmt.Println(v)

}

func get_method_info() {
	u := reflect.TypeOf(Student{}) // 结构体 方法名小写不会统计    接口类型不管方法名大小写都会统计 方法

	x := u.Method(0)

	fmt.Println(x.Name)
	fmt.Println(x.Type)
	fmt.Println(x.IsExported())
	fmt.Println(x.Type.NumIn())
	fmt.Println(x.Type.NumOut())

}

func add(a, b int) int {
	return a + b
}

func get_func_info() {
	funcType := reflect.TypeOf(add)

	fmt.Println(funcType.Kind() == reflect.Func) // true

	funNum := funcType.NumIn()     // 方法 输入 参数
	funNumOut := funcType.NumOut() // 方法 输出 参数

	fmt.Println(funNumOut)
	for i := 0; i < funNum; i++ {
		funcT := funcType.In(i) // i 个参数类型
		fmt.Println(funcT)
	}

}

func set_type() {

	userType := reflect.TypeOf(User{})
	studentType := reflect.TypeOf(Student{})

	// 是否可赋值
	fmt.Printf("student is = %t \n", userType.AssignableTo(studentType))
	fmt.Printf("userType is = %t \n", studentType.AssignableTo(userType))
	// 是否可转换
	fmt.Printf("student is = %t \n", userType.ConvertibleTo(studentType))
	fmt.Printf("userType is = %t \n", studentType.ConvertibleTo(userType))

	// 同类型  自身判断
	user2 := reflect.TypeOf(User{})
	fmt.Printf("userType is = %t \n", user2.AssignableTo(userType))
	fmt.Printf("user2 is = %t \n", userType.ConvertibleTo(user2))
}

type People interface {
	Examine(a int, b float64) (byte, string)
}

func impl() {
	// 借口类型不能直接获取  需要转换成 指针
	peopleType := reflect.TypeOf((*People)(nil)).Elem()

	userType := reflect.TypeOf(User{})
	studentType := reflect.TypeOf(&Student{})

	fmt.Println(userType.Implements(peopleType))    // userType是否实在people接口
	fmt.Println(studentType.Implements(peopleType)) // studentType是否实在people接口
}

func ReflectFunc() {

	// reflect.typeof 获取类型 参数 方法

	//get_type()

	//get_struct()

	//get_field()

	//get_method_info()

	//get_func_info()

	// 数据转换,判断类型是否可转换
	//set_type()

	// 获取 interface 方法
	impl()
	// reflect.value
}
