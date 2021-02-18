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
	nums := []int{2, 1}
	//[1 2 2 3 3 4 5 5 6]
	largest := findKthLargest_ByFast(nums, 2)
	fmt.Println(largest)
	copy := []int{2, 1}
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
	// 想要正序第几个元素
	if nums == nil || len(nums) == 0 || len(nums) < k {
		return k
	}

	// kk是k的正序索引值
	kk := len(nums) - k
	return findKthLargest_ByFast_sub(nums, 0, len(nums)-1, kk)
}

func findKthLargest_ByFast_sub(nums []int, start int, end int, kk int) int {
	if end-start < 0 {
		return -1
	}
	pivotIndex := start
	l, r := start, end
	for l != r {
		// 如果lv大于rv
		rv := nums[r]
		lv := nums[l]
		/*
			如果没有发生交换 则继续走原指针
			如果发生交换则要指针变换 左变右 或右变左
		*/
		// pivot == r 表示刚才遍历的是左指针 (不可能遍历pivot一侧)
		isLeft := pivotIndex == r
		if rv < lv {
			nums[l] = rv
			nums[r] = lv
			if isLeft {
				r--
			} else {
				l++
			}
		} else {
			if isLeft {
				l++
			} else {
				r--
			}
		}

	}

	// 快排扫描一次后，pivotIndex==k说明本次扫描恰好找到了index=kk的值
	if kk == pivotIndex {
		return nums[pivotIndex]
	} else if kk < pivotIndex {
		// 如果kk比本次扫描到的pivotIndex小，则仅扫描pivot左边的数组,反之仅扫描右边的数组
		return findKthLargest_ByFast_sub(nums, start, pivotIndex-1, kk)
	} else if kk > pivotIndex {
		return findKthLargest_ByFast_sub(nums, pivotIndex+1, end, kk)
	}
	return -2
}
