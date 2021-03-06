package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/nabetani/marshaljsonf64"
)

type banana struct {
	Foo float32
}

type cherry struct {
	Foo float32
}

// MarshalJSON この関数を実装すると json.Marshal 内で呼ばれる
func (o cherry) MarshalJSON() ([]byte, error) {
	return marshaljsonf64.Impl(&o, reflect.TypeOf(o))
}

func main() {
	b := banana{Foo: 2415919104}
	c := cherry{Foo: 2415919104}
	jb, _ := json.Marshal(b)
	jc, _ := json.Marshal(c)
	fmt.Println("banana:", string(jb))
	//=> banana: {"Foo":2415919000}

	fmt.Println("cherry:", string(jc))
	//=> cherry: {"Foo":2415919104}
}
