package sockethandler

import (
	"fmt"
	"log"
	"strings"

	"github.com/barokurniawan/websocket/tructure"
	"github.com/novalagung/gubrak"
)

//Connections a bucket for every connection
var Connections = make([]*tructure.WebSocketConnection, 0)

//HandleIO handler input and output
func HandleIO(currentConn *tructure.WebSocketConnection, Connections []*tructure.WebSocketConnection) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR", fmt.Sprintf("%v", r))
		}
	}()

	for {
		payload := tructure.SocketPayload{}
		err := currentConn.ReadJSON(&payload)
		if err != nil {
			if strings.Contains(err.Error(), "websocket: close") {
				ejectConnection(currentConn)
				return
			}

			log.Println("ERROR", err.Error())
			continue
		}

		broadcastMessage(currentConn, payload.Message)
	}
}

//broadcastMessage sent message through every connection
func broadcastMessage(currentConn *tructure.WebSocketConnection, message string) {
	for _, eachConn := range Connections {
		if eachConn == currentConn {
			continue
		}

		//just send message to the same channel
		if eachConn.Channel == currentConn.Channel {
			eachConn.WriteJSON(tructure.SocketResponse{
				Message: message,
			})
		}
	}
}

//ejectConnection there is some error or unwanted connection? eject!!!
func ejectConnection(currentConn *tructure.WebSocketConnection) {
	filtered, _ := gubrak.Reject(Connections, func(each *tructure.WebSocketConnection) bool {
		return each == currentConn
	})

	Connections = filtered.([]*tructure.WebSocketConnection)
}
