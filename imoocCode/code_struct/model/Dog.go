/*

   @author #Kkk
   @File model
   @Description:
   @version 0.1
   @date 2022/8/9 9:35

*/

package model

type Dog struct {
	Name string
	Age  int
}

type dog struct {
	Name  string `json:"user_name"` // json 重命名
	Age   int
	Color // 继承
}
type Color struct {
	TypeName string
}

func NewDog(name string, age int, color string) *dog {
	return &dog{
		Name: name,
		Age:  age,
		Color: Color{
			TypeName: color,
		},
	}
}

func (d *Dog) GetDogName() string {
	return d.Name
}
func (d *Dog) SetDogName(name string) {
	d.Name = name
}
