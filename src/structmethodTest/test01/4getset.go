package test01

import "fmt"

type Dog struct{
    name string
    sex int
}

func(d *Dog)SetName(name string){
    d.name=name
    //return d.name
}
func(d *Dog)GetName()string{
    return d.name
}

func (d *Dog) bit(){
    fmt.Println("",d.name)
}

func Tester01(){
    d:=Dog{"erha",11}
    d.bit()
}

func main() {
    Tester01()
}
