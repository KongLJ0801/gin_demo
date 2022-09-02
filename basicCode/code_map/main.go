/*

   @author #Kkk
   @File code_map
   @Description:
   @version 0.1
   @date 2022/8/8 11:11

*/

package main

import "fmt"

type user struct {
	Name string
	age  int
}

func (u *user) getName() string {
	return u.Name
}

func initMap() {

	mp := make(map[string]int)

	mu := make(map[string]user)

	mp["id"] = 1
	mp["id2"] = 2

	mu["hole"] = user{
		Name: "kkk",
		age:  0,
	}

	mu["jack"] = user{
		Name: "jjj",
		age:  18,
	}

	for ele := range mu {
		fmt.Printf("%+v", mu[ele])
		u := mu[ele]
		x := u.getName()
		fmt.Printf("%s", x)
	}

}

func deleteMap() {
	md := make(map[int]int, 1)
	md[0] = 0
	fmt.Printf("%+v, %p %d\n", md, md, len(md))

	md[1] = 1
	md[2] = 2
	delete(md, 1)
	fmt.Printf("%+v, %p %d\n", md, md, len(md))

}

func main() {
	//initMap()

	deleteMap()
}
