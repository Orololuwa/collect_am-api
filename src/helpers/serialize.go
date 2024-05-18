package helpers

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

// SerializeStruct serializes the data from srcStruct into dstStruct based on tags.
func SerializeStruct(srcStruct interface{}, dstStructPtr interface{}) error {

	// if srcStruct == nil {
	// 	log.Println("@here")
	// 	dstStructPtr = nil
	// 	return nil
	// }

    srcValue := reflect.ValueOf(srcStruct)
    if srcValue.Kind() == reflect.Ptr {
        srcValue = srcValue.Elem() // Dereference the pointer
    }

    dstStruct := reflect.Indirect(reflect.ValueOf(dstStructPtr))
    if srcValue.Kind() != reflect.Struct || dstStruct.Kind() != reflect.Struct  {
        return errors.New("srcStruct and dstStructPtr must be structs")
    }

    srcType := srcValue.Type()
    for i := 0; i < srcType.NumField(); i++ {
        srcField := srcType.Field(i)
        srcFieldValue := srcValue.Field(i)
        dstField, ok := dstStruct.Type().FieldByName(srcField.Name)
        if !ok {
            continue // Skip if field not found in destination struct
        }

        tag := dstField.Tag.Get("json")
        if tag == "" || tag == "-" || tag == "omitempty" {
            continue // Skip if tag is empty or "-"
        }

        dstStructField := dstStruct.FieldByName(dstField.Name)
        if !dstStructField.CanSet() {
            continue // Skip if field cannot be set (unexported or not addressable)
        }

		// special handling for bool strings
		dataTag := dstField.Tag.Get("dataType")
        if dataTag != "" && dataTag == "bool"{
			// Convert src field value to bool
			var boolValue bool
			switch srcFieldValue.Kind() {
			case reflect.String:
				// Convert string to bool
				boolValue, err := StringToBool(srcFieldValue.String())
				if err != nil {
					return fmt.Errorf("error converting string to bool: %v", err)
				}
				dstStructField.SetBool(boolValue)
			case reflect.Bool:
				// Already a bool value, just set it
				boolValue = srcFieldValue.Bool()
				dstStructField.SetBool(boolValue)
			default:
				return fmt.Errorf("unexpected type for bool conversion: %v", srcFieldValue.Kind())
			}
			continue
        }


		// Special handling for time.Time fields
		if srcFieldValue.Type() == reflect.TypeOf(time.Time{}) {
			if !srcFieldValue.Interface().(time.Time).IsZero() {
				dstStructField.Set(srcFieldValue)
			}
			continue
		}

        switch srcFieldValue.Kind() {
			case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
				reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
				reflect.Float32, reflect.Float64, reflect.String:
				dstStructField.Set(srcFieldValue)
			case reflect.Struct:
				if dstStructField.Kind() != reflect.Struct {
					continue // Skip if destination field is not a struct
				}
				if err := SerializeStruct(srcFieldValue.Interface(), dstStructField.Addr().Interface()); err != nil {
					return err // Propagate error if nested struct serialization fails
				}
			default:
				// Add support for other types if needed
				dstStructField.Set(srcFieldValue)
        }
    }

    return nil
}

