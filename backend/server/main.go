package server

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"github.com/gofiber/websocket"
	"github.com/sai-prasad-1/Project_Video-call-app/handlers"
	w "github.com/sai-prasad-1/Project_Video-call-app/pkg/webrtc"
)

var (
	addr = flag.String("addr", ":"+os.Getenv("PORT"), "")
	cert = flag.String("cert", "", "TLS certificate file")
	key  = flag.String("key", "", "TLS key file")
)

func Run() error {

	flag.Parse()
	if *addr == "" {
		*addr = ":8080"
	}

	engine := html.New("./views", ".html")
	// Create a new server
	app := fiber.New(fiber.Config{Views: engine})
	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/", handlers.Welcome)
	app.Get("/room/create", handlers.RoomCreate)
	app.Get("/room/:id", handlers.Room)
	app.Get("/room/:id/ws", websocket.New(handlers.RoomWS, websocket.Config{
		HandshakeTimeout: 10 * time.Second,
	}))
	// app.Get("/room/:id/chat", handlers.RoomChat)
	app.Get("/room/:id/chat/ws", websocket.New(handlers.RoomChatWS, websocket.Config{}))
	// app.Get("/room/:id/viewers", handlers.RoomViewers)
	// app.Get("/stream/:id", handlers.Stream)
	app.Get("/stream/:id/ws", websocket.New(handlers.StreamWS, websocket.Config{}))
	app.Get("/stream/:id/chat/ws", websocket.New(handlers.StreamChatWS, websocket.Config{}))
	app.Get("/stream/:id/viewer/ws", websocket.New(handlers.StreamViewerWS, websocket.Config{}))
	app.Static("/", "./assets")

	w.Rooms = make(map[string]*w.Room)
	w.Streams = make(map[string]*w.room)
	go dispatchKeyframes()

	fmt.Println("Server started on port", *addr)
	if *cert != "" {
		app.ListenTLS(*addr, *cert, *key)

	}
	return app.Listen(*addr)

}

func dispatchKeyframes() {
	for range time.NewTicker(3 * time.Second).C {
		for _, room := range w.Rooms {
			room.Peers.DispatchKeyFrames()
		}

	}
}
