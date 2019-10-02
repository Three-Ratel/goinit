package main

import "fmt"

type MyInt int

func Add(a,b MyInt) MyInt{
	return a+b
}

func(a MyInt) Add (b MyInt)  MyInt{
	return a + b
 }

type Person struct {
	name string
	sex string
	age int
}

func (p Person)PrintInfo(){
	fmt.Println(p.name,p.sex,p.age)
}

//值语义与引用语义
func (p *Person)SetPrintInfo(){
	p.name="ls"
}

func (p Person)SetPrintValue(){
	p.name="wang"
}

type Student struct {
	Person
}

func(p *Student) PrintInfo(){
	fmt.Println("Student",p.name)
}


func main() {
	var a MyInt=1
	var b MyInt=1
	fmt.Println(Add(a,b))
	fmt.Println(a.Add(b))
	p:=Person{"zs","male",18}
	p.PrintInfo()

	(&p).SetPrintInfo()
	(p).SetPrintValue()
	fmt.Println(p)

	s:=Student{Person{"zs","male",18}}
	s.SetPrintValue()
	s.PrintInfo()
	//方法值
	pFunc:=p.PrintInfo
	pFunc()
	//方法表达式
	pFunc2:=(*Person).PrintInfo
	pFunc2(&p)

}