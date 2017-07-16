package model

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Task struct {
	id string
	repeat string //repeat type，dail|monthly|weekly
	content string 
}

func main() {
	session, err := mgo.Dial("")
	if err != nil {
		panic(err)
	}
	defer session.Close()
}