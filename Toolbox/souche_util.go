package main

import (
	"./file"
	"./strutil"
	"fmt"
)

func main() {
	findPhoneNumber()
}

func findPhoneNumber() {
	dataPath := "./Toolbox/souche_util.txt"
	str := file.ReadStrFromFile(dataPath)
	phones := strutil.ExtractPhoneNumFromStr(str)
	for _, phone := range phones {
		fmt.Println(phone)
	}

}
