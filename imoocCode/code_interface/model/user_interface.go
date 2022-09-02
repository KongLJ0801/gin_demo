/*

   @author #Kkk
   @File model
   @Description:
   @version 0.1
   @date 2022/8/9 13:58

*/

package model

import "fmt"

type UserInfo struct {
	Name string
	Age  int
}
type DogInfo struct {
	Name string
	Age  int
}

type Info interface {
	GetInfo() string
}

func (u *UserInfo) GetInfo() string {
	fmt.Println("userInfo")
	return "userinfo"
}

func (d *DogInfo) GetInfo() string {
	fmt.Println("DogInfo")
	return "DogInfo"
}
