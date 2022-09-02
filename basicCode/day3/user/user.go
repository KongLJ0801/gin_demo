package user

import (
	"fmt"
	"reflect"
	"time"
)

/**
 *
 * @date 2022/7/15 11:40
 * @version 0.1
 * @author KongLingJie
 *	struct 结构体初始化
 */

//  挎包(pkg)访问问题 大小写区分成员变量也存在大小的区分
//  不挎包可访问
type User struct {
	id         int
	Score      float32
	time       time.Time
	name, addr string
}

//func (User) hello(m string) {
//	fmt.Println("hi: " + m + ")
//}

func (u User) hello(m string) {
	u.name = "杰克" + m
	fmt.Println("hi: " + m + ",fuck : " + u.name)
}

func (u *User) hello2(m string) {
	u.name = "杰克" + m
}

func initType() {
	var u User

	fmt.Printf("id:=%d,score:=%f,time:=%v,name:=%s,addr:=%s \n", u.id, u.Score, u.time, u.name, u.addr)

	//初始化值,可省略key,但必须写全所有值
	u = User{1, 15.12, time.Now(), "Rose", "beijing"}
	// 成员函数
	u.hello("tom")

	fmt.Println(u)
}

func hideType() {
	type People struct {
		name string
		age  int
	}
	//var abc struct {
	//	name string
	//	age  int
	//}
	//abc.name = "jack"
	//abc.age = 1

	var abc People
	abc.age = 1

	var abcd People
	abcd.age = 1
}

// 任意类型
type UserMap map[int]User

func (p UserMap) Get(id int) User {
	return p[id]
}
func (p UserMap) Del(id int) {
	delete(p, id)
}

// 匿名字段只能使用一次
type Student struct {
	Name string
	int
	string
	float32
}

// 结构体指针  (避免值拷贝)
func NewStudent(name, city string, age int) *Student {
	return &Student{name, age, city, 12.5}
}

// 结构体嵌套
type paper struct {
	user User // 命名嵌套
	User      // 匿名嵌套
	name string
}

func UserMain() {
	u := User{
		1,
		12.3,
		time.Now(),
		"Jack",
		"beijing",
	}
	fmt.Println(u)
	//u.hello2("tom")  // 传递指针值 会修改外部值,
	u.hello("tom") // 值传递 方法内部接收值时会 拷贝一份值 不影响外部值
	fmt.Println(u)

	x1 := NewStudent("tom", "ZH", 12)
	v2 := *x1

	fmt.Println(x1)
	fmt.Println(v2)
	//initType()

	//s := Student{"rose", 1, "12", 12.13}
	s := Student{Name: "rose", string: "12", float32: 12.13, int: 1}
	fmt.Println(s)

	v := reflect.ValueOf(s)
	count := v.NumField()

	for i := 0; i < count; i++ {
		f := v.Field(i)
		switch f.Kind() {
		case reflect.String:
			fmt.Println(f.String())
		case reflect.Int:
			fmt.Println(f.Int())
		case reflect.Float32:
			fmt.Println(f.Float())
		}
	}
}
