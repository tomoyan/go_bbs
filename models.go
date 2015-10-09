package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strings"
	"time"

	// Install Command: go get github.com/mattn/go-sqlite3
	_ "github.com/tomoyan/go_bbs/Godeps/_workspace/src/github.com/mattn/go-sqlite3"
)

func connectDB() (*sql.DB, error) {
	//fmt.Println("### Connect to DB ###")
	Debug.Println("# Start connectDB")
	Debug.Println("# Connect to DB")
	return sql.Open("sqlite3", "db/bbs.db")
}

func createTable() {
	//fmt.Println("### Create Table ###")
	Debug.Println("# Start createTable")
	Debug.Println("# Create Table")
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

func deleteMessages() {
	db.Exec(
		`DELETE FROM message;`)
}

func insertSampleMessages() {
	stmt, err := db.Prepare(
		`INSERT INTO message
		(name, email, title, message, created)
		values
		(?,?,?,?,?)`)
	checkErr(err)

	name := "Go掲示板のテスト名"
	email := "go-bbs@jp.go"
	title := "Go掲示板のテストタイトル"
	message := "Go掲示板のてすとメッセージ!!!"
	created := time.Now().Format(DATE_TIME)
	_, err = stmt.Exec(name, email, title, message, created)
	checkErr(err)

}
func getMessage() (*sql.Rows, error) {
	//fmt.Println("### Select from DB ###")
	Debug.Println("# Start getMessage")
	Debug.Println("# Select from DB")
	return db.Query("SELECT * FROM message order by message_id desc")
}

func insertMessage(r *http.Request) {
	//fmt.Println("### Insert into DB ###")
	Debug.Println("# Start insertMessage")
	Debug.Println("# Insert into DB")
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
	created := time.Now().Format(DATE_TIME)
	_, err = stmt.Exec(name, email, title, message, created)
	checkErr(err)
}

// Form Input Validation
func (msg *PostMessage) validate() bool {
	Debug.Println("# Start validate")
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

func render(w http.ResponseWriter, filename string, data interface{}) {
	Debug.Println("# Start render")
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
