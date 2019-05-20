package main

import "fmt"

type Monster struct {
	Name string
	Age int
}

type E struct {
	Monster
	int
	n int
}

func main() {
	var e E
	e.Monster.Name = "钢铁侠"
	e.Age = 56
	e.int = 5
	e.n = 40
	fmt.Println(e)
}
