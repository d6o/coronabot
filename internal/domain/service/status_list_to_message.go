package service

import (
	"fmt"

	"github.com/disiqueira/coronabot/internal/domain/model"
)

type (
	StatusListToMessage struct {
	}
)

func NewStatusListToMessage() *StatusListToMessage {
	return &StatusListToMessage{}
}

func (s StatusListToMessage) Convert(statusList []model.Status) model.Message {
	result := ":biohazard_sign: *|COVID-19|*\n"
	result += "_50 Most affected countries:_\n"
	result += "```"
	result += fmt.Sprintf("|%20s|%10s|%7s|%10s|\n", "Country", "Confirmed", "Deaths", "Recovered")
	result += fmt.Sprintf("|%20s|%10s|%7s|%10s|\n", "", "", "", "")
	i := 0
	for _, status := range statusList {
		if i++; i <= 50 {
			result += fmt.Sprintf("|%20s|%10d|%7d|%10d|\n", status.Country(), status.Confirmed(), status.Deaths(), status.Recovered())
		}
	}
	result += "```"
	return model.NewMessage(result)
}
