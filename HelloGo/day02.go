package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// 检查变更的类型
	var a = "true"
	b := 1
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(a))
	// 类型转换 strconv
	parseBool, _ := strconv.ParseBool(a)
	fmt.Println(parseBool)
	fmt.Println(strconv.FormatBool(true))
	// 获取变量的内存地址
	fmt.Println(&a)

}
