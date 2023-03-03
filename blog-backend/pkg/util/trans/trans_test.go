package trans

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Trans(t *testing.T) {
	str := String("tea")
	strVal := StringValue(str)
	assert.Equal(t, "tea", strVal)
	assert.Equal(t, "", StringValue(nil))

	strSlice := StringSlice([]string{"tea"})
	strSliceVal := StringSliceValue(strSlice)
	assert.Equal(t, []string{"tea"}, strSliceVal)
	assert.Nil(t, StringSlice(nil))
	assert.Nil(t, StringSliceValue(nil))

	b := Bool(true)
	bVal := BoolValue(b)
	assert.Equal(t, true, bVal)
	assert.Equal(t, false, BoolValue(nil))

	bSlice := BoolSlice([]bool{false})
	bSliceVal := BoolSliceValue(bSlice)
	assert.Equal(t, []bool{false}, bSliceVal)
	assert.Nil(t, BoolSlice(nil))
	assert.Nil(t, BoolSliceValue(nil))

	f64 := Float64(2.00)
	f64Val := Float64Value(f64)
	assert.Equal(t, 2.00, f64Val)
	assert.Equal(t, float64(0), Float64Value(nil))

	f32 := Float32(2.00)
	f32Val := Float32Value(f32)
	assert.Equal(t, float32(2.00), f32Val)
	assert.Equal(t, float32(0), Float32Value(nil))

	f64Slice := Float64Slice([]float64{2.00})
	f64SliceVal := Float64ValueSlice(f64Slice)
	assert.Equal(t, []float64{2.00}, f64SliceVal)
	assert.Nil(t, Float64Slice(nil))
	assert.Nil(t, Float64ValueSlice(nil))

	f32Slice := Float32Slice([]float32{2.00})
	f32SliceVal := Float32ValueSlice(f32Slice)
	assert.Equal(t, []float32{2.00}, f32SliceVal)
	assert.Nil(t, Float32Slice(nil))
	assert.Nil(t, Float32ValueSlice(nil))

	i := Int(1)
	iVal := IntValue(i)
	assert.Equal(t, 1, iVal)
	assert.Equal(t, 0, IntValue(nil))

	i8 := Int8(int8(1))
	i8Val := Int8Value(i8)
	assert.Equal(t, int8(1), i8Val)
	assert.Equal(t, int8(0), Int8Value(nil))

	i16 := Int16(int16(1))
	i16Val := Int16Value(i16)
	assert.Equal(t, int16(1), i16Val)
	assert.Equal(t, int16(0), Int16Value(nil))

	i32 := Int32(int32(1))
	i32Val := Int32Value(i32)
	assert.Equal(t, int32(1), i32Val)
	assert.Equal(t, int32(0), Int32Value(nil))

	i64 := Int64(int64(1))
	i64Val := Int64Value(i64)
	assert.Equal(t, int64(1), i64Val)
	assert.Equal(t, int64(0), Int64Value(nil))

	iSlice := IntSlice([]int{1})
	iSliceVal := IntValueSlice(iSlice)
	assert.Equal(t, []int{1}, iSliceVal)
	assert.Nil(t, IntSlice(nil))
	assert.Nil(t, IntValueSlice(nil))

	i8Slice := Int8Slice([]int8{1})
	i8ValSlice := Int8ValueSlice(i8Slice)
	assert.Equal(t, []int8{1}, i8ValSlice)
	assert.Nil(t, Int8Slice(nil))
	assert.Nil(t, Int8ValueSlice(nil))

	i16Slice := Int16Slice([]int16{1})
	i16ValSlice := Int16ValueSlice(i16Slice)
	assert.Equal(t, []int16{1}, i16ValSlice)
	assert.Nil(t, Int16Slice(nil))
	assert.Nil(t, Int16ValueSlice(nil))

	i32Slice := Int32Slice([]int32{1})
	i32ValSlice := Int32ValueSlice(i32Slice)
	assert.Equal(t, []int32{1}, i32ValSlice)
	assert.Nil(t, Int32Slice(nil))
	assert.Nil(t, Int32ValueSlice(nil))

	i64Slice := Int64Slice([]int64{1})
	i64ValSlice := Int64ValueSlice(i64Slice)
	assert.Equal(t, []int64{1}, i64ValSlice)
	assert.Nil(t, Int64Slice(nil))
	assert.Nil(t, Int64ValueSlice(nil))

	ui := Uint(1)
	uiVal := UintValue(ui)
	assert.Equal(t, uint(1), uiVal)
	assert.Equal(t, uint(0), UintValue(nil))

	ui8 := Uint8(uint8(1))
	ui8Val := Uint8Value(ui8)
	assert.Equal(t, uint8(1), ui8Val)
	assert.Equal(t, uint8(0), Uint8Value(nil))

	ui16 := Uint16(uint16(1))
	ui16Val := Uint16Value(ui16)
	assert.Equal(t, uint16(1), ui16Val)
	assert.Equal(t, uint16(0), Uint16Value(nil))

	ui32 := Uint32(uint32(1))
	ui32Val := Uint32Value(ui32)
	assert.Equal(t, uint32(1), ui32Val)
	assert.Equal(t, uint32(0), Uint32Value(nil))

	ui64 := Uint64(uint64(1))
	ui64Val := Uint64Value(ui64)
	assert.Equal(t, uint64(1), ui64Val)
	assert.Equal(t, uint64(0), Uint64Value(nil))

	uiSlice := UintSlice([]uint{1})
	uiValSlice := UintValueSlice(uiSlice)
	assert.Equal(t, []uint{1}, uiValSlice)
	assert.Nil(t, UintSlice(nil))
	assert.Nil(t, UintValueSlice(nil))

	ui8Slice := Uint8Slice([]uint8{1})
	ui8ValSlice := Uint8ValueSlice(ui8Slice)
	assert.Equal(t, []uint8{1}, ui8ValSlice)
	assert.Nil(t, Uint8Slice(nil))
	assert.Nil(t, Uint8ValueSlice(nil))

	ui16Slice := Uint16Slice([]uint16{1})
	ui16ValSlice := Uint16ValueSlice(ui16Slice)
	assert.Equal(t, []uint16{1}, ui16ValSlice)
	assert.Nil(t, Uint16Slice(nil))
	assert.Nil(t, Uint16ValueSlice(nil))

	ui32Slice := Uint32Slice([]uint32{1})
	ui32ValSlice := Uint32ValueSlice(ui32Slice)
	assert.Equal(t, []uint32{1}, ui32ValSlice)
	assert.Nil(t, Uint32Slice(nil))
	assert.Nil(t, Uint32ValueSlice(nil))

	ui64Slice := Uint64Slice([]uint64{1})
	ui64ValSlice := Uint64ValueSlice(ui64Slice)
	assert.Equal(t, []uint64{1}, ui64ValSlice)
	assert.Nil(t, Uint64Slice(nil))
	assert.Nil(t, Uint64ValueSlice(nil))

	tm := Time(time.Now())
	tmVal := TimeValue(tm)
	assert.Equal(t, time.Now(), tmVal)
	assert.Equal(t, time.Now(), TimeValue(nil))
}
