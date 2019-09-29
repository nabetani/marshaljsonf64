package marshaljsonf64_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"marshaljsonf64"
)

type Avocado struct {
	Foo int
	Bar float64
	Baz float64
}

type AvocadoF64 struct {
	Foo int
	Bar float64
	Baz float32
}

func (o AvocadoF64) MarshalJSON() ([]byte, error) {
	return marshaljsonf64.MarshalJSONF64(&o, reflect.TypeOf(o))
}

const F1 = 0xf1000000
const F2 = 0xf2000000
const F3 = 0xf3000000
const F4 = 0xf4000000

func TestSimple(t *testing.T) {
	a64 := AvocadoF64{Foo: F1, Bar: F2, Baz: F3}
	j, err := json.Marshal(a64)
	if err != nil {
		t.Errorf("err = %v, want nil", err)
		return
	}
	a := Avocado{}
	err = json.Unmarshal(j, &a)
	if err != nil {
		t.Errorf("err = %v, want nil", err)
		return
	}
	if float64(a.Foo) != F1 {
		t.Errorf("a.Foo=%v, want %v", a.Foo, F1)
	}
	if float64(a.Bar) != F2 {
		t.Errorf("a.Bar=%v, want %v", a.Bar, F2)
	}
	if float64(a.Baz) != F3 {
		t.Errorf("a.Baz=%v, want %v", a.Baz, F3)
	}
}
