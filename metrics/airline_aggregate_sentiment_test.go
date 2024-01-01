package metrics

import (
	"testing"

	"github.com/mich-chen/beehiveAI-RT-dashboard/messages"
)

func TestNewAggregatedSentiment(t *testing.T) {
	aggregated := NewAggregatedSentiment()
	if aggregated == nil {
		t.Error("NewAggregatedSentiment returned nil")
	}
	if len(aggregated) != 0 {
		t.Error("NewAggregatedSentiment should be empty")
	}
}

func TestAggregateSentiment(t *testing.T) {
	aggregated := NewAggregatedSentiment()

	validMessage := &messages.MessageData{
		TweetId:          570306133677760513,
		AirlineSentiment: "positive",
		Airline:          "Delta",
	}

	// Test adding new airline sentiment
	err := aggregated.AggregateSentiment(validMessage)
	if err != nil {
		t.Errorf("Error aggregating sentiment: %v", err)
	}

	sentiment, ok := aggregated["Delta"]
	if !ok {
		t.Error("Airline sentiment not added")
	}

	expectedSentiment := struct {
		Positive int `json:"positive"`
		Negative int `json:"negative"`
		Neutral  int `json:"neutral"`
	}{
		Positive: 1,
		Negative: 0,
		Neutral:  0,
	}
	if sentiment != expectedSentiment {
		t.Errorf("Incorrect sentiment counts: %+v", sentiment)
	}

	// Test updating existing sentiment
	validMessage2 := &messages.MessageData{
		TweetId:          570306133677760000,
		AirlineSentiment: "neutral",
		Airline:          "Delta",
	}
	aggregated.AggregateSentiment(validMessage2)

	sentiment2, ok2 := aggregated["Delta"]
	if !ok2 {
		t.Error("Airline sentiment not added")
	}

	expectedSentiment2 := struct {
		Positive int `json:"positive"`
		Negative int `json:"negative"`
		Neutral  int `json:"neutral"`
	}{
		Positive: 1,
		Negative: 0,
		Neutral:  1,
	}
	if sentiment2 != expectedSentiment2 {
		t.Errorf("Incorrect sentiment counts: %+v", sentiment)
	}
}
