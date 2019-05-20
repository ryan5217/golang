package main

import "fmt"

var age int = 50
var Name string = "ryan"

func main() {
	//var num int32 = 20
	num := 20
	test(&num)
	fmt.Println("main() num", num)

	var n int
	fmt.Println("请输入打印的金字塔层数")
	fmt.Scanln(&n)
	printPyramid(n)
	HeHe(n)

	printMulti(n)
}

func test(n1 *int) {
	*n1 = *n1 + 10
	fmt.Println("test() n1 =", *n1)
}

func printPyramid(total int) {
	for i := 1; i < total; i++ {
		for k := 1; k <= total-i; k++ {
			fmt.Print(" ")
		}

		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func HeHe(total int)  {
	for i := 1; i < total; i++ {
		//fmt.Println("*")
		for k := 1; k <= i; k++ {
			if k == i {
				fmt.Println("*")
			} else {
				fmt.Print("*")
			}
		}
	}
}

func printMulti(num int) {
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v \t",j,i,j*i)
		}
		fmt.Println()
	}
}