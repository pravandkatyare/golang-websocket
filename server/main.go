package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8181", nil))
}

func setupRoutes() {
	http.HandleFunc("/", landing)

}

func landing(w http.ResponseWriter, r *http.Request) {
	websocket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Websocket Connected")
	listen(websocket)
}

func listen(conn *websocket.Conn) {
	for {
		messageType, messageContent, err := conn.ReadMessage()
		timeReceived := time.Now()
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(string(messageContent))
		messageResponse := fmt.Sprintf("Message: %s. Received at %v.", messageContent, timeReceived)
		if err := conn.WriteMessage(messageType, []byte(messageResponse)); err != nil {
			log.Println(err)
			return
		}
	}
}
