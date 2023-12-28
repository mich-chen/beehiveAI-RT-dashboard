package main

import (
	"encoding/json"
	"log"
	"time"
)

type messagesStore map[int]struct {
	data      *messageData
	timestamp int64
}

type messageData struct {
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
}

func newMessagesMap() messagesStore {
	return make(messagesStore)
}

func (messages messagesStore) addMessage(msg []byte) {
	var msgData messageData
	if err := json.Unmarshal(msg, &msgData); err != nil {
		log.Println("JSON unmarshal err:", err)
	}

	// add new message and timestamp to message store
	timestamp := time.Now().UnixNano()
	messages[msgData.TweetId] = struct {
		data      *messageData
		timestamp int64
	}{
		data:      &msgData,
		timestamp: timestamp,
	}
	log.Println("message store:", messages)
}
