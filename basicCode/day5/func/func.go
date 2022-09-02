package _func

import "fmt"

/**
 *
 * @date 2022/7/18 16:04
 * @version 0.1
 * @author KongLingJie
 *
 */

func add(a, b *int) (sum int) {
	/**

	* 解析
	& 取址

	*/
	*b = 888 + *a
	sum = *a + *b
	return
}

func update_slice(arr []int) {
	arr[0] = 8           // 外部 len 范围内,所以 修改了默认值,会影响外部值
	arr = append(arr, 3) // 修改了 len 的长度 开辟了新的内存空间所以外部不受影响
	fmt.Println(arr, "update_slice")

}

func FuncMain() {
	a, b := 1, 2
	add(&a, &b)
	fmt.Printf("a = %d, b = %d \n", a, b)

	arr := make([]int, 2, 5)
	arr[0] = 1
	update_slice(arr)
	fmt.Println(arr)
}
