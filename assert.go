package assert

import "testing"
import "reflect"

var green = "\033[32m"
var red = "\033[31m"
var end = "\033[0m"

type A struct {
	T *testing.T
}

func (a A) Nil(i interface{}, msg ...interface{}) {
	if !isNil(i) {
		a.T.Fatalf(red+"%v"+end+" != nil", i)
		a.T.Fatal(msg...)
	}
}

func (a A) NotNil(i interface{}, msg ...interface{}) {
	if isNil(i) {
		a.T.Fatalf(red+"%v"+end+" == nil", i)
		a.T.Fatal(msg...)
	}
}

func (a A) Equal(x, y interface{}, msg ...interface{}) {
	if !equal(x, y) {
		a.T.Fatalf(red+"%v"+end+" != "+red+"%v"+end+" \n", x, y)
		a.T.Fatal(msg...)
	}
}

func (a A) NotEqual(x, y interface{}, msg ...interface{}) {
	if equal(x, y) {
		a.T.Fatalf(red+"%v"+end+" == "+red+"%v"+end+" \n", x, y)
		a.T.Fatal(msg...)
	}
}

func isNil(i interface{}) bool {
	if i == nil {
		return true
	}

	value := reflect.ValueOf(i)
	kind := value.Kind()

	if kind >= reflect.Chan && kind <= reflect.Slice && value.IsNil() {
		return true
	}

	return false
}

func equal(x, y interface{}) bool {
	if x == nil || y == nil {
		return x == y
	}

	return reflect.DeepEqual(x, y)
}
