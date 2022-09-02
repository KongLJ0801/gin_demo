package copy

import (
	"fmt"
)

/**
 *
 * @date 2022/7/15 17:40
 * @version 0.1
 * @author KongLingJie
 *
 */

type User struct {
	Name string
}
type Vedio struct {
	Length int
	Author User
}

type Vedio2 struct {
	Length int
	Author *User
}

func initCopy() {
	u := User{
		"tiger",
	}

	v := Vedio{
		1,
		u, // 值拷贝.(开辟了新的空间,修改操作不影响原先的内存)
	}
	v2 := Vedio2{
		1,
		&u, // 引用地址拷贝.(指针类型,还是原来的内存空间,修改操作直接作用在原内存空间上)
	}
	v.Author.Name = "jack"
	v2.Author.Name = "rose"

	fmt.Printf("u is user = %s,v is user.name = %s, v2 is name = %s\n", u, v.Author.Name, v2.Author.Name)
}
func upd_users(users []User) {
	users[0].Name = "rose"
	fmt.Printf("slice is user edit name : %s\n", users[0].Name)
}

func CopyMain() {
	//initCopy()

	u := make([]User, 10)
	u = append(u, User{Name: "jack"})
	fmt.Println(u)

	// 原本 slice 切片创建的本身就是 指针地址 传递值时也是 指针地址拷贝所以在函数体内修改会影响原有值
	//user := []User{{Name: "tom"}}
	//upd_users(user)
	//fmt.Printf("slice is user edit name : %s\n", user[0].Name)

	user := User{Name: "tom"}
	user2 := []User{user} // 值拷贝后 再存在 slice 里面 的指针地址所以不会影响 user的Name
	upd_users(user2)
	fmt.Printf("slice is user edit name : %s\n", user.Name)
	fmt.Printf("slice is user edit name : %s\n", user2[0].Name)

}
