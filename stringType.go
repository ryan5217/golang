package main

import (
	"fmt"
	"strconv"
)

func main() {
	var address string = "北京第三交通委提醒您 道路千万条 安全第一条 行车不规范 亲人两行泪"
	fmt.Println(address)

	var str = "hello " + "world"

	str += " ryan"

	fmt.Println(str)

	//var a int = 95
	var b byte = 'b'
	var c float64 = 2.2653
	//var d bool = false
	var e string

	e = strconv.FormatInt(int64(b),10)

	fmt.Printf("e type %T e = %q \n",e,e)

	//基本类型转换string类型
	e = strconv.FormatFloat(c,'f',10,64)
	//string类型转换为基本类型
	//strconv.ParseBool(str)
	//strconv.FormatBool()
	//strconv.ParseFloat()
	//e = fmt.Sprintf("%d",b)
	fmt.Printf("e type %T e = %q \n",e,e)

	// 基本类型之间转换 T(v) T就是类型 例float32 int64 但是变量本身的数值类型不会发生变化
	//var n1 float32 = float32(a)
	// 基本类型之间的转换 T(v) T就是类型 例float32 int64 但是变量本身的数值类型不会发生变化
	//var n1 float32 = float32(a)
	//var n2 float32
	//n2 = float32(a)


	//var num1 int64 = 999999
	//var num2 int8 = int8(num1)
	//fmt.Println("num2 = ",num2)

	// 说明： 'f'格式 10：表示小数位保留10位 64：表示这个额小数是float64
	//strconv.FormatFloat(num2,'f',10,64)


	var num5 int32
	num5 = int32(num)
	var num6 int64
	num6= int64(num)


}
