package main

//空接口

type Element interface {

}

type Persons struct{
	name string
	age int
}

func main() {
	list:=make([]Element,3)
	list[0]=1
	list[1]="Hello"
	list[2]=Persons{"张三",18}
	//循环
	/*
	for index,element:=range list{
		//类型断言:
		//value,ok=element.(T)
		if value,ok:=element.(int);ok{
			fmt.Println("int类型",index,value)

		}else if value,ok:=element.(string);ok{
			fmt.Println("是string类型",index,value)

		}else if value,ok :=element.(Persons);ok{
			fmt.Println("是persons类型",index,value)

		}else{
			fmt.Println("未知类型")
		}

	}*/
	for index.element:range list{
		switch value:=element.(type){

	}
	}

}