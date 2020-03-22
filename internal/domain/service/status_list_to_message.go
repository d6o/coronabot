package service

import (
	"fmt"

	"github.com/disiqueira/coronabot/internal/domain/model"
)

type (
	StatusListToMessage struct {
		limit int
	}
)

func NewStatusListToMessage(limit int) *StatusListToMessage {
	return &StatusListToMessage{
		limit: limit,
	}
}

func (s StatusListToMessage) Convert(statusList []model.Status) model.Message {
	result := ":biohazard_sign: *|COVID-19|*\n"
	result += fmt.Sprintf("_%d Most affected countries:_\n", s.limit)
	result += "```"
	result += fmt.Sprintf("|%20s|%10s|%7s|%10s|\n", "Country", "Confirmed", "Deaths", "Recovered")
	result += fmt.Sprintf("|%20s|%10s|%7s|%10s|\n", "", "", "", "")
	for i, status := range statusList {
		if i < s.limit {
			result += fmt.Sprintf("|%20s|%10d|%7d|%10d|\n", status.Country(), status.Confirmed(), status.Deaths(), status.Recovered())
		}
	}
	result += "```"
	return model.NewMessage(result)
}
