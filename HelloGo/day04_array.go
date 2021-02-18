package main

import (
	"fmt"
)

func main() {
	// 数组的命名
	arr0 := [5]int{}
	fmt.Println(arr0)

	var arr2 [5]int
	fmt.Println(arr2)

	l := 2
	// 创建一个长度为L的切片数组
	var arr3 = make([]int, l)
	arr3[0] = 1
	arr3[1] = 2
	// append在必要时会为数组扩容
	ints := append(arr3, 3, 4, 5, 6, 7)
	fmt.Println(len(ints), ints)

	// 切片： 取arr3中前1个元素创建新数组
	split1 := ints[:2]
	fmt.Println("取2个元素创建新数组", split1)
	split2 := ints[4:]
	fmt.Println("取第四个元素以后的所有数据 创建新数组", split2)
	// 在数组i后面 添加split2数组中的所有元素，创建为新数组
	union1 := append(split1, split2...)
	fmt.Println("组合两个切片生成的新数组union1:", union1)

	cparr := make([]int, 3)
	copy(cparr, arr3)
	fmt.Println("将arr3中的元素复制到cparr", cparr)

	// map
	m := make(map[string]int)
	m["apple"] = 4
	m["banana"] = 10
	m["mile"] = 6
	delete(m, "mile")
	fmt.Println("ge语言的map结构", m)

}
