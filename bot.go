package main

import(
  	"net"
    "fmt"
    "log"
    "io/ioutil"
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
