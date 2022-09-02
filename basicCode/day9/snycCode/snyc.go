package snycCode

import (
	"fmt"
	"sync"
	"time"
)

/**
 *
 * @date 2022/7/26 15:24
 * @version 0.1
 * @author KongLingJie
 *
 */

var wg sync.WaitGroup

// 全局方式退出worker
var exit bool

func worker() {
	defer wg.Done()
	for {
		fmt.Println("worker ")
		time.Sleep(time.Second)
		if exit {
			break
		}
	}
}

func workerChannel(ch <-chan bool) {
	defer wg.Done()
LABEL:
	for {
		select {
		case <-ch:
			break LABEL
		default:
			fmt.Println("worker ")
			time.Sleep(time.Second)
		}

	}
}

func SyncMain() {
	// 1.全局方式退出worker
	//wg.Add(1)
	//go worker()
	//time.Sleep(time.Second * 5)
	//exit = true
	//wg.Wait()
	//fmt.Printf("symc.waitGroup ")

	//2.channel
	ch := make(chan bool, 1)
	wg.Add(1)
	go workerChannel(ch)
	time.Sleep(time.Second * 5)
	ch <- true
	wg.Wait()
	fmt.Printf("symc.waitGroup ")
}
