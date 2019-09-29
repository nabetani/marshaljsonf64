package marshaljsonf64

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func isFloat32(t reflect.Type) bool {
	if t.Kind() == reflect.Ptr {
		return isFloat32(t.Elem())
	}
	return t.Kind() == reflect.Float32
}

func formatFloat32(v reflect.Value) string {
	if v.Type().Kind() == reflect.Ptr {
		if v.IsNil() || !v.IsValid() {
			return "null"
		}
		return formatFloat32(v.Elem())
	}
	f32 := v.Interface().(float32)
	return strconv.FormatFloat(float64(f32), 'f', -1, 64)
}

// Impl メンバに float32 型の値がある場合に、 精度多めで出力するためのメソッド
func Impl(o interface{}, t reflect.Type) ([]byte, error) {
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

	isFloat32Array := func(t reflect.Type) bool {
		return t.Kind() == reflect.Array && isFloat32(t.Elem())
	}

	isFloat32Slice := func(t reflect.Type) bool {
		return t.Kind() == reflect.Slice && isFloat32(t.Elem())
	}

	encodeF32List := func(v reflect.Value, f reflect.StructField, length int) string {
		items := []string{}
		for i := 0; i < length; i++ {
			items = append(items, formatFloat32(v.Index(i)))
		}
		return strings.Join(items, ",")
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
		v := oval.FieldByName(f.Name)
		name := nameFromField(f)
		if name == nil {
			continue
		} else if isFloat32(f.Type) {
			items = append(items, fmt.Sprintf("%q:%s", *name, formatFloat32(v)))
		} else if isFloat32Array(f.Type) {
			items = append(items, fmt.Sprintf("%q:[%s]", *name, encodeF32List(v, f, v.Type().Len())))
		} else if isFloat32Slice(f.Type) {
			items = append(items, fmt.Sprintf("%q:[%s]", *name, encodeF32List(v, f, v.Len())))
		} else {
			j, err := json.Marshal(v.Interface())
			if err != nil {
				return nil, err
			}
			items = append(items, fmt.Sprintf("%q:%s", *name, j))
		}
	}
	return []byte("{" + strings.Join(items, ",") + "}"), nil
}