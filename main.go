package main

import (
	"fmt"
	"net/http"

	"github.com/barokurniawan/websocket/context"
	"github.com/barokurniawan/websocket/controller"
)

func main() {
	context.LoadConfig()

	http.HandleFunc("/socket", controller.SocketHandler)
	fmt.Println("Server starting at " + context.Config.Address + context.Config.Port)
	http.ListenAndServe(context.Config.Address+context.Config.Port, nil)
}
