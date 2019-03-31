package main

import "fmt"

func main() {

	for i := 1 ; i <= 10; i++ {
		fmt.Println("你好 我是邹忠豪",i)
	}

	var j int = 10

	for j <= 18 {
		fmt.Println("你好 我是jjj",j)
		j++
	}

	// while
	k := 1
	for {
		if k <= 10 {
			fmt.Println("ok ~",k)
		} else {
			break
		}
		k++
	}
	s := 1
	// do while
	for {
		s++
		if s <= 10 {
			fmt.Println("ss ~",s)
		} else {
			break
		}
	}

	//var str string = "hello world上海"
	//
	//for i := 0; i < len(str); i++ {
	//	fmt.Printf("%c \n",str[i])
	//}
	//
	//// 常用的遍历 记住 肯定常用
	//for index, val := range str{
	//	fmt.Printf("index = %d,val = %c \n",index,val)
	//}
	//
	//var count uint64
	//var sum uint64 = 0
	//var i uint64 = 1
	//for ; i <= 100; i++ {
	//	if i%9 == 0 {
	//		count++
	//		sum += i
	//	}
	//}
	//
	//fmt.Printf("count = %v sum = %v \n",count,sum)



	var classNum int = 3
	var stuNum int = 5
	var totalSum float64 = 0.0

	for j := 1; j <= classNum; j++ {
		sum := 0.0

		for i := 1; i <= stuNum; i++ {
			var score float64
			fmt.Printf("请输入第%d班 第%d个学生的成绩 \n",j,i)
			fmt.Scanln(&score)

			sum += score
		}

		fmt.Printf("第%d个班级的平均分是%v\n",j,sum / float64(stuNum))

		totalSum += sum
	}

	fmt.Printf("各个班级的总成绩%v 所有班级平均分是%v\n",totalSum,totalSum / float64(stuNum))



}
