package main

import (
	"fmt"
	"time"
)

func for2 (){
	for i:=1 ; i<5;i++{
		fmt.Println("123")
		time.Sleep(1*time.Second)
	}
}

func for3(){
	s:="abc"
	for i :=range s{
		fmt.Printf("下标是%d,值%c",i,s[i])
	}
}

func tt3(args ...int){
	for _,n := range args{
		fmt.Println(n)}
}

func plus()(res int){
	sum:=0
	for i:=0;i<101;i++{
		fmt.Println(i)
		sum+=i

	}
	//fmt.Println(sum)
	return sum
}

func plus2(first   int)int {
	if first ==1{
		return 1}

	return  first+plus2(first-1)
}


func main() {
	//tt3(1,2)
	res:=plus2(100)
	fmt.Println(res)
}
