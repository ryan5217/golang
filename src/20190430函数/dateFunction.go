package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().Format("2006-01-02 15:04:05")

	fmt.Println(now)

	i := 0
	for {
		i++
		fmt.Println(i)
		time.Sleep(time.Microsecond * 100)
		if i == 100 {
			break
		}
	}

	fmt.Println(time.Now().Unix())

	//func (t Time) Unix() int64


}