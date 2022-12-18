package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket"
	guuid "github.com/google/uuid"
	w "github.com/sai-prasad-1/Project_Video-call-app/pkg/webrtc"
)

func RoomCreate(c *fiber.Ctx) error {
	return c.Redirect("/room/" + guuid.New().String())
}

func Room(c *fiber.Ctx) error {
	uuid := c.Params("id")
	if uuid == "" {
		return c.Redirect("/room/create")
	}
	_, _, room := createOrGetRoom(uuid)
	return c.Render("room", fiber.Map{
		"uuid": uuid,
		"room": room,
	})
}

func RoomWS(c *websocket.Conn) error {
	uuid := c.Params("id")
	if uuid == "" {
		return c.Close()
	}
	Uuid, Suuid, room := createOrGetRoom(uuid)
	w.RoomConnection(c, room.Peers)
	return c.Close()
}

func createOrGetRoom(uuid string) (string, string, *w.Room) {

}

func RoomViewerWS(c *websocket.Conn) error {
	uuid := c.Params("id")
	if uuid == "" {
		return c.Close()
	}
	// Uuid, Suuid, room := createOrGetRoom(uuid)
	return c.Close()
}

func RoomViewer(c *fiber.Ctx, p *w.Peers) error {
	uuid := c.Params("id")
	if uuid == "" {
		return c.Redirect("/room/create")
	}
	_, _, room := createOrGetRoom(uuid)
	return c.Render("room_viewer", fiber.Map{
		"uuid": uuid,
		"room": room,
	})
}

type websocketMessage struct {
	Event string `json:"event"`
	Data  string `json:"data"`
}
