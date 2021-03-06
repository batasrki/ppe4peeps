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

// The product schema.
type Product struct {
	// A unique product code.
	ProductCode string `json:"productCode"`
	// Quantity of the product ordered.
	Quantity int32 `json:"quantity"`
}

const ProductAvroCRC64Fingerprint = "0\a\xef\xb0֥UO"

func NewProduct() *Product {
	return &Product{}
}

func DeserializeProduct(r io.Reader) (*Product, error) {
	t := NewProduct()
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

func DeserializeProductFromSchema(r io.Reader, schema string) (*Product, error) {
	t := NewProduct()

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

func writeProduct(r *Product, w io.Writer) error {
	var err error
	err = vm.WriteString(r.ProductCode, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Quantity, w)
	if err != nil {
		return err
	}
	return err
}

func (r *Product) Serialize(w io.Writer) error {
	return writeProduct(r, w)
}

func (r *Product) Schema() string {
	return "{\"doc\":\"The product schema.\",\"fields\":[{\"doc\":\"A unique product code.\",\"name\":\"productCode\",\"type\":\"string\"},{\"doc\":\"Quantity of the product ordered.\",\"name\":\"quantity\",\"type\":\"int\"}],\"name\":\"product\",\"type\":\"record\"}"
}

func (r *Product) SchemaName() string {
	return "product"
}

func (_ *Product) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *Product) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *Product) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *Product) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *Product) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *Product) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *Product) SetString(v string)   { panic("Unsupported operation") }
func (_ *Product) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Product) Get(i int) types.Field {
	switch i {
	case 0:
		return &types.String{Target: &r.ProductCode}
	case 1:
		return &types.Int{Target: &r.Quantity}
	}
	panic("Unknown field index")
}

func (r *Product) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Product) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ *Product) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *Product) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *Product) Finalize()                        {}

func (_ *Product) AvroCRC64Fingerprint() []byte {
	return []byte(ProductAvroCRC64Fingerprint)
}
