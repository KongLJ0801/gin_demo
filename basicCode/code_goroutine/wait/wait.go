/*

   @author #Kkk
   @File wait
   @Description:
   @version 0.1
   @date 2022/8/5 14:31

*/

package wait

import (
	"fmt"
	"sync"
)

func Wait() {
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(num int) {
			fmt.Printf("Goroutine Test i:= %d \n", num)
			wg.Done()
		}(i)
	}
	defer wg.Wait()
}
