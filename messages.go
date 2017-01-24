package main

import (
	"strings"
	"time"

	"github.com/thoj/go-ircevent"
	//"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Message struct {
	Id      bson.ObjectId `bson:"_id"`
	Time    int64         `bson:"time"`
	User    string        `bson:"user"`
	Message string        `bson:"message"`
	Channel string        `bson:"channel"`
}

func createMessage(event *irc.Event, db string) {
	// Format message in form that the database can use
	mes := Message{Id: bson.NewObjectId(), Time: time.Now().Unix(), User: event.User,
		Message: event.Message(), Channel: event.Arguments[0]}

	err := dbInsert(mes, db, strings.TrimPrefix(event.Arguments[0], "#"))

	if err != nil {
		panic(err)
	}
}
