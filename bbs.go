package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
	// Install Command: go get github.com/mattn/go-sqlite3
	_ "github.com/mattn/go-sqlite3"
)

// URL: http://localhost:8080/home
const URL = ":8080"

// Global Variables for DB connection
var db *sql.DB
var err error

type Message struct {
	Name    string
	Email   string
	Title   string
	Message string
	Created string
}

type PostMessage struct {
	PostName    string
	PostEmail   string
	PostTitle   string
	PostMessage string
	PostErrors  map[string]string
}

func (msg *PostMessage) Validate() bool {
	msg.PostErrors = make(map[string]string)

	if strings.TrimSpace(msg.PostName) == "" {
		fmt.Println("Empty Name", msg.PostName)
		msg.PostErrors["PostName"] = "Please write a name"
	}

	re := regexp.MustCompile(".+@.+\\..+")
	matched := re.Match([]byte(msg.PostEmail))
	if matched == false {
		msg.PostErrors["PostEmail"] = "Please enter a valid email address"
	}

	if strings.TrimSpace(msg.PostTitle) == "" {
		msg.PostErrors["PostTitle"] = "Please write a title"
	}

	if strings.TrimSpace(msg.PostMessage) == "" {
		msg.PostErrors["PostMessage"] = "Please write a message"
	}

	return len(msg.PostErrors) == 0
}

func connectDB() (*sql.DB, error) {
	fmt.Println("### Connect to DB ###")
	return sql.Open("sqlite3", "db/bbs.db")
}

func createTable(db *sql.DB) {
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

func getMessage(db *sql.DB) (*sql.Rows, error) {
	fmt.Println("### Select from DB ###")
	return db.Query("SELECT * FROM message order by message_id desc")
}

func insertMessage(db *sql.DB, r *http.Request) {
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

func bbsHome(w http.ResponseWriter, r *http.Request) {
	// Check request method
	// "GET" when the page is displayed
	if r.Method == "GET" {
		// Create Message table
		createTable(db)

		// Get messages from DB
		rows, err := getMessage(db)
		checkErr(err)

		// Create Message Slice
		var messages []Message

		// Loop through Message Data
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

		// Parse template file
		t, err := template.ParseFiles("template/bbs_home.tpl")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = t.Execute(w, messages)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func postMessage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// print at server side
	fmt.Println("### Post Data ###")
	fmt.Println("Name:", template.HTMLEscapeString(r.Form.Get("name")))
	fmt.Println("Email:", template.HTMLEscapeString(r.Form.Get("email")))
	fmt.Println("Title:", template.HTMLEscapeString(r.Form.Get("title")))
	fmt.Println("Message:", template.HTMLEscapeString(r.Form.Get("message")))

	// PostMessage for Input Validation
	msg := &PostMessage{
		PostName:    r.Form.Get("name"),
		PostEmail:   r.Form.Get("email"),
		PostTitle:   r.Form.Get("title"),
		PostMessage: r.Form.Get("message"),
	}

	// Input Validation
	if msg.Validate() == false {
		render(w, "template/post_message.tpl", msg)
		return
	}

	insertMessage(db, r)

	// Parse template file
	t, err := template.ParseFiles("template/thank_you.tpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	// Connect to DB
	db, err = connectDB()
	checkErr(err)

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

func render(w http.ResponseWriter, filename string, data interface{}) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
