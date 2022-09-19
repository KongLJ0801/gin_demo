/*

   @author #Kkk
   @File code_syncMap
   @version 0.1
   @date 2022/9/5 15:33
   @Description:
   @Analysis:

*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	var sm sync.Map
	// 新增
	sm.Store(1, "Kkk")
	// 获取
	if vv, ok := sm.Load(1); ok {
		fmt.Println(vv)
	}
	// 获取或保存(不更新value),不存在则新增
	if vv, ok := sm.LoadOrStore(1, "c"); ok {
		fmt.Println(vv)
	}
	if vv, ok := sm.LoadOrStore(2, "c"); ok {
		fmt.Println(vv)
	}
	fmt.Println(sm)
	// 删除,不存在不做任何处理 (LoadAndDelete底层调用,无接收)
	sm.Delete(3)
	sm.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})
	// 删除,返回 value,bool
	if vv, ok := sm.LoadAndDelete(1); ok {
		fmt.Println(vv, ok)
	}

}
