package kafka

import (
	"encoding/json"
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	event "github.com/thyagopereira/full-cycle/eda/internal/events"
	"github.com/thyagopereira/full-cycle/eda/pkg/events"
)

type Consumer struct {
	ConfigMap *ckafka.ConfigMap
	Topics    []string
}

func NewConsumer(configMap *ckafka.ConfigMap, topics []string) *Consumer {
	return &Consumer{
		ConfigMap: configMap,
		Topics:    topics,
	}
}

func (c *Consumer) Consume(dispatcher events.EventDispatcher) error {
	consumer, err := ckafka.NewConsumer(c.ConfigMap)
	if err != nil {
		panic(err)
	}
	err = consumer.SubscribeTopics(c.Topics, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Consumming events ...")
	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			fmt.Println("Error consumming kafka", err)
			return err
		}

		switch *msg.TopicPartition.Topic {
		case "balances":
			var balanceUpdated event.BalanceUpdated
			json.Unmarshal(msg.Value, &balanceUpdated)
			fmt.Println(" we are getting events")
			dispatcher.Dispatch(&balanceUpdated)
		}

	}
}
