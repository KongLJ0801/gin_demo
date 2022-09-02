package exec

import (
	"errors"
	"fmt"
	"time"
)

/**
 *
 * @date 2022/7/19 11:28
 * @version 0.1
 * @author KongLingJie
 *
 */

type PathError struct {
	path       string
	op         string
	createTime string
	message    string
}

func NewPathError(path, op, message string) PathError {
	return PathError{
		path:       path,
		op:         op,
		createTime: time.Now().Format("2006-01-02"),
		message:    message,
	}
}

func (e PathError) Error() string {
	return e.createTime + "-" + e.path + "-" + e.op + "-" + e.message
}

func deletePath(path string) error {
	return NewPathError(path, "delete", "delete not exec")
}

func soo() {
	defer func() {
		if err := recover(); err != nil { // 捕获错误信息
			fmt.Println("发生了 panic")
		}
	}()

	panic(errors.New("errors"))
}

func ExecFunc() {
	soo()
	fmt.Println("33333")
	deletePath("/path/error")
}
