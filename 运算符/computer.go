package main

import "fmt"

func main() {
	//运算符
	// 运算符是一种特殊的符号 用以表示数据的运算和赋值和比较等
	// 运算符是一种特殊的符号 用以表示数据的运算和赋值和比较等
	// 算术运算符 赋值运算符 比较运算符 逻辑运算符 位运算符 其他运算符
	// 如果运算的数都是整数 那么除后 去掉小数部分 保留整数部分
	fmt.Println(10 / 4)

	var n1 float32 = 10 /4
	fmt.Println(n1)

	//如果希望保留小数部分 则需要有浮点数参加运算
	fmt.Println(10.0 / 4)

	//算术运算符使用的注意事项
	// 对于除号 / 它的整数除和小数除湿有区别的 整数之间做除法时 只保留整数部分 小数除看类型保留长度
	// 当对一个数 取模时 可以等价 a%b = a - a/b*b

	var i int = 10
	var n int = 2
	i++
	n += i
	fmt.Println(n)

	if i > 8 {
		fmt.Println("大于8了")
	}

	var days int = 97
	var week int = days / 7
	var day int = days % 7
	fmt.Printf("%d个星期零%d天\n",week,day)

	// 定义一个标量保
}
