package main

import (
	"fmt"
	"time"
)

func main() {
	go printOneByOne()
	fmt.Println("先打卬我")
	time.Sleep(time.Second * 11)
}

func printOneByOne() {
	for i := 0; i < 10; i++ {
		// sleep 阻塞线程
		time.Sleep(time.Second * 1)
		fmt.Println(i)
	}
}
