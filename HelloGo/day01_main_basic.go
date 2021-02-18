package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func main() {
	// 打印hello go
	var str = "hello,go! nice to meet you at first time"
	str = strings.ToUpper(str)
	var buf bytes.Buffer
	buf.WriteString(str)
	buf.WriteString("\n 这是java的stringbuffer")
	fmt.Print(buf.String())
	// char
	chara, charb := 'a', 98
	fmt.Println(string(chara), string(charb))
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
	sum := 0
	for i := 0; i < 10; i++ {
		sum++
	}
	fmt.Println(sum)
	fmt.Print("hello go,get it ")

}

/*
leetcode 1470. 重新排列数组
给你一个数组 nums ，数组中有 2n 个元素，按 [x1,x2,...,xn,y1,y2,...,yn] 的格式排列。
请你将数组按 [x1,y1,x2,y2,...,xn,yn] 格式重新排列，返回重排后的数组。

示例 1：

输入：nums = [2,5,1,3,4,7], n = 3
输出：[2,3,5,4,1,7]
解释：由于 x1=2, x2=5, x3=1, y1=3, y2=4, y3=7 ，所以答案为 [2,3,5,4,1,7]

*/
func TestForRange(t *testing.T) {
	t.Log("进入单测")
	reRange := arrayReRange()
	fmt.Println(len(reRange))
	fmt.Println(reRange)
	t.Log("单测结束")
}

func TestForHello(t *testing.T) {
	t.Log(9 / 2)
}

func shuffle(nums []int, n int) []int {
	// 取mid index
	var mid_indx = n
	ret := make([]int, 0, 10)
	for i := 0; i < mid_indx; i++ {
		j := i + mid_indx
		ret = append(ret, nums[i], nums[j])
	}
	return ret
}

func arrayReRange() [6]int {
	var arr = [6]int{2, 5, 1, 3, 4, 7}
	const length = len(arr)
	// 取mid index
	var mid_indx = length / 2
	var ret [6]int
	for i := 0; i < mid_indx; i++ {
		j := i + mid_indx
		ret[2*i] = arr[i]
		ret[2*i+1] = arr[j]
	}
	return ret
}
