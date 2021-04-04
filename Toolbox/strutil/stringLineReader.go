package strutil

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
)

func StrLineRead2SqlWhere() {
	str := strLineRead2JsonArray()
	fmt.Print("(" + str + ")")
}

func StrLineRead2JsonArray() {
	str := strLineRead2JsonArray()
	fmt.Print("[" + str + "]")
}

func strLineRead2JsonArray() (ret string) {
	dataPath := "./Toolbox/strutil/stringline.txt"
	abs, _ := filepath.Abs(dataPath)
	file, e := ioutil.ReadFile(abs)
	if e != nil {
		fmt.Printf("\n %c[%d;%d;%dm%s%c[0m\n\n", 0x1B, 1, 0, 31, e, 0x1B)
	}
	str := string(file)
	split := strings.Split(str, "\n")
	var sb strings.Builder
	for i, s := range split {
		sb.WriteString("\"")
		sb.WriteString(strings.TrimSpace(s))
		sb.WriteString("\"")
		if i < len(split)-1 {
			sb.WriteString(",")
		}
	}
	ret = sb.String()
	return
}

// 手机号正则匹配
const PHONE_NUM_REG = `^1[3|4|5|6|7|8|9][0-9]\\d{8}$`

func ExtractPhoneNumFromStr(text string) []string {
	Regexp := regexp.MustCompile(`([\d]{11})`)
	return Regexp.FindStringSubmatch(text)
}
