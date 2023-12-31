package metrics

import (
	"beehiveAI/messages"
	"fmt"
)

type AirlineAggregatedSentiment map[string]struct {
	Positive int `json:"positive"`
	Negative int `json:"negative"`
	Neutral  int `json:"neutral"`
}

func NewAggregatedSentiment() AirlineAggregatedSentiment {
	return make(AirlineAggregatedSentiment)
}

// map where key is airline with value of aggregated sentiment
func (aggregated AirlineAggregatedSentiment) AggregateSentiment(msg *messages.MessageData) error {
	sentiment, ok := aggregated[msg.Airline]
	if !ok {
		sentiment = struct {
			Positive int `json:"positive"`
			Negative int `json:"negative"`
			Neutral  int `json:"neutral"`
		}{
			Positive: 0,
			Negative: 0,
			Neutral:  0,
		}
	}

	switch msg.AirlineSentiment {
	case "positive":
		sentiment.Positive++
		break
	case "negative":
		sentiment.Negative++
		break
	case "neutral":
		sentiment.Neutral++
		break
	default:
		break
	}
	aggregated[msg.Airline] = sentiment

	_, okAdded := aggregated[msg.Airline]
	if !okAdded {
		return fmt.Errorf("Could not add new airline sentiment")
	}
	return nil
}
