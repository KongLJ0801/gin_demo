package math

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

/**
 *
 * @date 2022/7/22 15:19
 * @version 0.1
 * @author KongLingJie
 *
 */

func constant() {
	fmt.Println(math.E)
	fmt.Println(math.Pi)
	fmt.Println(math.MaxInt64)
	fmt.Printf("%b\n", math.MaxInt64) // 64-1
	fmt.Println(uint64(math.MaxUint64))
	fmt.Printf("%b\n", uint64(math.MaxUint64))

	// 精度
	fmt.Println(math.SmallestNonzeroFloat32)
	fmt.Println(math.SmallestNonzeroFloat64)

}

// not a number
func nan() float64 {
	if f, err := strconv.ParseFloat("123.5465464", 64); err == nil {
		return f
	} else {
		return math.NaN()
	}

}

func math_func() {
	fmt.Println(math.Ceil(2.5))   // 向上取整  3
	fmt.Println(math.Floor(2.5))  // 向下取整  2
	fmt.Println(math.Ceil(-2.5))  // -2
	fmt.Println(math.Floor(-2.5)) // -3
	fmt.Println(math.Trunc(2.5))  // 截取整数 2
	fmt.Println(math.Modf(2.5))   // 分割整数,小数 2 0.5
	fmt.Println(math.Abs(-2.5))   // 绝对值 2.5

	fmt.Println(math.Max(3, 7)) // 最大值 7
	fmt.Println(math.Min(3, 7)) // 3

	// x-y if x-y > 0  x  else  0
	fmt.Println(math.Dim(3, 7)) // 0
	fmt.Println(math.Dim(7, 3)) // 4

	// 平方
	fmt.Println(math.Sqrt(9))   //3
	fmt.Println(math.Pow(3, 2)) //9
}

func ran() {
	//rand.Seed(time.Now().UnixMicro())  // 解决 rand 种子问题
	//fmt.Println(rand.Int())      // 随机值
	//fmt.Println(rand.Int31n(26)) // 随机范围值
	//arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//// 乱序
	//rand.Shuffle(len(arr), func(i, j int) {
	//	arr[i], arr[j] = arr[j], arr[i]
	//})
	//fmt.Println(arr)

	seed := time.Now().UnixMicro()
	source := rand.NewSource(seed)
	rand1 := rand.New(source)
	fmt.Println(rand1.Intn(26), rand1.Intn(26), rand1.Intn(26), rand1.Intn(26), rand1.Intn(26))

	source.Seed(seed)
	rand2 := rand.New(source)

	fmt.Println(rand2.Intn(26), rand2.Intn(26), rand2.Intn(26), rand2.Intn(26), rand2.Intn(26))

}
func shuffle(arr *[]int) {
	rand.Seed(time.Now().UnixNano())
	n := len(*arr)
	for i := 0; i < 5*n; i++ {
		idx := rand.Intn(n)
		idj := rand.Intn(n)
		(*arr)[idx], (*arr)[idj] = (*arr)[idj], (*arr)[idx]

	}
}

func MathFunc() {
	//constant()

	//ret := nan()
	//if math.IsNaN(ret) {
	//	fmt.Println("解析出错了!")
	//} else {
	//	fmt.Println(ret)
	//}

	//math_func()

	//ran()

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	shuffle(&arr)
	fmt.Println(arr)
}
