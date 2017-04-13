package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"fmt"
)

// This file contains any functions that directly call the database.

// Insert a message into Maria
func mariaDbMessageInsert(message Message, db string, collection string) error {
	conn, _ := sql.Open("mysql", "bot:@/" + db)
	defer conn.Close()

	// insert
	stmt, err := conn.Prepare("INSERT message SET channel=?,user=?,message=?,timestamp=?")
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(message.Channel, message.User, strings.TrimSpace(string(message.Message)), message.Time)


	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}