package future

import (
	"errors"
	"fmt"
	"github.com/Trendyol/go-future/test/assert"
	"math/rand"
	"testing"
	"time"
)

func TestGroup_Get(t *testing.T) {
	// Given
	gr := Group[string]{}
	gr.Go(func() (string, error) {
		WaitRandomMillis(10)
		return "result", nil
	})

	gr.GoWithParams(func(p Params) (string, error) {
		WaitRandomMillis(10)
		str := p.GetStringParam(0)
		return str, nil
	}, Params{"param"})
	// When
	results, err := gr.Get()
	// Then
	assert.Nil(t, err)
	assert.Equal(t, []string{"result", "param"}, results)
}

func TestGroup_Get_WhenAnyErrorOccurred(t *testing.T) {
	// Given
	gr := Group[string]{}
	gr.Go(func() (string, error) {
		WaitRandomMillis(10)
		return "result", nil
	})

	gr.GoWithParams(func(p Params) (string, error) {
		WaitRandomMillis(10)
		str := p.GetStringParam(0)
		return str, errors.New("error")
	}, Params{"param"})
	// When
	results, err := gr.Get()
	// Then
	assert.NotNil(t, err)
	assert.Nil(t, results)
}

func TestGroup_WaitContinueOnError(t *testing.T) {
	// Given
	gr := Group[string]{}
	for i := 0; i < 100; i++ {
		gr.GoWithParams(func(params Params) (string, error) {
			index := params[0].(int)
			var err error
			if index%2 == 0 {
				err = fmt.Errorf("error:%d", index)
			}
			WaitRandomMillis(100)
			return fmt.Sprintf("result:%d", index), err
		}, Params{i})
	}
	// When
	errs := gr.WaitContinueOnError()

	// Then
	doneCount := 0
	futures := gr.GetFutures()
	for i := range futures {
		if futures[i].IsDone {
			doneCount++
		}
	}
	assert.Equal(t, 50, len(errs))
	assert.Equal(t, len(futures), doneCount)
}

func TestGroup_Wait(t *testing.T) {
	// Given
	gr := Group[string]{}
	for i := 0; i < 100; i++ {
		gr.GoWithParams(func(params Params) (string, error) {
			index := params.GetIntParam(0)
			var err error
			if index%2 == 0 {
				err = fmt.Errorf("error:%d", index)
			}
			WaitRandomMillis(100)
			return fmt.Sprintf("result:%d", index), err
		}, Params{i})
	}
	// When
	err := gr.Wait()

	// Then
	doneCount := 0
	futures := gr.GetFutures()
	for i := range futures {
		if futures[i].IsDone {
			doneCount++
		}
	}
	assert.NotNil(t, err)
	assert.NotEqual(t, len(futures), doneCount)
}

func WaitRandomMillis(maxWaitMillis int) {
	waitMillis := rand.Intn(maxWaitMillis)
	time.Sleep(time.Duration(waitMillis) * time.Millisecond)
}
