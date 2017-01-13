package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/thoj/go-ircevent"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	htmlFolder = "html/"

	filename = ".config/secret"
	// Twitch Variables
	bot_owner = "cmalloc"
	bot       = "cmalloc"
	channel   = "#cmalloc"
	server    = "irc.chat.twitch.tv"
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

// Generalized version of rendering code
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(htmlFolder + tmpl + ".html")
	t.Execute(w, p)
}

// Handles the index of our webserver
func indexHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/"):]

	// '_' handles any unhandled returned values. In this case, it's the 'err' return variable
	p, _ := loadPage(title)

	renderTemplate(w, "index", p)
}

func readOauthToken() string {
	creds, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Couldn't read credentials\n")
		panic(err.Error())
	}
	return string(creds)
}

func connectBot() {

}

func spawnBot() *Bot {
	return &Bot{serv: server + ":6667",
		port:    "6667",
		nick:    "cmallocbot",
		user:    "DefinitelyNotAndrew",
		chann:   "#herald_likem",
		auth:    readOauthToken(),
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

type Emoji struct {
	Id    bson.ObjectId `bson:"_id"`
	Emoji string        `bson:"emoji"`
	Count int           `bson:"count"`
}

func main() {
	// Connect to database, https://labix.org/mgo
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}

	defer session.Close()

	c := session.DB("TwitchEmoji").C("emoji")

	test := Emoji{Emoji: "JKanStyle", Count: 0}
	c.Insert(test)

	var emojis []Emoji
	c.Find(bson.M{}).All(&emojis)

	for _, val := range emojis {
		fmt.Println(val.Id, val.Emoji, val.Count)
	}

	ircBot := spawnBot()
	testConn := irc.IRC(ircBot.nick, ircBot.user)
	testConn.Password = ircBot.auth

	testConn.VerboseCallbackHandler = false
	testConn.Debug = false
	testConn.UseTLS = false
	testConn.AddCallback("001", func(e *irc.Event) { testConn.Join(ircBot.chann) })

	testConn.AddCallback("PRIVMSG", func(event *irc.Event) {
		fmt.Printf("%v:%v:%v\n", event.Arguments[0], event.User, event.Message())
		//if strings.Contains(event.Message(), "kappa") == true {
		//	testConn.Privmsg(ircBot.chann, "kappa")
		//}

		if strings.Contains(event.Message(), "!kojima") == true {
			testConn.Privmsg(ircBot.chann, "KOOOOOOOOOJIMA")
		}

		if strings.Contains(event.Message(), "!rick") == true {
			testConn.Privmsg(ircBot.chann, "Fuck off Kevin")
		}

		if strings.Contains(event.Message(), ":)") == true {
			// change := mgo.Change{
			// 	Update: bson.M{"$inc":bson.M{":)":1}},
			// 	ReturnNew: true,
			// }
		}
		//info, err := c.Find(M{"_id": id}).Apply(change, &doc)

		//event.Message() contains the message
		//event.Nick Contains the sender
		//event.Arguments[0] Contains the channel
	})

	// testConn.AddCallback("366", func(e *irc.Event) { })

	err = testConn.Connect(ircBot.serv)

	if err != nil {
		fmt.Printf("Err %s", err)
		return
	}
	// testConn.Privmsg(ircBot.chann, ircBot.message)

	testConn.Loop()

	// Register the handler
	//http.HandleFunc("/", indexHandler)

	// Wait for requests
	//http.ListenAndServe(":8080", nil)
}
