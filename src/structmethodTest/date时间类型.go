package main

import (
	"fmt"
	"time"
)

func main() {

	now:=time.Now()
	fmt.Println(now)
	//golang必须写死
	timeStr:=now.Format("2006/01/02 15:04:05")
	fmt.Println("time:%s\n",timeStr)
}
