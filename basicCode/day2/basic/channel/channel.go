package channel

import "fmt"

/**
 *
 * @date 2022/7/15 9:34
 * @version 0.1
 * @author KongLingJie
 * 管道 (容器)
 * channel 先存先取,cap超出初始长度不会开辟新的内存,后续无法执行.
 */

func initChannel() {
	//var cha chan int
	//fmt.Printf("channel is nil %t \n", cha == nil) // true
	//fmt.Printf("channel of len %d \n ", len(cha)) // 0

	cha := make(chan int, 10) // 初始化

	//fmt.Printf("channel of len %d \n ", len(cha)) // 0
	//fmt.Printf("channel of len %d \n ", cap(cha)) // 10

	X := cap(cha)

	for i := 0; i < 20; i++ {
		//cha <- i // 存值
		//fmt.Printf("i := %d , X := %d \n", i, X)
		if i < X {
			cha <- i // 存值
		} else {
			break
		}

	}
	fmt.Printf("channel of len %d  cha := %v\n", len(cha), cha) // 10
	fmt.Printf("channel of len %d \n", cap(cha))                // 10

	//<-cha
	//cha <- 3
	//i := <-cha // 取值
	//fmt.Printf("%T,%d \n", i, i)
	//fmt.Printf("channel of len %d \n", len(cha)) // 9
	//fmt.Printf("channel of len %d \n", cap(cha)) // 10

	// 1:
	//L := len(cha)   // 只读 cap(cha)长度 所以不会死锁   所以不用写close
	//for i := 0; i < L; i++ {
	//	ele := <-cha
	//	fmt.Printf("ele :=  %d \n", ele)
	//}
	//fmt.Printf("channel of len %d \n", len(cha)) // 10
	//fmt.Printf("channel of len %d \n", cap(cha)) // 10

	// 2:
	//close(cha) // for range ( 等待 可能存在其他线程写入/取出  会阻塞)   (close 后不能再写入数据)
	for ele := range cha {
		fmt.Println(ele)
		if ele == 4 {
			close(cha) // ???
			break
		}
	}
	fmt.Printf("channel of len %d \n", len(cha)) // 5
	fmt.Printf("channel of len %d \n", cap(cha)) // 10
}

func ChannelMain() {
	initChannel()
}
