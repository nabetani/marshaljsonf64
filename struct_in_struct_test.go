package marshaljsonf64_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"marshaljsonf64"
)

type Peach struct {
	Foo float64
}

type Lime struct {
	Foo Peach
}

type Guava struct {
	Foo Lime
}

type Grape struct {
	Foo Peach
	Bar Peach
	Baz Guava
}

type Peach64 struct {
	Foo float64
}

type Lime64 struct {
	Foo Peach64
}

type Guava64 struct {
	Foo Lime64
}

type Grape64 struct {
	Foo Peach64
	Bar Peach64
	Baz Guava64
}

func (o Peach64) MarshalJSON() ([]byte, error) {
	return marshaljsonf64.Impl(&o, reflect.TypeOf(o))
}

func TestStructInStruct(t *testing.T) {
	g64 := Grape64{
		Foo: Peach64{F1},
		Bar: Peach64{F2},
	}
	g64.Baz.Foo.Foo.Foo = F3
	j, err := json.Marshal(g64)
	if err != nil {
		t.Errorf("err = %v, want nil", err)
		return
	}
	g := Grape64{}
	err = json.Unmarshal(j, &g)
	if err != nil {
		t.Errorf("err = %v, want nil", err)
		return
	}
	if g.Foo.Foo != F1 {
		t.Errorf("g.Foo.Foo=%v, want %v", g.Foo.Foo, F1)
	}
	if g.Bar.Foo != F2 {
		t.Errorf("g.Bar.Foo=%v, want %v", g.Bar.Foo, F1)
	}
	if g.Baz.Foo.Foo.Foo != F3 {
		t.Errorf("g.Baz.Foo.Foo.Foo=%v, want %v", g.Baz.Foo.Foo.Foo, F3)
	}
}
