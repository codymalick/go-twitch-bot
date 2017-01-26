package main

import (
	"fmt"
	"strings"

	"github.com/thoj/go-ircevent"
)

func registerEvents(connection *irc.Connection, channel string, db string, user string) {

	// Register event handlers
	connection.AddCallback("001", func(e *irc.Event) { connection.Join(channel) })

	// Any message sent on the server, have to be followed on the server to respond
	connection.AddCallback("PRIVMSG", func(event *irc.Event) {
		//event.Message() contains the message
		//event.Nick Contains the sender
		//event.Arguments[0] Contains the channel

		// spawn thread, record message
		go createMessage(event, db)

		fmt.Printf("%v:%v:%v\n", event.Arguments[0], event.User, event.Message())

		if strings.Contains(event.Message(), "william") || strings.Contains(event.Message(), "willie") {
			connection.Privmsg(channel, "Also, BILL!")
		}

		if strings.Contains(event.Message(), "!kojima") {
			connection.Privmsg(channel, "KOOOOOOOOOJIMA")
		}

		if strings.Contains(event.Message(), "!kevin") {
			connection.Privmsg(channel, "Metal Gear is not a stealth game")
		}

		if strings.Contains(event.Message(), "hey "+user) {
			connection.Privmsg(channel, "B) Hello Dave")
		}
		if strings.Contains(event.Message(), "KappaHD ") ||
			strings.Contains(event.Message(), "MiniK ") ||
			strings.Contains(event.Message(), "Kappa ") ||
			strings.Contains(event.Message(), "kappa") {

			go incrementKappa(event.User, db)

		}

		if strings.Contains(event.Message(), "!mykappa") {
			go kappaCounter(event.User, db, connection, channel)
		}
		if strings.Contains(event.Message(), "!globalkappa") {
			go globalKappaCounter(event.User, db, connection, channel)
		}

	})

}

func botMain(user string, nick string, channel string, oauth string, db string,
	debug bool) {

	connection := irc.IRC(nick, user)
	connection.Password = oauth

	connection.VerboseCallbackHandler = false
	connection.Debug = debug
	connection.UseTLS = false

	registerEvents(connection, channel, db, user)

	err := connection.Connect(server)

	if err != nil {
		fmt.Printf("Err %s", err)
		return
	}

	connection.Loop()
}
