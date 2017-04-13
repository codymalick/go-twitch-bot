package main

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id         bson.ObjectId `bson:"_id"`
	Username   string        `bson:"username"`
	KappaCount int64         `bson:"kappacount"`
}

const (
	userCollection = "user"
)

func incrementKappa(username string, db string) {
	switch databaseType {
	case "mongo":
		if userCheck(username, db) {
			dbFindAndUpdateUser(username, db, userCollection)
		} else {
			user := User{bson.NewObjectId(), username, 1}
			mongoDbInsert(user, db, userCollection)

		}
		break
	case "maria":
		break
	}

}

func userCheck(username string, db string) bool {
	user := dbFindUser(username, db, userCollection)
	if len(user.Username) == 0 {
		return false
	}
	return true
}
