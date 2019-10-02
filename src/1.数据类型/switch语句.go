package main

import "fmt"

func grade(score int) string{
	g:=""
	switch  {
	case score <0 || score>100:
		g="输入错误"
	case  score <60:
		g="凉凉"
	case score <80:
		g="还行"
	}
	return g
}

func main() {
	fmt.Println(grade(79))
}
