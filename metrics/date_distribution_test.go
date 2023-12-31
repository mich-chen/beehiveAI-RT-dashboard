package metrics

import (
	"testing"

	"beehiveAI/messages"
)

func TestNewDateDistribution(t *testing.T) {
	distributions := NewDateDistribution()
	if distributions == nil {
		t.Error("NewDateDistribution returned nil")
	}
	if len(distributions) != 0 {
		t.Error("NewDateDistribution should be empty")
	}
}

func TestAggregateDateDistribution(t *testing.T) {
	distributions := NewDateDistribution()

	// Test successful aggregation
	validMessage := &messages.MessageData{
		TweetId:      570306133677760513,
		TweetCreated: "2015-02-24 11:35:52 -0800",
	}

	err := distributions.AggregateDateDistribution(validMessage)
	if err != nil {
		t.Errorf("Error aggregating date distribution: %v", err)
	}
	if count, ok := distributions["2015-02-24"]; !ok || count != 1 {
		t.Error("Date distribution not aggregated correctly")
	}

	// Test updating an existing date
	validMessage2 := &messages.MessageData{
		TweetId:      570306133677760000,
		TweetCreated: "2015-02-24 11:35:52 -0800",
	}
	distributions.AggregateDateDistribution(validMessage2)
	if count, ok := distributions["2015-02-24"]; !ok || count != 2 {
		t.Error("Date distribution not updated correctly")
	}

	// Test error handling for invalid time format
	InvalidTimeMessage := &messages.MessageData{
		TweetId:      570306133677760001,
		TweetCreated: "invalid time format",
	}
	err = distributions.AggregateDateDistribution(InvalidTimeMessage)
	if err == nil {
		t.Error("Expected time parse error to be returned")
	}
}
