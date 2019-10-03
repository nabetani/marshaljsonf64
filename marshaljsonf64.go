package marshaljsonf64

import (
	"encoding/json"
	"fmt"
	"reflect"
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
	i := v.Interface()
	return fmt.Sprintf("%.20g", i)
}

func nameFromField(f reflect.StructField) *string {
	tag := f.Tag.Get("json")
	if tag == "-" {
		return nil // ignore this field
	}
	if tag == "" {
		return &f.Name
	}
	return &tag
}

func isFloat32Array(t reflect.Type) bool {
	return t.Kind() == reflect.Array && isFloat32(t.Elem())
}

func isFloat32Slice(t reflect.Type) bool {
	return t.Kind() == reflect.Slice && isFloat32(t.Elem())
}

func encodeF32List(v reflect.Value, f reflect.StructField, length int) string {
	items := []string{}
	for i := 0; i < length; i++ {
		items = append(items, formatFloat32(v.Index(i)))
	}
	return strings.Join(items, ",")
}

func collectJSONItems(t reflect.Type, o reflect.Value) ([]string, error) {
	items := []string{}
	numf := t.NumField()
	for n := 0; n < numf; n++ {
		f := t.Field(n)
		exported := f.PkgPath == ""
		if !exported {
			continue
		}
		v := o.FieldByName(f.Name)
		name := nameFromField(f)
		if name == nil {
			continue
		} else if isFloat32(f.Type) {
			items = append(items, fmt.Sprintf("%q:%s", *name, formatFloat32(v)))
		} else if isFloat32Array(f.Type) {
			items = append(items, fmt.Sprintf("%q:[%s]", *name, encodeF32List(v, f, v.Type().Len())))
		} else if isFloat32Slice(f.Type) {
			items = append(items, fmt.Sprintf("%q:[%s]", *name, encodeF32List(v, f, v.Len())))
		} else if f.Anonymous && "" == f.Tag {
			embedded, err := collectJSONItems(f.Type, v)
			if err != nil {
				return nil, err
			}
			items = append(items, embedded...)
		} else {
			j, err := json.Marshal(v.Interface())
			if err != nil {
				return nil, err
			}
			items = append(items, fmt.Sprintf("%q:%s", *name, j))
		}
	}
	return items, nil
}

// Impl メンバに float32 型の値がある場合に、 精度多めで出力するためのメソッド
func Impl(o interface{}, t reflect.Type) ([]byte, error) {
	oval := reflect.ValueOf(o).Elem()
	items, err := collectJSONItems(t, oval)
	if err != nil {
		return nil, err
	}
	return []byte("{" + strings.Join(items, ",") + "}"), nil
}
