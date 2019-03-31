package main

import "fmt"

func test(char byte) byte {
	return char + 1
}

func main() {
	var a byte = 's'
	println(test(a))

	var i int = 1

	switch i {
		case 1:
			fmt.Println("我是1")
			fallthrough //在case语句中增加fallthrough 会穿透 后面的一个case语句会执行 也是就是“我是二会打印出来”
		case 2:
			fmt.Println("我是二")
		default:
			fmt.Println("你什么都不是")
	}
}
