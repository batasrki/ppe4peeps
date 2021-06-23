package handlers

import (
	"bytes"
	"fmt"
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
	log.WithField("order", order).Info("Received a new order")

	if err := validateOrder(order); err != nil {
		return nil, err
	}
	return order.BuildAvroMessage(), nil
}

func validateOrder(o *orders.Order) error {
	if len(o.Products) == 0 {
		return fmt.Errorf("there are no products in this order")
	}

	for i, pc := range o.Products {
		if len(pc.ProductCode) == 0 {
			return fmt.Errorf("the product code for order [%d] is blank", i)
		}

		if pc.Quantity == 0 {
			return fmt.Errorf("the quantity for order [%d] must be > 0", i)
		}
	}

	if len(o.Customer.Email) == 0 {
		return fmt.Errorf("the customer must have an e-mail")
	}

	if len(o.Customer.ShippingAddress.Address) == 0 {
		return fmt.Errorf("the customer's shipping address street and number is required")
	}

	if len(o.Customer.ShippingAddress.City) == 0 {
		return fmt.Errorf("the customer's shipping address city is required")
	}

	if len(o.Customer.ShippingAddress.PostalCode) == 0 {
		return fmt.Errorf("the customer's shipping address postal code is required")
	}

	if len(o.Customer.ShippingAddress.StateProvince) == 0 {
		return fmt.Errorf("the customer's shipping address state or province is required")
	}

	return nil
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
