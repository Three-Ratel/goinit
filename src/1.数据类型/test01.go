package __数据类型

import (
	"fmt"
	"reflect"
)

func enums2(){
	const(
		python=0
		java
		golang=1
		php
	)
	fmt.Println(python,java,golang,php)
}
//iota生成器   常量自增
func enums3(){
	const(
		python=iota
		java
		golang
	)
	fmt.Println(python,java,golang)
}

/*
func main(){
	enums3()
}
 */

func enums4(){
	//位运算
	//b等于1左移
	const(
	b=1<<(10*iota)
	kb
	mb
	gb )
	fmt.Println(b,kb,mb,gb)
}


func main() {
	//enums4()
	var ch byte
	ch='a'
	fmt.Println(ch)
	fmt.Printf("ch=%c\n",ch)

	var str string
	str="abc"
	fmt.Println(str[0])
	fmt.Println(reflect.TypeOf(str[0122]))
}



