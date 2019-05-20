package main

import "fmt"

func main() {
	//defer在函数执行完毕后，及时的释放资
	fmt.Println(sum(1,3))
}

func sum(n1 int, n2 int) int {
	// 当执行到defer时 暂时不执行 会将defer后面的语句压人到独立的栈（defer栈）
	// 当函数执行完毕后 在从defer栈 按照先入后出的方式执行
	defer fmt.Println("ok1 n1=",n1)
	defer fmt.Println("ok2 n2=",n2)

	res := n1 + n2
	fmt.Println("ok3 res=",res)
	return res
}

//func test() {
//	file = openfile('./index.txt')
//	defer file.close()
//}
