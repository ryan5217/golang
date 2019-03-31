package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var price = 15.26
	fmt.Print(price,"\n")
	fmt.Printf("price类型%T字节长度为%d",price,unsafe.Sizeof(price))
	fmt.Print("\n")

	var c1 byte = 'a'
	var c2 byte = '0'
	fmt.Printf("%c",c1)
	fmt.Println("c1 = ",c1)
	fmt.Println("c2 = ",c2)
	fmt.Printf("c1=%c c2=%c\n",c1,c2)

	var c3 = '北'
	fmt.Print(c3)
	fmt.Printf("c3=%c c3对应码值=%d",c3,c3)
}
