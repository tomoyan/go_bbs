package main

// Message type for messages in DB
type Message struct {
	Name    string
	Email   string
	Title   string
	Message string
	Created string
}

// PostMessage typ for Posted Message from Form
type PostMessage struct {
	PostName    string
	PostEmail   string
	PostTitle   string
	PostMessage string
	PostErrors  map[string]string
}
