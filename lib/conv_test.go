package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_StringToFloat32(t *testing.T) {
	assert.Equal(t, float32(3.14159), StringToFloat32("3.14159", 0.59))
}

func Test_StringToFloat32_ReturnDefaultValue(t *testing.T) {
	assert.Equal(t, float32(0.59), StringToFloat32("abc", 0.59))
}

func Test_Float32ToString(t *testing.T) {
	assert.Equal(t, "3.14159", Float32ToString(3.14159))
}

func Test_StringToFloat64(t *testing.T) {
	assert.Equal(t, float64(3.14159), StringToFloat64("3.14159", 0.59))
}

func Test_StringToFloat64_ReturnDefaultValue(t *testing.T) {
	assert.Equal(t, float64(0.59), StringToFloat64("abc", 0.59))
}

func Test_Float64ToString(t *testing.T) {
	assert.Equal(t, "3.14159", Float64ToString(3.14159))
}

func Test_StringToInt(t *testing.T) {
	assert.Equal(t, 31337, StringToInt("31337", 100))
}

func Test_StringToInt_ReturnDefaultValue(t *testing.T) {
	assert.Equal(t, 100, StringToInt("abc", 100))
}

func Test_IntToString(t *testing.T) {
	assert.Equal(t, "31337", IntToString(31337))
}

func Test_StringToInt8(t *testing.T) {
	assert.Equal(t, int8(20), StringToInt8("20", 10))
}

func Test_StringToInt8_ReturnDefaultValue(t *testing.T) {
	assert.Equal(t, int8(10), StringToInt8("aa", 10))
}

func Test_Int8ToString(t *testing.T) {
	assert.Equal(t, "20", Int8ToString(20))
}

func Test_StringToInt16(t *testing.T) {
	assert.Equal(t, int16(1337), StringToInt16("1337", 1000))
}

func Test_StringToInt16_ReturnDefaultValue(t *testing.T) {
	assert.Equal(t, int16(1000), StringToInt16("aabb", 1000))
}

func Test_Int16ToString(t *testing.T) {
	assert.Equal(t, "1337", Int16ToString(1337))
}

func Test_StringToInt32(t *testing.T) {
	assert.Equal(t, int32(31337), StringToInt32("31337", 13123))
}

func Test_StringToInt32_ReturnDefaultValue(t *testing.T) {
	assert.Equal(t, int32(13123), StringToInt32("aabb", 13123))
}

func Test_Int32ToString(t *testing.T) {
	assert.Equal(t, "31337", Int32ToString(31337))
}

func Test_StringToInt64(t *testing.T) {
	assert.Equal(t, int64(31337), StringToInt64("31337", 13123))
}

func Test_StringToInt64_ReturnDefaultValue(t *testing.T) {
	assert.Equal(t, int64(13123), StringToInt64("aabb", 13123))
}

func Test_Int64ToString(t *testing.T) {
	assert.Equal(t, "31337", Int64ToString(31337))
}

func Test_StringToUint(t *testing.T) {
	assert.Equal(t, uint(31337), StringToUint("31337", 13123))
}

func Test_StringToUint_ReturnDefaultValue(t *testing.T) {
	assert.Equal(t, uint(13123), StringToUint("aabb", 13123))
}

func Test_UintToString(t *testing.T) {
	assert.Equal(t, "31337", UintToString(31337))
}

func Test_StringToUint8(t *testing.T) {
	assert.Equal(t, uint8(128), StringToUint8("128", 100))
}

func Test_StringToUint8_ReturnDefaultValue(t *testing.T) {
	assert.Equal(t, uint8(100), StringToUint8("abc", 100))
}

func Test_Uint8ToString(t *testing.T) {
	assert.Equal(t, "128", Uint8ToString(128))
}

func Test_StringToUint16(t *testing.T) {
	assert.Equal(t, uint16(1337), StringToUint16("1337", 1000))
}

func Test_StringToUint16_ReturnDefaultValue(t *testing.T) {
	assert.Equal(t, uint16(1000), StringToUint16("aabb", 1000))
}

func Test_Uint16ToString(t *testing.T) {
	assert.Equal(t, "1337", Uint16ToString(1337))
}

func Test_StringToUint32(t *testing.T) {
	assert.Equal(t, uint32(31337), StringToUint32("31337", 13123))
}

func Test_StringToUint32_ReturnDefaultValue(t *testing.T) {
	assert.Equal(t, uint32(13123), StringToUint32("aabb", 13123))
}

func Test_Uint32ToString(t *testing.T) {
	assert.Equal(t, "31337", Uint32ToString(31337))
}

func Test_StringToUint64(t *testing.T) {
	assert.Equal(t, uint64(31337), StringToUint64("31337", 13123))
}

func Test_StringToUint64_ReturnDefaultValue(t *testing.T) {
	assert.Equal(t, uint64(13123), StringToUint64("aabb", 13123))
}

func Test_Uint64ToString(t *testing.T) {
	assert.Equal(t, "31337", Uint64ToString(31337))
}
