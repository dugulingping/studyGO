package main

import (
	"fmt"
)
func main () {
	//map(映射) 是 key - value 类型的数据结构，本身是无序的
	//map据说底层是使用hash实现的
	m := make(map[int]string, 10)
	m[123] = "B"
	fmt.Println(m)
	fmt.Printf("%v\n",m)

	//map可以套娃使用，使用前需要初始化，否则会报错
	//有两种初始化的方式
	//1. 直接初始化
	m1 := map[string]map[string]string {
		"1" :{
			"11":"A",
			"12":"B",
		},
		"2" :{
			"21":"C",
			"22":"D",
		},
	}
	fmt.Println(m1)
	//2. make方式
	m2 := make(map[string]map[string]string,10)
	m2["1"] = make(map[string]string,10)
	m2["1"]["11"] = "a"
	m2["1"]["12"] = "b"
	m2["2"] = make(map[string]string,10)
	m2["2"]["21"] = "c"
	m2["2"]["22"] = "d"
	fmt.Println(m2)
	
	//删除key
	//func delete(m map[Type]Type1, key Type)
	delete(m2["2"],"21")
	fmt.Println(m2)
}