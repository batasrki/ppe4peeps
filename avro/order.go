// Code generated by github.com/actgardner/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCE:
 *     order_received.avsc
 */
package avro

import (
	"github.com/actgardner/gogen-avro/v7/compiler"
	"github.com/actgardner/gogen-avro/v7/vm"
	"github.com/actgardner/gogen-avro/v7/vm/types"
	"io"
)

// The order schema.
type Order struct {
	// A unique ID of the order.
	Id int64 `json:"id"`
	// A Unix timestamp of the time the order was created.
	Timestamp int64 `json:"timestamp"`
	// A list of products in the order.
	Products []*Product `json:"products"`
	// Customer information.
	Customer *Information `json:"customer"`
}

const OrderAvroCRC64Fingerprint = "E\xc1u\xbe(\xe6\xa5\xd0"

func NewOrder() *Order {
	return &Order{}
}

func DeserializeOrder(r io.Reader) (*Order, error) {
	t := NewOrder()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func DeserializeOrderFromSchema(r io.Reader, schema string) (*Order, error) {
	t := NewOrder()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func writeOrder(r *Order, w io.Writer) error {
	var err error
	err = vm.WriteLong(r.Id, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.Timestamp, w)
	if err != nil {
		return err
	}
	err = writeArrayProduct(r.Products, w)
	if err != nil {
		return err
	}
	err = writeInformation(r.Customer, w)
	if err != nil {
		return err
	}
	return err
}

func (r *Order) Serialize(w io.Writer) error {
	return writeOrder(r, w)
}

func (r *Order) Schema() string {
	return "{\"doc\":\"The order schema.\",\"fields\":[{\"doc\":\"A unique ID of the order.\",\"name\":\"id\",\"type\":\"long\"},{\"doc\":\"A Unix timestamp of the time the order was created.\",\"name\":\"timestamp\",\"type\":\"long\"},{\"doc\":\"A list of products in the order.\",\"name\":\"products\",\"type\":{\"items\":{\"doc\":\"The product schema.\",\"fields\":[{\"doc\":\"A unique product code.\",\"name\":\"productCode\",\"type\":\"string\"},{\"doc\":\"Quantity of the product ordered.\",\"name\":\"quantity\",\"type\":\"int\"}],\"name\":\"product\",\"type\":\"record\"},\"type\":\"array\"}},{\"doc\":\"Customer information.\",\"name\":\"customer\",\"type\":{\"fields\":[{\"name\":\"firstName\",\"type\":\"string\"},{\"name\":\"lastName\",\"type\":\"string\"},{\"name\":\"email\",\"type\":\"string\"},{\"name\":\"shippingAddress\",\"type\":{\"fields\":[{\"name\":\"address\",\"type\":\"string\"},{\"name\":\"city\",\"type\":\"string\"},{\"name\":\"stateProvince\",\"type\":\"string\"},{\"name\":\"postalCode\",\"type\":\"string\"}],\"name\":\"addressLines\",\"type\":\"record\"}}],\"name\":\"information\",\"type\":\"record\"}}],\"name\":\"order\",\"type\":\"record\"}"
}

func (r *Order) SchemaName() string {
	return "order"
}

func (_ *Order) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *Order) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *Order) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *Order) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *Order) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *Order) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *Order) SetString(v string)   { panic("Unsupported operation") }
func (_ *Order) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Order) Get(i int) types.Field {
	switch i {
	case 0:
		return &types.Long{Target: &r.Id}
	case 1:
		return &types.Long{Target: &r.Timestamp}
	case 2:
		r.Products = make([]*Product, 0)

		return &ArrayProductWrapper{Target: &r.Products}
	case 3:
		r.Customer = NewInformation()

		return r.Customer
	}
	panic("Unknown field index")
}

func (r *Order) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Order) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ *Order) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *Order) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *Order) Finalize()                        {}

func (_ *Order) AvroCRC64Fingerprint() []byte {
	return []byte(OrderAvroCRC64Fingerprint)
}
