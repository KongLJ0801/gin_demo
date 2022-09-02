/*

   @author #Kkk
   @File context_go
   @Description:
   @version 0.1
   @date 2022/8/5 14:42

*/

package context_go

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func ContextGo() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	var runTime = 0
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("goroutine Done")
				return
			default:
				fmt.Printf("goroutine Test, runTime := %d\n", runTime)
				runTime = runTime + 1
			}
			if runTime > 5 {
				cancel()
				wg.Done()
			}
		}
	}(ctx)
	defer wg.Wait()
}
