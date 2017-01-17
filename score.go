package main

import (
  	"strings"
    "fmt"
  	//"time"

  	"github.com/thoj/go-ircevent"
  	//"gopkg.in/mgo.v2"
  	//"gopkg.in/mgo.v2/bson"
)

func userPoints(plusminus bool, event irc.Event) {
	if plusminus {

	} else {

	}
  message := strings.Split(event.Message()," ")

  var user string

  for _,val := range message {
    if strings.HasSuffix(val, "++") {
      user = strings.TrimRight(val, "++")
    }
  }

  fmt.Println(user)
  //
	// // Connect to database, https://labix.org/mgo
	// session, err := mgo.Dial("localhost:27017")
	// if err != nil {
	// 	panic(err)
	// }
  //
	// // Add connection close to the stack
	// defer session.Close()
  //
	// // Access the appropriate collection
	// c := session.DB("TwitchEmoji").C(strings.TrimPrefix(event.Arguments[0], "#"))
	// err = c.Insert(mes)
  //
	// if err != nil {
	// 	panic(err)
	// }
}