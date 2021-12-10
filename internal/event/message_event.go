package event

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"errors"
	"github.com/nicolascancino/event-message/internal/dto"
	"log"
	"sync"
)

type DataLoaderPublishEvent struct {
	*pubsub.Topic
	context.Context
}

type DataLoaderListenEvent struct {
	*pubsub.Subscription
	context.Context
}

func NewDataLoaderPublishEvent(topic *pubsub.Topic, ctx context.Context) *DataLoaderPublishEvent {
	return &DataLoaderPublishEvent{topic, ctx}
}

func NewDataLoaderListenEvent(subscription *pubsub.Subscription, ctx context.Context) *DataLoaderListenEvent {
	return &DataLoaderListenEvent{subscription, ctx}
}

func (d DataLoaderPublishEvent) Publish(message []byte, attributes map[string]string) error {

	if id, err := d.Topic.Publish(d.Context, &pubsub.Message{Data: message, Attributes: attributes}).Get(d.Context); err != nil {
		log.Printf("A error has occurred publishing the message with id %v", id)
		return errors.New("A error has occurred publish the message" + err.Error())
	}
	return nil
}

func (d *DataLoaderListenEvent) ReceiveMessage() (*dto.Out, error) {
	var mu sync.Mutex

	err := d.Receive(d.Context, func(ctx context.Context, msg *pubsub.Message) {

		data := make(map[string]interface{})
		json.Unmarshal(msg.Data, &data)

		toPrint := &dto.Out{
			ID:          msg.ID,
			Attributes:  msg.Attributes,
			Data:        data,
			PublishTime: msg.PublishTime.String(),
		}

		mu.Lock()
		defer mu.Unlock()
		log.Printf("Got message: %q\n", toPrint)
		msg.Ack()

	},
	)

	return nil, err

}
