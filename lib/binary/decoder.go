package binary

import (
	"encoding/binary"
	"io"
	"math"
)

func NewDecoder(input io.Reader) *Decoder {
	return &Decoder{
		input: input,
	}
}

type Decoder struct {
	input   io.Reader
	scratch [16]byte
}

func (decoder *Decoder) Bool() (bool, error) {
	v, err := decoder.ReadByte()
	if err != nil {
		return false, err
	}
	return v == 1, nil
}

func (decoder *Decoder) Uvarint() (uint64, error) {
	return binary.ReadUvarint(decoder)
}

func (decoder *Decoder) Int8() (int8, error) {
	v, err := decoder.ReadByte()
	if err != nil {
		return 0, err
	}
	return int8(v), nil
}

func (decoder *Decoder) Int16() (int16, error) {
	v, err := decoder.UInt16()
	if err != nil {
		return 0, err
	}
	return int16(v), nil
}

func (decoder *Decoder) Int32() (int32, error) {
	v, err := decoder.UInt32()
	if err != nil {
		return 0, err
	}
	return int32(v), nil
}

func (decoder *Decoder) Int64() (int64, error) {
	v, err := decoder.UInt64()
	if err != nil {
		return 0, err
	}
	return int64(v), nil
}

func (decoder *Decoder) UInt8() (uint8, error) {
	v, err := decoder.ReadByte()
	if err != nil {
		return 0, err
	}
	return uint8(v), nil
}

func (decoder *Decoder) UInt16() (uint16, error) {
	buf := decoder.scratch[:2]
	if _, err := decoder.input.Read(buf); err != nil {
		return 0, err
	}
	return uint16(buf[0]) | uint16(buf[1])<<8, nil
}

func (decoder *Decoder) UInt32() (uint32, error) {
	buf := decoder.scratch[:4]
	if _, err := decoder.input.Read(buf); err != nil {
		return 0, err
	}
	return uint32(buf[0]) | uint32(buf[1])<<8 | uint32(buf[2])<<16 | uint32(buf[3])<<24, nil
}

func (decoder *Decoder) UInt64() (uint64, error) {
	buf := decoder.scratch[:8]
	if _, err := decoder.input.Read(buf); err != nil {
		return 0, err
	}
	return uint64(buf[0]) | uint64(buf[1])<<8 | uint64(buf[2])<<16 | uint64(buf[3])<<24 |
		uint64(buf[4])<<32 | uint64(buf[5])<<40 | uint64(buf[6])<<48 | uint64(buf[7])<<56, nil
}

func (decoder *Decoder) Float32() (float32, error) {
	v, err := decoder.UInt32()
	if err != nil {
		return 0, err
	}
	return math.Float32frombits(v), nil
}

func (decoder *Decoder) Float64() (float64, error) {
	v, err := decoder.UInt64()
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(v), nil
}

func (decoder *Decoder) Fixed(ln int) ([]byte, error) {
	if reader, ok := decoder.input.(FixedReader); ok {
		return reader.Fixed(ln)
	}
	buf := make([]byte, ln)
	if _, err := decoder.input.Read(buf); err != nil {
		return nil, err
	}
	return buf, nil
}

func (decoder *Decoder) String() (string, error) {
	strlen, err := decoder.Uvarint()
	if err != nil {
		return "", err
	}
	str, err := decoder.Fixed(int(strlen))
	if err != nil {
		return "", err
	}
	return string(str), nil
}

func (decoder *Decoder) ReadByte() (byte, error) {
	if _, err := decoder.input.Read(decoder.scratch[:1]); err != nil {
		return 0x0, err
	}
	return decoder.scratch[0], nil
}

type FixedReader interface {
	Fixed(ln int) ([]byte, error)
}