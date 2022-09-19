/*

   @author #Kkk
   @File Analysis
   @version 0.1
   @date 2022/9/5 11:43
   @Description:
   @Analysis:
	Go 1.9之前的解决方案

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
	sync.RWMutex
}

func (u *User) Add(name string, mark int) {
	u.Lock()
	defer u.Unlock()
	u.Marks[name] = mark
}

func (u *User) Get(name string) int {
	u.RLock()
	defer u.RUnlock()
	if mark, ok := u.Marks[name]; ok {
		return mark
	}
	return -1
}

func main() {

	//syncRWMutex()

	syncMapFunc()
}

func syncRWMutex() {
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
}

var syncMap sync.Map

func syncMapFunc() {
	for i := 0; i < 10000; i++ {
		w.Add(1)
		go func(mark int) {
			key := strconv.Itoa(mark)
			syncMap.Store(key, mark)
			if vv, ok := syncMap.Load(key); ok {
				fmt.Printf("k=:%v,v:=%v\n", key, vv)
			}
			w.Done()
		}(i)
	}
	w.Wait()
}
