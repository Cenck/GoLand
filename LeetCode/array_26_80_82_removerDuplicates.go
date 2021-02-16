package main

import (
	"fmt"
)

/*
给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。



示例 1:

给定数组 nums = [1,1,2],

函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。

你不需要考虑数组中超出新长度后面的元素。
示例  2:

给定 nums = [0,0,1,1,1,2,2,3,3,4],

函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。

你不需要考虑数组中超出新长度后面的元素。

*/
func main() {
	var arr = []int{1, 1, 1, 2, 2, 3}
	element := removeDuplicates2(arr)
	fmt.Println(element)
	fmt.Println(arr)

	TestForDelLn()
}

/*
leet 26
给定 nums = [0,0,1,1,1,2,2,3,3,4],

函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。

你不需要考虑数组中超出新长度后面的元素。

*/
func removeDuplicates(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}
	k := nums[0]
	j := 1
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num == k {
		} else {
			nums[j] = num
			j++
			k = num
		}
	}
	return j
}

/*
leet80
重复元素不得出现2次

输入：nums = [1,1,1,2,2,3]
输出：5, nums = [1,1,2,2,3]
解释：函数应返回新长度 length = 5, 并且原数组的前五个元素被修改为 1, 1, 2, 2, 3 。 你不需要考虑数组中超出新长度后面的元素。

*/
// 常规思想 记录重复次数
func removeDuplicates2(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}

	j := 1
	// 记录当前与上个元素不重复的元素值
	k := nums[0]
	// 记录 k 重复的次数
	n := 1
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num == k {
			n++
		} else {
			k = num
			n = 1
		}
		if n <= 2 {
			// 没有重复或者重复的次数不够
			nums[j] = num
			j++
		}
	}
	return j
}

// 更好的办法
func removeDuplicates2_2(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}
	l := 0
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		// 拿当前元素和当前元素下标-2的元素比较 如果相同 说明重复出现2次以上，否则不重复或重复小于二次 可以赋值到新数组
		if l < 2 || nums[l-2] != num {
			//
			nums[l] = num
			l++
		}
	}
	return l
}

/*
leet 82
给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。

示例 1:

输入: 1->2->3->3->4->4->5
输出: 1->2->5
示例 2:

输入: 1->1->1->2->3
输出: 2->3

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestForDelLn() {
	a, b, c, d, e, f, g := &ListNode{Val: 1}, &ListNode{Val: 2}, &ListNode{Val: 3}, &ListNode{Val: 3}, &ListNode{Val: 4}, &ListNode{Val: 4}, &ListNode{Val: 5}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	e.Next = f
	f.Next = g
	if g.Next == nil {
		fmt.Println("g的尾部为空")
	}
	duplicates := deleteDuplicates(&ListNode{1, &ListNode{1, nil}})
	fmt.Println(duplicates)
}

/**
遍历链表 拿当前元素和前后都比较 只有不相同才加入新链表
*/
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := (*ListNode)(nil)
	newHead := (*ListNode)(nil)
	last := (*ListNode)(nil)
	for i := head; i != nil; i = i.Next {
		val := i.Val
		if (i.Next == nil || i.Next.Val != val) && (last == nil || last.Val != val) {
			if cur == nil {
				cur = &ListNode{val, nil}
				newHead = cur
			} else {
				cur.Next = &ListNode{val, nil}
				cur = cur.Next
			}
		}
		last = i
	}
	return newHead
}
