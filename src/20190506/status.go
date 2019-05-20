package main

import (
	"20190506/model"
	"fmt"
)

func main() {
	//go语言的特性

	p := model.NewPerson("smith")

	p.SetAge(18)
	p.SetSal(5626.5)
	fmt.Println(p)
	fmt.Println(p.Name, " age =", p.GetAge(), " sal = ", p.GetSal())
}
