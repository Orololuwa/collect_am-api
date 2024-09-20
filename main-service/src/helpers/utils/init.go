package utils

import (
	"errors"
	"fmt"
	"reflect"
	"unicode"
)

func Capitalize(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func CamelToSnakeCase(s string) string {
	var result []rune
	for i, r := range s {
		if unicode.IsUpper(r) {
			// Add an underscore before the uppercase letter, if it's not the first character
			if i > 0 {
				result = append(result, '_')
			}
			// Convert the letter to lowercase
			result = append(result, unicode.ToLower(r))
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}

// Expected field type information
type FieldInfo struct {
	Type reflect.Kind
}

func ValidateMap(body map[string]interface{}, fields map[string]FieldInfo, convertToSnakeCase bool) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	for key, value := range body {
		if fieldInfo, ok := fields[key]; ok {
			valueType := reflect.TypeOf(value).Kind()

			// Check type and handle int to float64 conversion
			if valueType != fieldInfo.Type {
				if fieldInfo.Type == reflect.Float64 && (valueType == reflect.Int || valueType == reflect.Int8 || valueType == reflect.Int16 || valueType == reflect.Int32 || valueType == reflect.Int64) {
					// Convert int to float64
					value = float64(reflect.ValueOf(value).Int())
				} else {
					msg := fmt.Sprintf("Invalid type for key %s: expected %s, got %s\n", key, fieldInfo.Type, valueType)
					return result, errors.New(msg)
				}
			}

			// Optionally convert key to snake case
			resKey := key
			if convertToSnakeCase {
				resKey = CamelToSnakeCase(key)
			}

			if resKey != "" {
				result[resKey] = value
			}
		}
	}

	return result, nil
}
