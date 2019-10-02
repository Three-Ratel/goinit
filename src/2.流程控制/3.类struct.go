package main

import "fmt"

type Student struct {
	id int
	name string
	sex byte
	age int
	addr string
}

func main() {
	var s1 Student =Student{1,"约翰",'f',20,"shahe"}
	//2
	s2:=Student{id:2,age:20}
	//3结构体作为指针变量初始化
	var s3 *Student=&Student{3,"abc",'m',25,"changping"}

	fmt.Println(s1,s2,s3)
	fmt.Println(s1.id,)
	fmt.Println((*s3).id)
	fmt.Println(s3.id)

	fmt.Println("传递非指针")

	temStudent2(&s1)
	fmt.Println("main",s1)
}

//非指针传递
//func temStudent(tmp Student){
//	tmp.id=250
//	fmt.Println(tmp)
//}

//指针传递
func temStudent2(p *Student){
	p.id=300
	fmt.Println(p)
}

