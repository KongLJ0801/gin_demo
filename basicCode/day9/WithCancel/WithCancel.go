package WithCancel

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/**
 *
 * @date 2022/7/26 17:44
 * @version 0.1
 * @author KongLingJie
 *
 */
var wg sync.WaitGroup

type UserName string
type UserId string

func worker(ctx context.Context) {
	key := UserName("user_name")
	idKey := UserId("user_id")
	u, ok := ctx.Value(key).(string)
	if !ok {
		fmt.Println("err user name", ok)
	}
	id, ok := ctx.Value(idKey).(int64)
	if !ok {
		fmt.Println("err user id", ok)
	}
	fmt.Printf("username = %s, user id = %d \n", u, id)
LOOP:
	for {
		time.Sleep(time.Millisecond * 1)
		select {
		case <-ctx.Done():
			break LOOP
		default:

		}
	}

	wg.Done()
}

func WithCancelMain() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	// WithValue :注意key 要是自定义类型
	ctx = context.WithValue(ctx, UserName("user_name"), "hole")
	ctx = context.WithValue(ctx, UserId("user_id"), int64(6545645))
	wg.Add(1)
	go worker(ctx)
	time.Sleep(5 * time.Second)
	cancel()
	wg.Wait()
}
