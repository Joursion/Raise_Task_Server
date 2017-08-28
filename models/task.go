package model

import (
	"fmt"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Task struct {
	Id string `bson:"_id"`
	Creator string `bson:"creator"`
	Repeat string `bson:"repeat"` //repeat type，daily|monthly|weekly
	Content string `bson:"content"`
	CreatedAt string `bson:"createdAt"`
	Reward int32 `bson:"reward"`
	FinishedAt string `bson:"finishedAt"`
	LastUpdate string `bson:"lastUpdate"` //
}

func IncRaise(s *mgo.Session, id string, count int32) {
	session := s.Copy()
	defer session.Close()

	db := session.DB(dbName).C(userCollection)

	data := bson.M{"$inc": bson.M{"raise": count}}
	return db.Update(bson.M{_id: id}, data)
}

func FindTaskById(s *mgo.Session, id string) *Task{
	session := s.Copy()
	defer session.Close()

	db := session.DB(dbName).C(taskCollection)

	var task Task
	_, err = db.Find(bson.M{_id: id}).One(task)
	return task
}

func CreateTask(s *mgo.Session, userId string, task Task) {
	session := s.Copy()
	defer session.Close()

	db := session.DB(dbName).C(taskCollection)

	return  db.Insert(bson.M(&Task(task)))
}

func DelTask(s *mgo.Session, userId string, taskId string) {
	session := s.Copy()
	defer session.Close()

	db := session.DB(dbName).C(taskCollection)
	return db.Remove(bson.M{_id: taskId})
}

func FindTasks(s *mgo.Session, userId string, page int32, limit int32) {
	session := s.Copy()
	defer session.Close()

	db := session.DB(dbName).C(taskCollection)

	var tasks []Task
	db.Find(bson.M{creator: userId}).Skip(page * limit).Limit(limit).Sort(bson.M{createdAt: -1}).All(&tasks)
	return tasks
}

func FindHistoryTasks(s *mgo.Session, userId string, page int32, limit int32) {
	session := s.Copy()
	defer session.Close()

	db := session.DB(dbName).C(taskCollection)

	res, err := db.Find(bson.M{userId: userId}).Skip(page * limit).Limit(limit)
	return res, err
}

func UpdateTask(s *mgo.Session, taskId string, task Task) {
	session := s.Copy()
	defer session.Close()

	db := session.DB(dbName).C(taskCollection)

	res, err := db.Update(bson.M{taskId: taskId}, bson.M{$set:{bson.M{task}}})
	return res, err
}

