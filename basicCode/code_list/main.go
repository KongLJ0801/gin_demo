/*

   @author #Kkk
   @File code_list
   @Description:
   @version 0.1
   @date 2022/8/20 14:44

*/

package main

import (
	"container/list"
	"fmt"
	"strconv"
)

func main() {
	listJson()
}

type People struct {
	Name string
	Age  int
}

func listJson() {

	// Create a new list and put some numbers in it.

	l := list.New()

	l.PushBack(People{Name: "zjw"})
	l.PushBack(People{Name: "kkk"})

	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {

		//fmt.Printf("%d\n", e.Value.(int))
		//e := e.Front()
		//
		p, ok := (e.Value).(People)

		if ok {

			fmt.Println("Name:" + p.Name)

			fmt.Println("Age:" + strconv.Itoa(p.Age))

		} else {

			fmt.Println("e is not an People")

		}

	}
}
