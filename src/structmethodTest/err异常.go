package main

import (
	"errors"
	"fmt"
)

//系统异常

func test01() {
	a := [5]int{0, 1, 2, 3, 4}
	index := 10
	a[index] = 123
}

func main() {
	//test01()
	//getCirleArea(-5)
	//test03()
	test04()
	ret, err := getCircleArea2(-5)
	if err != nil {
		fmt.Println(err)

	} else{
		fmt.Println(ret)
	}
}

//自己抛异常
func test02() {

}
func getCirleArea(radius float32) (area float32) {
	if radius < 0 {
		panic("半径不能为负")
	}
	return 3.14 * radius * radius
}

func test03() {
	//延迟到函数结束
	//异常发生时
	defer func() {
		//recover返回程序为什么挂了
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	getCirleArea(-5)
	fmt.Println(123)
}

func test04() {
	test03()
	fmt.Println(4)
	//bytes,e:=ioutil.ReadFile()
}

func getCircleArea2(radius float32) (ret float32, err error) {
	if radius < 0 {
		err = errors.New("半径不能为负")
		return
	}
	ret = 3.14 * radius * radius
	return
}
