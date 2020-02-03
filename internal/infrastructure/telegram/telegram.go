package telegram

import (
	"context"

	"github.com/disiqueira/coronabot/internal/domain/model"

	"gopkg.in/tucnak/telebot.v2"
)

type (
	Telegram struct {
		bot *telebot.Bot

		recipients map[int64]telebot.Recipient
	}
)

func New(bot *telebot.Bot) *Telegram {
	recipients := make(map[int64]telebot.Recipient)

	bot.Handle("/start", func(m *telebot.Message) {
		recipients[m.Chat.ID] = m.Chat
	})
	bot.Handle("/stop", func(m *telebot.Message) {
		delete(recipients, m.Chat.ID)
	})

	go bot.Start()

	return &Telegram{
		bot: bot,

		recipients: recipients,
	}
}

func (t *Telegram) Send(ctx context.Context, m model.Message) error {
	for _, recipient := range t.recipients {
		_, err := t.bot.Send(recipient, m.Text(), telebot.ModeMarkdown)
		if err != nil {
			return err
		}
	}
	return nil
}
