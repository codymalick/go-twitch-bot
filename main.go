package main

import (
	"flag"
	"io/ioutil"
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

	user := readUserVariables("config/user")

	// Identical to user, twitch doesn't care about nicknames, but we need one
	nick := user

	oauth := readUserVariables("config/secret")
	channel := "#cmalloc"
	debugFlag := false

	db := "go-twitch-bot"

	// Handle channel flag
	cmdChannel := flag.String("c", channel, "Usage: TwitchEmoji [<channel>]")

	// cmdDebug := flag.String("d", channel, "Usage")
	flag.Parse()

	*cmdChannel = "#" + *cmdChannel

	if *cmdChannel != channel {
		channel = *cmdChannel
	}

	// if cmdDebug != nil  {
	// 	debugFlag = true
	// }

	// user, nick, channel, debug
	botMain(user, nick, channel, oauth, db, debugFlag)
}
