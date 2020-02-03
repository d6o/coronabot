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
	result := "```\n"
	result += fmt.Sprintf("|%20s|%9s|%6s|%9s|\n", "Country", "Confirmed", "Deaths", "Recovered")
	result += fmt.Sprintf("|%20s|%9s|%6s|%9s|\n", "", "", "", "")
	for _, status := range statusList {
		result += fmt.Sprintf("|%20s|%9d|%6d|%9d|\n", status.Country(), status.Confirmed(), status.Deaths(), status.Recovered())
	}
	result += "```"
	return model.NewMessage(result)
}
