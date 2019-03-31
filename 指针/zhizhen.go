package main

import "fmt"

func main() {
	var i int = 90
	fmt.Println("i的地址是多少",&i)

	// * 获取指针类型所指向的值
	// * 获取指针类型所指向的值
	var ptr *int = &i
	fmt.Printf("ptr=%v\n",ptr)
	fmt.Printf("ptr 的地址=%v",&ptr)
	fmt.Printf("ptr 指向的值=%v",*ptr)

	//指针的使用细节
	//1)值类型 都有对应的指针类型 形式为 *数据类型 比如int 对应的指针就是*int float32对应的执政类型就是 *float32 以此类推
	//2)值类型包括 基本数据类型 int 系列 float 系列 bool string 数组和结构体struct


}
