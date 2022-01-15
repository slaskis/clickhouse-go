package column

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2/lib/binary"
)

type offset struct {
	values   UInt64
	scanType reflect.Type
}

type Array struct {
	chType   Type
	values   Interface
	offsets  []*offset
	scanType reflect.Type
}

func (col *Array) parse(t Type) (_ Interface, err error) {
	col.chType = t
	var (
		depth   int
		typeStr = string(t)
	)
parse:
	for _, str := range strings.Split(typeStr, "Array(") {
		switch {
		case len(str) == 0:
			depth++
		default:
			typeStr = str[:len(str)-depth]
			break parse
		}
	}
	if depth != 0 {
		if col.values, err = Type(typeStr).Column(); err != nil {
			return nil, err
		}
		offsetScanTypes := make([]reflect.Type, 0, depth)
		col.offsets, col.scanType = make([]*offset, 0, depth), col.values.ScanType()
		for i := 0; i < depth; i++ {
			col.scanType = reflect.SliceOf(col.scanType)
			offsetScanTypes = append(offsetScanTypes, col.scanType)
		}
		for i := len(offsetScanTypes) - 1; i >= 0; i-- {
			col.offsets = append(col.offsets, &offset{
				scanType: offsetScanTypes[i],
			})
		}
		return col, nil
	}
	return &UnsupportedColumnType{
		t: t,
	}, nil
}

func (col *Array) Type() Type {
	return col.chType
}

func (col *Array) ScanType() reflect.Type {
	return col.scanType
}

func (col *Array) Rows() int {
	if len(col.offsets) != 0 {
		return len(col.offsets[0].values)
	}
	return 0
}

func (col *Array) Row(i int) interface{} {
	return col.make(uint64(i), 0)
}

func (col *Array) ScanRow(dest interface{}, row int) error {
	elem := reflect.Indirect(reflect.ValueOf(dest))
	if elem.Type() != col.scanType {
		return &ColumnConverterErr{
			op:   "ScanRow",
			to:   fmt.Sprintf("%T", dest),
			from: string(col.chType),
		}
	}
	{
		elem.Set(reflect.ValueOf(col.make(uint64(row), 0)))
	}
	return nil
}

func (col *Array) Append(v interface{}) (nulls []uint8, err error) {
	switch v := v.(type) {
	case []time.Time:

	default:
		return nil, &ColumnConverterErr{
			op:   "Append",
			to:   string(col.chType),
			from: fmt.Sprintf("%T", v),
		}
	}
	return
}

func (col *Array) AppendRow(v interface{}) error {
	elem := reflect.Indirect(reflect.ValueOf(v))
	if elem.Type() != col.scanType {
		return &ColumnConverterErr{
			op:   "AppendRow",
			to:   fmt.Sprintf("%T", v),
			from: string(col.chType),
		}
	}
	return col.append(elem, 0)
}

func (col *Array) append(elem reflect.Value, level int) error {
	if elem.Kind() == reflect.Slice {
		offset := uint64(elem.Len())
		if ln := len(col.offsets[level].values); ln != 0 {
			offset += col.offsets[level].values[ln-1]
		}
		col.offsets[level].values = append(col.offsets[level].values, offset)
		for i := 0; i < elem.Len(); i++ {
			if err := col.append(elem.Index(i), level+1); err != nil {
				return err
			}
		}
		return nil
	}
	return col.values.AppendRow(elem.Interface())
}

func (col *Array) Decode(decoder *binary.Decoder, rows int) error {
	for _, offset := range col.offsets {
		if err := offset.values.Decode(decoder, rows); err != nil {
			return err
		}
		rows = int(offset.values[len(offset.values)-1])
	}
	return col.values.Decode(decoder, rows)
}

func (col *Array) Encode(encoder *binary.Encoder) error {
	for _, offset := range col.offsets {
		if err := offset.values.Encode(encoder); err != nil {
			return err
		}
	}
	return col.values.Encode(encoder)
}

func (col *Array) make(row uint64, level int) interface{} {
	offset := col.offsets[level]
	var (
		end   = offset.values[row]
		start = uint64(0)
	)
	if row > 0 {
		start = offset.values[row-1]
	}
	slice := reflect.MakeSlice(offset.scanType, 0, int(end-start))
	for i := start; i < end; i++ {
		var value interface{}
		switch {
		case level == len(col.offsets)-1:
			value = col.values.Row(int(i))
		default:
			value = col.make(i, level+1)
		}
		slice = reflect.Append(slice, reflect.ValueOf(value))
	}
	return slice.Interface()
}

var _ Interface = (*Date)(nil)