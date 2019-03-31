package main

import (
	"fmt"
	"strings"
)

func init()  {
	fmt.Println("init()...")
}

func main() {
	//var da int = 132;
	fmt.Println("main(),,...")

	f := addUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(20))

	f2 := mackSuffix(".jpg")
	fmt.Println("文件处理后",f2("win"))
	fmt.Println("文件处理后",f2("hello.jpg"))
	fmt.Println("文件处理后",f2("hello2"));
}

func addUpper() func(int) int {
	var n int = 12
	var str = "hello"
	return func(x int) int {
		n = n + x
		str += string(36)
		fmt.Println("str =",str)
		return n
	}
}

func mackSuffix(suffix string) func(string) string {
	// 闭包的理解 就是常驻内存
	return func(name string) string {
		if !strings.HasSuffix(name,suffix) {
			return name + suffix
		}
		return name
	}
}
