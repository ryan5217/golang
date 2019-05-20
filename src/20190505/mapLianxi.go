package main

import "fmt"

func main() {
	users := make(map[string]map[string]string, 10)
	users["smith"] = make(map[string]string, 2)
	users["smith"]["pwd"] = "99999"
	users["smith"]["nickname"] = "小花猫"

	users["ryan"] = make(map[string]string,2)
	users["ryan"]["pwd"] = "1312939"
	users["ryan"]["nickname"] = "呵呵哒"

	modifyUser(users, "ryan")
	modifyUser(users, "tony")

	fmt.Println(users)
}

func modifyUser(users map[string]map[string]string, name string)  {
	//v, ok := users[name]

	if users[name] != nil {
		users[name]["pwd"] = "888888"
	} else {
		users[name] = make(map[string]string, 2)
		users[name]["pwd"] = "888888"
		users[name]["nickname"] = "昵称~"+ name
	}
}
