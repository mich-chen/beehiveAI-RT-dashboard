package messages

import (
	"time"
)

type MessagesStore map[int]struct {
	Data      *MessageData
	Timestamp int64
}

type MessageData struct {
	TweetId                   int    `json:"tweetId"`
	AirlineSentiment          string `json:"airlineSentiment"`
	AirlineSentimentConfident int    `json:"airlineSentimentConfidence"`
	NegativeReason            string `json:"negativereason"`
	NegativeReasonConfidence  int    `json:"negativereasonConfidence"`
	Airline                   string `json:"airline"`
	AirlineSentimentGold      string `json:"airlineSentimentGold"`
	Name                      string `json:"name"`
	NegativeReasonGold        string `json:"negativereasonGold"`
	RetweetCount              int    `json:"retweetCount"`
	Text                      string `json:"text"`
	TweetCord                 string `json:"tweetCord"`
	TweetCreated              string `json:"tweetCreated"`
	TweetLocation             string `json:"tweetLocation"`
	UserTimezone              string `json:"userLocation"`
}

func NewMessagesMap() MessagesStore {
	return make(MessagesStore)
}

func (messages MessagesStore) AddMessage(msg *MessageData) {
	// add new message and timestamp to message store
	timestamp := time.Now().UnixNano()
	messages[msg.TweetId] = struct {
		Data      *MessageData
		Timestamp int64
	}{
		Data:      msg,
		Timestamp: timestamp,
	}
}
