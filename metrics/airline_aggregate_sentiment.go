package metrics

import (
	"beehiveAI/messages"
)

type AirlineAggregatedSentiment map[string]struct {
	Total    int `json:"total"`
	Positive int `json:"positive"`
	Negative int `json:"negative"`
	Neutral  int `json:"neutral"`
}

func NewAggregatedSentiment() AirlineAggregatedSentiment {
	return make(AirlineAggregatedSentiment)
}

// map where key is airline with value of aggregated sentiment
func (aggregated AirlineAggregatedSentiment) AggregateSentiment(msg *messages.MessageData) {
	sentiment, ok := aggregated[msg.Airline]
	if ok {
		sentiment.Total += 1
	} else {
		sentiment = struct {
			Total    int `json:"total"`
			Positive int `json:"positive"`
			Negative int `json:"negative"`
			Neutral  int `json:"neutral"`
		}{
			Total:    1,
			Positive: 0,
			Negative: 0,
			Neutral:  0,
		}
	}

	switch msg.AirlineSentiment {
	case "positive":
		sentiment.Positive += 1
		break
	case "negative":
		sentiment.Negative += 1
		break
	case "neutral":
		sentiment.Neutral += 1
		break
	default:
		break
	}
	aggregated[msg.Airline] = sentiment
}
