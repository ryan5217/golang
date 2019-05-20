package tool

import "github.com/google/uuid"

var uniqueID = make(chan string, 200)

// 阻塞式 unique id
func NewUniqueIDBlocked() {
	for {
		guid := uuid.New().String()
		uniqueID <- guid
	}
}

func init() {
	NewUniqueIDAsync()
}

// 异步 unique id
func NewUniqueIDAsync() {
	go NewUniqueIDBlocked()
}

// return unique id
func GetUID() string {
	select {
	case uID, okay := <-uniqueID:
		if okay {
			return uID
		}
	}
	return ""
}
