package main

import (
	"flag"
	"io/ioutil"
	"fmt"
)

const (
	// Twitch Variables
	server   = "irc.chat.twitch.tv:6667"
	database = "localhost:27017"
)

func readUserVariables(file string) string {
	value, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err.Error())
	}
	return string(value)
}

func main() {
	// Handle flags
	cmdUsername := flag.String("user", "defaultUsername", "Usage: go-twitch-bot -user <botname>")
	cmdChannel := flag.String("c", "cmalloc", "Usage: go-twitch-bot -c <channel>")
	cmdDb := flag.String("db", "go-twitch-bot", "Usage: go-twitch-bot -db <database>")
	cmdOauth := flag.String("auth", "000000000", "Usage: go-twitch-bot -auth <oauth-token")
	cmdDebug := flag.Bool("debug", false, "Usage: go-twitch-bot -debug")

	flag.Parse()

	username := *cmdUsername
	// Identical to user, twitch doesn't care about nicknames, but we need one
	nick := username
	channel := "#" + *cmdChannel
	db := *cmdDb
	oauth := *cmdOauth
	debugFlag := *cmdDebug

	fmt.Printf("Username:%v\n", username)
	fmt.Printf("Channel:%v\n", channel)
	fmt.Printf("Database:%v\n", db)
	fmt.Printf("OauthToken:%v\n", oauth)





	// user, nick, channel, debug
	botMain(username, nick, channel, oauth, db, debugFlag)
}
