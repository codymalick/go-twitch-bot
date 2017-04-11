package main

import (
	"fmt"
	"strings"

	"github.com/thoj/go-ircevent"
)

type Bot struct {
	User string
	Channel string
	Db string
	Connection *irc.Connection
	Voting chan int
}

func (bot *Bot) registerEvents() {

	// Register event handlers
	bot.Connection.AddCallback("001", func(e *irc.Event) { bot.Connection.Join(bot.Channel) })

	// Any message sent on the server, have to be followed on the server to respond
	bot.Connection.AddCallback("PRIVMSG", func(event *irc.Event) {
		//event.Message() contains the message
		//event.Nick Contains the sender
		//event.Arguments[0] Contains the channel

		// spawn thread, record message
		go createMessage(event, bot.Db)

		fmt.Printf("%v:%v:%v\n", event.Arguments[0], event.User, event.Message())

		// Example response to "hey cmallocbot"
		if strings.Contains(event.Message(), "hey " + bot.User) {
			bot.Connection.Privmsg(bot.Channel, "B) Hello Dave")
		}

		// Kappa counter
		if strings.Contains(event.Message(), "KappaHD ") ||
			strings.Contains(event.Message(), "MiniK ") ||
			strings.Contains(event.Message(), "Kappa ") ||
			strings.Contains(event.Message(), "kappa") {

			go incrementKappa(event.User, bot.Db)

		}

		if strings.Contains(event.Message(), "!mykappa") {
			go kappaCounter(event.User, bot.Db, bot.Connection, bot.Channel)
		}
		if strings.Contains(event.Message(), "!globalkappa") {
			go globalKappaCounter(event.User, bot.Db, bot.Connection, bot.Channel)
		}
		if strings.Contains(event.Message(), "!vote") {
			go startVote(event.User, bot.Db, bot.Connection, bot.Channel, event.Message(), bot.Voting)
		}
		if strings.Contains(event.Message(), "!cast") {
			go castVote(event.User, bot.Db, bot.Connection, bot.Channel, event.Message())
		}
	})
}

func botMain(user string, nick string, channel string, oauth string, db string,
	debug bool) {

	bot := Bot{
		User: user,
		Channel: channel,
		Db: db,
		Connection: irc.IRC(nick,user),
	}
	bot.Connection.Password = oauth

	bot.Connection.VerboseCallbackHandler = false
	bot.Connection.Debug = debug
	bot.Connection.UseTLS = false

	bot.registerEvents()

	err := bot.Connection.Connect(server)

	if err != nil {
		fmt.Printf("Err %s", err)
		return
	}

	bot.Connection.Loop()
}
