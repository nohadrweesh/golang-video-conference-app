package server

import (
	"flag"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"github.com/gofiber/websocket/v2"
)

var(
	addr = flag.String("addr,":"", os.Getenv("PORT"),"")
	cert =flag.String("cert","","")
	key =flag.String("key","","")
)

func Run() error{

	if *addr ==":"{
		*addr=":8080"
	}

	engine := html.New("./views" , ".html")
	app := fiber.New(fiber.Config(Views: engine))
	app.Use(logger.New())
	app.Use(Cors.New())

	app.Get("/",handlers.Welcome)
	app.Get("/room/create",handlers.RoomCreate)
	app.Get("/room/:uuid",handlers.Room)
	app.Get("/room/:uuid/websocket",websocket.new(handlers.RoomWebsocket,websocket.Config{HandleTimeout: 10*time.Second}))

	app.Get("/room/:uuid/chat",handlers.RoomChat)
	app.Get("/room/:uuid/chat/websocket",handlers.RoomChatWebsocket)
	app.Get("/room/:uuid/viewer/websocket",handlers.RoomViewerWebsocket)


	app.Get("/room/:ssuid",handlers.Stream)
	app.Get("/room/:ssuid/websocket",)
	app.Get("/room/:ssuid/chat/websocket",)
	app.Get("/room/:ssuid/viewer/websocket",)

}