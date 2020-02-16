package main

import "fmt"

//声明全局变量
var (
	b int
	c string
	d bool
)
var ddd [5]int

func main() {
	//全局变量的赋值
	b = 100
	c = "aaa"
	d = true
	//局部变量
	a := 12
	var f = "你是谁"
	var sum int
	fmt.Println("Helloworld")
	a = a + 2*6
	//if判断语句的应用
	if b < 20 {
		fmt.Println(a)
	}
	//循环语句写累加的过程
	for n := 0; n < 11; n++ {
		sum += n
	}
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(sum)
	//fmt.Println(add(10), 222)
	add(10)
	ying(19)
	fmt.Print(f)
	//简单数组的运用
	for ppp := 0; ppp < 5; ppp++ {
		ddd[ppp] = ppp*2 - 9
	}
	for ppp := 0; ppp < 5; ppp++ {
		fmt.Print(ddd[ppp])
	}
}

//写一个累加的函数
func add(m int) {
	var o int
	for p := 1; p < m; p++ {
		o += p
	}
	fmt.Println(o)
}

//写一个递归的函数，求一个数的因子
func ying(n int) {
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
