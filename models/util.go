package model

import (
	"fmt"
)

const (
	dbName           = "raise"
	taskCollection   = "task"
	userCollection   = "user"
	wishCollection   = "wish"
	recentCollection = "recent"
)

func CheckError(err error, from string) {
	fmt.Println("error from %s: %s", from, err)
}
