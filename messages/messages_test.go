package messages

import (
	"testing"
)

func TestNewMessagesStore(t *testing.T) {
	messages := NewMessagesStore()
	if messages == nil {
		t.Error("NewMessagesStore returned nil")
	}
	if len(messages) != 0 {
		t.Error("NewMessagesStore should be empty")
	}
}

func TestAddMessage(t *testing.T) {
	messages := NewMessagesStore()

	// Test messages vars
	validMessage := &MessageData{
		TweetId:                   570306133677760513,
		AirlineSentiment:          "neutral",
		AirlineSentimentConfident: 1,
		NegativeReason:            "",
		NegativeReasonConfidence:  0,
		Airline:                   "Virgin America",
		AirlineSentimentGold:      "",
		Name:                      "cairdin",
		NegativeReasonGold:        "",
		RetweetCount:              0,
		Text:                      "@VirginAmerica What @dhepburn said.",
		TweetCord:                 "[40.74804263, -73.99295302]",
		TweetCreated:              "2015-02-24 11:35:52 -0800",
		TweetLocation:             "",
		UserTimezone:              "Eastern Time (US & Canada)",
	}

	// Test successful addition
	err := messages.AddMessage(validMessage)
	if err != nil {
		t.Errorf("Error adding valid message: %v", err)
	}
	if len(messages) != 1 || messages[0] != validMessage {
		t.Error("Valid message not added correctly")
	}
}
