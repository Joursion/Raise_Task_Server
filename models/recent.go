package model

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Recent struct {
	Id        string `bson:"_id"`
	Content   string `bson:"content"`
	CreatedAt string `bson:"createdAt"`
	Creator   string `bson:"creator"`
	Value     string `bson:"value"`
}

func CreateRecent(s *mgo.Session, userId string, content string, value int32) {
	session := s.Copy()
	defer session.Close()

	db := session.DB(dbName).C(recentCollection)
	createdAt := time.Now()
	return db.Insert(bson.M{creator: userId, content: content, value: value, createdAt: createdAt})
}

/** If userId exists , will remove all document, or remove only recent doc with id */
func DelRecent(s *mgo.Session, id string, userId string) {
	session := s.Copy()
	defer session.Close()

	db := session.DB(dbName).C(recentCollection)
	if userId {
		return db.RemoveAll(bson.M{creator: userId})
	}
	return db.RemoveId(bson.M{_id: id})
}
