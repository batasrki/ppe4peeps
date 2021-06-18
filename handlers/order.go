package handlers

import (
	"bytes"
	"net/http"

	orders "github.com/batasrki/ppe4peeps/builders"
	"github.com/batasrki/ppe4peeps/config"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	log "github.com/sirupsen/logrus"
)

func Orders(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func NewOrder(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)

	r.Body.Read(body)

	message, err := prepareMessage(body)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	err = produceMessage(*message)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusAccepted)
}

func prepareMessage(input []byte) (*bytes.Buffer, error) {
	order := orders.NewOrder()
	if _, err := order.BuildFromJson(input); err != nil {
		return nil, err
	}
	return order.BuildAvroMessage(), nil
}

func produceMessage(message bytes.Buffer) error {
	log.WithField("event", "orderReceived").Info("Attempting to produce the message to the cluster")
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": config.BrokerAddress()})
	if err != nil {
		return err
	}
	defer p.Close()

	deliveryChan := make(chan kafka.Event)

	topic := config.OrderReceivedTopicName
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          message.Bytes(),
	}, deliveryChan)

	e := <-deliveryChan
	msg := e.(*kafka.Message)

	if msg.TopicPartition.Error != nil {
		return msg.TopicPartition.Error
	}

	log.WithField("topic", msg.TopicPartition.Topic).
		WithField("partition", msg.TopicPartition.Partition).
		WithField("offset", msg.TopicPartition.Offset).
		Infof("Delivered message to topic")

	close(deliveryChan)

	return nil
}
