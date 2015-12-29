package ddb

import (
	"errors"
	"reflect"
	"strconv"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Unmarshal converts dynamodb attribute value map to map or struct
func Unmarshal(item map[string]*dynamodb.AttributeValue, v interface{}) error {
	t := reflect.TypeOf(v)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	} else {
		return errors.New("value must be a pointer")
	}

	isptr := false
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		isptr = true
	}

	if t.Kind() == reflect.Struct {
		dest := reflect.ValueOf(v)
		if isptr {
			dest = dest.Elem()
		}

		if dest.IsNil() {
			dest.Set(reflect.New(t))
		}

		destType := dest.Elem().Type()
		numOfField := destType.NumField()
		for i := 0; i < numOfField; i++ {
			f := destType.Field(i)

			if f.PkgPath != "" {
				continue
			}

			name, option := parseTag(f.Tag.Get("dynamodb"))
			if name == "-" {
				continue
			}
			if option == "omitifempty" {
				continue
			}
			if name == "" {
				name = f.Name
			}

			if value, ok := item[name]; ok {
				targetField := dest.Elem().FieldByIndex(f.Index)

				switch f.Type.Kind() {
				case reflect.String:
					if value.S != nil {
						targetField.SetString(*value.S)
					}
				case reflect.Bool:
					if value.BOOL != nil {
						targetField.SetBool(*value.BOOL)
					}
				case reflect.Int:
					targetField.SetInt(parseIntAttrValue(value, 0))
				case reflect.Int8:
					targetField.SetInt(parseIntAttrValue(value, 8))
				case reflect.Int16:
					targetField.SetInt(parseIntAttrValue(value, 16))
				case reflect.Int32:
					targetField.SetInt(parseIntAttrValue(value, 32))
				case reflect.Int64:
					targetField.SetInt(parseIntAttrValue(value, 64))
				case reflect.Uint:
					targetField.SetUint(parseUintAttrValue(value, 0))
				case reflect.Uint8:
					targetField.SetUint(parseUintAttrValue(value, 8))
				case reflect.Uint16:
					targetField.SetUint(parseUintAttrValue(value, 16))
				case reflect.Uint32:
					targetField.SetUint(parseUintAttrValue(value, 32))
				case reflect.Uint64:
					targetField.SetUint(parseUintAttrValue(value, 64))
				case reflect.Float32:
					targetField.SetFloat(parseFloatAttrValue(value, 32))
				case reflect.Float64:
					targetField.SetFloat(parseFloatAttrValue(value, 64))
				}
			}
		}

	}

	return nil
}

func parseIntAttrValue(value *dynamodb.AttributeValue, bitSize int) int64 {
	n, _ := strconv.ParseInt(*value.N, 10, bitSize)
	return n
}

func parseUintAttrValue(value *dynamodb.AttributeValue, bitSize int) uint64 {
	n, _ := strconv.ParseUint(*value.N, 10, bitSize)
	return n
}

func parseFloatAttrValue(value *dynamodb.AttributeValue, bitSize int) float64 {
	f, _ := strconv.ParseFloat(*value.N, bitSize)
	return f
}