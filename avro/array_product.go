// Code generated by github.com/actgardner/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCE:
 *     order_received.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v7/vm"
	"github.com/actgardner/gogen-avro/v7/vm/types"
)

func writeArrayProduct(r []*Product, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeProduct(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayProductWrapper struct {
	Target *[]*Product
}

func (_ *ArrayProductWrapper) SetBoolean(v bool)                { panic("Unsupported operation") }
func (_ *ArrayProductWrapper) SetInt(v int32)                   { panic("Unsupported operation") }
func (_ *ArrayProductWrapper) SetLong(v int64)                  { panic("Unsupported operation") }
func (_ *ArrayProductWrapper) SetFloat(v float32)               { panic("Unsupported operation") }
func (_ *ArrayProductWrapper) SetDouble(v float64)              { panic("Unsupported operation") }
func (_ *ArrayProductWrapper) SetBytes(v []byte)                { panic("Unsupported operation") }
func (_ *ArrayProductWrapper) SetString(v string)               { panic("Unsupported operation") }
func (_ *ArrayProductWrapper) SetUnionElem(v int64)             { panic("Unsupported operation") }
func (_ *ArrayProductWrapper) Get(i int) types.Field            { panic("Unsupported operation") }
func (_ *ArrayProductWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *ArrayProductWrapper) Finalize()                        {}
func (_ *ArrayProductWrapper) SetDefault(i int)                 { panic("Unsupported operation") }
func (r *ArrayProductWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r *ArrayProductWrapper) AppendArray() types.Field {
	var v *Product
	v = NewProduct()

	*r.Target = append(*r.Target, v)

	return (*r.Target)[len(*r.Target)-1]
}
