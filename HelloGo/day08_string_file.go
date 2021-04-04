package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"unsafe"
)

func main() {
	/*
		【string转byte数组 转char数组】；
		https://www.cnblogs.com/shuiyuejiangnan/p/9707066.html
	*/

	// string转byte数组 有复制
	s := "hello"
	var buf = []byte(s)
	fmt.Println("byte数组的104和h是否相等", 'h' == buf[0])
	for i := 0; i < len(buf); i++ {
		fmt.Println(buf[i])
	}

	// string转char数组

}

func readFile() {
	// 相对项目根目录的相对路径
	dataPath := "./HelloGo/day02.go"
	abs, _ := filepath.Abs(dataPath)
	file, e := ioutil.ReadFile(abs)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(string(file))
}

/*
高效 byte[]转string
*/
func bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s)) // 获取s的起始地址开始后的两个 uintptr 指针
	h := [3]uintptr{x[0], x[1], x[1]}      // 构造三个指针数组
	return *(*[]byte)(unsafe.Pointer(&h))
}
