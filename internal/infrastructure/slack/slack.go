package slack

import (
	"context"

	"github.com/nlopes/slack"

	"github.com/disiqueira/coronabot/internal/domain/model"
)

type (
	Slack struct {
		client  *slack.Client
		channel string
	}
)

func New(client *slack.Client, channel string) *Slack {
	return &Slack{
		client:  client,
		channel: channel,
	}
}

func (s *Slack) Send(_ context.Context, m model.Message) error {
	_, _, _, err := s.client.SendMessage(s.channel, slack.MsgOptionText(m.Text(), false))
	return err
}
