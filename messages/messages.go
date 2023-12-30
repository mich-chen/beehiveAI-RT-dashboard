package messages

import "fmt"

type MessagesStore []*MessageData

type MessageData struct {
	TweetId                   int    `json:"tweetId,omitempty"`
	AirlineSentiment          string `json:"airlineSentiment,omitempty"`
	AirlineSentimentConfident int    `json:"airlineSentimentConfidence,omitempty"`
	NegativeReason            string `json:"negativereason,omitempty"`
	NegativeReasonConfidence  int    `json:"negativereasonConfidence,omitempty"`
	Airline                   string `json:"airline,omitempty"`
	AirlineSentimentGold      string `json:"airlineSentimentGold,omitempty"`
	Name                      string `json:"name,omitempty"`
	NegativeReasonGold        string `json:"negativereasonGold,omitempty"`
	RetweetCount              int    `json:"retweetCount,omitempty"`
	Text                      string `json:"text,omitempty"`
	TweetCord                 string `json:"tweetCord,omitempty"`
	TweetCreated              string `json:"tweetCreated,omitempty"`
	TweetLocation             string `json:"tweetLocation,omitempty"`
	UserTimezone              string `json:"userTimezone,omitempty"`
}

func NewMessagesStore() MessagesStore {
	var messages MessagesStore
	return messages
}

func (messages *MessagesStore) AddMessage(msg *MessageData) error {
	before := len(*messages)

	*messages = append([]*MessageData{msg}, *messages...)

	if curr := len(*messages); curr <= before {
		return fmt.Errorf("Error adding new message to messages")
	}
	return nil
}
