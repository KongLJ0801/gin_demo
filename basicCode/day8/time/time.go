package time

import (
	"fmt"
	"time"
)

/**
 *
 * @date 2022/7/22 9:53
 * @version 0.1
 * @author KongLingJie
 *
 */
func basic() {
	// 时间戳
	//t := time.Now()
	//fmt.Println(t.Unix())
	//fmt.Println(t.UnixMilli())
	//fmt.Println(t.UnixMicro())
	//fmt.Println(t.UnixNano())
	//
	begin := time.Now()
	//
	//var sum int
	//
	//for i := 0; i < 10000000; i++ {
	//	sum += i
	//}
	//
	//end := time.Since(begin) // 计算开始时间到结束时间
	//
	//fmt.Println(end)
	//
	//dua := time.Duration(8 * time.Hour)
	//
	//ends := begin.Add(dua)
	//
	//fmt.Println(ends)

	fmt.Println(begin.Weekday())    // 星期
	fmt.Println(begin.Hour())       // 时
	fmt.Println(begin.Minute())     // 分
	fmt.Println(begin.Second())     // 秒
	fmt.Println(begin.Year())       // 年
	fmt.Println(begin.YearDay())    // 年的第几天
	fmt.Println(int(begin.Month())) // 月   (强制转换为 int , month 底层是   type Month int)
	if int(begin.Month()) < 10 {
		fmt.Printf("0%d \n", int(begin.Month()))
	} else {
		fmt.Println(int(begin.Month()))
	}

	fmt.Println(begin.Day()) // 月的第几天

}

const (
	TIME_FMT  = "2006-01-02 15:04:05"
	TIME_FMT2 = "2006-01-02"
	TIME_FMT3 = "15:04:05"
)

var Loc *time.Location

func init() {
	//Loc, _ = time.LoadLocation("Asia/shanghai") // 解决时区问题 (东巴区)
}

func time_fmt() {
	begin := time.Now()
	fmt.Println(begin.Format(TIME_FMT))
	fmt.Println(begin.Format(TIME_FMT2))
	fmt.Println(begin.Format(TIME_FMT3))

	if d, err := time.ParseInLocation(TIME_FMT, "2022-10-22 15:11:42", Loc); err == nil {
		fmt.Println(d.Year())
		fmt.Println(d.Month())
		//fmt.Println(int(d.Month()))
		fmt.Println(d.Day())

	}

}

func ticker_func() {
	// 定时执行
	tc := time.NewTicker(5 * time.Second)
	defer tc.Stop()
	for i := 0; i < 6; i++ {
		<-tc.C
		fmt.Println(time.Now().Unix())
	}
}
func timer_func() {
	// 1.延时执行,只执行一次 (管道)
	fmt.Println(time.Now().Unix())
	tc := time.NewTimer(time.Duration(3 * time.Second))
	defer tc.Stop()
	<-tc.C
	fmt.Println(time.Now().Unix(), '-')

	// 2.延时执行,Reset执行多次 (管道)
	for i := 0; i < 6; i++ {
		tc.Reset(1 * time.Second)
		<-tc.C
		fmt.Println(time.Now().Unix(), i)
	}

	//3.无需stop (管道)
	//<-time.After(3 * time.Second)
	//fmt.Println(time.Now().Unix())

	//4 sleep  延时执行
	time.Sleep(3 * time.Second)
	fmt.Println(time.Now().Unix(), 's')

}
func TimeFunc() {
	//basic()

	//time_fmt()

	//ticker_func()

	timer_func()
}
