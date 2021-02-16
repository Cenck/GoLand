package main

import "fmt"

func main() {
	var arr = []int{0, 1, 0, 3, 12}
	moveZeroes2(arr)
	fmt.Println(arr)
}

/**
方法一：双指针  遍历数组的同时，记录并更新数组中第一个0元素的下标j，其后的遍历，每遇到一个j后面不为0的元素都和j交换位置 并将J+1
时间复杂度n
*/
func moveZeroes(nums []int) {
	// j始终等于第一个0的下标
	if len(nums) == 0 || len(nums) == 1 {
		return
	}
	j := -1
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if num != 0 && j != -1 {
			// num不等于0 则第一个I走过的0的索引下标与num交换位置  每次交换完位置 j+1
			nums[i] = 0
			nums[j] = num
			j++
		} else if num == 0 && j == -1 {
			// 如果是第一个0位置 则j=该0的位置
			j = i
		}
	}

}

/*
方法二：将所有非0元素提到数组前面 ，然后将其余位置 以0填充
时间复杂度n
*/
func moveZeroes2(nums []int) {
	// j始终等于第一个0的下标
	if len(nums) == 0 || len(nums) == 1 {
		return
	}
	j := 0
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if num != 0 {
			nums[j] = num
			j++
		}
	}

	for i := j; i < len(nums); i++ {
		nums[i] = 0
	}
}
