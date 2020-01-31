package application

import (
	"context"

	"github.com/disiqueira/coronabot/internal/domain/model"
	"github.com/disiqueira/coronabot/internal/domain/service"
)

type (
	NotifyService struct {
		messageSender             model.MessageSender
		statusReporter            model.StatusReporter
		statusPerCountryToMessage *service.StatusListToMessage
	}
)

func NewNotifyService(messageSender model.MessageSender, statusReporter model.StatusReporter, statusPerCountryToMessage *service.StatusListToMessage) *NotifyService {
	return &NotifyService{
		messageSender:             messageSender,
		statusReporter:            statusReporter,
		statusPerCountryToMessage: statusPerCountryToMessage,
	}
}

func (n *NotifyService) Execute(ctx context.Context) error {
	statusList, err := n.statusReporter.StatusPerCountry()
	if err != nil {
		return err
	}

	message := n.statusPerCountryToMessage.Convert(statusList)
	return n.messageSender.Send(ctx, message)
}
