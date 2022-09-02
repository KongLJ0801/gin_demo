/*

   @author #Kkk
   @File code_slice
   @Description:  slice 切片
   @version 0.1
   @date 2022/8/8 10:20

	1.slice 为什么不是线程安全的 ?
	slice底层结构并没有使用加锁的方式,不支持并发读写

	2.数组和切片区别 ?
	数组是具有相同唯一类型的一组已编号且长度固定的数据项序列
    切片是内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大
*/

package main

import "fmt"

func initSlice() {
	var sl []int
	println(sl == nil) // 初始值是nil 所以是 true
	sl = append(sl, 1) // 添加元素
	// 内存扩容  版本1.7     容量 < 1024  = 容量*2    容量 > 1024  = 容量*1.25
	// 内存扩容  版本1.8     容量 < threshold  = 容量*2    容量 > threshold  = (旧容量 += (旧容量+3*threshold)/4)
	println(len(sl), cap(sl))
}

func makeSlice() {
	// len,cap =5    len长度,cap容量
	ms := make([]int, 5) // [0 0 0 0 0]
	for i, ele := range ms {
		println(i, ele)
	}
	println(len(ms), cap(ms))
	ms = append(ms, 1) // 添加元素
	fmt.Printf("ms:=%+v \n", ms)
	//ms[len(ms)-1]  获取最后一位元素
	println(len(ms), cap(ms), ms[len(ms)-1])

}
func sonAndMother() {
	ms := make([]int, 5, 5) // 父  cap =5
	sam := ms[0:3]          // 子   cap = 5
	sam = append(sam, 20)
	sam = append(sam, 10)
	sam = append(sam, 30) // cap 5*2  和父,不在有关联(超出父容量,产生新的空间)
	fmt.Printf("sam:=%+v len:=%d,cap:=%d, ms := %+v\n", sam, len(sam), cap(sam), ms)
	fmt.Printf("addr is sam :=%p, or ms := %p \n", sam, ms)

	sam = append(sam[:3], sam[3+1:]...) // 删除指定元素 20
	fmt.Printf("sam := %+v, addr sam := %p ,len:=%d,cap:=%d \n", sam, sam, len(sam), cap(sam))
	sam = append(sam[:0], sam[2+1:]...) // 删除前三个元素
	fmt.Printf("sam := %+v, addr sam := %p ,len:=%d,cap:=%d \n", sam, sam, len(sam), cap(sam))
}
func main() {
	//initSlice()
	//makeSlice()

	sonAndMother()
}
