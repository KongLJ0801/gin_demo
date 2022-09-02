package _array

import "fmt"

/**
 *
 * @date 2022/7/14 14:31
 * @version 0.1
 * @author KongLingJie
 *
 */

func basic() {
	var arr1 [5]int = [5]int{1, 2, 3, 4: 8}
	var arr2 = [...]int{1, 2, 4}
	fmt.Println(arr1)
	arr2[len(arr2)-1] = 8
	fmt.Println(arr2, len(arr2), arr2[0])
}

func mean(arr [5]int) float64 {
	var sum int
	//for i := 0; i < len(arr); i++ {
	//	sum += arr[i]
	//}
	for _, i2 := range arr {
		sum += i2
	}
	return float64(sum / len(arr))
}

func arrPoint(arr *[5]int) {
	arr[0] += 10 // 解析地址
	fmt.Println(arr[0])
	//(*arr)[0] = 8
}

func for_range() {
	//arr := [5][3]int{{123}, {1}, {122, 2, 3}}
	////arr1 := [...][3]int{{123}, {1}, {122, 2, 3}}
	////fmt.Println(arr)
	////fmt.Println(len(arr))
	////fmt.Println(cap(arr))
	////fmt.Println(arr1)
	////fmt.Println(len(arr1))
	////fmt.Println(cap(arr1))
	//
	//for i, row := range arr {
	//	fmt.Printf("%T \n", row)
	//	for j, col := range row {
	//		fmt.Println(col, i, j)
	//	}
	//
	//}

	arr := [...]int{1, 2, 3, 4, 5}
	for i := range arr {
		arr[i] += 8
	}
	fmt.Println(arr[len(arr)-1:])
}

func intPointer() {
	//arr := [5]int{1, 2, 3, 4}

	//m := mean(arr)
	//fmt.Println(m)

	// 数组指针, 修改原来地址指
	//arrPoint(&arr)
	//fmt.Println(arr)
	//fmt.Println(len(arr))
	//fmt.Println(cap(arr))

	//var a int = 9
	//var b *int
	//b = &a
	//*b = 3
	//fmt.Println(*b, a)

	var arr [5]int
	arrPoint(&arr)
	fmt.Println(arr)
}

func ArrayMain() {

	for_range()
	//intPointer()
}
