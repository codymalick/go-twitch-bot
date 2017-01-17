package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"

	"strings"

	"github.com/thoj/go-ircevent"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// http://stackoverflow.com/questions/13342128/simple-golang-irc-bot-keeps-timing-out
// Great thanks to jwesonga for the code
type Bot struct {
	serv          string
	port          string
	nick          string
	user          string
	chann         string
	auth          string
	pread, pwrite chan string
	conn          net.Conn
	message       string
}

func (bot *Bot) spawnBot(channel string, authFile string) *Bot {
	return &Bot{serv: server + ":6667",
		port:    "6667",
		nick:    "cmallocbot",
		user:    "DefinitelyNotAndrew",
		chann:   channel,
		auth:    bot.readOauthToken(authFile),
		message: "Test message, please ignore"}
}




func (bot *Bot) Connect() (conn net.Conn, err error) {

	fmt.Printf("Using auth token:%v\n", bot.auth)

	fmt.Printf("Connecting to irc server: %v\n", server)
	connString := server + ":6667"
	connection, err := net.Dial("tcp", connString)
	if err != nil {
		fmt.Printf("Connecting is hard: %v\n", err.Error())
		log.Fatal("Unable to connect to IRC server ", err)
	}

	bot.conn = connection
	log.Printf("Connected to IRC server %s (%s)\n", bot.serv, bot.conn.RemoteAddr())
	return bot.conn, nil

}

func (bot *Bot) readOauthToken(file string) string {
	creds, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("Couldn't read credentials\n")
		panic(err.Error())
	}
	fmt.Println(string(creds))
	return string(creds)
}



func botMain(channel string) {


	// if *cmdDebug != debugFlag {
	// 	debugFlag = *cmdDebug
	// }

	// Connect to database, https://labix.org/mgo
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}

	// Add connection close to the stack
	defer session.Close()

	// Access the appropriate collection
	c := session.DB("TwitchEmoji").C("emoji")

	var emojis = new([]Emoji)
	c.Find(bson.M{}).All(emojis)

	// This isn't pretty, need to find the correct way to instantiate an object
	ircBot := new(Bot)
	ircBot = ircBot.spawnBot(channel, filename)
	testConn := irc.IRC(ircBot.nick, ircBot.user)
	testConn.Password = ircBot.auth

	testConn.VerboseCallbackHandler = false
	testConn.Debug = false
	testConn.UseTLS = false

	// Register event handlers
	testConn.AddCallback("001", func(e *irc.Event) { testConn.Join(ircBot.chann) })

	// Any message sent on the server, have to be followed on the server to respond
	testConn.AddCallback("PRIVMSG", func(event *irc.Event) {
		//event.Message() contains the message
		//event.Nick Contains the sender
		//event.Arguments[0] Contains the channel

		// spawn thread, record message
		go createMessage(event)

		fmt.Printf("%v:%v:%v\n", event.Arguments[0], event.User, event.Message())

		if strings.Contains(event.Message(), "william") || strings.Contains(event.Message(), "willie") {
      testConn.Privmsg(ircBot.chann, "Also, BILL!")
		}

		if strings.Contains(event.Message(), "!kojima") {
			testConn.Privmsg(ircBot.chann, "KOOOOOOOOOJIMA")
		}

		if strings.Contains(event.Message(), "!rick") {
			testConn.Privmsg(ircBot.chann, "Fuck off Kevin")
		}

		if strings.Contains(event.Message(), "!kevin") {
			testConn.Privmsg(ircBot.chann, "Metal Gear is not a stealth game")
		}

		if strings.Contains(event.Message(), "hey cmallocbot") {
			testConn.Privmsg(ircBot.chann, "B) Hello Dave")
		}

		if strings.Contains(event.Message(), "++") {
			go userPoints(true, *event)
		}

		if strings.Contains(event.Message(), "--") {
			go userPoints(false, *event)
		}

	})


	err = testConn.Connect(ircBot.serv)

	if err != nil {
		fmt.Printf("Err %s", err)
		return
	}

	testConn.Loop()

}
