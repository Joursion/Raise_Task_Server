package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbName = "raise"
)

type User struct {
	Id        string `bson"_id"`
	Name      string `bson"name"`
	Email     string `bson:"email"`
	CreatedAt string `bson:"createdAt"`
	Raise     int32  `bson:"raise"`
}

func CreateUser(s *mgo.Session, user User) {
	session := s.Copy()
	defer session.Close()

	db := session.DB(dbName).C(userCollection)
	return db.Insert(&user)
}

func FindUserByName(s *mgo.Session, name string) {
	session := s.Copy()
	defer session.Close()

	db := session.DB(dbName).C(userCollection)
	user := User{}
	res, err = db.Find(bson.M{"name": name}).One(&user)
	checkErr(err)
	return user
}

func UpdateUserByname(s *mgo.Session, name string) {
	session := s.Copy()
	defer session.Close()

	// db := initDB(session)
	db := session.DB(dbName).C("user")

}

func initDB(session *mgo.Session) {
	return session.DB(dbName).C("user")
}

func IncRaise(s *mgo.Session, id string, count int32) {
	session := s.Copy()
	defer session.Close()

	db := session.DB(dbName).C("user")

	data := bson.M{"$inc": bson.M{"raise": count}}
	res, err := db.Update(bson.M{id: id}, data)
	checkErr(err)
	return res
}
