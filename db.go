package main

import (
	"gopkg.in/mgo.v2"
)

func dbInsert(object interface{}, db string, collection string) error {
	// Connect to database, https://labix.org/mgo
	session, err := mgo.Dial(database)
	if err != nil {
		panic(err)
	}

	// Add connection close to the stack
	defer session.Close()

	// Access the appropriate collection
	c := session.DB(db).C(collection)
	err = c.Insert(object)

	if err != nil {
		return err
	}
	return nil
}
