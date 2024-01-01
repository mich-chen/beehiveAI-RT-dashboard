package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/go-playground/validator/v10"
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

var mutex sync.Mutex
var validate *validator.Validate

// initiate a new Server
func newServer() *Server {
	return &Server{
		conns:                make(map[*websocket.Conn]*websocket.Conn),
		messages:             messages.NewMessagesStore(),
		aggregatedSentiments: metrics.NewAggregatedSentiment(),
		dateDistributions:    metrics.NewDateDistribution(),
	}
}

func (s *Server) handleWebsocket(w http.ResponseWriter, r *http.Request) {
	// upgrade the connection from HTTP to websocket to create new connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade err:", err)
		return
	}
	log.Println("Connection Successful!")
	log.Println("New incoming connection from client: ", conn.RemoteAddr())

	// add new connection to Server struct
	mutex.Lock()
	s.conns[conn] = conn
	mutex.Unlock()

	// Send current store of data to any new client connections
	mutex.Lock()
	responseData := &ResponseData{
		Messages: &s.messages,
		Metrics: &ResponseMetrics{
			AggregatedSentiments: &s.aggregatedSentiments,
			DateDistributions:    &s.dateDistributions,
		},
	}

	if err := conn.WriteJSON(responseData); err != nil {
		log.Println("Broadcasting write message err:", err)
		conn.Close()
	}
	mutex.Unlock()

	// read for any client connection closures
	for {
		if _, _, err := conn.NextReader(); err != nil {
			log.Println("Recived client close message --> removing connection")
			mutex.Lock()
			delete(s.conns, conn)
			mutex.Unlock()
			conn.Close()
			break
		}
	}
}

func (s *Server) handleWebhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse incoming data
	mutex.Lock()
	var data *messages.MessageData
	decode := json.NewDecoder(r.Body)
	decode.DisallowUnknownFields()
	if err := decode.Decode(&data); err != nil {
		log.Println("Parsing webhook data err:", err)
		http.Error(w, "Invalid data", http.StatusBadRequest)
		mutex.Unlock()
		return
	}
	defer r.Body.Close()
	// validate required fields
	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		log.Println("Invalid data, missing required fields", err)
		http.Error(w, "Invalid data, missing required fields", http.StatusBadRequest)
		mutex.Unlock()
		return
	}

	// Handle storing and aggregating data
	if err := s.messages.AddMessage(data); err != nil {
		log.Println("Add new message err:", err)
		http.Error(w, "Message could not be added", http.StatusInternalServerError)
		mutex.Unlock()
		return
	}
	if err := s.aggregatedSentiments.AggregateSentiment(data); err != nil {
		log.Println("Aggregating airline sentiment err:", err)
		http.Error(w, "Could not aggregate airline sentiment", http.StatusInternalServerError)
		mutex.Unlock()
		return
	}
	if err := s.dateDistributions.AggregateDateDistribution(data); err != nil {
		log.Println("Aggregating date distribution err:", err)
		http.Error(w, "Could not aggregate date distribution", http.StatusInternalServerError)
		mutex.Unlock()
		return
	}
	mutex.Unlock()

	responseData := &ResponseData{
		Messages: &s.messages,
		Metrics: &ResponseMetrics{
			AggregatedSentiments: &s.aggregatedSentiments,
			DateDistributions:    &s.dateDistributions,
		},
	}

	// Broadcast to WebSocket clients
	s.broadcast(responseData)

	mutex.Lock()
	resp, err := json.Marshal(map[string]bool{"successful": true})
	if err != nil {
		log.Println("Error making response confirmation", err)
	}
	mutex.Unlock()

	w.Write(resp)
}

func (s *Server) broadcast(data *ResponseData) {
	mutex.Lock()
	clients := s.conns
	mutex.Unlock()

	for ws := range clients {
		mutex.Lock()
		if err := ws.WriteJSON(data); err != nil {
			log.Println("Broadcasting write message err:", err)
			ws.Close()
		}
		mutex.Unlock()
	}
}

func main() {
	server := newServer()
	http.HandleFunc("/webhook", server.handleWebhook)
	http.HandleFunc("/websocket", server.handleWebsocket)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
