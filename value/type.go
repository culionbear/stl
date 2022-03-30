package value

import (
	"encoding/binary"
	"math"
)

type Byte byte

func (b Byte) HashCode() uint32 {
	return uint32(b)
}

type String string

func (str String) HashCode() uint32 {
	var sum uint32
	for _, v := range str {
		sum = sum<<5 - sum + uint32(v)
	}
	return sum
}

type Bytes []byte

func (buf Bytes) HashCode() uint32 {
	var sum uint32
	for _, v := range buf {
		sum = sum<<5 - sum + uint32(v)
	}
	return sum
}

type Number int

func (num Number) HashCode() uint32 {
	return uint32(num)
}

type Bool bool

func (flag Bool) HashCode() uint32 {
	if flag {
		return 1
	}
	return 0
}

type Double float64

func (double Double) HashCode() uint32 {
	bits := math.Float64bits(float64(double))
	buf := make(Bytes, 8, 8)
	binary.LittleEndian.PutUint64(buf, bits)
	return buf.HashCode()
}

type Float float32

func (float Float) HashCode() uint32 {
	bits := math.Float32bits(float32(float))
	buf := make(Bytes, 4, 4)
	binary.LittleEndian.PutUint32(buf, bits)
	return buf.HashCode()
}
