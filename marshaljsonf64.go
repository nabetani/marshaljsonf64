package marshaljsonf64

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// MarshalJSONF64 メンバに float32 型の値がある場合に、 精度多めで出力するためのメソッド
func MarshalJSONF64(o interface{}, t reflect.Type) ([]byte, error) {
	nameFromField := func(f reflect.StructField) *string {
		tag := f.Tag.Get("json")
		if tag == "-" {
			return nil // ignore this field
		}
		if tag == "" {
			return &f.Name
		}
		return &tag
	}

	isFloat32Ptr := func(t reflect.Type) bool {
		return t.Kind() == reflect.Ptr &&
			t.Elem().Kind() == reflect.Float32
	}

	numf := t.NumField()
	oval := reflect.ValueOf(o).Elem()
	items := []string{}
	for n := 0; n < numf; n++ {
		f := t.Field(n)
		exported := f.PkgPath == ""
		if !exported {
			continue
		}
		v := oval.FieldByName(f.Name).Interface()
		name := nameFromField(f)
		if name == nil {
			continue
		}
		if f.Type.Kind() == reflect.Float32 {
			s := strconv.FormatFloat(float64(v.(float32)), 'f', -1, 64)
			items = append(items, fmt.Sprintf("%q:%s", *name, s))
		} else if isFloat32Ptr(f.Type) {
			s := strconv.FormatFloat(float64(*v.(*float32)), 'f', -1, 64)
			items = append(items, fmt.Sprintf("%q:%s", *name, s))
		} else {
			j, err := json.Marshal(v)
			if err != nil {
				return nil, err
			}
			items = append(items, fmt.Sprintf("%q:%s", *name, j))
		}
	}
	return []byte("{" + strings.Join(items, ",") + "}"), nil
}
