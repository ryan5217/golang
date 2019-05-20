package main

import (
	_ "gocurd/api/database"
	"gocurd/api/router"
	orm "gocurd/api/database"
)

func main() {
	defer orm.Eloquent.Close()
	router := router.InitRouter()
	router.Run(":8081")
}
