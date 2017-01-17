package main

import(
"fmt"
//"html/template"
//"net/http"
// "strings"
"flag"

// "github.com/thoj/go-ircevent"
// "gopkg.in/mgo.v2"
// "gopkg.in/mgo.v2/bson"
)
const (
	htmlFolder = "html/"

	filename = ".config/secret"
	// Twitch Variables
	bot_owner = "cmalloc"
	bot       = "cmalloc"
	server    = "irc.chat.twitch.tv"
	emojifile = "emojiList"
)

func main() {

	channel := "#cmalloc"
	//debugFlag := false

	// Handle channel flag
	cmdChannel := flag.String("c", channel, "Usage: TwitchEmoji [<channel>]")
	flag.Parse()

	*cmdChannel = "#" + *cmdChannel
	fmt.Printf("Flag: %v\n", *cmdChannel)

	if *cmdChannel != channel {
		channel = *cmdChannel
	}
	botMain(channel)
}
