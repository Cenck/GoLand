package main

import "fmt"

/*
[对碰指针]
翻转元音字符
*/
func main() {
	// string转byte[]
	s := "aA"
	fmt.Println(reverseVowels(s))

}

/*
编写一个函数，以字符串作为输入，反转该字符串中的元音字母。
发现元音字母时停止 与另一边的元音字母比较

示例 1：

输入："hello"
输出："holle"
*/
func reverseVowels(s string) string {
	if len(s) < 2 {
		return s
	}
	l, r := 0, len(s)-1
	var buf = []byte(s)
	for l < r {
		if !isVowel(buf[l]) {
			l++
			continue
		}
		if !isVowel(buf[r]) {
			r--
			continue
		}
		lv := buf[l]
		buf[l] = buf[r]
		buf[r] = lv
		l++
		r--
	}
	return string(buf)
}

func isVowel(c uint8) bool {
	return 'a' == c || 'e' == c || 'i' == c || 'o' == c || 'u' == c || 'A' == c || 'E' == c || 'I' == c || 'O' == c || 'U' == c
}
