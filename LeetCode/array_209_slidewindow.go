package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 3, 1, 2, 4, 3}
	arrayLen := minSubArrayLen(7, arr)
	fmt.Println(arrayLen)

	test_longestOnes()
}

/*
[长度最小的子数组]
给定一个含有 n 个正整数的数组和一个正整数 target 。

找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。



示例 1：

输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。
*/
/*
滑动窗口 l和r永远不会倒回去遍历 也就是说时间复杂度一般小于2n
*/
func minSubArrayLen(target int, nums []int) int {
	minLen := 0
	// 用于保存窗口总值 与target比较
	l, r := 0, 0
	sum := nums[r]
	for l < len(nums) {
		if sum < target {
			if r >= len(nums)-1 {
				// 从l到数组尾部这个子串的所有子串不会再比target大了
				break
			}
			r++
			sum += nums[r]
		}
		if sum >= target {
			// l向右收缩1
			if minLen == 0 || r-l+1 < minLen {
				minLen = r - l + 1
			}
			sum -= nums[l]
			l++
		}
	}
	return minLen
}

// 暴力循环
func minSubArrayLen_byForce(target int, nums []int) int {
	minLen := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			if sum < target {
				sum += nums[j]
			}
			if sum >= target {
				if minLen == 0 || j-i < minLen {
					minLen = j - i + 1
				}
				break
			}
		}
	}
	return minLen
}

/*
offer-59-I

滑动窗口的最大值
给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。

示例:

输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7]

*/
func test_maxSlidingWindow() {
	fmt.Println("test_maxSlidingWindow")
	arr := []int{3, 1, -1, -3, 5, 3, 6, 7}
	sl := maxSlidingWindow(arr, 3)
	fmt.Println(sl)
	arr2 := []int{}
	fmt.Println(maxSlidingWindow(arr2, 0))
	arr3 := []int{1, -1}
	fmt.Println(maxSlidingWindow(arr3, 1))

}
func maxSlidingWindow(nums []int, k int) []int {
	length := len(nums) - k + 1
	if len(nums) == 0 || k == 0 {
		return []int{}
	}
	maxArr := make([]int, length)
	// maxa指的是0->k的最大值，maxb指的是1->k的值，maxb用于下一次
	max, maxindex := nums[0], 0
	for i := 1; i < k; i++ {
		if nums[i] > max {
			max = nums[i]
			maxindex = i
		}
	}
	maxArr[0] = max
	for i := 1; i < length; i++ {
		// 如果 maxindex!= i-1
		if maxindex == i-1 {
			if nums[i+k-1] > max {
				max = nums[i+k-1]
				maxindex = i + k - 1
			} else {
				// 需要遍历
				max = nums[i]
				maxindex = i
				for j := i; j < i+k; j++ {
					if nums[j] > max {
						max = nums[j]
						maxindex = j
					}
				}
			}
		} else {
			if nums[i+k-1] > max {
				max = nums[i+k-1]
				maxindex = i + k - 1
			}
		}
		maxArr[i] = max
	}
	return maxArr
}

/*
leet-1004
给定一个由若干 0 和 1 组成的数组 A，我们最多可以将 K 个值从 0 变成 1 。

返回仅包含 1 的最长（连续）子数组的长度。

示例 1：

输入：A = [1,1,1,0,0,0,1,1,1,1,0], K = 2
输出：6
解释：
[1,1,1,0,0,1,1,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 6。
示例 2：

输入：A = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], K = 3
输出：10
解释：
[0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 10。
*/
func test_longestOnes() {
	fmt.Println("test_longestOnes")
	arr := []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}
	ones := longestOnes(arr, 2)
	fmt.Println(arr, ones)
	arr2 := []int{1, 1}
	fmt.Println(arr2, longestOnes(arr2, 1))

	arr3 := []int{0, 0, 1, 1, 1, 0, 1, 1, 1, 1, 0}
	fmt.Println(arr3, longestOnes(arr3, 0))

}

func longestOnes(A []int, K int) int {
	// 窗口形成规则：窗口内元素为0的个数始终为k,因此当窗口向右滑动时，必须调整窗口使其满足这一规则
	l, r := 0, 0
	maxlen := 0
	for true {
		l, r = longestOnes_makeWindows(A, l, r, K)
		thisLen := r - l + 1
		if thisLen <= 0 {
			break
		} else if thisLen > maxlen {
			maxlen = thisLen
		}
	}
	return maxlen
}

/**

返回新窗口的左右下标，如果r-l<=0 说明无法生成新窗口，窗口滑动结束
*/
func longestOnes_makeWindows(A []int, start, end int, k int) (int, int) {
	kcount := 0
	l, r := 0, 0

	if k == 0 {
		// 从start下标开始找到最近连续为1的子数组作为窗口
		i := 0
		if start != 0 {
			i = end + 1
		}
		l, r = i, i
		has := false
		for ; i < len(A); i++ {
			if A[i] == 1 {
				if !has {
					r = i
					l = i
					has = true
				} else {
					r++
				}
			}
			if A[i] == 0 && has {
				break
			}
		}
		if !has {
			r = -1
		}
		return l, r
	}

	if start == 0 {
		l = 0
		kcount = 0
	} else {
		// 找窗口最左第一个0元素的下一位
		for i := start; i <= end; i++ {
			if A[i] == 0 {
				l = i + 1
				break
			}
		}
		kcount = 1
	}
	hasK := false
	for i := end + 1; i < len(A); i++ {
		if A[i] == 0 {
			kcount++
		}
		if kcount == k+1 {
			r = i - 1
			hasK = true
			break
		}
	}
	if !hasK {
		// 凑不够k个 新窗口无法组成
		r = -1
	}
	return 0, r
}
