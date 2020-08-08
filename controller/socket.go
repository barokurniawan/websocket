package controller

import (
	"net/http"
	"strings"

	"github.com/barokurniawan/websocket/context"
	"github.com/barokurniawan/websocket/helper"
	"github.com/barokurniawan/websocket/sockethandler"
	"github.com/barokurniawan/websocket/structure"
	"github.com/gorilla/websocket"
)

func SocketHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Origin") == "" {
		http.Error(w, "Invalid origin!", http.StatusBadRequest)
		return
	}

	originHost := strings.Split(r.Header.Get("Origin"), "://")
	host := originHost[1]
	if exists, _ := helper.InArray(host, context.Config.AllowedOrigins); exists == false {
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
}
