package marshaljsonf64_test

import (
	"encoding/json"
	"reflect"
	"strconv"
	"testing"

	"marshaljsonf64"
)

type Lemon struct {
	Foo []*float64
	Bar [3]*float64
}

type Lemon64 struct {
	Foo []*float32
	Bar [3]*float32
}

func (o Lemon64) MarshalJSON() ([]byte, error) {
	return marshaljsonf64.Impl(&o, reflect.TypeOf(o))
}

func TestArrayOfPtr(t *testing.T) {
	p := func(v float32) *float32 { return &v }
	l64 := Lemon64{
		Foo: []*float32{p(F1), p(F2), nil},
		Bar: [3]*float32{p(F3), p(F4), nil},
	}
	j, err := json.Marshal(l64)
	if err != nil {
		t.Errorf("err = %v, want nil", err)
		return
	}
	l := Lemon{}
	err = json.Unmarshal(j, &l)
	if err != nil {
		t.Errorf("err = %v, want nil", err)
		return
	}
	repr := func(p *float64) string {
		if p == nil {
			return "nil"
		}
		return strconv.FormatFloat(*p, 'f', -1, 64)
	}
	if len(l.Foo) != 3 || *l.Foo[0] != F1 || *l.Foo[1] != F2 {
		t.Errorf("l.Foo = [%s,%s,%s], want [%v,%v,nil]",
			repr(l.Foo[0]), repr(l.Foo[1]), repr(l.Foo[2]),
			F1, F2)
	}
	if *l.Bar[0] != F3 || *l.Bar[1] != F4 {
		t.Errorf("l.Bar = [%s,%s,%s], want [%v,%v,nil]",
			repr(l.Bar[0]), repr(l.Bar[1]), repr(l.Bar[2]),
			F3, F4)
	}
}
