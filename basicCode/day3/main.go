package main

import (
	"fmt"
	copy2 "gin_demo/basicCode/day3/copy"
)

type User struct {
	Name string
	age  int
}

func setUser() []*User {
	arr := make([]*User, 10)

	// 去掉 最后一个元素 也会存在问题.
	/**
	arr[len(arr)-1] = nil
	需要把最后一个元素处理成 nil 值 就不会返回了 ,也不会存在问题
	*/
	arr[len(arr)-1] = nil
	return arr[0 : len(arr)-1]
}

// 解决内存泄露问题
func foo() []int {
	arr := make([]int, 1<<20) // 1M, len==cap==1M
	//return arr[3:7] // 存在泄露问题 因为 母切片和子切片关联, 而只返回了 3:7 的元素,母切片还有很多元素没处理
	// 当垃圾回收机制 出发时, 母切片还有很多元素没有处理,所有导致了内存泄露

	// 处理方法   只获取 3:7 的元素 获取出来的元素和 arr 并不关联性, 开辟了新的空间 arr1
	arr1 := make([]int, 10)
	for i := 3; i < 7; i++ {
		arr1 = append(arr1, arr[i])
	}
	return arr1
}

func main() {
	//user.UserMain()
	copy2.CopyMain()

	arr := foo()

	fmt.Println(arr)

	return
}
