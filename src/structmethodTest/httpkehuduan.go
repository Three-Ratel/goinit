package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	res, err := http.Get("http://127.0.0.1:8000/go")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.Status)
	defer res.Body.Close()
	buf := make(
		[]byte, 1024)
	var tmp string
	for {
		n, err := res.Body.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		}
		if n==0{
			fmt.Println("读取内容结束")
			break
		}
		tmp+=string(buf[:n])
	}
	fmt.Println(tmp)
}
