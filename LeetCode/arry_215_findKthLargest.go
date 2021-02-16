package main

import (
	"./support"
	"fmt"
)

/*
在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

示例 1:

输入: [3,2,1,5,6,4] 和 k = 2
输出: 5
示例 2:

输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
输出: 4

*/

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	//[1 2 2 3 3 4 5 5 6]
	largest := findKthLargest(nums, 4)
	fmt.Println(largest)
	copy := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	support.FastSort(copy)
	fmt.Println(copy)
}

/*
选择排序 ，可以先判断k离数组头部近 还是尾部近 然后决定正排还是倒排来减少循环次数
*/
func findKthLargest(nums []int, k int) int {
	l := len(nums)
	// 想要正序第几个元素
	kk := l - k
	// 正序排还是倒序排
	isOrderAsc := k > kk

	ret := nums[0]
	if isOrderAsc {
		min := -1
		for i := 0; i <= kk; i++ {
			min = nums[i]
			minI := i
			for j := i; j < len(nums); j++ {
				b := nums[j]
				if b < min {
					min = b
					minI = j
				}
			}
			nums[minI] = nums[i]
		}
		ret = min
	} else {
		max := -1
		for i := l - 1; i >= kk; i-- {
			max = nums[i]
			maxI := i
			for j := 0; j < i; j++ {
				b := nums[j]
				if b > max {
					max = b
					maxI = j
				}
			}
			nums[maxI] = nums[i]
		}
		ret = max
	}
	return ret
}

/**
使用快速排序
有两个优点：1. 当某次扫描时，找到的pivot的index恰好等于k,则可以停止排序工作返回pivot,
*/
func findKthLargest_ByFast(nums []int, k int) int {
	l := len(nums)
	// 想要正序第几个元素
	kk := l - k
	// 正序排还是倒序排
	isOrderAsc := k > kk

	ret := nums[0]
	if isOrderAsc {
		min := -1
		for i := 0; i <= kk; i++ {
			min = nums[i]
			minI := i
			for j := i; j < len(nums); j++ {
				b := nums[j]
				if b < min {
					min = b
					minI = j
				}
			}
			nums[minI] = nums[i]
		}
		ret = min
	} else {
		max := -1
		for i := l - 1; i >= kk; i-- {
			max = nums[i]
			maxI := i
			for j := 0; j < i; j++ {
				b := nums[j]
				if b > max {
					max = b
					maxI = j
				}
			}
			nums[maxI] = nums[i]
		}
		ret = max
	}
	return ret
}
