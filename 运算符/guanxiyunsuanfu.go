package main

import "fmt"

func main() {
	var n1 int = 9
	var n2 int = 8
	//flag := n1 > n2
	var flag bool = n1 > n2
	fmt.Println("flag",flag)

	a := 9
	b := 2
	t := a
	a = b
	b = t

	fmt.Printf("交换后的情况是a = %v ,b = %v \n",a,b)

	var a int = 10
	var b int = 20
	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("a = %v b = %v",a,b)
}
