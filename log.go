package main

import (
	"log"
	"os"
)

// Global Variables for error log
var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
	Debug   *log.Logger
)

func initLog() {
	logFile := "log/debug.log"
	os.Create(logFile)
	debugFile, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open debugFile :", err.Error())
	}

	Debug = log.New(debugFile,
		"[Debug] ",
		log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
}
