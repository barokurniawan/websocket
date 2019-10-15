package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/barokurniawan/websocket/sockethandler"
	"github.com/barokurniawan/websocket/tructure"
	"github.com/gorilla/websocket"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		content, err := ioutil.ReadFile("index.html")
		if err != nil {
			http.Error(w, "Could not open requested file", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "%s", content)
	})

	http.HandleFunc("/other", func(w http.ResponseWriter, r *http.Request) {
		content, err := ioutil.ReadFile("other.html")
		if err != nil {
			http.Error(w, "Could not open requested file", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "%s", content)
	})

	http.HandleFunc("/socket", func(w http.ResponseWriter, r *http.Request) {
		channel, ok := r.URL.Query()["channel"]
		if !ok || len(channel[0]) < 1 {
			http.Error(w, "missing changel parameter", http.StatusBadRequest)
			return
		}

		currentGorillaConn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
		if err != nil {
			http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		}

		currentConn := tructure.WebSocketConnection{Conn: currentGorillaConn, Channel: string(channel[0])}
		sockethandler.Connections = append(sockethandler.Connections, &currentConn)

		go sockethandler.HandleIO(&currentConn, sockethandler.Connections)
	})

	fmt.Println("Server starting at :8080")
	http.ListenAndServe(":8080", nil)
}
