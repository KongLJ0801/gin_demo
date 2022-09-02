/*

   @author #Kkk
   @File code_bufio
   @Description:
   @version 0.1
   @date 2022/8/19 15:02

*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	//ch := make(chan []byte, 1024)
	//go readFromT(ch)
	//go writeToT()
	//go writeToT1(ch)
	//time.Sleep(3 * time.Minute)
	readTest()
}

// 用for高速写入到test.txt文件中
func writeToT() {
	filePath := "./test.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	_, _ = file.Seek(0, 2)
	write := bufio.NewWriter(file)
	start := time.Now().Minute()
	fmt.Println("开始写入T")
	i := 0
	for {
		stop := time.Now().Minute()
		if stop-start >= 2 {
			break
		}
		write.WriteString(fmt.Sprint(i) + "\n")
		write.Flush()
		i++
	}
	fmt.Println("结束写入T")
}

// 阻塞获取channel通道的值，写入到test1.txt文件中
func writeToT1(ch chan []byte) {

	filePath := "./test1.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()
	_, _ = file.Seek(0, 2)
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	for i := range ch {
		write.WriteString(string(i) + "\n")
		write.Flush()
	}
}

// 先执行writeToT（），睡眠一分钟后，每隔10ms读取test.txt文件的变化内容（尾部），并放入到channel中
func readFromT(rc chan []byte) {
	file, err := os.Open("./test.txt")
	if err != nil {
		return
	}
	defer file.Close()
	//从末位开始逐行读取
	_, _ = file.Seek(0, 2)
	reader := bufio.NewReader(file)
	time.Sleep(1 * time.Minute)

	fmt.Println("读取T开始")
	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			time.Sleep(10 * time.Millisecond)
			continue
		} else if err != nil {
			panic(fmt.Sprintf("ReadBytes error:%s", err.Error()))
		}
		rc <- line[:len(line)-1]
	}
}

// 在test1.txt中，用后一行的值减去前一行的值，看看有没有数据丢失
func readTest() {
	file, err := os.Open("./test.txt")
	if err != nil {
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	a := readOneValue(reader)
	fmt.Println("测试开始")
	for {
		b := readOneValue(reader)
		if b == -1 {
			break
		}
		if b == -2 {
			continue
		}
		if b-a != 1 {
			fmt.Printf("%+v ------> %+v\n", a, b)
		}
		a = b
	}
	fmt.Println("测试结束", a)
}
func readOneValue(reader *bufio.Reader) int {
	line, err := reader.ReadBytes('\n')
	if err == io.EOF {
		return -1
	} else if err != nil {
		panic(fmt.Sprintf("ReadBytes error:%s", err.Error()))
	}
	str := strings.TrimRight(string(line), "\n")
	if str == "" {
		return -2
	}
	v, _ := strconv.Atoi(str)
	return v
}
