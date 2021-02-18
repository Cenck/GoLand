package main

import "fmt"

/*
给定两个排序后的数组 A 和 B，其中 A 的末端有足够的缓冲空间容纳 B。 编写一个方法，将 B 合并入 A 并排序。

初始化 A 和 B 的元素数量分别为 m 和 n。

示例:

输入:
A = [1,2,3,0,0,0], m = 3
B = [2,5,6],       n = 3

输出: [1,2,2,3,5,6]

*/
func main() {

	a1 := []int{4, 5, 6, 0, 0, 0}
	a2 := []int{1, 2, 3}
	mergeSortedArray(a1, 3, a2, 3)
	fmt.Println(a1)
}

/*
临时数组非原地
时间复杂度低
空间复杂度极高
*/
func merge(A []int, m int, B []int, n int) {
	l, r := 0, 0
	temp := make([]int, m+n)
	k := 0
	for l < m || r < n {
		if l >= m {
			// 数组A提前遍历完毕
			temp[k] = B[r]
			k++
			r++
			continue
		}
		if r >= n {
			// 数组B提前遍历完毕
			temp[k] = A[l]
			k++
			l++
			continue
		}

		if B[r] < A[l] {
			temp[k] = B[r]
			k++
			r++
		} else {
			// A中小于等于都先取A
			temp[k] = A[l]
			k++
			l++
		}

	}
	for i := 0; i < len(temp); i++ {
		A[i] = temp[i]

	}

}

/*
数组搬移 原地算法
*/
func mergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
	if n < 1 {
		return
	}
	l, r := 0, 0
	for l < m+n {
		/*
			1. 右数组没有遍历完时，如果比左数组指针值小 则左数组其后元素整体搬移1位，右指针元素移到左指针位置
			2. 如果左数组遍历完 而右数组没完，则需要把右数组剩余元素贴在左数组后面
			3. 如果右数组先遍历完 不需要做其他特殊处理
		*/
		if r >= n {
			break
		}
		if nums1[l] > nums2[r] {
			for i := m + r; i > l; i-- {
				nums1[i] = nums1[i-1]
			}
			nums1[l] = nums2[r]
			l++
			r++
		} else if l >= m+r {
			// 左数组已遍历完
			nums1[l] = nums2[r]
			l++
			r++
		} else {
			l++
		}

	}

}

func merge_2(A []int, m int, B []int, n int) {
	if n <= 0 {
		return
	}
	l, r := 0, 0
	for l < m+n {
		if r < n && B[r] < A[l] {
			// 把A数组从l到m的所有元素 向右迁移1位，空出的元素赋值为B[r] ,当前数组尾部元素不是m是m+r
			for i := m + r; i > l; i-- {
				A[i] = A[i-1]
			}
			A[l] = B[r]
			l++
			r++
		} else if l-r >= m {
			// A已遍历完了
			A[l] = B[r]
			r++
			l++
		} else {
			l++
		}

	}

}
