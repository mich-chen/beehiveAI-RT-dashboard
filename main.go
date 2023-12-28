package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"beehiveAI/messages"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// have a server to track all connections for broadcasting
type Server struct {
	conns    map[*websocket.Conn]*websocket.Conn
	messages *messages.messagesStore
}

// initiate a new Server
func newServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]*websocket.Conn),
	}
}

func (s *Server) handleWebsocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil) // upgrade the connection from HTTP to websocket to create new connection
	if err != nil {
		log.Println("Upgrade err:", err)
	}
	log.Println("Connection Successful!")
	log.Println("New incoming connection from client: ", conn.RemoteAddr())

	// add new connection to Server struct
	s.conns[conn] = conn

	defer conn.Close()

	// read message loop
	s.readLoop(conn)
}

func (s *Server) readLoop(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read message err:", err)
			break
		}

		// log recceived message
		log.Printf("Received message: %s", message)

		// store messages and format to write
		s.messages.addMessage(message)

		// write messages and broadcast
		s.broadcast(message)
	}
}

func (s *Server) broadcast(message []byte) {
	for ws := range s.conns {
		go func(ws *websocket.Conn) {
			if err := ws.WriteMessage(websocket.TextMessage, message); err != nil {
				fmt.Println("Broadcasting write message err:", err)
			}
		}(ws)
	}
}

func main() {
	server := newServer()
	server.messages = messages.NewMessagesMap()
	http.HandleFunc("/websocket", server.handleWebsocket)
	log.Fatal(http.ListenAndServe(":3001", nil))
}
