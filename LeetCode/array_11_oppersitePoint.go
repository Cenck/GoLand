package main

/*
对撞指针
*/
import "fmt"

func main() {
	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	area := maxArea(arr)
	fmt.Println(area)
}

/*
leetcode 11
理解：就是 两两元素一组，取 索引差值*两元素中最小值 最大的那一对
对撞指针 时间复杂度为O(n)
*/
func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	l, r := 0, len(height)-1
	var max int
	for l < r {
		/*
			双指针遍历 谁小前进谁，
			对于小的那个，后面的遍历不管谁和你配对 都不可能产生面积更大值了
		*/
		var min int
		x, y := l, r
		if height[l] <= height[r] {
			min = height[l]
			l++
		} else {
			min = height[r]
			r--
		}
		sum := (y - x) * min
		if sum > max {
			max = sum
		}

	}
	return max
}
