module beehiveAI/RT-webhook-response-dashboard

go 1.21.5

require (
	beehiveAI/messages v0.0.0-00010101000000-000000000000
	github.com/gorilla/websocket v1.5.1
)

require golang.org/x/net v0.17.0 // indirect

replace beehiveAI/messages => ./messages
