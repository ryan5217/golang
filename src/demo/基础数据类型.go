package main

import "fmt"

func main() {

	type Myint int
	var num1 Myint
	var num2 int
	num1 = 20
	num2 = int(num1)
	fmt.Println("num1=",num1,"num2=",num2)

	type myFuntype func(int, int) int

	res3 := myFun2(getSum,500,600)
}

func myFun2(funvar func(int, int),num1 int,num2 int) int {
	return funvar(num1,num2)
}

func getSum()  {
	fmt.Println("getSum")
}

func sum(n1 int,args ...int) int {
	sum := n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func getSumAndSub(n1 int,n2 int) (sum int,sub int) {
	sub = n1 - n2
	sum = n1 + n2
	return sub,sum
}