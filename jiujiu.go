package main

import "fmt"

func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			if i == j {
				fmt.Print(j,"*",i,"=",i*j,"\n")
			} else {
				fmt.Print(j,"*",i,"=",i*j," ")
			}
		}
	}
}
