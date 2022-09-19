/*

   @author #Kkk
   @File code_lo
   @version 0.1
   @date 2022/9/13 16:02
   @Description:
   @Analysis:

*/

package main

import (
	"fmt"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	"strconv"
	"strings"
)

func main() {
	//lUniq()
	//lUniqBy()
	//lGroupBy()
	//lFilter()
	//lMap()
	//lFilterMap()
	//lFlatMap()
	//lReduce()
	//lReduceRight()
	//lForEach()
	//lopForEach()
	//lTimes()
	//lChunk()
	//lPartitionBy()
}
func lPartitionBy() {
	// 数据分割   //  lop.PartitionBy
	partitionBy := lo.PartitionBy[int, int]([]int{-2, -1, 0, 1, 2, 3, 4, 5, 6}, func(x int) int {
		if x < 0 {
			return 0
		} else if x%2 == 0 {
			return 1
		}
		return 2
	})
	// [][]int{[]int{-2, -1}, []int{0, 2, 4, 6}, []int{1, 3, 5}}
	fmt.Printf("%#v", partitionBy)
	// [[-2 -1] [0 2 4 6] [1 3 5]]
	fmt.Println(partitionBy)
}

func lChunk() {
	// 多维slice
	vChunk := lo.Chunk[int]([]int{1, 2, 3, 4}, 3)
	fmt.Println(vChunk) // [][]int{{1,2,3},{4}}
}

func lopTimes() {
	vTime := lop.Times[string](3, func(i int) string {
		return strconv.FormatInt(int64(i), 10)
	})
	fmt.Println(vTime) // []string{"0", "1", "2"}
}

func lTimes() {
	vTime := lo.Times[string](3, func(i int) string {
		return strconv.FormatInt(int64(i), 10)
	})
	fmt.Printf("%#v \n", vTime) // []string{"0", "1", "2"}
}

func lopForEach() {
	// 并行处理：类似于 lo.ForEach()，但回调被称为 goroutine。
	lop.ForEach[string]([]string{"lop", "lp"}, func(x string, _ int) {
		fmt.Println(x) // lop \n lp
	})
}

func lForEach() {
	// 遍历
	lo.ForEach[string]([]string{"tom", "jack"}, func(x string, _ int) {
		fmt.Println(x) // tom \n  jack
	})
}

func lReduceRight() {
	// 右到左
	result := lo.ReduceRight[[]int, []int]([][]int{{1, 2}, {3, 4}, {5, 6}}, func(r []int, t []int, _ int) []int {
		return append(r, t...)
	}, []int{7, 8})
	fmt.Println(result) // []int{7,8,5,6,3,4,1,2}
}

func lReduce() {
	// 左到右
	// 将集合减少为单个值。通过累加器函数累加集合中运行每个元素的结果来计算该值。每次连续调用都会提供上一次调用返回的返回值。
	sum := lo.Reduce[int, int]([]int{1, 2, 3, 4}, func(r int, t int, _ int) int {
		return r + t
	}, 0) // 0 基于某个值开始累加
	fmt.Println(sum) // 10
}

func lFlatMap() {
	// 操作切片并将其转换并展平为另一种类型的切片
	vFlatMap := lo.FlatMap[int, string]([]int{1, 2, 3, 4}, func(x int, _ int) []string {
		return []string{
			strconv.Itoa(x),
			strconv.Itoa(x),
		}
	})
	fmt.Println(vFlatMap) // []int{1,1,2,2,3,3,4,4}
}

func lFilterMap() {
	// 过滤不包含,包含,并修改返回值
	matching := lo.FilterMap[string, string]([]string{"cpu", "gpu", "mouse", "keyboard"}, func(x string, _ int) (string, bool) {
		if strings.HasSuffix(x, "pu") {
			return "xpu", true
		}
		return "", false
	})
	fmt.Println(matching) // []string{"xpu","xpu"}
}

func lMap() {
	// 类型转换
	vMap := lo.Map[int64, string]([]int64{1, 2, 3, 4}, func(x int64, _ int) string {
		return strconv.FormatInt(x, 10)
	})
	fmt.Printf("%#v", vMap) // []string{"1", "2", "3", "4"}
}

func lFilter() {
	// 过滤
	even := lo.Filter[int]([]int{1, 2, 3, 4}, func(x int, _ int) bool {
		return x%2 == 0
	})
	fmt.Println(even) // [2,4]
}

func lGroupBy() {
	// lop.GroupBy
	groups := lo.GroupBy[int, int]([]int{0, 1, 2, 3, 4, 5}, func(i int) int {
		return i % 3
	})
	fmt.Println(groups) // map[0:[3] 1:[1 4] 2:[2 5]]
}

func lUniqBy() {
	// 取余,去重，并且排序
	vUniqBy := lo.UniqBy[int, int]([]int{1, 2, 3, 4, 5}, func(i int) int {
		return i % 3
	})
	fmt.Println(vUniqBy) // []int{1,2,3}
}

func lUniq() {
	// 去重
	names := lo.Uniq[string]([]string{"Samuel", "Marc", "Samuel"})
	fmt.Println(names) // ["Samuel", "Marc"]

	// 去重，并且排序
	vUniq := lo.Uniq[int]([]int{1, 2, 2, 1})
	fmt.Println(vUniq)
}
