package main

import "fmt"

func main() {
	a := 10
	b := 20
	swap(&a,&b)
	fmt.Println(a,b)
}

func swap(n1 *int, n2 *int) {
	t := *n1
	*n1 = *n2
	*n2 = t
}