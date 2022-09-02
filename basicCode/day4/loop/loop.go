package loop

import (
	"fmt"
	"math/rand"
	"time"
)

/**
 *
 * @date 2022/7/18 14:16
 * @version 0.1
 * @author KongLingJie
 *
 */

func basic() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("arr is val =[%d],%d\n", arr[i], i)
	}
	for i := 0; i < len(arr); {
		fmt.Printf("arr is val =[%d],%d\n", arr[i], i)
		i++
	}
	i := 0
	sum := 0
	for i < len(arr) {
		sum += arr[i]
		fmt.Printf("arr is val =[%d],%d\n", arr[i], i)
		i++
	}
	for sum, i := 0, 0; i < len(arr) && sum < 100; sum, i = sum+arr[i], i+1 {
		fmt.Printf("arr is val =[%d],i =%d, sum = %d\n", arr[i], i, sum)
	}
}

func basicRange() {
	//for range 元素数据拷贝 不影响外部数据
	arr := make([]int, 10)
	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr = append(arr, 10)
	for i, i2 := range arr {
		fmt.Printf("arr is val =[%d],%d\n", i2, i)
	}

	m := make(map[int]int, 10)
	m[0] = 1
	m[2] = 3
	for i, ele := range m {
		fmt.Printf("m is val =[%d],%d\n", ele, i)
	}

	ch := make(chan int, 10)
	for i := 0; i < cap(ch); i++ {
		ch <- i
	}
	close(ch)

	for ele := range ch {
		fmt.Printf("ch is val =[%d]\n", ele)
	}

}

func dead_loop() {
	sum := 0
	for {
		fmt.Println("zzz", sum)
		time.Sleep(1 * time.Second)
		sum += 1
		if sum > 10 {
			break
		}
	}
}

func genMat() [4][4]int {
	a := [4][4]int{}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			a[i][j] = rand.Int()
		}
	}
	return a
}

func basic_for(a, b [4][4]int) [4][4]int {
	c := [4][4]int{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			sum := 0
			for k := 0; k < 4; k++ {
				sum += a[i][k] * b[k][j]
			}
			c[i][j] = sum
		}
	}
	return c
}

func LoopMain() {
	//basic()

	//basicRange()

	//dead_loop()

	A := genMat()
	B := genMat()

	C := basic_for(A, B)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("%d ", C[i][j])
		}
		fmt.Println()
	}
}
