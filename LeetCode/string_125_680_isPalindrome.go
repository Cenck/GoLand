package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
【对碰指针】
*/
func main() {
	s := `0P`
	palindrome := isPalindrome(s)
	fmt.Println(palindrome)
	s2 := "ebcbbececabbacecbbcbe"
	fmt.Println(validPalindrome(s2))

}

// 使用系统函数
func isPalindrome_1_usesys(s string) bool {
	if len(s) < 2 {
		// 只有一个字符 认为是回文
		return true
	}
	s = strings.ToUpper(s)
	l, r := 0, len(s)-1
	for l < r {
		u := s[l]
		v := s[r]
		for !isComparedChar(u) && l+1 < len(s) {
			l++
			u = s[l]
		}
		for !isComparedChar(v) && r-1 >= 0 {
			r--
			v = s[r]
		}
		if l >= r {
			break
		} else if u != v {
			return false
		} else {
			l++
			r--
		}
	}

	return true
}

func isComparedChar(c uint8) bool {
	return (c >= 48 && c <= 57) || (c >= 65 && c <= 90)
}

// [标准答案]验证字符串是否为回文符
func isPalindrome(s string) bool {
	if len(s) < 2 {
		return true
	}
	l, r := 0, len(s)-1
	s = strings.ToLower(s)
	for l < r {
		u := s[l]
		v := s[r]
		if !unicode.IsLetter(rune(u)) && !unicode.IsDigit(rune(u)) {
			// 即不是字母 也不是数字
			l++
			continue
		}
		if !unicode.IsLetter(rune(v)) && !unicode.IsDigit(rune(v)) {
			r--
			continue
		}
		if u != v {
			return false
		} else {
			l++
			r--
		}
	}
	return true
}

/*
给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。

示例 1:
输入: "aba"
输出: True
*/

func validPalindrome(s string) bool {
	return validPalindrome_sub(s, 0, len(s)-1, 0)
}

func validPalindrome_sub(s string, start, end int, delCount int) bool {
	length := end - start + 1
	if length < 2 {
		return true
	}
	l, r := start, end
	for l < r {
		if s[l] != s[r] {
			// 出现不相同是 删左边 或者删除右边 然后只要有一种情况是回文符就满足
			if delCount >= 1 {
				return false
			}
			return validPalindrome_sub(s, l, r-1, delCount+1) || validPalindrome_sub(s, l+1, r, delCount+1)
		} else {
			r--
			l++
		}
	}

	return true
}
