package marshaljsonf64_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"marshaljsonf64"
)

type Durian struct {
	Foo  float64
	Bar  float64
	Baz  *float64
	Qux  *float64
	Quux *float64
}

type Durian64 struct {
	Foo  float64
	Bar  float32
	Baz  *float64
	Qux  *float32
	Quux *float32
}

func (o Durian64) MarshalJSON() ([]byte, error) {
	return marshaljsonf64.MarshalJSONF64(&o, reflect.TypeOf(o))
}

func TestPtrToF32(t *testing.T) {
	d64 := Durian64{
		Foo:  F1,
		Bar:  F2,
		Baz:  func() *float64 { var f float64 = F3; return &f }(),
		Qux:  func() *float32 { var f float32 = F4; return &f }(),
		Quux: nil,
	}
	j, err := json.Marshal(d64)
	if err != nil {
		t.Errorf("err = %v, want nil", err)
		return
	}
	d := Durian{}
	err = json.Unmarshal(j, &d)
	if err != nil {
		t.Errorf("err = %v, want nil", err)
		return
	}
	if (d.Foo) != F1 {
		t.Errorf("d.Foo=%v, want %v", d.Foo, F1)
	}
	if (d.Bar) != F2 {
		t.Errorf("d.Bar=%v, want %v", d.Bar, F2)
	}
	if (*d.Baz) != F3 {
		t.Errorf("*d.Baz=%v, want %v", *d.Baz, F3)
	}
	if (*d.Qux) != F4 {
		t.Errorf("*d.Qux=%v, want %v", *d.Qux, F4)
	}
	if d.Quux != nil {
		t.Errorf("*d.Quux=%v, want nil", *d.Quux)
	}
}
