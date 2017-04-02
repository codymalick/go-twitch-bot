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
	if userCheck(username, db) {
		dbFindAndUpdateUser(username, db, userCollection)
	} else {
		user := User{bson.NewObjectId(), username, 1}
		dbInsert(user, db, userCollection)
	}
}

func userCheck(username string, db string) bool {
	user := dbFindUser(username, db, userCollection)
	if len(user.Username) == 0 {
		return false
	}
	return true
}
