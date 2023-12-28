package messages

import (
	"encoding/json"
	"log"
	"time"
)

type MessagesStore map[int]struct {
	Data      *MessageData
	Timestamp int64
}

type MessageData struct {
	TweetId                   int    `json:"tweet_id"`
	AirlineSentiment          string `json:"airline_sentiment"`
	AirlineSentimentConfident int    `json:"airline_sentiment_confidence"`
	NegativeReason            string `json:"negativereason"`
	NegativeReasonConfidence  int    `json:"negativereason_confidence"`
	Airline                   string `json:"airline"`
	AirlineSentimentGold      string `json:"airline_sentiment_gold"`
	Name                      string `json:"name"`
	NegativeReasonGold        string `json:"negativereason_gold"`
	RetweetCount              int    `json:"retweet_count"`
	Text                      string `json:"text"`
	TweetCord                 string `json:"tweet_cord"`
	TweetCreated              string `json:"tweet_created"`
	TweetLocation             string `json:"tweet_location"`
	UserTimezone              string `json:"user_location"`
}

func NewMessagesMap() MessagesStore {
	return make(MessagesStore)
}

func ParseFromJSON(msg []byte) *MessageData {
	var msgData MessageData
	if err := json.Unmarshal(msg, &msgData); err != nil {
		log.Println("JSON unmarshal err:", err)
	}
	return &msgData
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
