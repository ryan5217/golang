package main

import (
	"encoding/json"
	"fmt"
)


type Cat struct {
	Name string
	Age int
	Color string
	Hobby string
	ptr *int
	slice []int
	map1 map[string]string
	float66 float64
}

type Point struct {
	x int
	y int
}

type Rect struct {
	leftUp, rightDown Point
}

type Rect2 struct {
	leftUp, rightDown *Point
}

type Monster struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Skill string `json:"skill"`
}

func (a Monster) test() {
	fmt.Println(a.Name)
}

func (a Monster) computer()  {
	res := 0
	for i := 1; i <= 100; i++ {
		res += i
	}
	fmt.Println(a.Skill, "计算的结果是=", res)
}

func (a Monster) getSum(n1 int, n2 int) int {
	return n1 + n2
}

func main() {
	//var cat1Name string = "小白"
	//cat1Age := 3
	//cat1Color := "白色"
	//
	//var cat2Name string = "小花"
	//var cat2Age int = 99
	//cat2Color := "花色"
	//
	//var catName [2]string = [...]string{"小花", "小白"}
	//var catAges [2]int = [...]int{3, 100}
	//var catColors [2]string = [...]string{"白色", "花色"}
	//
	//
	//cat := make(map[string]map[string]string, 10)
	//cat["cat01"] = make(map[string]string, 10)
	//cat["cat01"]["name"] = "小白"
	//cat["cat01"]["age"] = "3"
	//cat["cat01"]["color"] = "白色"
	//
	//cat["cat02"] = make(map[string]string, 10)
	//cat["cat02"]["name"] = "小花"
	//cat["cat02"]["age"] = "99"
	//cat["cat02"]["color"] = "花色"



	var cat1 Cat
	cat1.Name = "小白"
	cat1.Age = 3
	cat1.Color = "白色"
	cat1.Hobby = "吃》fish"

	fmt.Println("cat1 =", cat1)

	var cat2 Cat
	cat2.Name = "小花"
	cat2.Age = 99
	cat2.Color = "花色"
	//cat2.Hobby = "吃 beef"

	fmt.Println("cat2 =", cat2)

	cat3 := Cat{}
	fmt.Println("cat3 =",cat3)


	r1 := Rect{Point{1,2}, Point{3, 4}}
	fmt.Printf("r1.leftUp 地址=%p r1.leftUp.y 地址=%p r1.rightDown.x 地址=%p r1.rightDown.y", &r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y)

	fmt.Println("")

	monster := Monster{"钢铁侠",56, "灭霸666~"}

	jsonStr, err := json.Marshal(monster)

	if err != nil {
		fmt.Println("json 处理错误", err)
	}

	monster.test()
	monster.computer()
	var res int = monster.getSum(10, 20)

	fmt.Println(res)
	fmt.Println("jsonStr", string(jsonStr))
}
