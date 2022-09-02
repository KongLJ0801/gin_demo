/*

   @author #Kkk
   @File code_new
   @Description: new 指针类型
   @version 0.1
   @date 2022/8/8 16:52

*/

package main

import (
	"fmt"
	"reflect"
)

func main() {
	newMap()
}

func newMap() {
	nMap := new(map[int]string)
	mMap := make(map[int]string)

	fmt.Println(reflect.TypeOf(nMap)) //*map[int]string
	fmt.Println(reflect.TypeOf(mMap)) // map[int]string

}
