package metrics

import (
	"time"

	"github.com/mich-chen/beehiveAI-RT-dashboard/messages"
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
func (distributions DateDistribution) AggregateDateDistribution(msg *messages.MessageData) error {
	timeObj, err := time.Parse(fullLayout, msg.TweetCreated)
	if err != nil {
		return err
	}

	date := timeObj.UTC().Format(dateLayout)

	if _, ok := distributions[date]; ok {
		distributions[date]++
	} else {
		distributions[date] = 1
	}
	return nil
}
