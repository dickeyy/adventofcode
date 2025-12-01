package utils

import (
	"fmt"
	"reflect"
	"strings"
)

func Debug(data any) {
	fmt.Printf("== DEBUG ==\n%s\n== END DEBUG ==\n", formatData(data, 0))
}

func formatData(data any, indent int) string {
	if data == nil {
		return "nil\n"
	}

	v := reflect.ValueOf(data)
	indentStr := strings.Repeat("  ", indent)

	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		if v.Len() == 0 {
			return "[]\n"
		}

		result := "[\n"
		for i := 0; i < v.Len(); i++ {
			result += indentStr + "  " + formatValue(v.Index(i), indent+1)
			if i < v.Len()-1 {
				result += ","
			}
			result += "\n"
		}
		result += indentStr + "]"
		return result

	case reflect.Map:
		return formatMap(v, indent)

	default:
		return formatValue(v, indent) + "\n"
	}
}

func formatValue(v reflect.Value, indent int) string {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d", v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf("%d", v.Uint())
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%.2f", v.Float())
	case reflect.Bool:
		return fmt.Sprintf("%t", v.Bool())
	case reflect.String:
		return fmt.Sprintf("\"%s\"", v.String())
	case reflect.Struct:
		return formatStruct(v, indent)
	case reflect.Slice, reflect.Array:
		return formatSlice(v, indent)
	case reflect.Map:
		return formatMap(v, indent)
	case reflect.Ptr:
		if v.IsNil() {
			return "nil"
		}
		return "&" + formatValue(v.Elem(), indent)
	default:
		return fmt.Sprintf("%v", v.Interface())
	}
}

func formatStruct(v reflect.Value, indent int) string {
	t := v.Type()
	result := t.Name() + "{"

	for i := 0; i < v.NumField(); i++ {
		if i > 0 {
			result += ", "
		}
		fieldName := t.Field(i).Name
		fieldValue := formatValue(v.Field(i), indent)
		result += fieldName + ":" + fieldValue
	}

	result += "}"
	return result
}

func formatSlice(v reflect.Value, indent int) string {
	if v.Len() == 0 {
		return "[]"
	}

	// For small slices, format inline
	if v.Len() <= 10 && isSimpleType(v.Index(0)) {
		result := "["
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				result += ","
			}
			result += formatValue(v.Index(i), indent)
		}
		result += "]"
		return result
	}

	// For larger slices or complex types, format multiline
	result := "[\n" + strings.Repeat("  ", indent+1)
	for i := 0; i < v.Len(); i++ {
		if i > 0 {
			result += ", "
		}
		result += formatValue(v.Index(i), indent+1)
	}
	result += "\n" + strings.Repeat("  ", indent) + "]"
	return result
}

func formatMap(v reflect.Value, indent int) string {
	if v.Len() == 0 {
		return "{}\n"
	}

	result := "{\n"
	keys := v.MapKeys()
	indentStr := strings.Repeat("  ", indent+1)

	for _, key := range keys {
		result += indentStr + formatValue(key, indent) + ": " + formatValue(v.MapIndex(key), indent) + "\n"
	}
	result += strings.Repeat("  ", indent) + "}"
	return result
}

func isSimpleType(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64, reflect.Bool, reflect.String:
		return true
	default:
		return false
	}
}
