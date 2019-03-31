package fuxi

import "fmt"

func getNum(n1 int) string {
	fmt.Println("n1 = ",n1)
	return "你是谁哈哈哈"
}

func main() {
	getNum(12)

	var i int = 12
	fmt.Println(i)

	var num = 123
	fmt.Println(num)

	n1 := 132
	n2 := 13
	n3 := "呵呵哒"
	fmt.Println(n1,n2,n3)

	someone,sometow,somethree := 1,"呵呵哒",'猪'

	fmt.Println(someone,sometow,somethree)
}
