package main

import "fmt"

func main() {
	//var i int = 10
	//fmt.Println("i的地址",&i)
	//
	//var ptr *int = &i
	//fmt.Printf("ptr=%v\n",ptr)
	//fmt.Printf("ptr 的地址=%v\n",&ptr)
	//fmt.Printf("ptr 指向的值=%v\n",*ptr)

	var num int = 9
	fmt.Printf("num address=%v\n",&num)

	var ptr *int = &num

	*ptr = 11
	fmt.Println("num =",num)
}

//func some() {
//	var a int = 300
//	var b int = 400
//	var ptr *int = &a
//	*ptr = 100
//	ptr = &b
//	*ptr = 200
//	fmt.Printf("a=%d,b=%d,*ptr=%d",a,b,*ptr)
//}
