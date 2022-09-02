package slice

import (
	"fmt"
	"strings"
)

/**
 *
 * @date 2022/7/14 15:57
 * @version 0.1
 * @author KongLingJie
 *	切片
 *
 */

func initSlice() {
	var aa []int
	fmt.Println(len(aa), cap(aa))

	var a = make([][]int, 2, 3)
	fmt.Println(len(a), cap(a))
}

func appendSlice() {
	//a := []int{1, 2, 3, 4}

	// 超出  cap 空间 cap = a*2   cap>1024 ( cap = 1024 * 1.25 )
	a := make([]int, 2, 3)
	a = append(a, 100)
	a = append(a, 200)
	a = append(a, 300)
	a = append(a, 400)
	a = append(a, 500)
	fmt.Println(a, len(a), cap(a))
	a[1] = 10
	fmt.Println(a, len(a), cap(a))

}

func sliceSub() {
	// 子切片 超出 母切片的 CAP 着 两个切片不在有关联(内存分离),如没超出 修改 子母切片的任意值都会受影响

	a := make([]int, 3, 5)
	sub := a[1:2]
	fmt.Printf("a = %p ,sub =%p \n", &a[1], &sub[0])
	fmt.Println(len(a), cap(a), a)
	fmt.Println(len(sub), cap(sub), sub)
	a[0] = 8
	sub[0] = 18
	fmt.Println(len(a), cap(a), a)
	fmt.Println(len(sub), cap(sub), sub)
	sub = append(sub, 5)
	sub = append(sub, 9)
	sub = append(sub, 10)
	sub = append(sub, 11)
	sub = append(sub, 22)
	sub = append(sub, 33)
	sub = append(sub, 44)
	sub = append(sub, 55)
	a[0] = 100
	sub[0] = 28

	fmt.Println(len(a), cap(a), a)
	fmt.Println(len(sub), cap(sub), sub)
	fmt.Printf("a = %p ,sub= %p \n", &a[1], &sub[0])

}

func updateSlice(arr []int) {
	arr[0] = 100
}

func SliceMain() {
	//initSlice()
	//appendSlice()
	//sliceSub()
	arr := make([]int, 3, 5)
	sub := arr[1:3]
	updateSlice(sub)
	fmt.Println(arr, sub)

	msg := "\njack"
	if !(strings.LastIndex(msg, `\n`) < 0) {
		msg = strings.Split(msg, `\n`)[0]
		fmt.Println(msg)
	} else {
		fmt.Println(msg)
	}

}
