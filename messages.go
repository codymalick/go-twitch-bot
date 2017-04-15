package main

import (
	"strings"
	"time"

	"github.com/thoj/go-ircevent"
	//"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"fmt"
)

type Message struct {
	Id      int64 `bson:"_id"`
	Time    int64         `bson:"time"`
	Userid    int64        `bson:"userid"`
	Message string        `bson:"message"`
	Channelid int64        `bson:"channel"`
}

func createMessage(event *irc.Event, db string, cache *UserCache, user *User, chann *Channel) {
	// Check user cache for recent user
	// if !checkCache(event.User, &cache) {
	// 	cache = append(cache, User{})
	// }

	// Format message in form that the database can use
	mes := Message{Id: 0, Time: time.Now().Unix(), Userid: user.Id,
		Message: event.Message(), Channelid: chann.Id}

	switch databaseType {
	case "maria":
		err := mariaDbMessageInsert(mes, db, strings.TrimPrefix(event.Arguments[0], "#"))

		if err != nil {
			fmt.Println(err)
		}
		break
	case "mongo":
		err := mongoDbInsert(mes, db, strings.TrimPrefix(event.Arguments[0], "#"))

		if err != nil {
			panic(err)
		}
		break
	}

}
