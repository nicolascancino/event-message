package service

import (
	"encoding/json"
	"github.com/nicolascancino/event-message/internal/dto"
)

type DataLoaderPublishEvent interface {
	Publish(message []byte, attributes map[string]string) error
}

type DataLoaderListenEvent interface {
	ReceiveMessage() (*dto.Out, error)
}
type MessageService struct {
	publishEvent DataLoaderPublishEvent
	//listenEvent  DataLoaderListenEvent
}

func NewMessageService(publishEvent DataLoaderPublishEvent) *MessageService {
	return &MessageService{publishEvent: publishEvent}
}

func (ms MessageService) PublishMessageService(message *dto.Message) error {
	rawData, err := json.Marshal(message.Data)
	if err != nil {
		return err
	}

	var attributeMaps map[string]string
	rawAttributes, _ := json.Marshal(message.Attributes)
	json.Unmarshal(rawAttributes, &attributeMaps)
	if err := ms.publishEvent.Publish(rawData, attributeMaps); err != nil {
		return err
	}
	return nil
}
func (ms MessageService) ListenMessageService() {

}
