package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/barokurniawan/websocket/config"
	"github.com/barokurniawan/websocket/helper"
	"github.com/barokurniawan/websocket/sockethandler"
	"github.com/barokurniawan/websocket/structure"
	"github.com/gorilla/websocket"
)

func main() {
	AppConfig := config.App{
		Port:    ":3001",
		Address: "127.0.0.1",
	}

	allowedOrigin := []string{"localhost:3001", "103.93.161.49:3100"}

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
		if r.Header.Get("Origin") == "" {
			http.Error(w, "Invalid origin!", http.StatusBadRequest)
			return
		}

		originHost := strings.Split(r.Header.Get("Origin"), "://")
		host := originHost[1]
		if exists, _ := helper.InArray(host, allowedOrigin); exists == false {
			http.Error(w, "Invalid origin!", http.StatusForbidden)
			return
		}

		channel, ok := r.URL.Query()["channel"]
		if !ok || len(channel[0]) < 1 {
			http.Error(w, "missing chanel parameter", http.StatusBadRequest)
			return
		}

		currentGorillaConn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
		if err != nil {
			http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		}

		currentConn := structure.WebSocketConnection{Conn: currentGorillaConn, Channel: string(channel[0])}
		sockethandler.Connections = append(sockethandler.Connections, &currentConn)

		go sockethandler.HandleIO(&currentConn, sockethandler.Connections)
	})

	fmt.Println("Server starting at " + AppConfig.Address + AppConfig.Port)
	http.ListenAndServe(AppConfig.Address+AppConfig.Port, nil)
}
