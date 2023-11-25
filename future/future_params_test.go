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
