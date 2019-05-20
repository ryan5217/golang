package models

import (
	orm "gocurd/api/database"
)

type Test struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Age int64 `json:"age"`
}

func (Test) TableName() string {
	return "test"
}

var List []Test

func (test *Test) List() (lists []Test, err error){
	if err =  orm.Eloquent.Find(&lists).Error; err != nil {
		return
	}
	return
}
