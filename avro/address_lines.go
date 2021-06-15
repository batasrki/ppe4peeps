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

type AddressLines struct {
	Address string `json:"address"`

	City string `json:"city"`

	StateProvince string `json:"stateProvince"`

	PostalCode string `json:"postalCode"`
}

const AddressLinesAvroCRC64Fingerprint = "YI<\xc3i\xa0\x01)"

func NewAddressLines() *AddressLines {
	return &AddressLines{}
}

func DeserializeAddressLines(r io.Reader) (*AddressLines, error) {
	t := NewAddressLines()
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

func DeserializeAddressLinesFromSchema(r io.Reader, schema string) (*AddressLines, error) {
	t := NewAddressLines()

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

func writeAddressLines(r *AddressLines, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Address, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.City, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.StateProvince, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.PostalCode, w)
	if err != nil {
		return err
	}
	return err
}

func (r *AddressLines) Serialize(w io.Writer) error {
	return writeAddressLines(r, w)
}

func (r *AddressLines) Schema() string {
	return "{\"fields\":[{\"name\":\"address\",\"type\":\"string\"},{\"name\":\"city\",\"type\":\"string\"},{\"name\":\"stateProvince\",\"type\":\"string\"},{\"name\":\"postalCode\",\"type\":\"string\"}],\"name\":\"addressLines\",\"type\":\"record\"}"
}

func (r *AddressLines) SchemaName() string {
	return "addressLines"
}

func (_ *AddressLines) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *AddressLines) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *AddressLines) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *AddressLines) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *AddressLines) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *AddressLines) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *AddressLines) SetString(v string)   { panic("Unsupported operation") }
func (_ *AddressLines) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *AddressLines) Get(i int) types.Field {
	switch i {
	case 0:
		return &types.String{Target: &r.Address}
	case 1:
		return &types.String{Target: &r.City}
	case 2:
		return &types.String{Target: &r.StateProvince}
	case 3:
		return &types.String{Target: &r.PostalCode}
	}
	panic("Unknown field index")
}

func (r *AddressLines) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *AddressLines) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ *AddressLines) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *AddressLines) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *AddressLines) Finalize()                        {}

func (_ *AddressLines) AvroCRC64Fingerprint() []byte {
	return []byte(AddressLinesAvroCRC64Fingerprint)
}
