package main

import (
	//"gopkg.in/mgo.v2/bson"
	"strings"
)

type User struct {
	Id         int64 				 `bson:"_id"`
	Username   string        `bson:"username"`
//	KappaCount int64         `bson:"kappacount"`
}

// UserCache is used to reduce calls to the user table in mariadb
type UserCache struct {
	Recent []User
}

const (
	userCollection = "user"
	cacheSize = 100
)

func checkCache(username string, cache *UserCache) *User {
	for _,user := range cache.Recent {
		// strings.Compare() returns 0 if a == b
		if strings.Compare(user.Username, username) == 0 {
			// if the user exists in the cache, return nil
			return &user
		}
	}
	return nil

}

// func incrementKappa(username string, db string) {
// 	switch databaseType {
// 	case "mongo":
// 		if userCheck(username, db) {
// 			dbFindAndUpdateUser(username, db, userCollection)
// 		} else {
// 			user := User{0, username}
// 			mongoDbInsert(user, db, userCollection)
//
// 		}
// 		break
// 	case "maria":
// 		break
// 	}
//
// }

func userCheck(username string, db string) bool {
	switch databaseType {
	case "mongo":
		user := dbFindUser(username, db, userCollection)
		if len(user.Username) == 0 {
			return false
		}
		break
	case "maria":
		user := mariaFindUser(username, db)
		if len(user.Username) == 0 {
			return false
		}
		break
	}
	return true
}
