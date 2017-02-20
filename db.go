package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// This file contains any functions that directly call the database.

// Generic function to insert any object type into a collection
func dbInsert(object interface{}, db string, collection string) error {
	// Connect to database, https://labix.org/mgo
	// 'database' is a global const declared in main.go
	session, err := mgo.Dial(database)
	if err != nil {
		panic(err)
	}

	// Add connection close to function exit stack
	defer session.Close()

	// Access the appropriate collection
	c := session.DB(db).C(collection)
	err = c.Insert(object)

	if err != nil {
		return err
	}
	return nil
}


// User related functions
func dbFindUser(username string, db string, collection string) User {
	session, err := mgo.Dial(database)
	if err != nil {
		panic(err)
	}

	defer session.Close()

	user := User{}
	c := session.DB(db).C(collection)
	err = c.Find(bson.M{"username": username}).One(&user)

	if err != nil && err != mgo.ErrNotFound {
		panic(err)
	}

	return user
}

func dbFindAndUpdateUser(username string, db string, collection string) error {
	session, err := mgo.Dial(database)
	if err != nil {
		panic(err)
	}

	defer session.Close()

	user := User{}

	c := session.DB(db).C(collection)
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"kappacount": 1}},
		ReturnNew: false,
	}
	_, err = c.Find(bson.M{"username": username}).Apply(change, &user)

	if err != nil {
		panic(err)
	}
	return nil
}

func dbFindTotalKappa(db string, collection string) []User {
	session, err := mgo.Dial(database)
	if err != nil {
		panic(err)
	}

	defer session.Close()

	// have to use type matching document
	var kappas []User

	c := session.DB(db).C(collection)
	err = c.Find(nil).Select(bson.M{"kappacount": 1}).All(&kappas)

	if err != nil && err != mgo.ErrNotFound {
		panic(err)
	}

	return kappas
}
