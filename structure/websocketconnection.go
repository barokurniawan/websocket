package structure

import "github.com/gorilla/websocket"

type WebSocketConnection struct {
	*websocket.Conn
	Channel string
}
