package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/nicolascancino/event-message/internal/event"
	"github.com/nicolascancino/event-message/internal/handler"
	"github.com/nicolascancino/event-message/internal/service"
	"github.com/nicolascancino/event-message/pkg/router"
	"github.com/nicolascancino/event-message/pkg/server"
	"log"
	"os"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Panicf("Error getting enviroment variable  %v", err.Error())
	}

	var ctx = context.Background()

	pubSubInstance := getPubsubInstance(ctx, os.Getenv("PROJECT_ID"))
	topicInstance := getTopic(pubSubInstance, os.Getenv("TOPIC_ID"))
	subscriptionInstance := getSubscription(pubSubInstance, os.Getenv("SUBSCRIPTION_NAME"))

	publishEventInstance := event.NewDataLoaderPublishEvent(topicInstance, ctx)
	listenEventInstance := event.NewDataLoaderListenEvent(subscriptionInstance, ctx)
	serviceInstance := service.NewMessageService(publishEventInstance, listenEventInstance)
	handlerInstance := handler.NewMessageHandler(serviceInstance)
	muxInstance := mux.NewRouter()
	routerInstance := router.NewRouter(handlerInstance, muxInstance)
	serverInstance := server.NewServer(routerInstance)
	serverInstance.Start()

	log.Println("started publish and listener application")
}

func getTopic(pubSubInstance *pubsub.Client, topicId string) *pubsub.Topic {
	return pubSubInstance.Topic(topicId)
}

func getSubscription(pubSubInstance *pubsub.Client, subscriptionName string) *pubsub.Subscription {
	return pubSubInstance.Subscription(subscriptionName)
}

func getPubsubInstance(ctx context.Context, projectId string) *pubsub.Client {
	pubsubInstance, err := pubsub.NewClient(ctx, projectId)
	if err != nil {
		log.Panicf("Error starting pubsub client %v", err.Error())
	}
	return pubsubInstance
}
