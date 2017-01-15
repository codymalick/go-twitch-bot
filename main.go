package main

import (
	"fmt"
	//"html/template"
	//"net/http"
	"strings"
	"flag"

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
	server    = "irc.chat.twitch.tv"
)




// Generalized version of rendering code
// func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
// 	t, _ := template.ParseFiles(htmlFolder + tmpl + ".html")
// 	t.Execute(w, p)
// }

// Handles the index of our webserver
// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	title := r.URL.Path[len("/"):]
//
// 	// '_' handles any unhandled returned values. In this case, it's the 'err' return variable
// 	p, _ := loadPage(title)
//
// 	renderTemplate(w, "index", p)
// }



func connectBot() {

}

func main() {
	channel := "#cmalloc"
	//debugFlag := false

	// Handle channel flag
	cmdChannel := flag.String("c", channel, "Usage: TwitchEmoji [<channel>]")
	//cmdDebug := flag.Bool("d", debugFlag, "[-d]")
	flag.Parse()

	*cmdChannel = "#" + *cmdChannel
	fmt.Printf("Flag: %v\n", *cmdChannel)

	if *cmdChannel != channel {
		channel = *cmdChannel
	}

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
	//
	// test := Emoji{Id: bson.NewObjectId(), Emoji: "sodaNOPE", Count: 10}
	// err = c.Insert(test)
	// if err != nil {
	// 	panic(err)
	// }

	var emojis = new([]Emoji)
	c.Find(bson.M{}).All(emojis)


	// emojiList := []*Emoji{}

	// Note: the 'range' operator creates a copy of the data to iterate over. You
	// cannot alter data inside of that loop unless you use pointers to each
	// object
	// for _, val := range *emojis {
	// 	fmt.Println(val.Id, val.Emoji, val.Count)
	// 	emojiList = append(emojiList, &val)
	// }
	//
	// for _, val := range emojiList {
	// 	fmt.Println(val.Id, val.Emoji, val.Count)
	// }

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

	messages := 0
	// Any message sent on the server
	testConn.AddCallback("PRIVMSG", func(event *irc.Event) {
		messages++
		//event.Message() contains the message
		//event.Nick Contains the sender
		//event.Arguments[0] Contains the channel

		// spawn thread, record message
		go createMessage(event)
		fmt.Printf("%v:%v:%v\n", event.Arguments[0], event.User, event.Message())

		if event.User == "Atrum_Cordis" {
			fmt.Printf("BILL!")
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

		// Do this last
		// if messages > 25 {
		// 	for _, val := range *emojis {
		// 		fmt.Println(val.Emoji, val.Count)
		// 	}
		// 	messages = 0
		// }
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
