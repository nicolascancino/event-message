package event

import (
	"cloud.google.com/go/pubsub"
	"context"
	"errors"
	"log"
)

type DataLoaderEvent struct {
	*pubsub.Client
	*pubsub.Topic
	context.Context
}

func NewDataLoaderEvent(client *pubsub.Client, topic *pubsub.Topic, ctx context.Context) *DataLoaderEvent {
	return &DataLoaderEvent{client, topic, ctx}
}

func (d DataLoaderEvent) Publish(message []byte, attributes map[string]string) error {
	ctx := context.Background()
	if id, err := d.Topic.Publish(d.Context, &pubsub.Message{Data: message, Attributes: attributes}).Get(ctx); err != nil {
		log.Printf("A error has occurred publishing the message with id %v", id)
		return errors.New("A error has occurred publish the message" + err.Error())
	}

	log.Printf("The message with event id %v has been published successfully", attributes["eventId"])

	return nil
}
