package types

type AirlineAggregatedSentiment map[string]struct {
	Total    int `json:"total"`
	Positive int `json:"positive"`
	Negative int `json:"negative"`
	Neutral  int `json:"neutral"`
}

type DateDistribution map[string]int

type ResponseMetrics struct {
	AggregatedSentiments AirlineAggregatedSentiment `json:"aggregatedSentiments"`
	DateDistributions    DateDistribution           `json:"dateDistributions"`
}
