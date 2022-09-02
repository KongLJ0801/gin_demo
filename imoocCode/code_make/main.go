/*

   @author #Kkk
   @File code_make
   @Description: init make (slice,map,chan)
   @version 0.1
   @date 2022/8/8 16:35

	make :初始化,返回引用内型
*/

package main

import "fmt"

func main() {
	//makeSlice()
	//appendAndCopySlice()

	//makeMap()
	deleteMap()
	//makeChan()

}

// 创建 slice
func makeSlice() {
	mSlice := make([]string, 3)
	mSlice[0] = "dog"
	mSlice[1] = "cat"
	mSlice[2] = "tiger"
	fmt.Println(mSlice)
}

// slice->(append/copy)
func appendAndCopySlice() {
	mSlice := make([]int, 3, 5)
	mSlice = append(mSlice[:0], []int{1, 2, 3}...)
	// len 实际长度, cap 容量   1.8-version (1.cap < threshold = cap*2   2.cap+= (cap+3*threshold)/4)    1.7-version  cap > 1024 = cap*1.25     cap < 1024 = cap*2
	// 超出 cap  开辟新的内存   改变长度需要使用append
	fmt.Println(mSlice, len(mSlice), cap(mSlice))

	fmt.Printf("mSlice:= %p \n", mSlice)
	cSlice := make([]int, 2)
	fmt.Printf("cSlice:= %p \n", cSlice)
	// 容量不够,不会扩容
	copy(cSlice, mSlice)
	fmt.Println(cSlice)
}

// 创建 map
func makeMap() {
	mMap := make(map[int]string, 10)
	mMap[0] = "dog"
	mMap[1] = "cat"
	fmt.Println(mMap)
}

// map -> (delete)
func deleteMap() {
	mMap := make(map[int]string)
	mMap[0] = "dog"
	mMap[1] = "cat"
	delete(mMap, 0)
	fmt.Println(mMap)
}

// 创建 chan
func makeChan() {
	mChan := make(chan int)
	defer close(mChan)
	mChan <- 1
	fmt.Println(<-mChan)

}
