package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Wish struct {
	Id        string `bson:"_id"`
	Creator   string `bson:"creator"`
	Value     int32  `bson:"price"`
	Content   string `bson:"content"`
	CreatedAt string `bson:"createdAt"`
	Finished  bool   `bson:"finished"`
}

func FinishWish(s *mgo.Session, id string) {
	session := s.Copy()
	defer session.Close()

	db := session.DB(dbName).C(wishCollection)
	return db.Update(bson.M{_id: id}, bson.M{"$set": bson.M{finished: true}})
}

func UpdateWish(s *mgo.Session, wishId string, content string, value int) {
	session := s.Copy()
	defer session.Close()

	db := session.DB(dbName).C(wishCollection)
	return db.Update(bson.M{_id: id}, bson.M{"$set": bson.M{content: content, value: value}})
}

func DelWish(s *mgo.Session, wishId string) {
	session := s.Copy()
	defer session.Close()

	db := session.DB(dbName).C(wishCollection)

	return db.Remove(bson.M{id: wishId})
}

func GetWishes(s *mgo.Session, userId string, page int32, limit int32, isFinished bool) {
	session := s.Copy()
	defer session.Close()

	db := session.DB(dbName).C(wishCollection)

	wishes := []Wish{}
	db.Find(bson.M{creator: userId, finished: bson.M{"$exists": isFinished}}).Skip(page * limit).Limit(limit).Sort("createdAt").All(&wishes)
	return wishes

}

func GetWishById(s *mgo.Session, wishId string) {
	session := s.Copy()
	defer session.Close()

	db := session.DB(dbName).C(wishCollection)

	return db.Find(bson.M{_id: wishId})
}
