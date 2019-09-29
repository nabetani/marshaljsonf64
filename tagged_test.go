package marshaljsonf64_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"marshaljsonf64"
)

type Apricot struct {
	Foo float64 `json:"foo-tagged"`
	Bar float64 `json:"bar-tagged"`
}

type Apricot64 struct {
	Foo float32 `json:"foo-tagged"`
	Bar float64 `json:"bar-tagged"`
}

func (o Apricot64) MarshalJSON() ([]byte, error) {
	return marshaljsonf64.MarshalJSONF64(&o, reflect.TypeOf(o))
}

func TestWithNamedFields(t *testing.T) {
	a64 := Apricot64{Foo: F1, Bar: F2}
	j, err := json.Marshal(a64)
	if err != nil {
		t.Errorf("err = %v, want nil", err)
		return
	}
	a := Apricot{}
	err = json.Unmarshal(j, &a)
	if err != nil {
		t.Errorf("err = %v, want nil ( json=%q )", err, string(j))
		return
	}
	if float64(a.Foo) != F1 {
		t.Errorf("a.Foo=%v, want %v", a.Foo, F1)
	}
	if float64(a.Bar) != F2 {
		t.Errorf("a.Bar=%v, want %v", a.Bar, F2)
	}
}

// TODO: 埋め込み構造体
// TODO: *float32
// TODO: []float32
// TODO: []*float32
// TODO: 構造体内構造体
