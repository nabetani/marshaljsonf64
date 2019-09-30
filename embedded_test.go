package marshaljsonf64_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/nabetani/marshaljsonf64"
)

type Lychee struct {
	Foo float64
}

type Nectarine struct {
	Bar float64
}

type Mango struct {
	Nectarine
	Lychee `json:"lychee"`
	Qux    float64
}

type Lychee64 struct {
	Foo float32
}

type Mango64 struct {
	Nectarine
	Lychee64 `json:"lychee"`
	Qux      float32
}

func (o Lychee64) MarshalJSON() ([]byte, error) {
	return marshaljsonf64.Impl(&o, reflect.TypeOf(o))
}

func (o Mango64) MarshalJSON() ([]byte, error) {
	return marshaljsonf64.Impl(&o, reflect.TypeOf(o))
}

func TestEmbedded(t *testing.T) {
	m64 := Mango64{
		Nectarine: Nectarine{Bar: F4},
		Lychee64:  Lychee64{Foo: F1},
		Qux:       F2,
	}
	j, err := json.Marshal(m64)
	if err != nil {
		t.Errorf("err = %v, want nil", err)
		return
	}
	m := Mango{}
	err = json.Unmarshal(j, &m)
	if err != nil {
		t.Errorf("err = %v, want nil", err)
		return
	}
	if m.Foo != F1 {
		t.Errorf("m.Foo=%v, want %v", m.Foo, F1)
	}
	if m.Qux != F2 {
		t.Errorf("m.Qux=%v, want %v", m.Qux, F2)
	}
	if m.Bar != F4 {
		t.Errorf("m.Bar=%v, want %v", m.Bar, F4)
	}
}
