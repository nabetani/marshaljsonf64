package marshaljsonf64_test

import (
	"encoding/json"
	"log"
	"reflect"
	"testing"

	"github.com/nabetani/marshaljsonf64"
)

type Tangerine float64
type Plum [1]float64

type Lychee struct {
	Foo float64
}

type Nectarine struct {
	Bar float64
}

type Mandarine struct {
	Baz float64
}

type Mango struct {
	Tangerine
	Nectarine
	Mandarine `json:"Mandarine"`
	Plum
	Lychee `json:"lychee"`
	Qux    float64
}

type Tangerine64 float32
type Plum64 [1]float32

type Lychee64 struct {
	Foo float32
}
type Nectarine64 struct {
	Bar float32
}
type Mandarine64 struct {
	Baz float32
}

type Mango64 struct {
	Nectarine64
	Tangerine64 `json:"Tangerine"`
	Plum64      `json:"Plum"`
	Lychee64    `json:"lychee"`
	Mandarine64 `json:"Mandarine"`
	Qux         float32
}

func (o Lychee64) MarshalJSON() ([]byte, error) {
	return marshaljsonf64.Impl(&o, reflect.TypeOf(o))
}

func (o Mango64) MarshalJSON() ([]byte, error) {
	return marshaljsonf64.Impl(&o, reflect.TypeOf(o))
}

func (o Mandarine64) MarshalJSON() ([]byte, error) {
	return marshaljsonf64.Impl(&o, reflect.TypeOf(o))
}

func TestEmbedded(t *testing.T) {
	m64 := Mango64{
		Tangerine64: F3,
		Mandarine64: Mandarine64{Baz: F6},
		Nectarine64: Nectarine64{Bar: F4},
		Plum64:      [1]float32{F5},
		Lychee64:    Lychee64{Foo: F1},
		Qux:         F2,
	}
	j, err := json.Marshal(m64)
	log.Println(string(j))
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
	if m.Tangerine != F3 {
		t.Errorf("m.Tangerine=%v, want %v", m.Tangerine, F3)
	}
	if m.Bar != F4 {
		t.Errorf("m.Bar=%v, want %v", m.Bar, F4)
	}
	if m.Plum[0] != F5 {
		t.Errorf("m.Plum[0]=%v, want %v", m.Plum[0], F5)
	}
	if m.Baz != F6 {
		t.Errorf("m.Baz=%v, want %v", m.Baz, F6)
	}
}

type Pomegranate struct {
	Nectarine
}

type Pomegranate64 struct {
	Nectarine64
}

func (o Pomegranate64) MarshalJSON() ([]byte, error) {
	return marshaljsonf64.Impl(&o, reflect.TypeOf(o))
}

func TestEmbedded2(t *testing.T) {
	p64 := Pomegranate64{Nectarine64: Nectarine64{Bar: F4}}
	j, err := json.Marshal(p64)
	log.Println(string(j))
	if err != nil {
		t.Errorf("err = %v, want nil", err)
		return
	}
	p := Pomegranate{}
	err = json.Unmarshal(j, &p)
	if p.Bar != F4 {
		t.Errorf("p.Bar=%v, want %v", p.Bar, F4)
	}
}
