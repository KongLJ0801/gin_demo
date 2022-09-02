/*

   @author #Kkk
   @File code_goroutine
   @Description:
   @version 0.1
   @date 2022/8/11 10:09

*/

package main

import (
	"gin_demo/imoocCode/code_goroutine/model"
	"time"
)

/**
并发实现: 协程
	比线程更小,需要内存少4K-5K(初始内存)
多协程通行
	channel 管道
多协程同步
	sync.WaitGroup
*/

func main() {
	// go 并发:协程,如果不延时,很可能main函数执行完,协程都没跑完.
	go model.Create()
	time.Sleep(time.Second)
}
