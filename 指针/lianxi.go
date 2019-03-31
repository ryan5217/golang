package main

import "fmt"

func main() {
	var i int = 13
	fmt.Println(&i)

	var ptr *int = &i

	*ptr = 10
	fmt.Println("num = ",i)


	var a int = 300
	ptr = &a
	*ptr = 5000

	fmt.Println("num = ",ptr)
}
