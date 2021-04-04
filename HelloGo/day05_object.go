package main

import (
	"fmt"
	"strconv"
)

/*
go语言的自定义对象叫 结构体
和Java不同：结构体中只能声明变量 不能声明结构体方法，方法要声明与结构体关联的参数

go的接口只能有方法集 不可以存在任何变量 包括常量

*/
func main() {
	m2 := new(Movie)
	m2.Name = "唐探3"
	m2.Price = 45.5
	fmt.Println(m2.Price)

	m := Movie{"你好李焕英", 38.9, m2}
	fmt.Println(m.Name)

	// go语言变量默认是值传递和值引用
	copy_m := m
	fmt.Println("结构体实例的赋值是值的引用,两个实例相等但内存地址不同", copy_m == m)
	fmt.Printf("m的内存地址是：%p\n", &m)
	fmt.Printf("m的副本的内存地址：%p\n", &copy_m)
	copy_m.Name = "hello,Mom"

	// 方法： 依赖于结构体的函数 see这个方法只能被Movie结构体的实例调用
	m.CannotChangeName("你好，李焕英")
	m.see()
	if isMovieOk(&m) {
		a := m.Comments()[0]
		fmt.Println(a)
	}

	// 类的继承
	bm := EmotionMovie{m, true}
	fmt.Println("情感电影：", bm.Name)

	// switch 枚举
	var a = 5
	switch a {
	case DEV:
		fmt.Println("这是开发环境")
		break
	case STABLE:
		break
	}
}

// 情感电影是电影的一种 继承了Movie
type EmotionMovie struct {
	Movie
	// 看电影是否需要哭
	NeedCry bool
}

type Movie struct {
	Name  string
	Price float64
	// 结构体内的 类型引用 只能是指针引用吗？
	Recom *Movie
}

func (m *Movie) see() {
	fmt.Println("我刚看了" + m.Name + "这部电影, 评分还不错，有" + strconv.FormatFloat(m.Score(), 'f', 2, 64) + "分")
}

func (m Movie) CannotChangeName(name string) {
	// 因为是值传递的缘故 无法设置名称
	m.Name = name
}

func (m *Movie) Comments() []string {
	return []string{"这是好片", "什么玩意", "水军真牛批"}
}

func (m *Movie) Score() float64 {
	return 7
}

type MovieComment interface {
	// 评论
	Comments() []string
	// 评分
	Score() float64
}

/*
接口主要是用于泛型 代码 复用
*/
func isMovieOk(m MovieComment) bool {
	score := m.Score()
	return score >= 7
}

// 枚举
const (
	DEV = iota
	STABLE
	PREPUB
	PRO
)
