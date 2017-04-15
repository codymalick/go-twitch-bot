package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"fmt"
	"log"
)

// This file contains any functions that directly call the database.


func mariaFindUser(username string, db string) *User {
	var result User
	// TODO: Make db user a parameter
  conn, _ := sql.Open("mysql", "bot:@/" + db)
	defer conn.Close()

	// select
  rows, err := conn.Query("SELECT * FROM user WHERE username=\"" + username + "\"")
	if err != nil {
		panic(err)
	}

  for rows.Next() {
      var id int64
      var user string
      err = rows.Scan(&id, &user)
      if err != nil {
          log.Fatal(err)
      }
      result = User{id,user}
  }
	return &result
}

// This function will check for an existing user, and create a new one if none
// are found.
func mariaCheckOrInsert(username string, db string) *User {
	result := mariaFindUser(username, db)
	if result.Id != 0 {
		fmt.Println("User exists")
		return result
	}
	mariaAddUser(username, db)
	// We need the db generated id
	// TODO: track last id in state
	result = mariaFindUser(username, db)

	return result
}

func mariaAddUser(username string, db string) error {
	conn, _ := sql.Open("mysql", "bot:@/" + db)
	defer conn.Close()

	// select
	// insert
    stmt, err := conn.Prepare("INSERT user SET username=?")
    if err != nil {
  	  panic(err)
    }

    _, err = stmt.Exec(username)


    if err != nil {
  	  fmt.Println(err.Error())
  	  return err
    }
	return nil
}


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