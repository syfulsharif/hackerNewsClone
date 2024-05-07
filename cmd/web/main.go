package main

import (
	"log"
	"os"

	"github.com/CloudyKit/jet"
)

type application struct {
	appName string
	server  server
	debug   bool
	errLog  *log.Logger
	infoLog *log.Logger
	view    *jet.Set
}

type server struct {
	host string
	port string
	url  string
}

func main() {

	server := server{
		host: "localhost",
		port: "8009",
		url:  "http://localhost:8009",
	}
	app := &application{
		server:  server,
		appName: "HackerNews",
		debug:   true,
		infoLog: log.New(os.Stdout, "INFO\t", log.Ltime|log.Ldate|log.Lshortfile),
		errLog:  log.New(os.Stderr, "ERROR\t", log.Ltime|log.Ldate|log.Llongfile),
	}

	if err := app.listenAndServer(); err != nil {
		log.Fatal(err)
	}
}
