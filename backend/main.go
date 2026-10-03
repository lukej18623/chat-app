package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// upgrader to turn http request into websocket connection
var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,

	CheckOrigin: func(r *http.Request) bool { return true },
}


// listen for new messages being sent to websocket endpoint 
func reader(conn *websocket.Conn) {
	for {
	//read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
	// print out that message for clarity 
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}


// websocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	// upgrade to a websocket connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// listen for new messages on websocket connection
	reader(ws)
}


func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "simple server")
	})

	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Chat app v0.01")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}