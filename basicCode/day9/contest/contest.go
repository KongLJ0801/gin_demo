package contest

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/**
 *
 * @date 2022/7/26 16:22
 * @version 0.1
 * @author KongLingJie
 *
 */

var wg sync.WaitGroup

func worker(ctx context.Context) {
	defer wg.Done()
	go worker2(ctx)
LABEL:
	for {
		fmt.Println("worker...")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			break LABEL
		default:

		}
	}
}

func worker2(ctx context.Context) {
	defer wg.Done()
LABEL:
	for {
		fmt.Println("worker222...")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			break LABEL
		default:

		}
	}
}

func Gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	i := 1

	go func() {
		select {
		case <-ctx.Done():
			return
		case dst <- i:
			i++
			fmt.Println(i, "gen")
		}
	}()

	return dst
}

func ContextMain() {
	//ctx, cancel := contest.WithCancel(contest.Background())
	//
	//wg.Add(2)
	//
	//go worker(ctx)
	//time.Sleep(time.Second * 5)
	//cancel()
	//wg.Wait()
	//fmt.Println("over ")

	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)

	defer cancel()

	for n := range Gen(ctx) {
		fmt.Println(n, "range")
		if n == 5 {
			break
		}
	}

	wg.Wait()

}
