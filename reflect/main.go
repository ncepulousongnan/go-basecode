package main

import "fmt"

//结构体的使用
type shuzu struct {
	name  string
	age   int
	class []string
}
type shu struct {
	name  string
	age   int
	class []string
}

//定义函数
func givename(ssss string) {
	//	var ssss shuzu
	fmt.Println(ssss)
}

//定义方法
func (ssss shuzu) giveage() {
	//	var ssss shuzu
	fmt.Println(ssss.age)
}

//为什么用指针定义可以实现这个方法，而值就不行
func (ssss *shuzu) changeage(x int) {
	ssss.age = x
}

func main() {
	//结构类型的赋值
	bbb := shuzu{"lsn", 21, []string{}} //"university"}}
	fmt.Println(bbb)
	var aaa shuzu
	aaa.name = "song"
	aaa.age = 19
	aaa.class = []string{"sss", "fff"}
	(&aaa).changeage(22)
	fmt.Println(aaa.age)
	if aaa.age < 20 {
		fmt.Println(aaa.class)
	} else {
		fmt.Println(aaa.age)
	}
	//方法和函数
	qqq := bbb.name
	givename(qqq)
	givename(aaa.name)
	aaa.giveage()
	bbb.giveage()

	//匿名结构体，临时应用
	var z struct {
		lll int
		ch  string
	}
	z.lll = 15
	z.ch = "j"
	fmt.Println(z, z.lll, z.ch)

	//创建映射
	//1 使用make a:=make(map[键的数据类型]值的数据类型)
	s1 := make(map[int]string)
	//2 使用映射字面量声明空映射并赋值
	s2 := map[string]int{
		"hhh": 1,
		"nnn": 2,
		"ddd": 3,
	}
	//映射的赋值操作
	s1[5] = "这是五"
	s1[6] = "这是六"
	//删除操作，delete（映射名，键）
	delete(s1, 5)
	//查看映射用range
	for i, k := range s1 {
		fmt.Println(i, k)
	}
	for q, p := range s2 {
		fmt.Println(q, p)
	}
	//获取一个键对应的值，并判断键值对是否存在
	m := s2["ddd"]
	if m != 0 {
		fmt.Println(m)
	} else {
		fmt.Println("键值对不存在")
	}

}
