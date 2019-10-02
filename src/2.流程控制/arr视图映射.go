package main

import "fmt"

func main() {
	arr:=[...]int{0,1,2,3,4,5,6,7}
	arr2:=arr[:2]
	fmt.Println(len(arr2))
	fmt.Println(cap(arr2))
	fmt.Println(append(arr2,245))
	fmt.Println(arr)
	s5:=append(arr2,110,4,4,4,4,4,4,4,4,4,4,4,4,4,4)
	fmt.Println(s5)
	fmt.Println(arr)
	fmt.Println("=============")
	s1:=arr[8:]
	s2:=arr[:5]
	copy(s2,s1)
	fmt.Println(s2,s1,arr)
}
//内建函数copy   2个切片直接复制数据
