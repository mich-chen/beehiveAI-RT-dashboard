package messages

type MessagesStore []*MessageData

type MessageData struct {
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
	TweetCord                 string `json:"tweetCord"`
	TweetCreated              string `json:"tweetCreated"`
	TweetLocation             string `json:"tweetLocation"`
	UserTimezone              string `json:"userTimezone"`
}

func NewMessagesStore() MessagesStore {
	var messages MessagesStore
	return messages
}

func (messages *MessagesStore) AddMessage(msg *MessageData) {
	*messages = append([]*MessageData{msg}, *messages...)
}
