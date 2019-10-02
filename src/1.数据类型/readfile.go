package __数据类型

import (
	"fmt"

	"io/ioutil"
)


func main(){
	const filename="test01.go"
	contents,err:=ioutil.ReadFile(filename)
	if err!=nil{
		fmt.Println(err)
	}
	el (
		fmt.Printf("%s",contents))



}

