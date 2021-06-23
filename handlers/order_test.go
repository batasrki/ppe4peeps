package handlers

import (
	"testing"

	orders "github.com/batasrki/ppe4peeps/builders"
)

func TestValidateOrderWhenOrderIsValid(t *testing.T) {
	order := validOrder()
	if err := validateOrder(order); err != nil {
		t.Errorf("Valid order isn't valid!")
	}
}

func TestValidateOrderWhenThereAreNoProducts(t *testing.T) {
	order := validOrder()
	order.Products = (make([]*orders.Product, 0))
	var err error

	if err = validateOrder(order); err == nil {
		t.Errorf("Valid order isn't valid!")
	}
	errMessage := "there are no products in this order"
	if err.Error() != errMessage {
		t.Errorf("Expected '%s', got '%s'", errMessage, err.Error())
	}
}

func TestValidateOrderWhenAProductHasNoProductCode(t *testing.T) {
	order := validOrder()
	order.Products[0].ProductCode = ""
	var err error

	if err = validateOrder(order); err == nil {
		t.Errorf("Valid order isn't valid!")
	}

	errMessage := "the product code for order [0] is blank"
	if err.Error() != errMessage {
		t.Errorf("Expected '%s', got '%s'", errMessage, err.Error())
	}
}

func TestValidateOrderWhenAProductHasNoZeroQuantity(t *testing.T) {
	order := validOrder()
	order.Products[0].Quantity = 0
	var err error

	if err = validateOrder(order); err == nil {
		t.Errorf("Valid order isn't valid!")
	}

	errMessage := "the quantity for order [0] must be > 0"
	if err.Error() != errMessage {
		t.Errorf("Expected '%s', got '%s'", errMessage, err.Error())
	}
}

func TestValidateOrderMustHaveCustomerEmailAddress(t *testing.T) {
	order := validOrder()
	order.Customer.Email = ""
	var err error

	if err = validateOrder(order); err == nil {
		t.Errorf("Valid order isn't valid!")
	}

	errMessage := "the customer must have an e-mail"
	if err.Error() != errMessage {
		t.Errorf("Expected '%s', got '%s'", errMessage, err.Error())
	}
}

func TestValidateOrderMustHaveCustomerStreetAddress(t *testing.T) {
	order := validOrder()
	order.Customer.ShippingAddress.Address = ""
	var err error

	if err = validateOrder(order); err == nil {
		t.Errorf("Valid order isn't valid!")
	}

	errMessage := "the customer's shipping address street and number is required"
	if err.Error() != errMessage {
		t.Errorf("Expected '%s', got '%s'", errMessage, err.Error())
	}
}

func TestValidateOrderMustHaveCustomerShippingAddressCity(t *testing.T) {
	order := validOrder()
	order.Customer.ShippingAddress.City = ""
	var err error

	if err = validateOrder(order); err == nil {
		t.Errorf("Valid order isn't valid!")
	}

	errMessage := "the customer's shipping address city is required"
	if err.Error() != errMessage {
		t.Errorf("Expected '%s', got '%s'", errMessage, err.Error())
	}
}

func TestValidateOrderMustHaveCustomerShippingAddressPostalCode(t *testing.T) {
	order := validOrder()
	order.Customer.ShippingAddress.PostalCode = ""
	var err error

	if err = validateOrder(order); err == nil {
		t.Errorf("Valid order isn't valid!")
	}

	errMessage := "the customer's shipping address postal code is required"
	if err.Error() != errMessage {
		t.Errorf("Expected '%s', got '%s'", errMessage, err.Error())
	}
}

func TestValidateOrderMustHaveCustomerShippingAddressStateOrProvince(t *testing.T) {
	order := validOrder()
	order.Customer.ShippingAddress.StateProvince = ""
	var err error

	if err = validateOrder(order); err == nil {
		t.Errorf("Valid order isn't valid!")
	}

	errMessage := "the customer's shipping address state or province is required"
	if err.Error() != errMessage {
		t.Errorf("Expected '%s', got '%s'", errMessage, err.Error())
	}
}

func validOrder() *orders.Order {
	o := orders.NewOrder()
	o.Products = append(o.Products, &orders.Product{
		ProductCode: "12345",
		Quantity:    1,
	})
	o.Customer = &orders.Customer{
		FirstName: "Tester",
		LastName:  "Tom",
		Email:     "tester@tom.com",
		ShippingAddress: &orders.ShippingAddress{
			Address:       "1 testing way",
			City:          "Teston",
			PostalCode:    "l5a3l1",
			StateProvince: "ON",
		},
	}
	return o
}
