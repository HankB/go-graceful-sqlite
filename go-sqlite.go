package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	initDataBase("sqlite3", dbName)
	defer closeDataBase()
	fmt.Println("hit \"<ctrl>C\" now")
	time.Sleep(60 * time.Second)
}

var dbName = "test.db"

//var dbName = ":memory:"
var db *sql.DB
var err error

/* init the database or go up in flames
 */
func initDataBase(flavor, location string) {
	if location != ":memory:" {
		fmt.Println("Removing", location)
		os.Remove(location) // remove previous copy
	}

	// open the database
	db, err = sql.Open(flavor, location)
	if err != nil {
		fmt.Println("error opening", location)
		log.Fatal(err)
	}
	fmt.Println("opened", location)

	sqlCreateStmt := `create table data (value integer not null);`
	// create the table
	_, err = db.Exec(sqlCreateStmt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("created table \"data\"")
}

func closeDataBase() {
	fmt.Println("closing database")
	db.Close()
}
