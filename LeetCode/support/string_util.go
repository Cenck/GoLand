package support

const UP_CHAR_MINIS_LOWER = 'a' - 'A'

func Char2Upper(c uint8) uint8 {
	if c >= 97 && c <= 122 {
		// 是小写 需要转大写
		c -= UP_CHAR_MINIS_LOWER
	}
	return c
}

func IsCharAlphaOrNum(c uint8) bool {
	var isUpper = c >= 97 && c <= 122
	var isLower = c >= 65 && c <= 90
	var isNum = c >= 48 && c <= 57
	return isUpper || isLower || isNum
}
