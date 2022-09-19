/*

   @author #Kkk
   @File code_string
   @Description:
   @version 0.1
   @date 2022/8/1 10:16

*/

package main

import (
	"bytes"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"math"
	"strconv"
	"strings"
	"unicode"
)

var h string = "Hello"
var w string = "world"

// 字符串 比较
func compareStr() {
	//  0 : 成功, -1 失败
	println(strings.Compare(h, w)) // -1
	println(strings.Compare(h, h)) // 0

	// 计算 s 与 t 忽略字母大小写后是否相等。
	s := "Go"
	t := "go"
	println(strings.EqualFold(s, t)) // TRUE
	s = "壹"
	t = "一"
	println(strings.EqualFold(s, t)) // FALSE
}

// 字符串 遍历
func forStr() {
	// 注意事项 获取字符串中某个字节的地址的行为是非法的，例如：&h[i]。
	for i := 0; i < len(h); i++ {
		// 默认获取 h[i]  byte 类型 所以需要转成 string
		println(string(h[i]))
		println(h[i])
	}

}

// 是否存在某个字符或子串
func containsStr() {
	println(strings.Contains(h, "i")) // false

	println(strings.ContainsAny(h, "h & o")) // true

	// unicode 编码
	println(strings.ContainsRune(h, 72)) // TRUE

	//查看这三个函数的源码，发现它们只是调用了相应的 Index 函数（子串出现的位置），然后和 0 作比较返回 true 或 false。
	//func Contains(s, substr string) bool {
	//	return Index(s, substr) >= 0
	//}
}

// 子串出现次数 ( 字符串匹配 )
func countStr() {
	println(strings.Count(h, "l")) // 2

	println(strings.Count("中国", "中")) // 2

}

// 字符串分割为[]string
func fieldsStr() {
	str := "  foo   bar -    baz "
	fmt.Printf("fields : %q ", strings.Fields(str))

	f := func(c rune) bool {
		return unicode.IsSpace(c) || c == '.'
	}
	fmt.Printf("fields : %q ", strings.FieldsFunc("We are humans. We are social animals.", f))
	//实际上，Fields 函数就是调用 FieldsFunc 实现的：
	//空格的定义是unicode.IsSpace
	//func Fields(s string) []string {
	//	return FieldsFunc(s, unicode.IsSpace)
	//}
	for i, s := range strings.FieldsFunc("We are humans. We are social animals.", f) {
		println(i, s)
	}

}

//  字符串分割
func sliceStr() {
	//Split(s, sep) 和 SplitN(s, sep, -1) 等价；SplitAfter(s, sep) 和 SplitAfterN(s, sep, -1) 等价。
	fmt.Printf("%q\n", strings.Split("foo/bar/baz", "/"))      // 不保留 sep(,)       [foo bar baz]
	fmt.Printf("%q\n", strings.SplitAfter("foo,bar,baz", ",")) // 保留 sep(,)   [foo, bar, baz]

	for i, s := range strings.Split("foo,bar,baz", ",") {
		println(i, s)
	}
	// 表示返回的 slice 中最多只有 n 个元素，其中，最后一个元素不会分割
	fmt.Printf("%q\n", strings.SplitN("foo,bar,baz,kkk,", ",", 3)) // ["foo" "bar,baz"]

}

// 字符串是否有某个前缀或后缀
func hasStr() {

	println(strings.HasPrefix(h, "h")) // false   区分大小写

	println(strings.HasSuffix(w, "d")) // true

	//// s 中是否以 prefix 开始
	//func HasPrefix(s, prefix string) bool {
	//	return len(s) >= len(prefix) && s[0:len(prefix)] == prefix
	//}
	//// s 中是否以 suffix 结尾
	//func HasSuffix(s, suffix string) bool {
	//	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
	//}
}

// 字符串查找
func indexStr() {
	println(strings.Index(h, "o"))     // 4
	println(strings.LastIndex(h, "0")) // 1
}

// 字符串 拼接
func joinStr() {

	// 字符串拼接符 +   (生成一个新的字符串)
	hw := h + "," + w
	println(hw, "1") // hello,world

	// 使用函数 strings.Join()
	sj := strings.Join([]string{h, w}, ",")
	println(sj, "2") // hello,world

	//使用 bytes.Buffer
	var buff bytes.Buffer
	buff.WriteString(h)
	buff.WriteString(",")
	buff.WriteString(w)
	println(buff.String(), "3") // hello,world

	// fmt.Sprintf
	println(fmt.Sprintf("%s,%s", h, w), 4) // hello,world

	//strings.Builder{}
	var str strings.Builder
	str.WriteString("hello")
	str.WriteString(",")
	str.WriteString("world")
	println(str.String(), 5) // hello,world
}

// 字符串重复几次
func repeatStr() {
	println("ba" + strings.Repeat("na", 2)) // banana
}

// 字符替换
func mapStr() {
	mapping := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z': // 大写字母转小写
			return r + 32
		case r >= 'a' && r <= 'z': // 小写字母不处理
			return r
		case unicode.Is(unicode.Han, r): // 汉字换行
			return '\n'
		}
		return -1 // 过滤所有非字母、汉字的字符
	}
	println(strings.Map(mapping, "Hello你#￥%……\\n（'World\\n,好Hello^(&(*界gopher..."))
	/**
	hello
	nworldn
	hello
	gopher
	*/
}

// 字符串替换
func replaceStr() {
	// 如果 n < 0，则不限制替换次数，即全部替换
	println(strings.Replace("hole hole hole", "le", "kk", 2)) // hokk hokk hole

	// 该函数内部直接调用了函数 Replace(s, old, new , -1)
	println(strings.ReplaceAll("hole hole hole", "hole", "kkk"))
}

// 字符串大小写转换
func toStr() {
	// 转小写
	println(strings.ToLower("Hole World"))
	println(strings.ToLowerSpecial(unicode.TurkishCase, "Önnek İş"))
	// 转大写
	println(strings.ToUpper("Hole World"))
	println(strings.ToUpperSpecial(unicode.TurkishCase, "örnek iş"))

	println(strings.Title("hElLo wOrLd"))
	println(strings.ToTitle("hElLo wOrLd"))
	println(strings.ToTitleSpecial(unicode.TurkishCase, "hElLo wOrLd"))

	src := []string{
		"hello world!",
		"i with dot",
		"'n ijsberg",
		"here comes O'Brian",
	}
	for _, c := range []cases.Caser{
		cases.Lower(language.Und),
		cases.Upper(language.Turkish),
		cases.Title(language.Dutch), //**
		cases.Title(language.Und, cases.NoLower),
	} {
		println()
		for _, s := range src {
			println(c.String(s))
		}
	}
}

// 字符串(首尾)去除
func trimStr() {
	// 将 s 左侧和右侧中匹配 cutset 中的任一字符的字符去掉
	//func Trim(s string, cutset string) string
	// 将 s 左侧的匹配 cutset 中的任一字符的字符去掉
	//func TrimLeft(s string, cutset string) string
	// 将 s 右侧的匹配 cutset 中的任一字符的字符去掉
	//func TrimRight(s string, cutset string) string
	// 如果 s 的前缀为 prefix 则返回去掉前缀后的 string , 否则 s 没有变化。
	//func TrimPrefix(s, prefix string) string
	// 如果 s 的后缀为 suffix 则返回去掉后缀后的 string , 否则 s 没有变化。
	//func TrimSuffix(s, suffix string) string
	// 将 s 左侧和右侧的间隔符去掉。常见间隔符包括：'\t', '\n', '\v', '\f', '\r', ' ', U+0085 (NEL)
	//func TrimSpace(s string) string
	// 将 s 左侧和右侧的匹配 f 的字符去掉
	//func TrimFunc(s string, f func(rune) bool) string
	// 将 s 左侧的匹配 f 的字符去掉
	//func TrimLeftFunc(s string, f func(rune) bool) string
	// 将 s 右侧的匹配 f 的字符去掉
	//func TrimRightFunc(s string, f func(rune) bool) string

	x := "!!!@@@你!@@好,!@#$ Gophers###$$$"
	println(strings.Trim(x, "@#$!%^&*()_+=-"))
	println(strings.TrimLeft(x, "@#$!%^&*()_+=-"))
	println(strings.TrimRight(x, "@#$!%^&*()_+=-"))
	println(strings.TrimSpace(" \t\n Hello, Gophers \n\t\r\n"))
	println(strings.TrimPrefix(x, "!!!"))
	println(strings.TrimSuffix(x, "$"))

	//f := func(r rune) bool {
	//	return !unicode.Is(unicode.Han, r) // 非汉字返回 true
	//}
	//println(strings.TrimFunc(x, f))
	//println(strings.TrimLeftFunc(x, f))
	//println(strings.TrimRightFunc(x, f))
	//
	//println(strings.TrimRight("222G", "G"))

}

func strFindTrim() {
	s := "\\u0000\\u0000\\u0000\\ufffd\\u0000C{\\\"call_back_key\\\":\\\"scriptExecution\\\",\\\"time_out\\\":30000,\\\"data\\\":[{\\\"destination\\\":\\\"192.168.101.128\\\",\\\"gateway\\\":\\\"192.168.101.5\\\",\\\"netmask\\\":\\\"192.168.101.128\\\"},{\\\"destination\\\":\\\"192.168.101.128\\\",\\\"gateway\\\":\\\"192.168.101.5\\\",\\\"netmask\\\":\\\"192.168.101.128\\\"}]}\\ufffd\n"
	fmt.Println(s[:len(s)-7])
	//fmt.Println(strings.Index(s, "{"))
	//fmt.Println(strings.LastIndex(s, "}"))
	//x := strings.Index(s, "{")
	//y := strings.LastIndex(s, "}")
	//s = s[x : y+1]
	//fmt.Println(s)

	l := func(r rune) bool {
		return !(string(r) == "{") // 非汉字返回 true
	}
	x := strings.TrimLeftFunc(s, l)
	r := func(r rune) bool {
		return !(string(r) == "}") // 非汉字返回 true
	}
	fmt.Println(strings.TrimRightFunc(x, r))

}

// Replacer 类型
func replacerStr() {
	/**
	这是一个结构，没有导出任何字段，
	实例化通过 func NewReplacer(oldnew ...string) *Replacer 函数进行，
	其中不定参数 oldnew 是 old-new 对，即进行多个替换。
	如果 oldnew 长度与奇数，会导致 panic.
	*/
	r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
	println(r.Replace("This is <b>HTML</b>!"))
}

func float64Pow() {
	var flt float64 = 15.856418258
	fmt.Println(Decimal(flt))
	i := 0
	flashStr := strings.TrimSpace("")
	numberPow, _ := strconv.ParseFloat(flashStr, 64)
	for {
		if numberPow < math.Pow(2, float64(i)) {
			break
		}
		i++
	}
	size := strconv.FormatFloat(math.Pow(2, float64(i)), 'f', -1, 64)
	fmt.Printf("%s", size)
}
func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.4f", value), 64)
	return value
}

// str
/**
2022-08-22 16:17:03.629   807   944 W ProcessCpuTracker: Skipping unknown process pid 7279
2022-08-22 16:17:03.629   807   944 W ProcessCpuTracker: Failed to stat(/proc/7280): android.system.ErrnoException: stat failed: ENOENT (No such file or directory)
2022-08-22 16:17:03.629   807   944 W ProcessCpuTracker: Skipping unknown process pid 7280
2022-08-22 16:17:03.629   807   944 W ProcessCpuTracker: Failed to stat(/proc/7281): android.system.ErrnoException: stat failed: ENOENT (No such file or directory)
2022-08-22 16:17:03.629   807   944 W ProcessCpuTracker: Skipping unknown process pid 7281
2022-08-22 16:17:03.629   807   944 W ProcessCpuTracker: Failed to stat(/proc/7282): android.system.ErrnoException: stat failed: ENOENT (No such file or directory)
2022-08-22 16:17:03.630   807   944 W ProcessCpuTracker: Skipping unknown process pid 7282
2022-08-22 16:17:46.356   690  8433 E ResolverController: No valid NAT64 prefix (100, <unspecified>/0)
2022-08-22 16:17:46.360   690  8434 E ResolverController: No valid NAT64 prefix (100, <unspecified>/0)
2022-08-22 16:17:49.336   690  8439 E ResolverController: No valid NAT64 prefix (100, <unspecified>/0)
*/

type systemInfo struct {
	Level      string `json:"level"`
	CreateTime string `json:"create_time"`
	Detail     string `json:"detail"`
}

func strList() {
	str := `2022-08-22 16:17:03.629   807   944 W ProcessCpuTracker: Skipping unknown process pid 7279
	2022-08-22 16:17:03.629   807   944 W ProcessCpuTracker: Failed to stat(/proc/7280): android.system.ErrnoException: stat failed: ENOENT (No such file or directory)
	2022-08-22 16:17:03.629   807   944 W ProcessCpuTracker: Skipping unknown process pid 7280
	2022-08-22 16:17:03.629   807   944 W ProcessCpuTracker: Failed to stat(/proc/7281): android.system.ErrnoException: stat failed: ENOENT (No such file or directory)
	2022-08-22 16:17:03.629   807   944 W ProcessCpuTracker: Skipping unknown process pid 7281
	2022-08-22 16:17:03.629   807   944 W ProcessCpuTracker: Failed to stat(/proc/7282): android.system.ErrnoException: stat failed: ENOENT (No such file or directory)
	2022-08-22 16:17:03.630   807   944 W ProcessCpuTracker: Skipping unknown process pid 7282
	2022-08-22 16:17:46.356   690  8433 E ResolverController: No valid NAT64 prefix (100, <unspecified>/0)
	2022-08-22 16:17:46.360   690  8434 E ResolverController: No valid NAT64 prefix (100, <unspecified>/0)
	2022-08-22 16:17:49.336   690  8439 E ResolverController: No valid NAT64 prefix (100, <unspecified>/0)`
	strList := strings.Split(str, "\n")
	mps := make([]systemInfo, 20)
	if len(strList) >= 6 {
		for _, ele := range strList {
			rs := systemInfo{}
			wInx := strings.Index(ele, "W")
			eInx := strings.Index(ele, "E")
			fInx := strings.Index(ele, "F")
			if wInx > 0 {
				rs.Level = "waring"
				rs.Detail = ele[wInx+2:]
			} else if eInx > 0 {
				rs.Level = "error"
				rs.Detail = ele[eInx+2:]
			} else if fInx > 0 {
				rs.Level = "critical"
				rs.Detail = ele[fInx+2:]
			} else {
				rs.Level = "waring"
			}
			dataEle := strings.Fields(ele)
			time := dataEle[0] + dataEle[1]
			rs.CreateTime = time
			//fmt.Printf("%+v \n", rs)
			mps = append(mps, rs)
		}
		fmt.Printf("%+v \n", mps)
	}

}

func jsonStr() {
	str := `2022-08-26 11:36:46.155 32004 32004 I dist    : type=1400 audit(0.0:223605): avc: denied { append } for name=\"go_agent.log\" dev=\"dm-9\" ino=54341 scontext=u:r:start-ssh:s0 tcontext=u:object_r:system_data_file:s0 tclass=file permissive=1\n2022-08-26 11:36:46.155 32004 32004 I dist    : type=1400 audit(0.0:223606): avc: denied { create } for scontext=u:r:start-ssh:s0 tcontext=u:r:start-ssh:s0 tclass=tcp_socket permissive=1\n2022-08-26 11:36:46.159 32004 32004 I dist    : type=1400 audit(0.0:223607): avc: denied { bind } for scontext=u:r:start-ssh:s0 tcontext=u:r:start-ssh:s0 tclass=tcp_socket permissive=1\n2022-08-26 11:36:46.159 32004 32004 I dist    : type=1400 audit(0.0:223608): avc: denied { node_bind } for saddr=::1 scontext=u:r:start-ssh:s0 tcontext=u:object_r:node:s0 tclass=tcp_socket permissive=1\n2022-08-26 11:36:46.159 32004 32004 I dist    : type=1400 audit(0.0:223609): avc: denied { read } for name=\"somaxconn\" dev=\"proc\" ino=65941 scontext=u:r:start-ssh:s0 tcontext=u:object_r:proc_net:s0 tclass=file permissive=1\n2022-08-26 11:38:02.755 32137 32137 I dist    : type=1400 audit(0.0:223652): avc: denied { connect } for scontext=u:r:start-ssh:s0 tcontext=u:r:start-ssh:s0 tclass=tcp_socket permissive=1\n2022-08-26 11:38:02.755 32137 32137 I dist    : type=1400 audit(0.0:223653): avc: denied { name_connect } for dest=18080 scontext=u:r:start-ssh:s0 tcontext=u:object_r:port:s0 tclass=tcp_socket permissive=1\n2022-08-26 11:38:55.220   690 32210 E ResolverController: No valid NAT64 prefix (100, <unspecified>/0)\n2022-08-26 11:38:55.224   690 32212 E ResolverController: No valid NAT64 prefix (100, <unspecified>/0)\n2022-08-26 11:38:58.172   690 32213 E ResolverController: No valid NAT64 prefix (100, <unspecified>/0)\n`

	s := strings.Split(str, `\n`)
	for i, ele := range s {
		fmt.Printf("i:%d ,%+v  \n", i, ele)

	}

}

func strCut() {
	str := "aabbcc=ddee=ffgg"
	// 分割
	cut, after, found := strings.Cut(str, "=")
	if found {
		fmt.Println(cut, after, found) // 头,尾,真/假
	}
}

func main() {
	strCut()
	//strList()
	//strFindTrim()
	//replacerStr()
	//trimStr()
	//toStr()
	//mapStr()
	//repeatStr()
	//joinStr()
	//compareStr()
	//forStr()
	//containsStr()
	//countStr()
	//fieldsStr()
	//sliceStr()
	//hasStr()
	//indexStr()

	//float64Pow()

	//jsonStr()

}
