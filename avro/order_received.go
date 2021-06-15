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

type OrderReceived struct {
	// Event metadata.
	Metadata *Metadatum `json:"metadata"`

	Order *Order `json:"Order"`
}

const OrderReceivedAvroCRC64Fingerprint = "\xf2g\x88\x8e\x89\x9a\xc5m"

func NewOrderReceived() *OrderReceived {
	return &OrderReceived{}
}

func DeserializeOrderReceived(r io.Reader) (*OrderReceived, error) {
	t := NewOrderReceived()
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

func DeserializeOrderReceivedFromSchema(r io.Reader, schema string) (*OrderReceived, error) {
	t := NewOrderReceived()

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

func writeOrderReceived(r *OrderReceived, w io.Writer) error {
	var err error
	err = writeMetadatum(r.Metadata, w)
	if err != nil {
		return err
	}
	err = writeOrder(r.Order, w)
	if err != nil {
		return err
	}
	return err
}

func (r *OrderReceived) Serialize(w io.Writer) error {
	return writeOrderReceived(r, w)
}

func (r *OrderReceived) Schema() string {
	return "{\"fields\":[{\"doc\":\"Event metadata.\",\"name\":\"metadata\",\"type\":{\"fields\":[{\"doc\":\"A unique ID representing the event itself.\",\"name\":\"id\",\"type\":\"long\"},{\"doc\":\"A Unix timestamp of the event.\",\"name\":\"timestamp\",\"type\":\"long\"}],\"name\":\"metadatum\",\"type\":\"record\"}},{\"name\":\"Order\",\"type\":{\"doc\":\"The order schema.\",\"fields\":[{\"doc\":\"A unique ID of the order.\",\"name\":\"id\",\"type\":\"long\"},{\"doc\":\"A Unix timestamp of the time the order was created.\",\"name\":\"timestamp\",\"type\":\"long\"},{\"doc\":\"A list of products in the order.\",\"name\":\"products\",\"type\":{\"items\":{\"doc\":\"The product schema.\",\"fields\":[{\"doc\":\"A unique product code.\",\"name\":\"productCode\",\"type\":\"string\"},{\"doc\":\"Quantity of the product ordered.\",\"name\":\"quantity\",\"type\":\"int\"}],\"name\":\"product\",\"type\":\"record\"},\"type\":\"array\"}},{\"doc\":\"Customer information.\",\"name\":\"customer\",\"type\":{\"fields\":[{\"name\":\"firstName\",\"type\":\"string\"},{\"name\":\"lastName\",\"type\":\"string\"},{\"name\":\"email\",\"type\":\"string\"},{\"name\":\"shippingAddress\",\"type\":{\"fields\":[{\"name\":\"address\",\"type\":\"string\"},{\"name\":\"city\",\"type\":\"string\"},{\"name\":\"stateProvince\",\"type\":\"string\"},{\"name\":\"postalCode\",\"type\":\"string\"}],\"name\":\"addressLines\",\"type\":\"record\"}}],\"name\":\"information\",\"type\":\"record\"}}],\"name\":\"order\",\"type\":\"record\"}}],\"name\":\"OrderReceived\",\"type\":\"record\"}"
}

func (r *OrderReceived) SchemaName() string {
	return "OrderReceived"
}

func (_ *OrderReceived) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *OrderReceived) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *OrderReceived) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *OrderReceived) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *OrderReceived) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *OrderReceived) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *OrderReceived) SetString(v string)   { panic("Unsupported operation") }
func (_ *OrderReceived) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *OrderReceived) Get(i int) types.Field {
	switch i {
	case 0:
		r.Metadata = NewMetadatum()

		return r.Metadata
	case 1:
		r.Order = NewOrder()

		return r.Order
	}
	panic("Unknown field index")
}

func (r *OrderReceived) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *OrderReceived) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ *OrderReceived) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *OrderReceived) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *OrderReceived) Finalize()                        {}

func (_ *OrderReceived) AvroCRC64Fingerprint() []byte {
	return []byte(OrderReceivedAvroCRC64Fingerprint)
}
