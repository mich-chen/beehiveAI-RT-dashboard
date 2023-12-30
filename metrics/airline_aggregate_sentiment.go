package metrics

import "beehiveAI/messages"

type AirlineAggregatedSentiment map[string]struct {
	Total    int `json:"total,omitempty"`
	Positive int `json:"positive,omitempty"`
	Negative int `json:"negative,omitempty"`
	Neutral  int `json:"neutral,omitempty"`
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
			Total    int `json:"total,omitempty"`
			Positive int `json:"positive,omitempty"`
			Negative int `json:"negative,omitempty"`
			Neutral  int `json:"neutral,omitempty"`
		}{
			Total:    1,
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
}
