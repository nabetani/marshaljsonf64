package marshaljsonf64_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/nabetani/marshaljsonf64"
)

type Watermelon struct {
	Foo []float64
	Bar [2]float64
}

type Watermelon64 struct {
	Foo []float32
	Bar [2]float32
}

func (o Watermelon64) MarshalJSON() ([]byte, error) {
	return marshaljsonf64.Impl(&o, reflect.TypeOf(o))
}

func TestArray(t *testing.T) {
	w64 := Watermelon64{
		Foo: []float32{F1, F2},
		Bar: [2]float32{F3, F4},
	}
	j, err := json.Marshal(w64)
	if err != nil {
		t.Errorf("err = %v, want nil", err)
		return
	}
	w := Watermelon{}
	err = json.Unmarshal(j, &w)
	if err != nil {
		t.Errorf("err = %v, want nil", err)
		return
	}
	if len(w.Foo) != 2 || w.Foo[0] != F1 || w.Foo[1] != F2 {
		t.Errorf("w.Foo = %v, want [%v,%v]", w.Foo, F1, F2)
	}
	if w.Bar[0] != F3 || w.Bar[1] != F4 {
		t.Errorf("w.Bar = %v, want [%v,%v]", w.Bar, F3, F4)
	}
}
