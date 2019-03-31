package main

import "fmt"

func main() {
	var n1 int32 = 12
	var n2 int64
	var n3 int8

	n2 = int64(n1) + 20
	n3 = int8(n1) + 20

	fmt.Println(n2,n3)
}
