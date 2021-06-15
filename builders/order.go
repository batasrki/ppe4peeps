package models

import (
	"bytes"
	"encoding/json"
	"io"
	"time"

	"github.com/actgardner/gogen-avro/v7/soe"
	"github.com/batasrki/ppe4peeps/avro"
)

type Order struct {
	ID        int `json:"id"`
	CreatedAt int
	Products  []*Product `json:"products"`
	Customer  *Customer  `json:"customer"`
}

type Product struct {
	ProductCode string `json:"productCode"`
	Quantity    int    `json:"quantity"`
}

type Customer struct {
	FirstName       string           `json:"firstName"`
	LastName        string           `json:"lastName"`
	Email           string           `json:"emailAddress"`
	ShippingAddress *ShippingAddress `json:"shippingAddress"`
}

type ShippingAddress struct {
	Address       string `json:"line1"`
	City          string `json:"city"`
	StateProvince string `json:"state"`
	PostalCode    string `json:"postalCode"`
}

func NewOrder() *Order {
	return &Order{
		CreatedAt: int(time.Now().Unix()),
	}
}

func NewCustomer() *Customer {
	return &Customer{}
}

func (o *Order) BuildFromJson(orderInput string) *Order {

	if err := json.Unmarshal([]byte(orderInput), o); err != nil {
		panic(err)
	}
	return o
}

func (o *Order) BuildAvroMessage() *bytes.Buffer {
	metadata := &avro.Metadatum{
		Id:        1,
		Timestamp: time.Now().Unix(),
	}

	serializedCustomer := &avro.Information{
		FirstName:       o.Customer.FirstName,
		LastName:        o.Customer.LastName,
		Email:           o.Customer.Email,
		ShippingAddress: (*avro.AddressLines)(o.Customer.ShippingAddress),
	}

	serializedProducts := make([]*avro.Product, len(o.Products))

	for i := 0; i < len(o.Products); i++ {
		serializedProducts[i] = &avro.Product{
			ProductCode: o.Products[i].ProductCode,
			Quantity:    int32(o.Products[i].Quantity),
		}
	}

	serializedOrder := &avro.OrderReceived{
		Metadata: metadata,
		Order: &avro.Order{
			Id:        int64(o.ID),
			Timestamp: int64(o.CreatedAt),
			Products:  serializedProducts,
			Customer:  serializedCustomer,
		},
	}

	buffer := &bytes.Buffer{}

	err := soe.WriteRecord(buffer, serializedOrder)
	if err != nil {
		panic(err)
	}
	return buffer
}

func (o *Order) BuildFromAvroMessage(message *bytes.Buffer) *Order {
	_, err := soe.ReadHeader(message)
	if err != nil {
		panic(err)
	}

	avroOrderReceived, err := avro.DeserializeOrderReceived(message)
	if err != nil {
		if err != io.EOF {
			panic(err)
		}
	}

	avroOrder := avroOrderReceived.Order

	o.ID = int(avroOrder.Id)
	o.CreatedAt = int(avroOrder.Timestamp)
	o.Products = make([]*Product, len(avroOrder.Products))

	for i := 0; i < len(avroOrder.Products); i++ {
		o.Products[i] = &Product{
			ProductCode: avroOrder.Products[i].ProductCode,
			Quantity:    int(avroOrder.Products[i].Quantity),
		}
	}

	o.Customer = NewCustomer()
	o.Customer.FirstName = avroOrder.Customer.FirstName
	o.Customer.LastName = avroOrder.Customer.LastName
	o.Customer.Email = avroOrder.Customer.Email
	o.Customer.ShippingAddress = (*ShippingAddress)(avroOrder.Customer.ShippingAddress)

	return o
}
