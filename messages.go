package main

import (
	"strings"
	"time"

	"github.com/thoj/go-ircevent"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Message struct {
	Id      bson.ObjectId `bson:"_id"`
	Time    int64         `bson:"time"`
	User    string        `bson:"user"`
	Message string        `bson:"message"`
	Channel string        `bson:"channel"`
}

func createMessage(event *irc.Event) {
	mes := Message{Id: bson.NewObjectId(), Time: time.Now().Unix(), User: event.User,
		Message: event.Message(), Channel: event.Arguments[0]}

	// Connect to database, https://labix.org/mgo
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}

	// Add connection close to the stack
	defer session.Close()

	// Access the appropriate collection
	c := session.DB("TwitchEmoji").C(strings.TrimPrefix(event.Arguments[0], "#"))
	err = c.Insert(mes)

	if err != nil {
		panic(err)
	}

}
