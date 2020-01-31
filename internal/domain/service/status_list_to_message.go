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
	result := "```"
	result += fmt.Sprintf("|%20s|%10s|%7s|%10s|\n", "Country", "Confirmed", "Deaths", "Recovered")
	result += fmt.Sprintf("|%20s|%10s|%7s|%10s|\n", "", "", "", "")
	for _, status := range statusList {
		result += fmt.Sprintf("|%20s|%10d|%7d|%10d|\n", status.Country(), status.Confirmed(), status.Deaths(), status.Recovered())
	}
	result += "```"
	return model.NewMessage(result)
}
