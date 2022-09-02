package WithDeadline

import (
	"context"
	"fmt"
	"time"
)

/**
 *
 * @date 2022/7/26 17:12
 * @version 0.1
 * @author KongLingJie
 *
 */

func WithDeadlineMain() {
	d := time.Now().Add(50 * time.Millisecond) // 50毫秒
	ctx, cancel := context.WithDeadline(context.Background(), d)
	x := time.After(1 * time.Second)
	fmt.Printf("t : %v", x)
	defer cancel()
	select {
	case <-time.After(1 * time.Second): // 等待一秒取值才会打印
		fmt.Println("time.After")

	case <-ctx.Done(): // 立即执行
		fmt.Println(ctx.Err())
	}
}
