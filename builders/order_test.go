package models

import (
	"testing"
)

func TestBuildFromJson(t *testing.T) {
	o := NewOrder()
	input := staticInput()

	testOrder := o.BuildFromJson(input)

	// basics
	if testOrder.ID != 6 {
		t.Errorf("Expected ID of value 6, got %d", testOrder.ID)
	}

	// customer
	customer := testOrder.Customer

	if customer.FirstName != "Tom" {
		t.Errorf("Expected customer's first name to be Tom, got %s", customer.FirstName)
	}

	if customer.Email != "tom.hardy@email.com" {
		t.Errorf("Expected customer's e-mail of tom.hardy@email.com, got %s", customer.Email)
	}

}

func TestBuildAvroMessage(t *testing.T) {
	o := NewOrder()
	testOrder := o.BuildFromJson(staticInput())

	testAvroMessage := testOrder.BuildAvroMessage()

	o1 := NewOrder()
	deserializedMessage := o1.BuildFromAvroMessage(testAvroMessage)

	// basics
	if deserializedMessage.ID != 6 {
		t.Errorf("Expected ID of value 6, got %d", deserializedMessage.ID)
	}

	// products
	products := deserializedMessage.Products
	if products[0].ProductCode != "12345" {
		t.Errorf("Expected product code field with value 12345, got %s", products[0].ProductCode)
	}
	if len(products) != len(testOrder.Products) {
		t.Errorf("Expected number of products %d, got %d", len(testOrder.Products), len(products))
	}

	// customer
	customer := deserializedMessage.Customer
	if customer.FirstName != "Tom" {
		t.Errorf("Expected customer's name Tom, got %s", customer.FirstName)
	}
	if customer.Email != "tom.hardy@email.com" {
		t.Errorf("Expected email tom.hardy@email.com, got %s", customer.Email)
	}
}

func staticInput() string {
	return `
	{
		"id":6,
		"products":[
			{
				"productCode":"12345",
				"quantity":2
			}
		],
		"customer":{
			"firstName":"Tom",
			"lastName":"Hardy",
			"emailAddress":"tom.hardy@email.com",
			"shippingAddress":{
				"line1":"123 Anywhere St",
				"city":"Anytown",
				"state":"AL",
				"postalCode":"12345"
			}
		}
	}
	`
}
