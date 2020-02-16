package main

import "fmt"

func main() {
	//简单数组内容
	s1 := [10]int{5, 8, 3, 23, 58, 2, 35, 88, 96, 47}
	s2 := s1
	s2[0] = 26
	fmt.Println(s1, s2)
	//数组的遍历方法
	for i := 0; i < len(s1); i++ {
		fmt.Println(s1[i])
	}
	for k, h := range s1 {
		fmt.Println(k, h)
	}
	//切片的内容
	//直接创造一个切片,var slice 数据类型
	//切片为应用类型无法直接比较，判断切片是否为空应该用len()==0
	var p1 []int
	fmt.Println(p1, len(p1), cap(p1))
	fmt.Println(p1 == nil)
	s3 := []int{}
	fmt.Println(s3)
	fmt.Println(s3 == nil)
	fmt.Println(len(s3) == 0)
	s3 = []int{2, 1, 63, 9}
	fmt.Println(s3, len(s3), cap(s3))
	//从数组中切片，三种方式
	s6 := s1[:]
	s7 := s1[:5]
	s8 := s1[4:]
	s4 := s1[2:8]
	fmt.Println(s4, len(s4), cap(s4))
	fmt.Println(s6, len(s6), cap(s6))
	fmt.Println(s7, len(s7), cap(s7))
	fmt.Println(s8, len(s8), cap(s8))
	//切片的长度就是切片包含的元素个数，切片的容量是切片的第一个元素到底层数组的最后一个元素
	//切片中的元素会随着底层数组的变化而变化
	s1[3] = 100
	fmt.Println(s4, len(s4), cap(s4))
	//切片再切片
	s5 := s4[:4]
	fmt.Println(s5, len(s5), cap(s5))
	//利用make来构造切片，可以指定切片的长度和容量。make([]数据类型，长度，容量)
	a1 := make([]int, 4, 7)
	fmt.Println(a1 == nil, len(a1) == 0)
	//对构造的切片赋值
	a1[2] = 5
	a1[1] = 9
	fmt.Println(a1, len(a1), cap(a1))
	//利用append方法对切片增长
	a2 := []int{236, 857, 456, 12, 47, 56, 9, 7, 1, 23}
	a3 := a2[2:9:9]
	fmt.Println(a2, a3)
	//append可以让切片长度增加，但切片是引用类型，
	//要让200加到切片的下一个必须要让底层数组的对应位置更改
	//这会影响下面a4的切片和扩充，所以要用三个索引创造切片的方法让a3的长度和容量一致，
	//a3在扩充时会创造新的底层数组，不会影响到原数组a2
	a3 = append(a3, 200)
	fmt.Println(a2, a3)
	//原来的长度和容量为7，在增加数的过程中创造新的数组，让容量变为14
	fmt.Println(len(a3), cap(a3))
	//若切片无可用容量使用append是会自动创建新的底层数组并扩大来满足要求
	a4 := a2[8:]
	fmt.Println(a4)
	a4 = append(a4, 300)
	fmt.Println(a2, a4)

}
