package marshaljsonf64_test

import (
	"encoding/json"
	"math"
	"math/rand"
	"reflect"
	"testing"

	"github.com/nabetani/marshaljsonf64"
)

type Persimmon struct {
	Foo float32
}

type Persimmon64 struct {
	Foo float64
}

func (o Persimmon) MarshalJSON() ([]byte, error) {
	return marshaljsonf64.Impl(&o, reflect.TypeOf(o))
}

func TestRandom(t *testing.T) {
	rng := rand.New(rand.NewSource(1))
	for i := 0; i < 1000; {
		u := rng.Uint32()
		f := math.Float32frombits(u)
		if math.IsNaN(float64(f)) || math.IsInf(float64(f), 0) {
			continue
		}
		i++
		p := Persimmon{f}
		j, err := json.Marshal(p)
		if err != nil {
			t.Errorf("err = %v, want nil", err)
			return
		}
		q := Persimmon64{}
		err = json.Unmarshal(j, &q)
		if err != nil {
			t.Errorf("err = %v, want nil", err)
			return
		}
		if q.Foo != float64(p.Foo) {
			t.Errorf("q.Foo=%v, want %v, json=%v", q.Foo, p.Foo, string(j))
		}
	}
}
