package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 打印hello go
	fmt.Print("hello,go! nice to meet you for first time")
	// 基本变量有默认值
	var a int
	fmt.Println(a)
	// 类型推断
	var num = 5
	num2 := 6 //var num2=6 的简写形式
	fmt.Println(num + num2)
	// 多变量声明
	var b, c int16
	var (
		d bool
		e string
	)
	// 常量/多变量 命名+赋值
	const f, g, h = 11.2, true, "hehe"
	// string转化
	fmt.Println(strconv.Itoa(int(b+c)) + strconv.FormatFloat(f, 'f', 2, 64) + strconv.FormatBool(d) + e + h)
	fmt.Println(g)
	// go的加减乘除以及大小比较等运算符与java等其他语言的一致
	// 分支
	if !g {
		fmt.Println("不会走到这里的")
	} else if f > 100 {
		fmt.Println("f小于100 不打印")
	} else {
		fmt.Println("这是if分支测试")
	}

	switch num {
	case 4:
		//
		break
	case 5:
		fmt.Println("switch case")
		break
	default:

	}
	//循环
	sum := 0;
	for i:=0 ;i<10;i++ {
		sum++;
	}
	fmt.Println(sum)
	fmt.Print("hello go,get it ")

}
