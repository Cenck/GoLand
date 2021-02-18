package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {
	// 感知ioutil的系统异常
	readStr, err := ioReadError("")
	if err == nil {
		fmt.Println(readStr)
	} else {
		fmt.Println(err.Error())
	}

	for i := 0; i < 100000000; i++ {
		fmt.Println(i)
		if i == 1000 {
			// 在生产环境的服务端 使用panic的可能性并不高，仅当程序卡死无法执行时才可能使用panic来中止
			panic("强制关闭程序")

		}
	}

}

type illageArgumentError struct {
	Msg string
}

func (e illageArgumentError) Error() string {
	fmt.Println(e.Msg)
	return e.Msg
}
func ioReadError(path string) (str string, e error) {
	if len(path) == 0 {
		return "", illageArgumentError{"文件路径不能为空"}
	}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		e = err
		fmt.Println(err.Error())
	} else {
		var buf bytes.Buffer
		buf.Write(file)
		str = buf.String()
		fmt.Println(str)
	}
	return
}
