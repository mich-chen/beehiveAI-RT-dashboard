package metrics

import (
	"log"
	"time"

	"beehiveAI/messages"
)

type DateDistribution map[string]int

const (
	fullLayout = "2006-01-02 15:04:05 -0800"
	dateLayout = "2006-01-02"
)

func NewDateDistribution() DateDistribution {
	return make(DateDistribution)
}

// map of tweet distribution by date formatted "YYYY-MM-DD" from UTC time
func (distributions DateDistribution) AggregateDateDistribution(msg *messages.MessageData) {
	timeObj, err := time.Parse(fullLayout, msg.TweetCreated)
	if err != nil {
		log.Println("time parsing err:", err, msg.TweetCreated)
	}

	date := timeObj.UTC().Format(dateLayout)

	if _, ok := distributions[date]; ok {
		distributions[date]++
	} else {
		distributions[date] = 1
	}
}
