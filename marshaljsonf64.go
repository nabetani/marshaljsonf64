package marshaljsonf64

import (
	"encoding/json"
	"fmt"
	"math"
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

func errorUnlessFinite(v reflect.Value) error {
	f := v.Float()
	if math.IsInf(f, 0) || math.IsNaN(f) {
		return &json.UnsupportedValueError{v, strconv.FormatFloat(f, 'g', -1, 64)}
	}
	return nil
}

func formatFloat32(v reflect.Value) (string, error) {
	if v.Type().Kind() == reflect.Ptr {
		if v.IsNil() || !v.IsValid() {
			return "null", nil
		}
		return formatFloat32(v.Elem())
	}
	if err := errorUnlessFinite(v); err != nil {
		return "", err
	}
	return fmt.Sprintf("%.20g", v.Float()), nil
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

func encodeF32List(v reflect.Value, f reflect.StructField, length int) (string, error) {
	items := []string{}
	for i := 0; i < length; i++ {
		str, err := formatFloat32(v.Index(i))
		if err != nil {
			return "", err
		}
		items = append(items, str)
	}
	return strings.Join(items, ","), nil
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
			str, err := formatFloat32(v)
			if err != nil {
				return nil, err
			}
			items = append(items, fmt.Sprintf("%q:%s", *name, str))
		} else if isFloat32Array(f.Type) {
			str, err := encodeF32List(v, f, v.Type().Len())
			if err != nil {
				return nil, err
			}
			items = append(items, fmt.Sprintf("%q:[%s]", *name, str))
		} else if isFloat32Slice(f.Type) {
			str, err := encodeF32List(v, f, v.Len())
			if err != nil {
				return nil, err
			}
			items = append(items, fmt.Sprintf("%q:[%s]", *name, str))
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
