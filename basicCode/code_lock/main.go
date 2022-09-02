/*

   @author #Kkk
   @File code_lock
   @Description:
   @version 0.1
   @date 2022/8/20 16:28

*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	lock()
}
func lock() {
	var mutex sync.Mutex
	count := 0

	for r := 0; r < 20; r++ {
		go func() {

			if r%5 == 0 {
				mutex.Lock()
				count += 1
				mutex.Unlock()
				fmt.Println(r%5 == 0, r%5, r, "true")
			} else {
				fmt.Println(r%5 == 0, r%5, r)
			}
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("the count is : ", count)
}
