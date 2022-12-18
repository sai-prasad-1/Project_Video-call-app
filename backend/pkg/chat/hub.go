package chat

type Hub struct {
	// Registered connections.
	Clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}
