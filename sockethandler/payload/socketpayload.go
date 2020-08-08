package payload

type SocketPayload struct {
	Message          string
	Channel          string
	IsPrivateMessage bool
}
