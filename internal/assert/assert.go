package assert

import (
	"reflect"
	"testing"
)

// Equal compares non-comparable types like slices and maps.
func Equal[T any](t *testing.T, got, want T) {
	t.Helper()

	if !isEqual(got, want) {
		t.Errorf("got: %v; want: %v", got, want)
	}
}

// True is a helper to check that a value is true.
func True(t *testing.T, got bool) {
	t.Helper()

	if !got {
		t.Errorf("got: %t; want: true", got)
	}
}

// Nil is a helper to check that a value is nil, using the isNil() function below.
func Nil(t *testing.T, got any) {
	t.Helper()

	if !isNil(got) {
		t.Errorf("got: %v; want: nil", got)
	}
}

func isEqual[T any](got, want T) bool {
	// First check if both values are nil using the isNil() function below.
	if isNil(got) && isNil(want) {
		return true
	}

	// Otherwise use reflect.DeepEqual to check if they are the same.
	return reflect.DeepEqual(got, want)
}

func isNil(v any) bool {
	// Returns true if v equals nil.
	if v == nil {
		return true
	}

	// Use reflection to check the underlying type of v, and return true if it
	// is a nullable type (e.g. pointer, map or slice) with a value of nil.
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice, reflect.UnsafePointer:
		return rv.IsNil()
	}

	// Other types like string, bool, int are never nil
	return false
}
