package __数据类型

import (
	"fmt"
	"reflect"
)

func main() {
	type myint int
	var i myint = 100
	fmt.Println(i)
	fmt.Println(reflect.TypeOf(i))
}
