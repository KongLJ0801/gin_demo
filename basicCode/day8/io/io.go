package io

import (
	"bufio"
	"compress/zlib"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

/**
 *
 * @date 2022/7/22 17:11
 * @version 0.1
 * @author KongLingJie
 *
 */

func std() {
	// 接收两个值 空格分开
	//fmt.Println("please input Scan string")
	//var s string
	//fmt.Scan(&s)
	//var ss string
	//fmt.Scan(&ss)
	//fmt.Printf("you input %s,%s\n", s, ss)

	fmt.Println("please input Scanf int")
	var i int
	fmt.Scanf("%d", &i)
	var i2 int
	fmt.Scanf("%d", &i2)
	fmt.Printf("you input %d,%d\n", i, i2)

}

// 读文件
func read_file() {
	exPath, _ := os.Getwd() // 获取文件完整路径
	fmt.Println("expath", exPath)
	if f, err := os.Open(exPath + "/day8/io/io.go"); err != nil {
		fmt.Println(err)
	} else {
		defer f.Close()
		buffer := strings.Builder{}
		for {
			bs := make([]byte, 10) //1024
			if n, err := f.Read(bs); err != nil {
				fmt.Println(err)
				if err == io.EOF {
					break
				}
			} else {
				fmt.Printf("从文件中读取了 %d \n", n)
				buffer.WriteString(string(bs))
			}
		}
		fmt.Println(buffer.String())

	}
}

// 读文件 工作常用
func read_file2() {
	exPath, _ := os.Getwd() // 获取文件完整路径
	if f, err := os.Open(exPath + "/day8/io/io.go"); err != nil {
		fmt.Println(err)
	} else {
		defer f.Close()
		reader := bufio.NewReader(f)
		i := 0
		for {
			if n, err := reader.ReadString('\n'); err != nil {
				fmt.Println(err, "eof")
				if err == io.EOF { //EOF end of file 文件末尾
					break
				}
			} else {
				fmt.Printf("从文件中读取了 %d 行\n", i)
				//fmt.Print(n)
				fmt.Println(strings.Trim(n, "\n"))
				i++
			}
		}
	}
}

// 写文件
func write_file() {
	// os.O_CREATE 文件存在就直接写, 不存在就创建
	// os.O_TRUNC  清空之前内容,在写入新的内存
	// os.O_WRONLY 只写不读
	// os.O_APPEND 不覆盖 直接追加内容
	if wf, err := os.OpenFile("hole.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666); err != nil {
		fmt.Println(err)
	} else {
		defer wf.Close()
		wf.Write([]byte("hello \n"))
	}
}

// 写文件
func write_file2() {
	// os.O_CREATE 文件存在就直接写, 不存在就创建
	// os.O_TRUNC  清空之前内容,在写入新的内存
	// os.O_WRONLY 只写不读
	// os.O_APPEND 不覆盖 直接追加内容
	if wf, err := os.OpenFile("hole.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666); err != nil {
		fmt.Println(err)
	} else {
		defer wf.Close()
		// bufio 是把内容写在缓冲区
		writer := bufio.NewWriter(wf)
		writer.WriteString("rose")
		writer.WriteString("\n")
		writer.Flush() // 一次IO操作写入文件
	}
}

// 文件操作,文件信息
func file_op() {
	//os.Remove("hole.txt")      // 删除文件
	//os.Mkdir("p1", 0666)       // 创建目录
	//os.MkdirAll("p1/p2", 0666) // 创建多层目录
	//os.Create("p1/p2/sql.txt") // 创建文件 指定目录下
	//os.RemoveAll("p1")         // 删除文件
	//os.Create("old.txt")
	os.Rename("old.txt", "new.txt")
	exPath, _ := os.Getwd() // 获取文件完整路径
	file, _ := os.Open(exPath + "/day8/math/math.go")
	fmt.Println(file.Name(), "file") // 文件全路径名
	info, _ := file.Stat()
	fmt.Println(info.Name(), "info") // 文件名
	fmt.Println(info.Size())         // 文件大小
	fmt.Println(info.IsDir())        // 是否是个目录
	fmt.Println(info.Mode())         // 权限
	fmt.Println(info.ModTime())      //修改时间

}

// 遍历目录和文件
func welk_dir(path string) error {
	if file, err := ioutil.ReadDir(path); err != nil {
		return err
	} else {
		for _, info := range file {
			fmt.Println(info.Name())
			if info.IsDir() {
				if err := welk_dir(filepath.Join(path, info.Name())); err != nil {
					return err
				}
			}

		}
	}
	return nil
}

// go 系统默认日志
func logger() {
	if file, err := os.OpenFile("hole.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666); err != nil {
		fmt.Println(err)
	} else {
		defer file.Close()
		logWriter := log.New(file, "[BIZ_NAME]", log.Ldate|log.Lmicroseconds)
		logWriter.Println("hello")
		logWriter.Println("hole")
	}
}

// shell
func sys_call() {
	if cmd_path, err := exec.LookPath("npm"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(cmd_path)
		cmd := exec.Command("npm", "-v")

		if output, err := cmd.Output(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Print(string(output))
		}
	}

	cmd := exec.Command("rm", "new.txt")
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		fmt.Println("123")
	} else {
		fmt.Println("456")
	}
}

// 读取文件->创建文件->压缩文件    解压文件
func compress() {
	//exPath, _ := os.Getwd()
	//fin, err := os.Open(exPath + "/day8/io/io.go")
	//if err != nil {
	//	fmt.Println(err)
	//	panic(err)
	//}
	//defer fin.Close()
	////创建目录
	////os.Mkdir("test", os.ModePerm)
	//
	//filOut, e := os.OpenFile("./day8/io/go.io.zlib", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	//if e != nil {
	//	fmt.Println(e)
	//	panic(err)
	//
	//}
	//defer filOut.Close()
	//writer := zlib.NewWriter(filOut)
	//for {
	//	bs := make([]byte, 1024)
	//	n, err := fin.Read(bs)
	//	if err != nil {
	//		if err == io.EOF {
	//			break
	//		} else {
	//			fmt.Println(e)
	//			panic(err)
	//		}
	//	} else {
	//		writer.Write(bs[:n])
	//	}
	//}
	//writer.Flush()

	//解压文件
	finr, err := os.Open("./day8/io/go.io.zlib")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	reader, e := zlib.NewReader(finr)
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
	io.Copy(os.Stdout, reader)
	defer finr.Close()
	defer reader.Close()

}

func IoFunc() {
	//std()
	//read_file()
	//read_file2()

	//write_file()
	//write_file2()

	//file_op()

	//exPath, _ := os.Getwd() // 获取文件完整路径
	//welk_dir(exPath + "")

	//logger()

	//sys_call()

	compress()
}
