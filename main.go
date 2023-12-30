package main

import (
	"encoding/json"
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

type ResponseData struct {
	Messages *messages.MessagesStore `json:"messages,omitempty"`
	Metrics  *ResponseMetrics        `json:"metrics,omitempty"`
}

type ResponseMetrics struct {
	AggregatedSentiments *metrics.AirlineAggregatedSentiment `json:"aggregatedSentiments,omitempty"`
	DateDistributions    *metrics.DateDistribution           `json:"dateDistributions,omitempty"`
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
		return
	}
	log.Println("Connection Successful!")
	log.Println("New incoming connection from client: ", conn.RemoteAddr())

	// add new connection to Server struct
	s.conns[conn] = conn

	// handle close and removing client connections
	defer func(conn *websocket.Conn, server *Server) {
		delete(server.conns, conn)
	}(conn, s)
	defer conn.Close()

	// read messages
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read message err:", err)
			break
		}

		log.Println("/Websocket received message from client: ", message)
	}

	// broadcast current store of messages when new clients connect
	responseData := &ResponseData{
		Messages: &s.messages,
		Metrics: &ResponseMetrics{
			AggregatedSentiments: &s.aggregatedSentiments,
			DateDistributions:    &s.dateDistributions,
		},
	}

	s.broadcast(responseData)
}

func (s *Server) handleWebhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse incoming data
	var data *messages.MessageData
	decode := json.NewDecoder(r.Body)
	decode.DisallowUnknownFields()
	if err := decode.Decode(&data); err != nil {
		log.Println("Parsing webhook data err:", err)
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	// Handle storing and aggregating data
	if err := s.messages.AddMessage(data); err != nil {
		log.Println("Add new message err:", err)
		http.Error(w, "Message could not be added", http.StatusInternalServerError)
		return
	}
	if err := s.aggregatedSentiments.AggregateSentiment(data); err != nil {
		log.Println("Aggregating airline sentiment err:", err)
		http.Error(w, "Could not aggregate airline sentiment", http.StatusInternalServerError)
		return
	}
	if err := s.dateDistributions.AggregateDateDistribution(data); err != nil {
		log.Println("Aggregating date distribution err:", err)
		http.Error(w, "Could not aggregate date distribution", http.StatusInternalServerError)
		return
	}

	responseData := &ResponseData{
		Messages: &s.messages,
		Metrics: &ResponseMetrics{
			AggregatedSentiments: &s.aggregatedSentiments,
			DateDistributions:    &s.dateDistributions,
		},
	}

	// Broadcast to WebSocket clients
	s.broadcast(responseData)

	resp, err := json.Marshal(map[string]bool{"successful": true})
	if err != nil {
		log.Println("Error making response confirmation", err)
	}
	w.Write(resp)
}

func (s *Server) broadcast(data *ResponseData) {
	for ws := range s.conns {
		go func(ws *websocket.Conn) {
			if err := ws.WriteJSON(data); err != nil {
				log.Println("Broadcasting write message err:", err)
				ws.Close()
				return
			}
		}(ws)
	}
}

func main() {
	server := newServer()
	server.messages = messages.NewMessagesStore()
	server.aggregatedSentiments = metrics.NewAggregatedSentiment()
	server.dateDistributions = metrics.NewDateDistribution()
	http.HandleFunc("/webhook", server.handleWebhook)
	http.HandleFunc("/websocket", server.handleWebsocket)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
