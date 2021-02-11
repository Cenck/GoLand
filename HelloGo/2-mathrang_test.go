package main

import (
	"fmt"
	"testing"
)

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


func shuffle(nums []int, n int) []int {
	// 取mid index
	var mid_indx = n
	ret := make([]int,0,10)
	for i := 0; i < mid_indx; i++ {
		j := i + mid_indx
		ret = append(ret,nums[i],nums[j] )
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
