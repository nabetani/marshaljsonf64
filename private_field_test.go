package marshaljsonf64_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/nabetani/marshaljsonf64"
)

type Strawberry struct {
	Foo int
	foo int // private
	Bar float64
	bar float64 // private
}

type Strawberry64 struct {
	Foo int
	foo int // private
	Bar float32
	bar float32 // private
}

func (o Strawberry64) MarshalJSON() ([]byte, error) {
	return marshaljsonf64.Impl(&o, reflect.TypeOf(o))
}

func TestWithPrivateFields(t *testing.T) {
	s64 := Strawberry64{Foo: F1, foo: F2, Bar: F3, bar: F4}
	j, err := json.Marshal(s64)
	if err != nil {
		t.Errorf("err = %v, want nil", err)
		return
	}
	s := Strawberry{}
	err = json.Unmarshal(j, &s)
	if err != nil {
		t.Errorf("err = %v, want nil", err)
		return
	}
	if float64(s.Foo) != F1 {
		t.Errorf("s.Foo=%v, want %v", s.Foo, F1)
	}
	if float64(s.Bar) != F3 {
		t.Errorf("s.Bar=%v, want %v", s.Bar, F3)
	}
}
