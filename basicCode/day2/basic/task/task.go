package task

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
)

/**
 *
 * @date 2022/7/15 10:53
 * @version 0.1
 * @author KongLingJie
 *
 */

func genSlice(n int) []int {
	// make 类型 : slice   len : 0 初始化 最好是0 (如果>0 , 里面会有0填充) cap (需要的容器,避免append 扩充)
	arr := make([]int, 0, n)
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(10))
	}
	return arr
}

func countUniq(arr []int) int {
	// make 类型 : map[int]bool   key int 8字节, value bool 1字节 ,
	// len(arr) 如果不写第三个参数 默认第二个参数和第三个参数一样
	// map 去重  key 是会覆盖
	m := make(map[int]bool, len(arr))
	for _, b := range arr {
		m[b] = true
	}
	return len(m)
}

func concatString(arr []int) string {

	fmt.Println(reflect.TypeOf(arr))
	sb := strings.Builder{}
	for i, i2 := range arr {
		fmt.Printf("x[%d] 类型为int,内容为%d\n", i, i2)
		//strconv.Ito 整形转字符串
		sb.WriteString(strconv.Itoa(i2))
		//sb.WriteString(strconv.FormatInt(int64(i2), 10))
		sb.WriteString(" ")
	}
	// 去除末尾空格
	// 1.
	// strings.Trim(sb.String(), " ")
	// 2. 存在内存泄露 (子切片)
	// s := sb.String()
	// s = s[0 : len(s)-1]

	return strings.Trim(sb.String(), " ")
}

func TaskMain() {
	// 1: 100随机数 去重

	arr := genSlice(10)

	//fmt.Println(arr, len(arr), cap(arr))
	//
	//cut := countUniq(arr)
	//
	//fmt.Println(cut)

	// 2. 切换转字符串 且数据类型问题
	//x := concatString([]int{1, 3, 3})
	x := concatString(arr)
	fmt.Printf("x:= %s, \n", x)

}
