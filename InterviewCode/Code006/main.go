/*

   @author #Kkk
   @File Code006
   @version 0.1
   @date 2022/9/2 13:46
   @Description: 006: json包变量不加tag会怎么样?
   @Analysis:
	结构体成员首字母小写或tag(有,无)不能转;
	结构体成员首字母大写(有tag 则转换成tag,无tag 则json的字段名一致才可转换)

*/

package main

import (
	"encoding/json"
	"fmt"
)

type T struct {
	name string
	age  int `json:"age_tag"`
	Sex  int `json:"sex_tag"`
	Addr string
}

func main() {
	t := T{
		name: "name",
		age:  0,
		Sex:  1,
		Addr: "addr",
	}
	fmt.Printf("json.Marshal : %+v \n", t)
	jsonMarshal, err := json.Marshal(t)
	if err != nil {
		fmt.Println("json.Marshal error : ", err)
	}
	fmt.Printf("json.Marshal : %+v \n", string(jsonMarshal))

}
