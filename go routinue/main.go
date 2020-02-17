package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func hello(f chan int) {
	v := <-f
	fmt.Println("helloworld", v)
	defer wg.Done()
	return
}
func ying(h chan int) {
	defer wg.Done()
	var max int
	max = 0
	n := <-h
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Println(i)
			if i != 0 && i != n {
				max = i
			}
		}
	}
	if max != 1 {
		fmt.Println("最大公因数是", max)
	} else {
		fmt.Println("输入的数为质数")
	}
}
func main() {
	runtime.GOMAXPROCS(1)
	//使用无缓冲的通道进行传值，运行goroutine
	f := make(chan int)
	h := make(chan int)
	//使用有缓冲的通道进行传值
	g := make(chan int, 2)
	wg.Add(4)
	//x := 2365874
	//无缓冲通道
	go ying(h)
	go hello(f)
	//有缓冲通道,先使用第一个传入的值，顺序传值
	go ying(g)
	go ying(g)
	f <- 2365874
	h <- 2
	g <- 6
	g <- 98
	wg.Wait()
	fmt.Println("hello")
}
