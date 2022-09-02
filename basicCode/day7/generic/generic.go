package generic

/**
 *
 * @date 2022/7/20 9:46
 * @version 0.1
 * @author KongLingJie
 *
 */

func addint4(a, b int) int {
	return a + b
}

func addfloat64(a, b float64) float64 {
	return a + b
}

type A[T int | float32 | float64 | string] struct {
}

// 方法可以使用类型定义中的形参 T
func (receiver A[T]) Add(a T, b T) T {
	return a + b
}

func GenericFunc() {
	var a, b int
	var c, d string
	var e, f float64

	addint4(a, b)
	//addfloat64(c, d)
	addfloat64(e, f)

	var g A[string]
	g.Add(c, d)
}
