package main

import "fmt"

func main() {
	/*
	一个养鸡场有6个鸡 他们体重分别为3kg 5kg 1kg 请问这六只鸡的总体重是多少？
	平均体重？
	 */
	 hen1 := 3.0
	 hen2 := 5.0
	 hen3 := 1.0
	 hen4 := 3.4
	 hen5 := 50.0

	 total := hen1 + hen2 + hen3 + hen4 + hen5
	 avg := fmt.Sprintf("%.2f", total / 6)
	 fmt.Printf("total=%v avg=%v",total,avg)
}

func mapCheck() {
	var hens [7] float64

	total := 0.0

	for i := 0; i < len(hens); i++ {
		total += hens[i]
	}

	heros := [...]string{"钢铁侠","浩克","黑寡妇","黑豹"}

	for i, v := range heros {
		fmt.Printf("i=%v v=%v\n", i, v)
	}

	for _, v := range heros {
		fmt.Printf("元素的值=%v\n", v)
	}


	var slice  []float64 = make([]float64, 5, 10)
	slice[1] = 10
	slice[3] = 20
	fmt.Println(slice)
	fmt.Println("slice的size=",len(slice))
	fmt.Println("slice的cap=", cap(slice))

	var slice4 []int = []int{1,2,3,4}
	var slice5 = make([]int, 10, 19)
	copy(slice5,slice4)


	str := "heerwer@dad"
	slice6 := str[6:]
	fmt.Println("slice=", slice6)

	var a map[string]string
	a = make(map[string]string, 10)
	a["no1"] = "钢铁侠"
	a["no2"] = "黑寡妇"
	a["no3"] = "美国队长"

	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"

	heross := map[string]string{
		"her" : "dasd",
		"er2" : "ssss",
		"dasd" : "das",
	}

	studentMap := make(map[string]map[string]string)

	studentMap["std01"] = make(map[string]string, 3)
	studentMap["std01"]["name"] = "tom"
	studentMap["std01"]["sex"] = "男"
	studentMap["std01"]["address"] = "北京长安街"

	studentMap["std02"] = make(map[string]string, 3)
	studentMap["std02"]["name"] = "ryan"
	studentMap["std02"]["sex"] = "男"
	studentMap["std02"]["address"] = "上海黄浦区"

	delete(studentMap, "std01")

	//val, ok := studentMap["std02"]
	students := make(map[string]Stu, 10)


}
