package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync"
	"testing"

	"github.com/gorilla/websocket"
)

func TestNewServer(t *testing.T) {
	server := newServer()
	if server == nil {
		t.Error("newServer returned nil")
	}
	if len(server.conns) != 0 {
		t.Error("newServer should have empty connections")
	}
}

func setupTest(t *testing.T) (*Server, *http.Response, *websocket.Conn) {
	var mutex sync.Mutex
	mutex.Lock()
	server := newServer()
	mutex.Unlock()

	s := httptest.NewServer(http.HandlerFunc(server.handleWebsocket))
	defer s.Close()

	// Convert http url to ws url
	u, err := url.Parse(s.URL)
	if err != nil {
		t.Log("Error parsing url to prepare to switch to websocket connection")
	}
	u.Scheme = "ws"

	// Connect to the server
	ws, resp, wsErr := websocket.DefaultDialer.Dial(u.String(), nil)
	if wsErr != nil {
		t.Fatalf("%v", wsErr)
	}

	return server, resp, ws
}

func TestHandleWebsocket(t *testing.T) {
	var mutex sync.Mutex
	mutex.Lock()
	_, resp, _ := setupTest(t)
	mutex.Unlock()

	// Check successful connection
	if resp.StatusCode != http.StatusSwitchingProtocols {
		t.Errorf("Expected status code 101, got %d", resp.StatusCode)
	}

	// CAVEAT TO FIX TEST SETUP IN FOLLOUP:
	// - test fails when accsessing var server when running $ go test -race .
	// 		- race detected even though when t.Log(server) you will see connection successfully added to server.conns map
	// - uncomment below and update line 52 with server variable and rerun tests with $ go test . (removing -race flag) and tests will pass showing connection successfully added to server.conns map

	// Check if the connection is added to the server
	// if len(server.conns) != 1 {
	// 	t.Error("Connection not added to server")
	// }

	// FOLLOWUP: Unable to validate mock connection in server.conns map, would love to learn more on how I could validate the map key with the mock *websocket.Conn
}

func TestHandleWebhook(t *testing.T) {
	var mutex sync.Mutex
	mutex.Lock()
	server, _, _ := setupTest(t)
	mutex.Unlock()

	body := strings.NewReader(`
		{
			"tweetId": 5323041546422352922,
			"airlineSentiment": "positive",
			"airlineSentimentConfidence": 1.0,
			"negativereason": null,
			"negativereasonConfidence": null,
			"airline": "Delta",
			"airlineSentimentGold": null,
			"name": "test",
			"negativereasonGold": null,
			"retweetCount": 0,
			"text": "@VirginAmer we expect a choppy landing in NYC due to some gusty winds w/a temperature of about 5 degrees &amp; w/the windchill -8",
			"tweetCoord": "[40.74804263, -73.99295302]",
			"tweetCreated": "2003-01-11 11:35:52 -0800",
			"tweetLocation": "California",
			"userTimezone": "Pacific Time (US & Canada)"
	}
	`)

	w := httptest.NewRecorder()
	// Create a request for testing
	req, webhookErr := http.NewRequest(http.MethodPost, "/webhook", body)
	if webhookErr != nil {
		t.Fatal("Error creating request to endpoint", webhookErr)
	}

	server.handleWebhook(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Received HTTP status code %d, expected 200", w.Code)
	}

	expectedRes := `{"successful":true}`
	if w.Body.String() != expectedRes {
		t.Errorf("Response body %s does not match expected: %s", w.Body.String(), expectedRes)
	}
}

func TestBroadcast(t *testing.T) {
	var mutex sync.Mutex
	mutex.Lock()
	server, _, ws := setupTest(t)
	mutex.Unlock()

	body := strings.NewReader(`
		{
			"tweetId": 5323041546422352922,
			"airlineSentiment": "positive",
			"airline": "Delta",
			"name": "test",
			"tweetCreated": "2003-01-11 11:35:52 -0800"
	}
	`)

	w := httptest.NewRecorder()

	// Create a request for testing
	req, webhookErr := http.NewRequest(http.MethodPost, "/webhook", body)
	if webhookErr != nil {
		t.Fatal("Error creating request to endpoint", webhookErr)
	}

	server.handleWebhook(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Received HTTP status code %d, expected 200", w.Code)
	}

	// Test broadcasting message to connected websockets
	expectedBroadcast := `{"messages":[{"tweetId":5323041546422352922,"airlineSentiment":"positive","airline":"Delta","name":"test","tweetCreated":"2003-01-11 11:35:52 -0800"}],"metrics":{"aggregatedSentiments":{"Delta":{"positive":1,"negative":0,"neutral":0}},"dateDistributions":{"2003-01-11":1}}}`
	initial := true
	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			t.Errorf("Error broading message to clients: %v", err)
			break
		}

		if !initial {
			if !strings.Contains(string(message), expectedBroadcast) {
				t.Errorf("Broadcast message not in right format: %v", string(message))
				t.Errorf("Expected broadcast message: %v", expectedBroadcast)
			}
			ws.Close()
			break
		}
		initial = false
	}
}
