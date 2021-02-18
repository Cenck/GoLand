package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 具名函数
	sum := plus(3, 4)
	fmt.Println(sum)
	// 函数式编程 ：
	//将函数作为值传递
	isOdd := func(a int) bool { return a%2 == 1 }
	if isOdd(3) {
		fmt.Println("3是一个奇数")
	}
	// 将函数作为另一个函数的参数
	doCalculate(3, 4, plus)
	// defer : 让你的函数在返回前 执行另一个函数
	tryDefer()
	fmt.Println("我应该出现在tryDefer内部所有打印的后面")

}

func tryDefer() {
	/*
		通过控制台 可以猜想
		1. defer语句是在 母函数执行完毕 （但返回前）才开始执行的
		2. defer语句的顺序是倒序执行的，也就是说所有在母方法中声明的defer函数被放在一个栈中，后来居上，然后在母方法执行完毕后从栈顶挨个执行
	*/
	defer fmt.Println("我是1")
	defer fmt.Println("我是2")
	defer fmt.Println("我是3")
	fmt.Println("结束了")
}

/*
具名函数：在函数返回前就准备好并赋值了返回结果
*/
func plus(a, b int) (sum int) {
	// todosometing
	// 两值相加 已得到函数需要的返回结果，但整个函数的业务逻辑还没有处理完
	sum = a + b

	// todoOtherThings
	fmt.Println(sum)
	return
}

/*
传入a,b两个数，执行f函数操作
*/
func doCalculate(a, b int, f func(a, b int) int) {
	ret := f(a, b)
	fmt.Println("计算的结果为：" + strconv.Itoa(ret))
}
