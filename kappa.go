package main

import (
	"github.com/thoj/go-ircevent"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

func kappaCounter(username string, db string, connection *irc.Connection, channel string) {
	user := dbFindUser(username, db, userCollection)
	if len(user.Username) == 0 {
		// Add the user!
		newUser := User{bson.NewObjectId(), username, 0}
		dbInsert(newUser, db, userCollection)
		message := username + "! You've been add to kappa counter! Your kappa count is 0"
		connection.Privmsg(channel, message)
	} else {
		kappaCount := strconv.FormatInt(user.KappaCount, 10)
		message := username + "! Your kappa count is " + kappaCount
		connection.Privmsg(channel, message)
	}
}

func globalKappaCounter(username string, db string, connection *irc.Connection, channel string) {
	kappas := dbFindTotalKappa(db, userCollection)
	var total int64

	for _, value := range kappas {
		total += value.KappaCount
	}

	// Tell the user
	stotal := strconv.FormatInt(total, 10)
	message := username + "! The global kappa count is " + stotal
	connection.Privmsg(channel, message)
}
