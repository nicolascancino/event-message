package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/nicolascancino/event-message/internal/dto"
	"github.com/nicolascancino/event-message/internal/event"
	"github.com/nicolascancino/event-message/internal/service"
	"log"
	"net/http"
	"os"
)

type MessageService interface {
	FormatMessageService(message *dto.Message) error
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Panicf("Error getting enviroment variable  %v", err.Error())
	}

	http.HandleFunc("/publish", publishMessage)
	http.ListenAndServe(":8080", nil)

}

func publishMessage(w http.ResponseWriter, r *http.Request) {

	var message *dto.Message
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		log.Printf("Error decoding message %v", err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}

	ctx := context.Background()

	pubsubInstance, err := pubsub.NewClient(ctx, os.Getenv("PROJECT_ID"))

	if err != nil {
		log.Panicf("Error starting pubsub client %v", err.Error())
	}

	topic := pubsubInstance.Topic(os.Getenv("TOPIC_ID"))

	defer func() {
		if err := pubsubInstance.Close(); err != nil {
			log.Printf("Error closing pubsub client %v", err.Error())
		}
	}()

	ev := event.NewDataLoaderEvent(pubsubInstance, topic, ctx)
	s := service.NewMessageService(ev)

	s.FormatMessageService(message)

}
