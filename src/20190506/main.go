package main

import (
	"20190506/model"
	"fmt"
)

func main() {
	var stu = model.Student{
		Name: "tom",
		Score: 78.2,
	}

	fmt.Println(stu)

	var clazz = model.NewClazz("toem",56.6)

	fmt.Println(clazz)
}
