package main

import "fmt"

func main() {

	//map定义四种方式

	var m1 map[int]string

	m2 := map[int]string{}

	m3 := make(map[int]string)

	//指定大小
	m4 := make(map[int]string, 10)

	//1.定义初始化
	var m11 map[int]string = map[int]string{1: "abc", 2: "ddd"}
	//2
	m22 := map[int]string{1: "abc", 2: "cc"}
	fmt.Println(m1, m2, m3, m4, m11, m22)

	//键值操作
	mx := map[int]string{1: "北京", 2: "上海", 3: "广州"}

	mx[4]="杭州"
	mx[1]="北京"
	fmt.Println(mx)
    value,ok:=mx[1]
    fmt.Println(value,ok)
    delete(mx,2)
    fmt.Println(mx)
}


