package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	// Install Command: go get github.com/mattn/go-sqlite3
	_ "github.com/mattn/go-sqlite3"
)

// URL: http://localhost:8080/home
const URL = ":8080"

const (
	DATE_TIME = "2006-01-02 15:04:05"
)

// Global Variables for DB connection
var db *sql.DB
var err error

// Global Variables for error log
var (
	//Trace   *log.Logger
	//Info    *log.Logger
	//Warning *log.Logger
	//Error   *log.Logger
	Debug *log.Logger
)

func init() {
	debugFile, err := os.OpenFile("log/debug.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open traceFile :", err.Error())
	}

	Debug = log.New(debugFile,
		"[Debug] ",
		log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

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

	// Start HTTP server listening port 8080
	err := http.ListenAndServe(URL, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
