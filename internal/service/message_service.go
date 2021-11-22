package service

import (
	"encoding/json"
	"github.com/nicolascancino/event-message/internal/dto"
)

type MessageEvent interface {
	Publish(message []byte, attributes map[string]string) error
}
type MessageService struct {
	event MessageEvent
}

func NewMessageService(event MessageEvent) *MessageService {
	return &MessageService{event: event}
}

func (ms MessageService) FormatMessageService(message *dto.Message) error {

	var attributeMaps map[string]string
	rawAttributes, _ := json.Marshal(message.Attributes)
	json.Unmarshal(rawAttributes, &attributeMaps)

	rawData, _ := json.Marshal(message.Data)

	if err := ms.event.Publish(rawData, attributeMaps); err != nil {
		return err
	}

	return nil
}
