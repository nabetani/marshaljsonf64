# marshaljsonf64

float32 型の値を "encoding/json" で Marshal すると 10進数に丸められてしまう。

```go
type Hoge struct {
    Foo float32
    Bar float64
}

func main() {
    hoge := Hoge{Foo: 0x90000000, Bar: 0x90000000}
    j, _ := json.Marshal(hoge)
    fmt.Println(string(j)) //=> {"Foo":2415919000,"Bar":2415919104}
}
```

float64 である Bar の値 `2415919104` は、 `0x90000000` と、ぴったり。
float32 である Foo の値 `2415919000` は、 `0x8fffff98` と、代表値ではない値になる。困る。

困るので、対策を書いた。

## 使い方

```go
package main

import (
	"encoding/json"
	"fmt"
	"marshaljsonf64"
	"reflect"
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
	b := banana{Foo: 0x90000000}
	c := cherry{Foo: 0x90000000}
	jb, _ := json.Marshal(b)
	jc, _ := json.Marshal(c)
	fmt.Println("banana:", string(jb)) //=> banana: {"Foo":2415919000}
	fmt.Println("cherry:", string(jc)) //=> cherry: {"Foo":2415919104}
}
```

## TODO List

* ✅ `float32`
* ✅ `*float32`
* ✅ `[]float32`
* ✅ `[]*float32`
* 🙅 `map[sometype]float32`
* 🙅 `map[sometype]*float32`
* ✅ 埋め込み構造体
* ✅ 構造体内構造体
