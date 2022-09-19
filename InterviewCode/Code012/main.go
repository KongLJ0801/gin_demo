/*

   @author #Kkk
   @File Code012
   @version 0.1
   @date 2022/9/5 11:18
   @Description: 012: 代码有什么问题? (map并发读写)
   @Analysis:
	官方的faq已经提到内建的map不是线程(goroutine)安全的。

*/

package main

import (
	"fmt"
	"strconv"
	"sync"
)

var w sync.WaitGroup

type User struct {
	Marks map[string]int
	sync.Mutex
}

func (u *User) Add(name string, mark int) {
	u.Lock()
	defer u.Unlock()
	u.Marks[name] = mark
}
func (u *User) Get(name string) int {
	if mark, ok := u.Marks[name]; ok {
		return mark
	}
	return -1
}

func main() {
	u := &User{
		Marks: make(map[string]int),
	}
	for i := 0; i < 10000; i++ {
		w.Add(1)
		go func(mark int) {
			key := strconv.Itoa(mark)
			u.Add(key, mark)
			fmt.Printf("k=:%v,v:=%v\n", key, u.Get(key))
			w.Done()
		}(i)
	}
	w.Wait()

	// fatal error: concurrent map read and map write
}
