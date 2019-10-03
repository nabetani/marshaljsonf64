package marshaljsonf64_test

import (
	"encoding/json"
	"math"
	"strings"
	"testing"
)

func testNoFinite(t *testing.T, v float32, valtext string) {
	f := func(mod func(m *Mango64)) {
		m := Mango64{}
		mod(&m)
		j, err := json.Marshal(m)
		if err == nil {
			t.Errorf("err = nil, want non-nil( j = %v )", string(j))
			return
		}
		_, ok := err.(*json.MarshalerError)
		if !ok {
			t.Errorf("type of err is %T, want *json.MarshalerError", err)
		}
		prefix := "json: error calling MarshalJSON for type "
		if !strings.HasPrefix(err.Error(), prefix) {
			t.Errorf("err.Error()=%q, but it should have prefix %q", err.Error(), prefix)
		}

		suffix := "json: unsupported value: " + valtext
		if !strings.HasSuffix(err.Error(), suffix) {
			t.Errorf("err.Error()=%q, but it should have sufffix %q", err.Error(), suffix)
		}
	}
	f(func(m *Mango64) { m.Foo = v })
	f(func(m *Mango64) { m.Bar = v })
	f(func(m *Mango64) { m.Baz = v })
	f(func(m *Mango64) { m.Qux = v })
	f(func(m *Mango64) { m.Tangerine64 = Tangerine64(v) })
	f(func(m *Mango64) { m.Nectarine64.Bar = v })
	f(func(m *Mango64) { m.Plum64[0] = v })
	f(func(m *Mango64) { m.Lychee64.Foo = v })
	f(func(m *Mango64) { m.Mandarine64.Baz = v })
}

func TestNaN(t *testing.T) {
	testNoFinite(t, float32(math.NaN()), "NaN")
}
func TestPositiveInf(t *testing.T) {
	testNoFinite(t, float32(math.Inf(1)), "+Inf")
}

func TestNegativeInf(t *testing.T) {
	testNoFinite(t, float32(math.Inf(-1)), "-Inf")
}
