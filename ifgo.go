package main

import "fmt"

func main() {
	var score int

	fmt.Println("请输入成绩：")
	fmt.Scanln(&score)

	if score == 100 {
		fmt.Println("奖励一台BMW")
	} else if score > 80 && score <= 99 {
		fmt.Println("奖励一台iphone7")
	} else if score > 60 && score <= 79 {
		fmt.Println("奖励一个Ipad")
	} else {
		fmt.Println("回来跪搓衣板")
	}
}
