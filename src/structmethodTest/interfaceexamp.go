package main

import "fmt"

//接口
type Humaner interface {
	//方法
	Say()
}
//结合多态
type Stu struct{
	name string
	score int
}

func(s *Stu)Say(){
	fmt.Println(s.name,s.score)
}

type Teacher struct {
	name string
	group string
}

func(t *Teacher)Say(){
	fmt.Println(t.name,t.group)
}

//自定义类型
type Mystr string
func(str Mystr)Say(){
	fmt.Println("今晚加班到后天",str)
}
//多态的核心
//参数为接口类型
func WhoSay(i Humaner){
	i.Say()
}


func main() {
	s:=&Stu{"stu",88}
	t:=&Teacher{"tea","golang"}
	var tmp Mystr="abc"
	//基本调用
	s.Say()
	t.Say()
	tmp.Say()
	//多态
	WhoSay(s)
	WhoSay(t)
	WhoSay(tmp)
}