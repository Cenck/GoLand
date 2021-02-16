package main

import (
	"./support"
	"fmt"
)

/*
给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。



示例 1：

输入：nums = [2,0,2,1,1,0]
输出：[0,0,1,1,2,2]
示例 2：

输入：nums = [2,0,1]
输出：[0,1,2]

*/

func main() {
	arr := []int{2, 0, 2, 1, 1, 0}
	support.MergeSort(arr)
	fmt.Println(arr)
}

/*
插入排序
*/
func sortColors(nums []int) {
	if nums == nil || len(nums) == 1 {
		return
	}

	for i := 0; i < len(nums); i++ {
		// 此时 i之前的元素已经是有序的了
		for j := i; j > 0; j-- {
			b := nums[j]
			a := nums[j-1]
			if a > b {
				nums[j] = a
				nums[j-1] = b
			} else {
				break
			}
		}
	}

}

/*
选择排序
和插入排序相反 选择是从未排序区间选择一个最小值放在已排区间末尾
因为选择排序 必定要遍历整个未排区间，而插入排序在理想情况下，只需要遍历已排区间的1个元素，因此选择排序常常被插入排序替代
*/
func selectSort(nums []int) {
	if nums == nil || len(nums) == 1 {
		return
	}

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		// 从未排序区间 选择一个最小值与num交换位置
		min := num
		minIndex := i
		for j := i; j < len(nums); j++ {
			a := nums[j]
			if a < min {
				minIndex = j
				min = a
			}
		}
		nums[minIndex] = num
		nums[i] = min
	}

}

/*

 */
func fastSort(nums []int) {

}
