package main

import (
	"io"
	"fmt"
	"html/template"
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/gorilla/websocket"
)

type Message struct {
	ID string
	Name string
	Message string
	CreatedAt string
}

// TODOS
// Same the messages creating on websockets
// Load it when new client is open
// Add timestamps
// Channels and rooms

type Template struct {
	templates *template.Template
}

var clients = make(map[*websocket.Conn]bool)

var messages = []Message {
	{ID: "1", Name: "Test", Message: "Hello", CreatedAt: "test"},
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Home(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", "/")
}

func handleWebSocket(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	fmt.Println("Client connected")

	clients[conn] = true

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return err
		}

		fmt.Println(messageType, p)
		receivedMessage := string(p)
		fmt.Printf("Received message: %s\n", receivedMessage)

		// Broadcast the message to all connected clients
		for clientConn := range clients {
			err := clientConn.WriteMessage(messageType, p)
			if err != nil {
				fmt.Println(err)
				return err
			}
		}
	}
}

// func sendMessage(c echo.Context) error {
// 	message := c.FormValue("message")
// 	newMessage := Message {
// 		ID:        "2", 
// 		Name:      "User", 
// 		Message:   message,
// 		CreatedAt: "now", 
// 	}
// 	messages = append(messages, newMessage)
// 	return c.Render(http.StatusOK, "chat.html", newMessage)
// }

func getMessages(c echo.Context) error {
	return c.Render(http.StatusOK, "chats.html", messages)
}

func main() {
	e := echo.New()
	t := &Template{
    templates: template.Must(template.ParseGlob("web/templates/*.html")),
	}
	e.Renderer = t
	e.Static("/connection", "web/connection")

	e.GET("/ws", handleWebSocket)
	e.GET("/", Home)
	e.GET("/getMessages", getMessages)
	// e.POST("/sendMessage", sendMessage)
	e.Logger.Fatal(e.Start(":8080"))
}