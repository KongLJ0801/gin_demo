/*

   @author #Kkk
   @File code_channel
   @Description:
   @version 0.1
   @date 2022/8/8 15:14

*/

package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex

func func1() {
	var count int
	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			// 在 Lock 和 Unlock 方法之间的代码段称为资源的临界区，这一区间的代码是严格被锁保护的，是线程安全的，任何一个时间点最多只能有一个goroutine在执行。
			// 当某一goroutine执行了mutex.lock()方法后，如果有其他的goroutine来执行上锁操作，会被阻塞，直到当前的goroutine执行mutex.unlock()方法释放锁后其他的goroutine才会继续抢锁执行。
			count = count + 1
			mu.Unlock()
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(count)
}

func main() {
	syncWaitGroup()
	//func1()

}

func syncWaitGroup() {
	ch := make(chan int, 10)

	for i := 0; i < cap(ch); i++ {
		ch <- i
	}
	// 任务编排
	var wg sync.WaitGroup
	// 子任务完毕
	defer wg.Wait()
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(num int) {
			select {
			case c := <-ch:
				fmt.Printf("ch := %d num := %d \n", c, num)
			default:

			}
			wg.Done()
		}(i)
	}
}

func code() {
	ch := make(chan int)

	go func() {
		for {
			println(<-ch)
		}
	}()

	ch <- 1
	ch <- 2

	defer close(ch)
}
