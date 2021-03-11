package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type error interface {
	Error() string
}

type commonError struct {
	errorCode int
	errorMsg  string
}

func (ce *commonError) Error() string {
	return ce.errorMsg
}

type MyError struct {
	err error
	mgs string
}

func (e *MyError) Error() string {
	return e.err.Error() + e.mgs
}

func add(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, &commonError{
			errorCode: 1,
			errorMsg: "a或者b不能为负数",
		}
	} else {
		return a + b, nil
	}
}

func ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	stats, statsErr := f.Stat()
	if statsErr != nil {
		return nil, statsErr
	}
	bytes := make([]byte, stats.Size())

	buffer := bufio.NewReader(f)
	_, err = buffer.Read(bytes)
	return bytes, err
}

func connectMySQL(ip, username, password string)  {
	if ip == "" {
		panic("ip不能为空")
	}
}

func main() {
	i, err := strconv.Atoi("a")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}

	sum, err1 := add(-1, 2)
	if cm, ok := err1.(*commonError); ok {
		fmt.Println("错误代码为:", cm.errorCode, ", 错误信息为:", cm.errorMsg)
	} else {
		fmt.Println(sum)
	}

	newErr := MyError{err, "数据上传问题"}

	e := errors.New("原始错误e")
	w := fmt.Errorf("Wrap了一个错误:%w", e)

	fmt.Println(newErr)
	fmt.Println(w)

	fmt.Println(errors.Unwrap(w))

	fmt.Println(errors.Is(w, e))

	var cm *commonError
	if errors.As(err, &cm) {
		fmt.Println("错误代码为:", cm.errorCode, ", 错误信息为:", cm.errorMsg)
	} else {
		fmt.Println(sum)
	}

	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
	}()
	connectMySQL("", "root", "123456")
}