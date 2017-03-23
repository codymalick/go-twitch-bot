package main

import (
	"github.com/thoj/go-ircevent"
	"strings"
)

//vote usage !vote name topic1 topic2 topic3
func startVote(username string, db string, connection *irc.Connection, channel string, message string) {
	//if no topics do nothing
	var voteTopics []string
	voteTopics = strings.Split(message, " ")
	if len(voteTopics) <= 2 {
		return
	}

	//name vote
	//var voteName string = voteTopics[1]

	//store topics
	//loop through the rest of the topics and store them
}

//Each user can cast a vote for an ongoing vote.
func castVote(username string, db string, connection *irc.Connection, channel string, message string) {
	var voteInfo []string = strings.Split(message, " ")
	if len(voteInfo) <= 1 {
		return //we didn't get a thing to vote for
	}
	//check if a vote is happening for user at channel. if it is{
	//store voteInfo[2] for the vote with name voteInfo[1] for that user. users can only have one vote
	//}
}

//check how many votes have been cast for that user with that channel. Needs a name as well
func checkVotes(username string, db string, connection *irc.Connection, channel string, message string) {
	return
}

//stop the voting for a certain user and channel with that name, assuming that the vote is still going.
func stopVote(username string, db string, connection *irc.Connection, channel string, message string) {
	return
}
