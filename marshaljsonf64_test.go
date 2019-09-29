package marshaljsonf64_test

import (
	"encoding/json"
	"log"
	"reflect"
	"testing"

	"marshaljsonf64"
)

type Avocado struct {
	Foo int
	Bar float64
	Baz float32
}

type AvocadoF64 struct {
	Foo int
	Bar float64
	Baz float32
}

func (o AvocadoF64) MarshalJSON() ([]byte, error) {
	return marshaljsonf64.MarshalJSONF64(&o, reflect.TypeOf(o))
}

type Strawberry struct {
	Foo int
	foo int // private
	Bar float32
	bar float32 // private
}

type Strawberry64 struct {
	Foo int
	foo int // private
	Bar float32
	bar float32 // private
}

func (o Strawberry64) MarshalJSON() ([]byte, error) {
	return marshaljsonf64.MarshalJSONF64(&o, reflect.TypeOf(o))
}

type Apricot struct {
	Foo float32 `json:"foo-tagged"`
	Bar float64 `json:"bar-tagged"`
}

type Apricot64 struct {
	Foo float32 `json:"foo-tagged"`
	Bar float64 `json:"bar-tagged"`
}

func (o Apricot64) MarshalJSON() ([]byte, error) {
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

func TestWithNamedFields(t *testing.T) {
	a64 := Apricot64{Foo: F1, Bar: F2}
	j, err := json.Marshal(a64)
	if err != nil {
		t.Errorf("err = %v, want nil", err)
		return
	}
	log.Println(string(j))
	a := Apricot{}
	err = json.Unmarshal(j, &a)
	if err != nil {
		t.Errorf("err = %v, want nil ( json=%q )", err, string(j))
		return
	}
	if a.Foo != F1 {
		t.Errorf("a.Foo=%v, want %v", a.Foo, F1)
	}
	if a.Bar != F2 {
		t.Errorf("a.Bar=%v, want %v", a.Bar, F2)
	}
}

// TODO: 埋め込み構造体
// TODO: *float32
// TODO: []float32
// TODO: []*float32
// TODO: 構造体内構造体
// TODO: named fields
