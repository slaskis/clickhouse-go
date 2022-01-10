//go:build !(amd64 || arm64)

// Code generated by make codegen DO NOT EDIT.
// source: lib/column/codegen/column_safe.tpl

package column

import (
	"github.com/ClickHouse/clickhouse-go/lib/binary"
)

func (col *Float32) Decode(decoder *binary.Decoder, rows int) error {
	for i := 0; i < rows; i++ {
		v, err := decoder.Float32()
		if err != nil {
			return err
		}
		*col = append(*col, v)
	}
	return nil
}

func (col *Float32) Encode(encoder *binary.Encoder) error {
	for _, v := range *col {
		if err := encoder.Float32(v); err != nil {
			return err
		}
	}
	return nil
}

func (col *Float64) Decode(decoder *binary.Decoder, rows int) error {
	for i := 0; i < rows; i++ {
		v, err := decoder.Float64()
		if err != nil {
			return err
		}
		*col = append(*col, v)
	}
	return nil
}

func (col *Float64) Encode(encoder *binary.Encoder) error {
	for _, v := range *col {
		if err := encoder.Float64(v); err != nil {
			return err
		}
	}
	return nil
}

func (col *Int8) Decode(decoder *binary.Decoder, rows int) error {
	for i := 0; i < rows; i++ {
		v, err := decoder.Int8()
		if err != nil {
			return err
		}
		*col = append(*col, v)
	}
	return nil
}

func (col *Int8) Encode(encoder *binary.Encoder) error {
	for _, v := range *col {
		if err := encoder.Int8(v); err != nil {
			return err
		}
	}
	return nil
}

func (col *Int16) Decode(decoder *binary.Decoder, rows int) error {
	for i := 0; i < rows; i++ {
		v, err := decoder.Int16()
		if err != nil {
			return err
		}
		*col = append(*col, v)
	}
	return nil
}

func (col *Int16) Encode(encoder *binary.Encoder) error {
	for _, v := range *col {
		if err := encoder.Int16(v); err != nil {
			return err
		}
	}
	return nil
}

func (col *Int32) Decode(decoder *binary.Decoder, rows int) error {
	for i := 0; i < rows; i++ {
		v, err := decoder.Int32()
		if err != nil {
			return err
		}
		*col = append(*col, v)
	}
	return nil
}

func (col *Int32) Encode(encoder *binary.Encoder) error {
	for _, v := range *col {
		if err := encoder.Int32(v); err != nil {
			return err
		}
	}
	return nil
}

func (col *Int64) Decode(decoder *binary.Decoder, rows int) error {
	for i := 0; i < rows; i++ {
		v, err := decoder.Int64()
		if err != nil {
			return err
		}
		*col = append(*col, v)
	}
	return nil
}

func (col *Int64) Encode(encoder *binary.Encoder) error {
	for _, v := range *col {
		if err := encoder.Int64(v); err != nil {
			return err
		}
	}
	return nil
}

func (col *UInt8) Decode(decoder *binary.Decoder, rows int) error {
	for i := 0; i < rows; i++ {
		v, err := decoder.UInt8()
		if err != nil {
			return err
		}
		*col = append(*col, v)
	}
	return nil
}

func (col *UInt8) Encode(encoder *binary.Encoder) error {
	for _, v := range *col {
		if err := encoder.UInt8(v); err != nil {
			return err
		}
	}
	return nil
}

func (col *UInt16) Decode(decoder *binary.Decoder, rows int) error {
	for i := 0; i < rows; i++ {
		v, err := decoder.UInt16()
		if err != nil {
			return err
		}
		*col = append(*col, v)
	}
	return nil
}

func (col *UInt16) Encode(encoder *binary.Encoder) error {
	for _, v := range *col {
		if err := encoder.UInt16(v); err != nil {
			return err
		}
	}
	return nil
}

func (col *UInt32) Decode(decoder *binary.Decoder, rows int) error {
	for i := 0; i < rows; i++ {
		v, err := decoder.UInt32()
		if err != nil {
			return err
		}
		*col = append(*col, v)
	}
	return nil
}

func (col *UInt32) Encode(encoder *binary.Encoder) error {
	for _, v := range *col {
		if err := encoder.UInt32(v); err != nil {
			return err
		}
	}
	return nil
}

func (col *UInt64) Decode(decoder *binary.Decoder, rows int) error {
	for i := 0; i < rows; i++ {
		v, err := decoder.UInt64()
		if err != nil {
			return err
		}
		*col = append(*col, v)
	}
	return nil
}

func (col *UInt64) Encode(encoder *binary.Encoder) error {
	for _, v := range *col {
		if err := encoder.UInt64(v); err != nil {
			return err
		}
	}
	return nil
}
