package future

import (
	"github.com/Trendyol/go-future/test/assert"
	"testing"
)

func TestGetParam_ShouldReturnValue(t *testing.T) {
	// Given
	type StructType struct {
		field string
	}

	sliceParam := []int{1, 2, 3, 4, 5}
	mapParam := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	structParam := StructType{field: "test"}
	var nilParam *string
	params := Params{"paramStr", 10, true, sliceParam, mapParam, structParam, &structParam, nilParam}

	// When-Then
	t.Run("Get Param As Any", func(t *testing.T) {
		actual, err := GetParam[any](params, 0)
		assert.Nil(t, err)
		assert.Equal(t, "paramStr", actual)
	})

	t.Run("String Param", func(t *testing.T) {
		actual, err := GetParam[string](params, 0)
		assert.Nil(t, err)
		assert.Equal(t, "paramStr", actual)
	})

	t.Run("Int Param", func(t *testing.T) {
		actual, err := GetParam[int](params, 1)
		assert.Nil(t, err)
		assert.Equal(t, 10, actual)
	})

	t.Run("Bool Param", func(t *testing.T) {
		actual, err := GetParam[bool](params, 2)
		assert.Nil(t, err)
		assert.Equal(t, true, actual)
	})

	t.Run("Slice Param", func(t *testing.T) {
		actual, err := GetParam[[]int](params, 3)
		assert.Nil(t, err)
		assert.Equal(t, sliceParam, actual)
	})

	t.Run("Map Param", func(t *testing.T) {
		actual, err := GetParam[map[string]int](params, 4)
		assert.Nil(t, err)
		assert.Equal(t, mapParam, actual)
	})

	t.Run("Struct Param", func(t *testing.T) {
		actual, err := GetParam[StructType](params, 5)
		assert.Nil(t, err)
		assert.Equal(t, structParam, actual)
	})

	t.Run("Struct Pointer Param", func(t *testing.T) {
		actual, err := GetParam[*StructType](params, 6)
		assert.Nil(t, err)
		assert.Equal(t, &structParam, actual)
	})

	t.Run("Nil Param", func(t *testing.T) {
		actual, err := GetParam[*string](params, 7)
		assert.Nil(t, err)
		assert.Nil(t, actual)
	})
}

func TestGetParam_ShouldReturnError_When_InvalidTypeConversion(t *testing.T) {
	// Given
	params := Params{"paramStr", 1}

	// When
	value, err := GetParam[int](params, 0)
	// Then
	assert.Equal(t, ErrParamTypeConversionFailed, err)
	assert.Equal(t, value, 0)
}

func TestParams_GetParam(t *testing.T) {
	params := Params{10, int32(20), int64(30), uint(40), uint8(50), uint32(60), uint64(70), float32(80.5), float64(90.5), "hello", true, []byte("world")}

	t.Run("GetIntParam", func(t *testing.T) {
		actual := params.GetIntParam(0)
		assert.Equal(t, 10, actual)
	})

	t.Run("GetInt32Param", func(t *testing.T) {
		actual := params.GetInt32Param(1)
		assert.Equal(t, int32(20), actual)
	})

	t.Run("GetInt64Param", func(t *testing.T) {
		actual := params.GetInt64Param(2)
		assert.Equal(t, int64(30), actual)
	})

	t.Run("GetUIntParam", func(t *testing.T) {
		actual := params.GetUIntParam(3)
		assert.Equal(t, uint(40), actual)
	})

	t.Run("GetUInt8Param", func(t *testing.T) {
		actual := params.GetUInt8Param(4)
		assert.Equal(t, uint8(50), actual)
	})

	t.Run("GetUInt32Param", func(t *testing.T) {
		actual := params.GetUInt32Param(5)
		assert.Equal(t, uint32(60), actual)
	})

	t.Run("GetUInt64Param", func(t *testing.T) {
		actual := params.GetUInt64Param(6)
		assert.Equal(t, uint64(70), actual)
	})

	t.Run("GetFloat32Param", func(t *testing.T) {
		actual := params.GetFloat32Param(7)
		assert.Equal(t, float32(80.5), actual)
	})

	t.Run("GetFloat64Param", func(t *testing.T) {
		actual := params.GetFloat64Param(8)
		assert.Equal(t, float64(90.5), actual)
	})

	t.Run("GetStringParam", func(t *testing.T) {
		actual := params.GetStringParam(9)
		assert.Equal(t, "hello", actual)
	})

	t.Run("GetBoolParam", func(t *testing.T) {
		actual := params.GetBoolParam(10)
		assert.Equal(t, true, actual)
	})

	t.Run("GetBytesParam", func(t *testing.T) {
		actual := params.GetBytesParam(11)
		assert.Equal(t, []byte("world"), actual)
	})
}
