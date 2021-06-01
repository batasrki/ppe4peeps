package main

import (
	"net/http"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func main() {
	server := &http.Server{
		Addr: "0.0.0.0:8082",
	}
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/produce", produceMessage)

	server.ListenAndServe()

}

func produceMessage(w http.ResponseWriter, r *http.Request) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	defer p.Close()
	topic := "orders-received"
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte("This is an order from an HTTP request"),
	}, nil)
	p.Flush(1 * 100)

	w.Write([]byte("Order event produced to the topic"))
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong"))
}
