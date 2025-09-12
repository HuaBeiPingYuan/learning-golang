package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

//Reading: handleConnections per client → sends to channel.

// Writing: handleMessages → reads from channel → broadcasts to all clients.

// Channel: connects reading and writing goroutines safely.

// Mutex: protects shared clients map during writes.

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
		//        return r.Header.Get("Origin") == "https://yourdomain.com"
	},
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan string)

//This creates a new mutex (mutual exclusion lock) and stores a pointer to it in the variable mutex.
//sync.Mutex is a type provided by Go’s standard library for protecting shared data across goroutines.

var mutex = &sync.Mutex{}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websockets.html")
	})
	http.HandleFunc("/ws", handleConnections)
	go handleMessages()

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	//This *websocket.Conn represents a single, specific client connection. The connection opens immediately with:
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	//when this function (handleConnections) ends, close this WebSocket.
	defer ws.Close()

	mutex.Lock()
	clients[ws] = true
	mutex.Unlock()

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			// client disconnected
			fmt.Println("read error:", err)
			mutex.Lock()
			delete(clients, ws)
			mutex.Unlock()
			break
		}

		broadcast <- string(msg)
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		mutex.Lock()
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				fmt.Println("write error:", err)
				client.Close()
				delete(clients, client)
			}
		}
		mutex.Unlock()
	}
}
