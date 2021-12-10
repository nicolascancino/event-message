package handler

import (
	"encoding/json"
	"github.com/nicolascancino/event-message/internal/dto"
	"log"
	"net/http"
)

type MessageService interface {
	PublishMessageService(message *dto.Message) error
	ListenMessageService()
}

type MessageHandler struct {
	messageService MessageService
}

func NewMessageHandler(messageService MessageService) *MessageHandler {
	return &MessageHandler{messageService: messageService}
}

func (messageHandler MessageHandler) PublishMessageHandler(w http.ResponseWriter, r *http.Request) {

	data, err := toDTOIn(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	if err := messageHandler.messageService.PublishMessageService(data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
}

func toDTOIn(r *http.Request) (*dto.Message, error) {
	var message *dto.Message
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		log.Printf("Error decoding message %v", err.Error())
		return nil, err
	}
	return message, nil
}
