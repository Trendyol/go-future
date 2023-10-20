package future_test

import (
	"errors"
	"github.com/Trendyol/go-future/future"
	"github.com/Trendyol/go-future/test/assert"
	"math/rand"
	"testing"
	"time"
)

func createNFuture(futureCount int, hasError bool) []*future.Future[string] {
	fn := func() (string, error) {
		waitMillis := rand.Intn(100)
		time.Sleep(time.Duration(waitMillis) * time.Millisecond)
		return "expectedVal", nil
	}
	fnWithError := func() (string, error) {
		waitMillis := rand.Intn(100)
		time.Sleep(time.Duration(waitMillis) * time.Millisecond)
		return "", errors.New("error occurred")
	}
	futures := make([]*future.Future[string], 0)
	for i := 0; i < futureCount; i++ {
		if hasError && i == futureCount/2 {
			futures = append(futures, future.Run[string](fnWithError))
		} else {
			futures = append(futures, future.Run[string](fn))
		}
	}
	return futures
}

func Test_Future_Returns_Value(t *testing.T) {
	// Given
	expectedVal := "future-test"
	fn := func() (string, error) {
		return expectedVal, nil
	}

	// When
	f := future.Run[string](fn)
	result, err := f.Get()
	result2 := f.GetResult()

	// Then
	assert.Nil(t, err)
	assert.Equal(t, expectedVal, result)
	assert.Equal(t, result, result2)
}

func Test_Future_Returns_Pointer_Ref(t *testing.T) {
	// Given
	expectedVal := "future-test"
	fn := func() (*string, error) {
		return &expectedVal, nil
	}

	// When
	f := future.Run[*string](fn)
	result, err := f.Get()

	// Then
	assert.Nil(t, err)
	assert.Equal(t, "future-test", *result)
}

func Test_Future_Returns_Error(t *testing.T) {
	// Given
	fn := func() (*string, error) {
		return nil, errors.New("an error")
	}
	// When
	f := future.Run[*string](fn)
	result, err := f.Get()

	// Then
	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func Test_Future_RunWithParam_Returns_Value(t *testing.T) {
	// Given
	fn := func(myParam string) (string, error) {
		return "future-test-" + myParam, nil
	}

	// When
	f := future.RunWithParam[string, string](fn, "withParam")
	result, err := f.Get()

	// Then
	assert.Nil(t, err)
	assert.Equal(t, "future-test-withParam", result)
}

func Test_Future_RunWithParam_Returns_Pointer_Ref(t *testing.T) {
	// Given
	expectedVal := "param-is-null"
	fn := func(myParam *int) (*string, error) {
		if myParam == nil {
			return &expectedVal, nil
		}
		return nil, nil
	}

	// When
	f := future.RunWithParam[*string, *int](fn, nil)
	result, err := f.Get()

	// Then
	assert.Nil(t, err)
	assert.Equal(t, expectedVal, *result)
}

func Test_Future_RunWithParam_Returns_Error(t *testing.T) {
	// Given
	fn := func(myParam int) (*string, error) {
		return nil, errors.New("an error")
	}
	// When
	f := future.RunWithParam[*string, int](fn, 10)
	result, err := f.Get()

	// Then
	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func Test_Future_RunWithParam_When_Pass_Loop_Variable(t *testing.T) {
	// Given
	futures := make([]*future.Future[int], 0)
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, value := range numbers {
		f := future.RunWithParam(func(v int) (int, error) {
			println(v)
			return v, nil
		}, value)
		futures = append(futures, f)
	}

	// When
	result, err := future.GetAll(futures)

	// Then
	assert.Nil(t, err)
	assert.Equal(t, numbers, result)
}

func Test_Future_Get_All(t *testing.T) {
	// Given
	futures := createNFuture(100, false)

	expectedResult := make([]string, 0)
	for i := 0; i < 100; i++ {
		expectedResult = append(expectedResult, "expectedVal")
	}

	// When
	result, err := future.GetAll(futures)

	// Then
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}

func Test_Future_Get_All_When_Any_Error_Occurred(t *testing.T) {
	// Given
	futures := createNFuture(100, true)
	// When
	result, err := future.GetAll(futures)
	// Then
	assert.NotNil(t, err)
	assert.Equal(t, 0, len(result))
}

func Test_Future_WaitAllSilently(t *testing.T) {
	// Given
	futures := createNFuture(100, true)
	// When
	future.WaitAllSilently(futures)

	// Then
	doneCount := 0
	for i := range futures {
		if futures[i].IsDone {
			doneCount++
		}
	}
	assert.Equal(t, len(futures), doneCount)
}

func Test_Future_WaitFor(t *testing.T) {
	// Given
	f1 := future.Run[any](func() (any, error) {
		return "expectedVal", nil
	})

	f2 := future.Run[any](func() (any, error) {
		return 1, nil
	})

	// When
	err := future.WaitFor(f1, f2)
	result1 := future.GetResult[string](f1)
	result2 := future.GetResult[int](f2)

	// Then
	assert.Nil(t, err)
	assert.Equal(t, "expectedVal", result1)
	assert.Equal(t, 1, result2)
}
