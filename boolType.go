package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//fmt.Print("dadsa")
	var b = false

	fmt.Println("b = ",b)

	fmt.Println("b 的占用空间 = ",unsafe.Sizeof(b))
}
