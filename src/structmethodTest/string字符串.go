package main

import "fmt"

func main() {
	var str="hello"
	fmt.Println(str[0])
	fmt.Println(len(str))
	//转为字节切片
	var bs[]byte= []byte(str)
	bs[0]='a'
	str=string(bs)
	fmt.Println(str)

	str1:="hello你好"
	fmt.Println(len(str1))
	var r =[]rune(str1)
	fmt.Println(len(r))
}
