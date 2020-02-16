package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func hello() {
	defer wg.Done()
	fmt.Println("helloworld")
}
func ying(n int) {
	defer wg.Done()
	var max int
	max = 0
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
	wg.Add(2)
	x := 2365874
	go ying(x)
	go ying(x)
	wg.Wait()
	fmt.Println("hello")
}
