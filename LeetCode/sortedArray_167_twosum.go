package main

import "fmt"

func main() {
	arr := []int{2, 7, 11, 15}
	sum := twoSum2(arr, 18)
	fmt.Println(sum)
}

/*
[两数之和II]
给定一个已按照 升序排列  的整数数组 numbers ，请你从数组中找出两个数满足相加之和等于目标数 target 。

函数应该以长度为 2 的整数数组的形式返回这两个数的下标值。numbers 的下标 从 1 开始计数 ，所以答案数组应当满足 1 <= answer[0] < answer[1] <= numbers.length 。

你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。


示例 1：

输入：numbers = [2,7,11,15], target = 9
输出：[1,2]
解释：2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。

*/
// 对撞指针
func twoSum2(numbers []int, target int) []int {
	ret := make([]int, 2)
	if len(numbers) < 2 {
		return ret
	}
	l, r := 0, len(numbers)-1
	for l != r {
		lv := numbers[l]
		rv := numbers[r]
		if lv+rv < target {
			l++
		} else if lv+rv > target {
			r--
		} else {
			ret[0] = l + 1
			ret[1] = r + 1
			return ret
		}
	}
	return ret
}

// 暴力双重循环
func twoSum(numbers []int, target int) []int {
	ret := make([]int, 2)
	if len(numbers) < 2 {
		return ret
	}
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			pv := numbers[j] + numbers[i]
			if pv > target {
				break
			} else if pv == target {
				ret[0] = i + 1
				ret[1] = j + 1
				return ret
			}
		}

	}
	return ret
}
