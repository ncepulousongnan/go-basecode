package main

import "fmt"

//多态和接口
//定义两个结构体，结构体就是一组变量的集合
type cat struct {
	move string
	eat  string
}
type bird struct {
	move string
	eat  string
}

//内部嵌入类型
type bcat struct {
	cat
	color string
}

//定义一个接口，接口就是一组方法的集合
type animal interface {
	walk()
	chi()
}

//分别写出接口里的方法
//接口实现时候的方法集
//当接收者为值时，在main里放入方法中的值可以是值，也可以是指针，当接收者为指针时
//在main里放入方法中的值只能是指针
func (c cat) walk() {
	fmt.Println("猫的走路方式是", c.move)
}
func (c cat) chi() {
	fmt.Println("猫的食物是", c.eat)
}
func (b bird) walk() {
	fmt.Println("鸟的走路方式是", b.move)
}
func (b bird) chi() {
	fmt.Println("鸟的食物是", b.eat)
}

//外部类型的方法声明
func (b bcat) walk() {
	fmt.Println("黑猫的走路方式是", b.move)
}
func (b bcat) chi() {
	fmt.Println("黑猫的食物是", b.eat)
}

//通过sendmessage函数来实现接口
func sendmessage(a animal) {
	a.walk()
	a.chi()
}
func main() {
	//多态
	var s1 cat
	s1.move = "行走"
	s1.eat = "鱼"
	sendmessage(&s1)
	s2 := bird{
		move: "飞",
		eat:  "虫子"}
	sendmessage(&s2)
	//内部嵌入类型的使用
	s3 := bcat{
		cat:   cat{"奔跑", "老鼠"},
		color: "black"}
	sendmessage(s3)
	sendmessage(s3.cat)
	s3.chi()
	s3.cat.chi()

}
