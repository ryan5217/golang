package main

import "fmt"

func main() {

	//for i := 0; i < 9; i++ {
	//	if i % 3 == 0 {
	//		fmt.Print("\n*")
	//	} else {
	//		fmt.Print("*")
	//	}
	//}

	//for i := 1; i <= 3; i++ {
	//	switch i {
	//	case 1:
	//		fmt.Println("*")
	//	case 2:
	//		fmt.Println("**")
	//	case 3:
	//		fmt.Println("***")
	//	}
	//}

	var totalLevel int = 4

	for i := 1; i <= totalLevel; i++ {
		for k := 1; k <= totalLevel-i; k++ {
			fmt.Print(" ")
		}

		for j := 1; j <= 2*i-1; j++ {
			if j == 2*i-1 {
				fmt.Println("*")
			} else {
				fmt.Print("*")
			}
		}
	}
}
