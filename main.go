package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"beehiveAI/messages"
	"beehiveAI/metrics"
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
	conns                map[*websocket.Conn]*websocket.Conn
	messages             messages.MessagesStore
	aggregatedSentiments metrics.AirlineAggregatedSentiment
	dateDistributions    metrics.DateDistribution
}

type ResponseMetrics struct {
	AggregatedSentiments *metrics.AirlineAggregatedSentiment `json:"aggregatedSentiments"`
	DateDistributions    *metrics.DateDistribution           `json:"dateDistributions"`
}

// initiate a new Server
func newServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]*websocket.Conn),
	}
}

func ParseFromJSON(msg []byte) *messages.MessageData {
	var msgData messages.MessageData
	if err := json.Unmarshal(msg, &msgData); err != nil {
		log.Println("JSON unmarshal err:", err)
	}
	return &msgData
}

func ParseToMessage(msgData *messages.MessageData, sentimentsData metrics.AirlineAggregatedSentiment, distributionData metrics.DateDistribution) []byte {
	var formatted = struct {
		Message *messages.MessageData `json:"message"`
		Metrics *ResponseMetrics      `json:"metrics"`
	}{
		Message: msgData,
		Metrics: &ResponseMetrics{
			AggregatedSentiments: &sentimentsData,
			DateDistributions:    &distributionData,
		},
	}
	log.Println("main.go: formattedResponse:", formatted)

	responseMsg, err := json.Marshal(formatted)
	if err != nil {
		log.Println("JSON marshal err:", err)
	}
	return responseMsg
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
		log.Printf("Received message: %s, %T", message, message)

		// parse received message
		messageData := ParseFromJSON(message)

		// store messages
		s.messages.AddMessage(messageData)
		log.Println("message store:", s.messages)

		// aggregated metrics
		s.aggregatedSentiments.AggregateSentiment(messageData)
		log.Println("main.go: aggregated in server", s.aggregatedSentiments)

		s.dateDistributions.AggregateDateDistribution(messageData)
		log.Println("main.go: date distributions in server", s.dateDistributions)

		// format final message
		responseMessage := ParseToMessage(
			messageData,
			s.aggregatedSentiments,
			s.dateDistributions,
		)

		// write messages and broadcast
		s.broadcast(responseMessage)
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
	server.aggregatedSentiments = metrics.NewAggregatedSentiment()
	server.dateDistributions = metrics.NewDateDistribution()
	http.HandleFunc("/websocket", server.handleWebsocket)
	log.Fatal(http.ListenAndServe(":3001", nil))
}
