package main

import (
	"io"
	"html/template"
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/gorilla/websocket"
)
// ***Implement first the chat app***
// Then proceeds to channels and user creation

type Message struct {
	ID string
	Name string
	Message string
	CreatedAt string
}

type Template struct {
	templates *template.Template
}

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

// func handleWebSocket(c echo.Context) error {
// 	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
// 	if err != nil {
// 		return err
// 	}
// 	defer conn.Close()

// 	fmt.Println("Client connected")

// 	for {
// 		messageType, p, err := conn.ReadMessage()
// 		if err != nil {
// 			fmt.Println(err)
// 			return err
// 		}

// 		err = conn.WriteMessage(messageType, p)
// 		if err != nil {
// 			fmt.Println(err)
// 			return err
// 		}
// 	}
// }

func sendMessage(c echo.Context) error {
	message := c.FormValue("message")
	newMessage := Message {
		ID:        "2", 
		Name:      "User", 
		Message:   message,
		CreatedAt: "now", 
	}
	messages = append(messages, newMessage)
	return c.Render(http.StatusOK, "chat.html", messages)
}

func main() {
	e := echo.New()
	t := &Template{
    templates: template.Must(template.ParseGlob("web/templates/*.html")),
	}
	e.Renderer = t
	
	// e.GET("/ws", handleWebSocket)
	e.GET("/", Home)
	e.POST("/sendMessage", sendMessage)
	e.Logger.Fatal(e.Start(":8080"))
}