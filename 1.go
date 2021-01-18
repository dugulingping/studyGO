package main

import (
	"fmt"
	"math"
)
func swap(x string ,y string) (string, string) {	//golang不同于绝大多数语言，她支持函数的多返回值
	return y, x
}
//不同于绝大部分语言，如果在函数头定义了返回头，那么Golang为无参数的return语句自动添加了返回值
//Golang中的直接返回，没有参数的 return 语句返回已命名的返回值
func multiply_add(x, y int) (sum, result int)  {
	sum = x*y
	result = x+y
	return
}
func main()  {
	fmt.Println("hello",math.Pi)
	//:=为Golang的简短声明变量，作用等于var a, b string = swap("123", "asd")
	//不过它有一个限制，那就是它只能用在函数内部；在函数外部使用则会无法编译通过，所以一般还是用var方式来定义全局变量
	a, b := swap("123", "asd")		
	fmt.Println(a,b)
	
	//直接返回
	x, y := multiply_add(2,3)
	fmt.Println("multiply:",x,"add:",y)

	//defer 语句会将函数推迟到外层函数返回之后执行。
	//推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。
	defer fmt.Println("zxcvb")
	fmt.Println("nm,./")

	//Golang中数组切片的用法
	//s包含了2,3,4,即下标为1-4的元素
	s := []int {1,2,3,4,5,6}
	//var s[] int = array[1:4]
	//fmt.Println(s)

	//完整的数组
	fmt.Println(s)
	//切片长度与容量
	s = s[:0]
	fmt.Println("空切片:",s,"\t长度为",len(s),"\t容量为",cap(s))
	//拓展为3
	s = s[:3]
	fmt.Println("拓展为3",s,"\t长度为",len(s),"\t容量为",cap(s))
	//舍弃前2个
	s= s[2:]
	fmt.Println("舍弃前2个:",s,"\t长度为",len(s),"\t容量为",cap(s))

	//向切片内添加元素
	s = append(s, 4)
	fmt.Println("末尾添加4:",s,"\t长度为",len(s),"\t容量为",cap(s))

	//range遍历
	var pow = []int{1,2,4,8,16,32,64,128}
	for i, v:=range pow{
		fmt.Println(i,v)
	}
	//忽略索引
	for _, v:=range pow{
		fmt.Println(v)
	}
}