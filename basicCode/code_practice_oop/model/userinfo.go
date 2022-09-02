/*

   @author #Kkk
   @File model
   @Description:
   @version 0.1
   @date 2022/8/5 10:04

*/

package model

// 大小写区分

type UserInfo struct {
	Name     string
	Age      int
	Height   float64
	Schooled string
	Hobby    []string
	MoreInfo map[string]interface{}
}

type userInfo struct {
	Name     string
	Age      int
	Height   float64
	Schooled string
	Hobby    []string
	MoreInfo map[string]interface{}
}

// 工厂模式
func NewModelUserInfo(Name string,
	Age int,
	Height float64,
	Schooled string,
	Hobby []string,
	MoreInfo map[string]interface{}) *userInfo {
	return &userInfo{
		Name,
		Age,
		Height,
		Schooled,
		Hobby,
		MoreInfo,
	}
}
