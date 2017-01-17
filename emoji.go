package main

import (
	"gopkg.in/mgo.v2/bson"
	// "io/ioutil"
	// "fmt"
	// "strings"
)

type Emoji struct {
	Id    bson.ObjectId `bson:"_id"`
	Emoji string        `bson:"emoji,omitempty"`
	Count int           `bson:"count"`
}

// func readEmojiFile(file string) map[string]int {
// 	emojis, err := ioutil.ReadFile(file)
// 	if err != nil {
// 		panic(err)
// 	}
// 	emojisList := strings.Split(string(emojis), "\n")
// 	emojisMap := make(map[string]int)
// 	for _,val := range emojisList {
// 		emojisMap[val] = 0
// 	}
//
// 	for key, value := range emojisMap {
// 		fmt.Printf("Key: %v | Value: %v\n",key, value)
// 	}
// 	fmt.Printf("%v\n", emojisMap)
// 	return emojisMap
// }