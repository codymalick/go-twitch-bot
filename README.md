[![Go Report Card](https://goreportcard.com/badge/github.com/codymalick/go-twitch-bot)](https://goreportcard.com/report/github.com/codymalick/go-twitch-bot) 
[![Build Status](https://travis-ci.org/codymalick/go-twitch-bot.svg?branch=master)](https://travis-ci.org/codymalick/go-twitch-bot)
# Go Twitch Bot
go-twitch-bot is an open source bot written in Go, designed to fill basic needs for streamers, as well as add fun features to the stream such as voting, emoji tracking, and more.

Current complete features:
- Chat log
- Per user kappa emoji use
- Custom command response (example: "!mykappa" returns the number of times the user has said kappa or used any of the kappa emojis)

Current in development features include:
- Add CLI flags to change username/oauth token
- Proper Oauth2 integration
- Upvoting users
- Tracking emoji use
- Adding greeting messages
- Message of the day
- Add db credentials

If you're interesting in contributing, feel free to take a look at the issues list, or open new issues if you've found any bugs. 

# Compilation and Use:
At the moment, there are no pre-compiled binaries, but Go is cross-platform. You can download this repo and compile it with
`go build` or `go install go-twitch-bot`

Dependencies:
- Golang 1.7
- Mongodb


This software is released as-is under the MIT license. See the LICENSE file for more information


# Useful Links
Twitch Oauth2 generator: http://twitchapps.com/tmi/
