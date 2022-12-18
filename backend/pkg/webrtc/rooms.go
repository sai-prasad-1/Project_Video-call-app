package webrtc

import (
	"sync"

	"github.com/gofiber/websocket"
	"github.com/sai-prasad-1/Project_Video-call-app/pkg/chat"
)

type Room struct {
	Peers Peers
	Hub   *chat.Hub
}

func RoomConnection(c *websocket.Conn, peers *Peers) error {
	var config websocket.Config
	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		return err
	}
	newPeer := PeerConnectionState{
		Connection: peerConnection,
		WebSocket:  &ThreadSafeWriter{},
		Conn:       c,
		Mutex:      sync.Mutex{},
	}
	return nil
}
