package math

import (
	"encoding/binary"
	"math"
	"unsafe"
)

const (
	FLOAT64_MAX      = math.MaxFloat64
	UINT_MIN    uint = 0
	UINT_MAX         = ^uint(0)
	INT_MAX          = int(^uint(0) >> 1)
	INT_MIN          = ^INT_MAX
)

func FloatFloor(f float64, prec int) float64 {
	l := math.Pow10(prec)
	return math.Floor(f*l) / l
}

func Max[T byte | int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64](a, b T) T {
	return T(math.Max(float64(a), float64(b)))
}

func Min[T byte | int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64](a, b T) T {
	return T(math.Min(float64(a), float64(b)))
}

func Int2Byte(data int) (ret []byte) {
	var len uintptr = unsafe.Sizeof(data)
	ret = make([]byte, len)
	var tmp int = 0xff
	var index uint = 0
	for index = 0; index < uint(len); index++ {
		ret[index] = byte((tmp << (index * 8) & data) >> (index * 8))
	}
	return ret
}

func Byte2Int(data []byte) int {
	var ret int = 0
	var len int = len(data)
	var i uint = 0
	for i = 0; i < uint(len); i++ {
		ret = ret | (int(data[i]) << (i * 8))
	}
	return ret
}

func ByteToUint64(data []byte) uint64 {
	return binary.BigEndian.Uint64(data)
}

func Uint64ToByte(i uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, i)
	return b[:]
}

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}

func Float32ToByte(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)

	return bytes
}

func ByteToFloat32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	return math.Float32frombits(bits)
}

func Float64ToByte(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

func ByteToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	return math.Float64frombits(bits)
}
