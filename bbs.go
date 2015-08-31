package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
	// Install: go get github.com/mattn/go-sqlite3
	_ "github.com/mattn/go-sqlite3"
)

// URL: http://localhost:8080
const URL = ":8080"

type Message struct {
	Name    string
	Email   string
	Title   string
	Message string
	Created string
}

func db_connect() (*sql.DB, error) {
	fmt.Println("### Connect to DB ###")
	return sql.Open("sqlite3", "db/bbs.db")
}

func create_table(db *sql.DB) {
	fmt.Println("### Create Table ###")
	db.Exec(
		`CREATE TABLE IF NOT EXISTS message (
				message_id INTEGER PRIMARY KEY AUTOINCREMENT,
				name       VARCHAR(64) NULL,
				email      VARCHAR(64) NULL,
				title      VARCHAR(64) NULL,
				message    VARCHAR(64) NULL,
				created    VARCHAR(64) NULL
);`)
}

func get_message(db *sql.DB) (*sql.Rows, error) {
	fmt.Println("### Select from DB ###")
	return db.Query("SELECT * FROM message order by message_id desc")
}

func insert_message(db *sql.DB, r *http.Request) {
	fmt.Println("### Insert into DB ###")
	stmt, err := db.Prepare(
		`INSERT INTO message
		(name, email, title, message, created)
		values
		(?,?,?,?,?)`)
	checkErr(err)
	name := r.Form.Get("name")
	email := r.Form.Get("email")
	title := r.Form.Get("title")
	message := r.Form.Get("message")
	created := time.Now().Format("2006-01-02 15:04:05")
	_, err = stmt.Exec(name, email, title, message, created)
	checkErr(err)
}

func bbs_home(w http.ResponseWriter, r *http.Request) {
	// Check request method
	// "GET" when the page is displayed
	if r.Method == "GET" {
		// Connect to DB
		db, err := db_connect()
		checkErr(err)

		// Create Message table
		create_table(db)

		// Get messages from DB
		rows, err := get_message(db)
		checkErr(err)

		// Create Message Slice
		var messages []Message

		// Looping through Message Data
		for rows.Next() {
			id, name, email, title, message, created := 0, "", "", "", "", ""
			err = rows.Scan(&id, &name, &email, &title, &message, &created)
			checkErr(err)

			// Filling Message Slice
			messages = append(messages,
				Message{
					Name:    name,
					Email:   email,
					Title:   title,
					Message: message,
					Created: created})
		}

		t, _ := template.ParseFiles("template/bbs_home.tpl")
		t.Execute(w, messages)
	}
}

func post_message(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// print at server side
	fmt.Println("### Post Data ###")
	fmt.Println("Name:", template.HTMLEscapeString(r.Form.Get("name")))
	fmt.Println("Email:", template.HTMLEscapeString(r.Form.Get("email")))
	fmt.Println("Title:", template.HTMLEscapeString(r.Form.Get("title")))
	fmt.Println("Message:", template.HTMLEscapeString(r.Form.Get("message")))

	// Connect to DB
	db, err := db_connect()
	checkErr(err)
	insert_message(db, r)
	t, _ := template.ParseFiles("template/thank_you.tpl")
	t.Execute(w, nil)
}

func main() {
	// This is for images directory
	http.Handle("/", http.FileServer(http.Dir(".")))

	http.HandleFunc("/home", bbs_home)
	http.HandleFunc("/post_message", post_message)

	// Start HTTP server listening port 8080
	err := http.ListenAndServe(URL, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
