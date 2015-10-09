package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	// Install Command: go get github.com/mattn/go-sqlite3
	_ "github.com/tomoyan/go_bbs/Godeps/_workspace/src/github.com/mattn/go-sqlite3"
)

const (
	// URL: http://localhost:8080/home
	PORT      = ":8080"
	DATE_TIME = "2006-01-02 15:04:05"
)

// Global Variables for DB connection
var db *sql.DB
var err error

func init() {
	initLog()

	// Connect to DB
	db, err = connectDB()
	checkErr(err)

	// Setting up DB
	// Create Message table
	createTable()
	deleteMessages()
	insertSampleMessages()
}

func main() {
	Debug.Println("# Start main")

	// images directory
	// HTTP requests with the contents of the file system
	http.Handle("/", http.FileServer(http.Dir(".")))

	http.HandleFunc("/home", bbsHome)
	http.HandleFunc("/post", postMessage)

	// Start HTTP server listening port 8080(Default)
	// if there is no PORT parameter
	if port := os.Getenv("PORT"); port == "" {
		err = http.ListenAndServe(PORT, nil)
	} else {
		err = http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	}

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
