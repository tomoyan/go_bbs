package main

import (
	"html/template"
	"net/http"
)

func bbsHome(w http.ResponseWriter, r *http.Request) {
	Debug.Println("# Start bbsHome")
	// Check request method
	// "GET" when the page is displayed
	if r.Method == "GET" {
		// Get messages from DB
		rows, err := getMessage()
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
	Debug.Println("# Start postMessage")
	r.ParseForm()

	// print at server side
	//fmt.Println("### Post Data ###")
	//fmt.Println("Name:", template.HTMLEscapeString(r.Form.Get("name")))
	//fmt.Println("Email:", template.HTMLEscapeString(r.Form.Get("email")))
	//fmt.Println("Title:", template.HTMLEscapeString(r.Form.Get("title")))
	//fmt.Println("Message:", template.HTMLEscapeString(r.Form.Get("message")))

	Debug.Println("# Post Data")
	Debug.Printf(
		"# Name:%v Email:%v Title:%v Message:%v",
		template.HTMLEscapeString(r.Form.Get("name")),
		template.HTMLEscapeString(r.Form.Get("email")),
		template.HTMLEscapeString(r.Form.Get("title")),
		template.HTMLEscapeString(r.Form.Get("message")))

	// PostMessage for Input Validation
	msg := &PostMessage{
		PostName:    r.Form.Get("name"),
		PostEmail:   r.Form.Get("email"),
		PostTitle:   r.Form.Get("title"),
		PostMessage: r.Form.Get("message"),
	}

	// Input Validation
	if msg.validate() == false {
		render(w, "template/post_message.tpl", msg)
		return
	}

	insertMessage(r)

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
