package main

import (
	"fmt"
	"strconv"
)

func main() {
	hh := "hello呗"
	fmt.Println("str len=",len(hh))

	r := []rune(hh)

	for i := 0; i < len(r); i++ {
		fmt.Printf("字符串=%c\n", r[i])
	}

	n, err := strconv.Atoi("hello")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转换成果是", n)
	}

	str := strconv.Itoa(33123)
	fmt.Printf("str=%v, str=%T", str, str)

}
