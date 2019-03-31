package main

import (
	"fmt"
	"utils"
)

func main() {
	var n1 int32 = 26
	n3 := utils.Test(n1)
	fmt.Println(n3)

	fmt.Println(utils.Hehe())
}
//package main
//
//import (
//	"fmt"
//	"unsafe"
//)
//
//func main() {
//	fmt.Print("姓名\t年龄\t籍贯\t住址\t\n")
//	fmt.Print("ryan\t18\t江西\t九江\t\n")
//	var num = 1 + 1 * 9
//	fmt.Print(num,"\n")
//
//	types := "nan"
//
//	fmt.Print(types,"\n")
//
//	var name int = 2
//	name = 18
//	fmt.Print(name,"\n")
//
//	var n2 int = 10
//	fmt.Printf("n2 的类型 %T n2占用的字节数是%d",n2,unsafe.Sizeof(n2))
//}

//func main() {
//	fmt.Println("姓名\t年龄\t籍贯\t住址\t\n")
//	fmt.Println("ryan\t18\t江西\t九江\t\n")
//	var num = 1 + 1 +* 9
//	fmt.Println(num,"\n")
//	types := "nan"
//	types string = "nan"
//	fmt.Println(types)
//	var name int = 2
//	name = 18
//	fmt.Println(name,"\n")
//	var n2 int = 10
//	fmt.Printf("n2的类型%T n2占用的字节数是%d",n2,unsafe.Sizeof(n2))
//}
