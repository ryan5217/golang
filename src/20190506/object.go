package main

import "fmt"

type Student struct {
	Name string
	Age int
}

func (stu Student) String() string {
	str := fmt.Sprintf("Name=[%v] Age=[%v]", stu.Name, stu.Age)
	return str
}

type ChengFa struct {

}

func (stu ChengFa) Print()  {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}


func main() {
	var stu Student
	stu.Name = "tom"
	stu.Age = 20

	fmt.Println(stu)

	cheng := ChengFa{}
	cheng.Print()

	var stu1 = Student{"小明",12}
	stu2 := Student{"小红",18}

	var stu3 = Student{
		Name : "jack",
		Age : 20,
	}

	stu4 := Student{
		Age: 20,
		Name: "mary",
	}

	fmt.Println(stu1, stu2, stu3, stu4)
}
