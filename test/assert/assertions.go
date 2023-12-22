package assert

import (
	"bytes"
	"reflect"
	"testing"
)

func Nil(t *testing.T, object interface{}) bool {
	if isNil(object) {
		return true
	}
	t.Helper()
	t.Errorf("Expected nil, but got: %#v", object)
	return true
}

func NotNil(t *testing.T, object interface{}) bool {
	if !isNil(object) {
		return true
	}
	t.Helper()
	t.Errorf("Expected value not to be nil.")
	return true
}

func Equal(t *testing.T, expected, actual interface{}) bool {
	if ObjectsAreEqual(expected, actual) {
		return true
	}
	t.Helper()
	t.Errorf("Not equal: \n expected: %#v\n actual  : %#v", expected, actual)
	return true
}

func NotEqual(t *testing.T, expected, actual interface{}) bool {
	t.Helper()
	if ObjectsAreEqual(expected, actual) {
		t.Errorf("Should not be: %#v\n", actual)
	}
	return true
}

func ObjectsAreEqual(expected, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}

	exp, ok := expected.([]byte)
	if !ok {
		return reflect.DeepEqual(expected, actual)
	}

	act, ok := actual.([]byte)
	if !ok {
		return false
	}
	if exp == nil || act == nil {
		return exp == nil && act == nil
	}
	return bytes.Equal(exp, act)
}

func containsKind(kinds []reflect.Kind, kind reflect.Kind) bool {
	for i := 0; i < len(kinds); i++ {
		if kind == kinds[i] {
			return true
		}
	}

	return false
}

func isNil(object interface{}) bool {
	if object == nil {
		return true
	}

	value := reflect.ValueOf(object)
	kind := value.Kind()
	isNilableKind := containsKind(
		[]reflect.Kind{
			reflect.Chan, reflect.Func,
			reflect.Interface, reflect.Map,
			reflect.Ptr, reflect.Slice},
		kind)

	if isNilableKind && value.IsNil() {
		return true
	}

	return false
}
