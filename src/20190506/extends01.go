package main

import "fmt"

type Student struct {
	Name string
	Age int
	Score int
}

func (stu *Student) ShowInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v\n", stu.Name, stu.Age, stu.Score)
}

func (stu *Student) SetScore(score int) {
	stu.Score = score
}

//小学生
type Pupil struct {
	Student
}

func (p *Pupil) testing() {
	fmt.Println("小学生正在考试中...")
}

type Graduate struct {
	Student
}

