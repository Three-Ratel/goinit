//package main
//
//import (
//	"fmt"
//)
//
////循环字符串
//func main()  {
//	var a = "alexdsb"
//	for i,c := range a{
//		fmt.Printf("下标%d,值为%c\n",i,c)
//
//		if i == 2 {
//			fmt.Println("结束了")
//			break
//		}
//	}
//}



package main

import (
	"fmt"
	"time"
)

func main(){
	defer fmt.Println("a")
	defer fmt.Println("b")
	defer test(0)
	defer fmt.Println("c")
}
func test(x int) {
	time.Sleep(1*time.Second)
	fmt.Println(100 / x)

}
